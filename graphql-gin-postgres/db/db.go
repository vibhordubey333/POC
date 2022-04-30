package database

import(
	"fmt"
	"github.com/jinzhu/gorm"
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
	user:     "postgres",
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
	if !db.HasTable(&model.Ques)
}
