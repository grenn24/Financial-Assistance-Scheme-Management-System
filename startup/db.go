package startup

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	var err error

	conn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		os.Getenv("DB_user"),
		os.Getenv("DB_password"),
		os.Getenv("DB_host"),
		os.Getenv("DB_port"),
		os.Getenv("DB_name"))

	//Reference the db pointer to sql.DB instance
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})


	//check for connection string verification error
	if err != nil {
		log.Fatal("Error connecting to the database: ", err.Error())
	}
	fmt.Println("Database connection successful!")
	return db
}
