package products

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

type Products struct {
	ProductId   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	RetailPrice float64 `json:"retail_price"`
}

type JsonResponse struct {
	Data   []Products `json:"data"`
	Source string     `json:"source"`
}

func GetProducts() (*JsonResponse, error) {
	redisClient := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               "localhost:6379",
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	//ctx := context.Background()

	productsCache, errorResponse := redisClient.Get("products").Bytes()

	jsonResponse := JsonResponse{}

	if errorResponse != nil {
		productsDBResponse, errorResponse := fetchFromDB()
		if errorResponse != nil {
			return nil, errorResponse
		}

		cachedProducts, errorResponse := json.Marshal(productsDBResponse)
		if errorResponse != nil {
			return nil, errorResponse
		}

		redisErrorResponse := redisClient.Set("products", cachedProducts, 10*time.Second).Err()
		if redisErrorResponse != nil {
			return nil, redisErrorResponse
		}
		jsonResponse = JsonResponse{Data: productsDBResponse, Source: "Postgresql"}
		return &jsonResponse, redisErrorResponse
	}

	productsObject := []Products{}
	errorResponse = json.Unmarshal(productsCache, &productsObject)

	if errorResponse != nil {
		return nil, errorResponse
	}

	jsonResponse = JsonResponse{Data: productsObject, Source: "RedisCache"}

	return &jsonResponse, nil
}

func fetchFromDB() ([]Products, error) {

	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "sample_company"

	conString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", conString)

	if err != nil {
		return nil, err
	}

	queryString := `select
                     product_id,
                     product_name,
                     retail_price
                 from products`

	rows, err := db.Query(queryString)

	if err != nil {
		return nil, err
	}

	var records []Products

	for rows.Next() {

		var p Products
		err = rows.Scan(&p.ProductId, &p.ProductName, &p.RetailPrice)
		records = append(records, p)
		if err != nil {
			return nil, err
		}
	}

	return records, nil
}
