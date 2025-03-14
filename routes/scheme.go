package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grenn24/financial-assistance-scheme-management-system/controllers"
	"github.com/grenn24/financial-assistance-scheme-management-system/services"
	"gorm.io/gorm"
)

func SchemeRoutes(router *gin.RouterGroup, db *gorm.DB) {
	schemeRouter := router.Group("/schemes")
	schemeController := &controllers.SchemeController{Db: db, SchemeService: &services.SchemeService{Db: db}}

	schemeRouter.GET("/", func(context *gin.Context) {
		schemeController.GetAllSchemes(context)
	})

	schemeRouter.GET("/:ID", func(context *gin.Context) {
		schemeController.GetSchemeByID(context)
	})

	schemeRouter.GET("/eligible", func(context *gin.Context) {
		schemeController.GetEligibleSchemes(context)
	})

	schemeRouter.POST("/", func(context *gin.Context) {
		schemeController.CreateScheme(context)
	})

	schemeRouter.PUT("/:ID", func(context *gin.Context) {
		schemeController.UpdateScheme(context)
	})

	schemeRouter.DELETE("/", func(context *gin.Context) {
		schemeController.DeleteAllSchemes(context)
	})

	schemeRouter.DELETE("/:ID", func(context *gin.Context) {
		schemeController.DeleteSchemeByID(context)
	})

}
