package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/grenn24/financial-assistance-scheme-management-system/models"
	"github.com/grenn24/financial-assistance-scheme-management-system/services"
	"gorm.io/gorm"
)

type ApplicationController struct {
	Db                 *gorm.DB
	ApplicationService *services.ApplicationService
}

func (applicationController *ApplicationController) GetAllApplications(context *gin.Context) {
	applications, err := applicationController.ApplicationService.GetAllApplications()
	if err != nil {
		context.JSON(500, gin.H{"message": err.Error(), "status": "INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, applications)
}

func (applicationController *ApplicationController) GetApplicationByID(context *gin.Context) {
	id := context.Param("ID")
	// Validate id
	err := uuid.Validate(id)
	if err != nil {
		context.JSON(404, gin.H{"status": "VALIDATION_ERROR", "message": "INVALID_ID_FORMAT"})
		return
	}
	application, err := applicationController.ApplicationService.GetApplicationByID(id)
	if err != nil {
		context.JSON(500, gin.H{"message": err.Error(), "status": "INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, application)
}

func (applicationController *ApplicationController) CreateApplication(context *gin.Context) {
	application := new(models.Application)
	hasMultiple := context.Query("multiple")
	if (hasMultiple == "true") {
		applicationController.CreateMultipleApplications(context)
		return
	}

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

func (applicationController *ApplicationController) CreateMultipleApplications(context *gin.Context) {
	applicationRequest := new(models.CreateMultipleApplicationsRequest)
	applications := make([]*models.Application, 0)
	// Bind http request body into struct
	if err := context.ShouldBind(applicationRequest); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	// Validate http request body
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(applicationRequest); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			context.JSON(http.StatusBadRequest, gin.H{"status": "VALIDATION_ERROR",
				"message": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	for _, schemeID := range applicationRequest.SchemeID {

		application, err := applicationController.ApplicationService.CreateApplication(&models.Application{
			ApplicantID: applicationRequest.ApplicantID,
			SchemeID:    schemeID,
			Status:      applicationRequest.Status,
		})

		if err != nil {
			context.JSON(500, gin.H{"status": "INTERNAL_SERVER_ERROR",
				"message": err.Error()})
			return
		}
		applications = append(applications, application)
	}

	context.JSON(200, applications)

}

func (applicationController *ApplicationController) UpdateApplication(context *gin.Context) {
	id := context.Param("ID")
	// Validate id

	if err := uuid.Validate(id); err != nil {
		context.JSON(404, gin.H{"status": "VALIDATION_ERROR", "message": "INVALID_ID_FORMAT"})
		return
	}
	application := new(models.UpdateApplicationRequest)

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

	updatedApplication, err := applicationController.ApplicationService.UpdateApplication(application, id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(404, gin.H{"message": "Application not found"})
			return
		}
		context.JSON(500, gin.H{"message": err.Error(), "status": "INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, updatedApplication)

}

func (applicationController *ApplicationController) DeleteApplicationByID(context *gin.Context) {
	id := context.Param("ID")
	// Validate id
	err := uuid.Validate(id)
	if err != nil {
		context.JSON(404, gin.H{"status": "VALIDATION_ERROR", "message": "INVALID_ID_FORMAT"})
		return
	}
	application, err := applicationController.ApplicationService.DeleteApplicationByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(404, gin.H{"message": "Application not found"})
			return
		}
		context.JSON(500, gin.H{"message": err.Error(), "status": "INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, application)
}

func (applicationController *ApplicationController) DeleteAllApplications(context *gin.Context) {
	applicationsDeleted, err := applicationController.ApplicationService.DeleteAllApplications()
	if err != nil {
		context.JSON(500, gin.H{"message": err.Error(), "status": "INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, gin.H{
		"applicationsDeleted": applicationsDeleted,
	})
}
