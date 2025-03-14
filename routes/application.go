package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grenn24/financial-assistance-scheme-management-system/controllers"
	"github.com/grenn24/financial-assistance-scheme-management-system/services"
	"gorm.io/gorm"
)

func ApplicantionRoutes(router *gin.RouterGroup, db *gorm.DB) {
	applicationRouter := router.Group("/applications")

	applicationController := &controllers.ApplicationController{Db: db, ApplicationService: &services.ApplicationService{Db: db}}

	applicationRouter.GET("/", func(context *gin.Context) {
		applicationController.GetAllApplications(context, db)
	})

	applicationRouter.POST("/", func(context *gin.Context) {
		applicationController.CreateApplication(context, db)
	})

	applicationRouter.DELETE("/:ID", func(context *gin.Context) {
		applicationController.DeleteApplicationByID(context, db)
	})
}
