package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grenn24/financial-assistance-scheme-management-system/controllers"
	"gorm.io/gorm"
)

func ApplicantRoutes(router *gin.RouterGroup, db *gorm.DB) {
	applicantRouter := router.Group("/applicants")

	applicantController := &controllers.ApplicantController{Db:db}

	applicantRouter.GET("/", func(context *gin.Context) {
		applicantController.GetAllApplicants(context, db)
	})

	applicantRouter.POST("/", func(context *gin.Context) {
		applicantController.CreateApplicant(context, db)
	})

	applicantRouter.DELETE("/:ID", func(context *gin.Context) {
		applicantController.DeleteApplicantByID(context, db)
	})
}
