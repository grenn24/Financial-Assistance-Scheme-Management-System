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

type ApplicantController struct {
	Db               *gorm.DB
	ApplicantService *services.ApplicantService
}

func (applicantController *ApplicantController) GetAllApplicants(context *gin.Context) {
	applicants, err := applicantController.ApplicantService.GetAllApplicants()
	if err != nil {
		context.JSON(500, gin.H{"message": err.Error(), "status":"INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, applicants)
}

func (applicantController *ApplicantController) GetApplicantByID(context *gin.Context) {
	id := context.Param("ID")
	// Validate id
	err := uuid.Validate(id)
	if err != nil {
		context.JSON(404, gin.H{"status": "VALIDATION_ERROR", "message": "INVALID_ID_FORMAT"})
		return
	}
	applicant, err := applicantController.ApplicantService.GetApplicantByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(404, gin.H{"status":"BAD_REQUEST", "message": "Applicant not found"})
			return
		}
		context.JSON(500, gin.H{"message": err.Error(), "status":"INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, applicant)
}

func (applicantController *ApplicantController) CreateApplicant(context *gin.Context) {
	applicant := new(models.Applicant)

	// Bind http request body into struct
	if err := context.ShouldBind(applicant); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	// Validate http request body
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(applicant); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			context.JSON(http.StatusBadRequest, gin.H{"status": "VALIDATION_ERROR",
				"message": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	applicant, err := applicantController.ApplicantService.CreateApplicant(applicant)

	if err != nil {
		context.JSON(500, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}
	context.JSON(200, applicant)
}

func (applicantController *ApplicantController) UpdateApplicant(context *gin.Context) {
	id := context.Param("ID")
	// Validate id

	if err := uuid.Validate(id) ; err != nil {
		context.JSON(404, gin.H{"message": "INVALID_ID_FORMAT"})
		return
	}
	applicant := new(models.UpdateApplicantRequest)

	// Bind http request body into struct
	if err := context.ShouldBind(applicant); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	// Validate http request body
	validate := validator.New()
	if err := validate.Struct(applicant); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			context.JSON(http.StatusBadRequest, gin.H{"status": "VALIDATION_ERROR",
				"message": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	updatedApplicant, err := applicantController.ApplicantService.UpdateApplicant(applicant, id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(404, gin.H{"status": "VALIDATION_ERROR","message": "Applicant not found"})
			return
		}
		context.JSON(500, gin.H{"message": err.Error(), "status":"INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, updatedApplicant)
}

func (applicantController *ApplicantController) DeleteApplicantByID(context *gin.Context) {
	id := context.Param("ID")
	// Validate id
	err := uuid.Validate(id)
	if err != nil {
		context.JSON(404, gin.H{"message": "INVALID_ID_FORMAT"})
		return
	}
	applicant, err := applicantController.ApplicantService.DeleteApplicantByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(404, gin.H{"status": "VALIDATION_ERROR","message": "Applicant not found"})
			return
		}
		context.JSON(500, gin.H{"message": err.Error(), "status":"INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, applicant)
}

func (applicantController *ApplicantController) DeleteAllApplicants(context *gin.Context) {
	applicantsDeleted, err := applicantController.ApplicantService.DeleteAllApplicants()
	if err != nil {
		context.JSON(500, gin.H{"message": err.Error(), "status":"INTERNAL_SERVER_ERROR"})
		return
	}
	context.JSON(200, gin.H{
		"applicantsDeleted": applicantsDeleted,
	})
}
