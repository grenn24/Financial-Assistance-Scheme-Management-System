package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/grenn24/financial-assistance-scheme-management-system/models"
	"github.com/grenn24/financial-assistance-scheme-management-system/services"
	"gorm.io/gorm"
)

type ApplicationController struct {
	Db                 *gorm.DB
	ApplicationService *services.ApplicationService
}

func (applicationController *ApplicationController) GetAllApplications(context *gin.Context, db *gorm.DB) {
	applications, err := applicationController.ApplicationService.GetAllApplications()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, applications)
}

func (applicationController *ApplicationController) GetApplicationByID(context *gin.Context, db *gorm.DB) {
	ID := context.Param("ID")
	application, err := applicationController.ApplicationService.GetApplicationByID(ID)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, application)
}

func (applicationController *ApplicationController) CreateApplication(context *gin.Context, db *gorm.DB) {
	application := new(models.Application)

	// Bind http request body into struct
	if err := context.ShouldBind(application); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	// Validate http request body
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(application); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			context.JSON(http.StatusBadRequest, gin.H{"status": "VALIDATION_ERROR",
				"message": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	application, err := applicationController.ApplicationService.CreateApplication(application)

	if err != nil {
		context.JSON(500, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}
	context.JSON(200, application)

}

func (applicationController *ApplicationController) DeleteApplicationByID(context *gin.Context, db *gorm.DB) {
	ID := context.Param("ID")
	application, err := applicationController.ApplicationService.GetApplicationByID(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(404, gin.H{"error": "Application not found"})
			return
		}
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, application)
}
