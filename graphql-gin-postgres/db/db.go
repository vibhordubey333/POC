package database

import(
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/vibhordubey333/POC/graphql-gin-postgres/graphql_poc/api/models"
)

type dbConfig struct{
	host string
	port int
	user string 
	dbname string 
	password string
}

var config = dbConfig{
	host:     "localhost",
	port:     5432,
	user:     "postgres1",
	dbname:   "test",
	password: "root",
}

func getDatabaseUrl() string{
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",config.host,config.port,config.user,config.dbname,config.password)
}

func GetDatabase() (*gorm.DB,error){
	db,errorResponse := gorm.Open("postgres",getDatabaseUrl())
	return db,errorResponse
}

func RunMigration(db *gorm.DB){
	if !db.HasTable(&models.Question{}){
		db.CreateTable(&models.Question{})
	}
	if !db.HasTable(&models.Choice{}){
		db.CreateTable(&models.Choice{})
		db.Model(&models.Choice{}).AddForeignKey("question_id,", "questions(id)","CASCADE","CASCADE")
	}
}
