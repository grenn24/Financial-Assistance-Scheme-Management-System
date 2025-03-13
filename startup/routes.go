package startup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grenn24/financial-assistance-scheme-management-system/routes"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {
	// Declare api router group
	apiRouter := router.Group("/api")

	// Applicant Related Routes
	routes.ApplicantRoutes(apiRouter, db)

	// Application Related Routes
	routes.ApplicantionRoutes(apiRouter, db)

	// Scheme Related Routes
	routes.SchemeRoutes(apiRouter, db)

	// Missed Routes
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"Status":    "error",
			"ErrorCode": "INVALID_API_ROUTE",
			"Message":   "This route does not exist on the api",
		})
	})
}
