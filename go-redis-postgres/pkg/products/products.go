package products

import (
	"github.com/go-redis/redis"
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

func getProducts() (*JsonResponse, error) {
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
}
