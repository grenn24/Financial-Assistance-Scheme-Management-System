package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grenn24/financial-assistance-scheme-management-system/controllers"
	"github.com/grenn24/financial-assistance-scheme-management-system/services"
	"gorm.io/gorm"
)

func ApplicantRoutes(router *gin.RouterGroup, db *gorm.DB) {
	applicantRouter := router.Group("/applicants")

	applicantController := &controllers.ApplicantController{Db: db, ApplicantService: &services.ApplicantService{Db: db}}

	applicantRouter.GET("/", func(context *gin.Context) {
		applicantController.GetAllApplicants(context)
	})

	applicantRouter.GET("/:ID", func(context *gin.Context) {
		applicantController.GetApplicantByID(context)
	})

	applicantRouter.POST("/", func(context *gin.Context) {
		applicantController.CreateApplicant(context)
	})

	applicantRouter.PUT("/:ID", func(context *gin.Context) {
		applicantController.UpdateApplicant(context)
	})

	applicantRouter.DELETE("/:ID", func(context *gin.Context) {
		applicantController.DeleteApplicantByID(context)
	})

	applicantRouter.DELETE("/", func(context *gin.Context) {
		applicantController.DeleteAllApplicants(context)
	})
}
