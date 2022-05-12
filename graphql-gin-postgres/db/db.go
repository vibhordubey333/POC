package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/vibhordubey333/POC/graphql-gin-postgres/graphql_poc/api/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
	sslmode  string
}

var config = dbConfig{
	host:     "localhost",
	port:     5432,
	user:     "postgres1",
	dbname:   "test",
	password: "root",
	sslmode:  "disable",
}

func getDatabaseUrl() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.host, config.port, config.user, config.dbname, config.password,config.sslmode)
}

func GetDatabase() (*gorm.DB, error) {
	connStr := "host=localhost port=5432 user=postgres dbname=test password=postgres sslmode=disable"
	db, errorResponse := gorm.Open("postgres", connStr)
	return db, errorResponse
}

func RunMigration(db *gorm.DB) {
	if !db.HasTable(&models.Question{}) {
		db.CreateTable(&models.Question{})
	}
	if !db.HasTable(&models.Choice{}) {
		db.CreateTable(&models.Choice{})
		db.Model(&models.Choice{}).AddForeignKey("question_id,", "questions(id)", "CASCADE", "CASCADE")
	}
}
