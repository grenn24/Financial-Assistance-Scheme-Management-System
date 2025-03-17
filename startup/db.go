package startup

import (
	"fmt"
	"log"
	"os"

	"github.com/grenn24/financial-assistance-scheme-management-system/models"
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

	//Reference the db pointer to gorm.DB instance
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	// Check for connection error
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}
	fmt.Println("Database connection successful!")
	

	// Create enum types
	db.Exec(`
		DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'marital_status') THEN
				CREATE TYPE marital_status AS ENUM ('single', 'married', 'widowed', 'divorced');
			END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'sex') THEN
				CREATE TYPE sex AS ENUM ('male', 'female');
			END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'relation') THEN
				CREATE TYPE relation AS ENUM ('husband', 'wife', 'son', 'daughter', 'brother', 'sister');
			END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
				CREATE TYPE status AS ENUM ('pending', 'approved', 'rejected');
			END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'school_level') THEN
				CREATE TYPE school_level AS ENUM ('primary', 'secondary', 'tertiary');
			END IF;
		END $$;

		CREATE EXTENSION IF NOT EXISTS pgcrypto;
	`)

	// Create tables if they do not exist
	err = db.AutoMigrate(&models.Scheme{}, &models.Applicant{}, &models.Application{}, &models.SchemeBenefit{}, &models.SchemeCriteria{}, &models.HouseholdMember{})
	if err != nil {
		log.Fatal("Error migrating database schema: ", err.Error())
	}

	return db
}
