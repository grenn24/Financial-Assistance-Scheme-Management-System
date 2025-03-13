package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/grenn24/financial-assistance-scheme-management-system/startup"

	"gorm.io/gorm"
)

// Pointer to database struct
var db *gorm.DB

func main() {
	// Load env variables from .env
	startup.Env()

	// Create a pointer to gorm instance
	db := startup.Db()

	//Create a pointer to a gin.Engine instance
	var router *gin.Engine = gin.Default()
	startup.Routes(router, db)

	// Start the backend server
	fmt.Printf("Server is running on %v%v:%v\n", os.Getenv("PROTOCOL"),os.Getenv("DOMAIN_NAME"), os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}
