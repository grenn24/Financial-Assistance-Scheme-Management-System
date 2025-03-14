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

type SchemeController struct {
	Db            *gorm.DB
	SchemeService *services.SchemeService
}

func (schemeController *SchemeController) GetAllSchemes(context *gin.Context) {
	schemes, err := schemeController.SchemeService.GetAllSchemes()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, schemes)
}

func (schemeController *SchemeController) GetSchemeByID(context *gin.Context) {
	id := context.Param("ID")
	// Validate id
	err := uuid.Validate(id)
	if err != nil {
		context.JSON(404, gin.H{"error": "INVALID_ID_FORMAT"})
		return
	}
	scheme, err := schemeController.SchemeService.GetSchemeByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(404, gin.H{"error": "Scheme not found"})
			return
		}
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, scheme)
}

func (schemeController *SchemeController) CreateScheme(context *gin.Context) {
	scheme := new(models.Scheme)

	// Bind http request body into struct
	if err := context.ShouldBind(scheme); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	// Validate http request body
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(scheme); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			context.JSON(http.StatusBadRequest, gin.H{"status": "VALIDATION_ERROR",
				"message": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}

	scheme, err := schemeController.SchemeService.CreateScheme(scheme)

	if err != nil {
		context.JSON(500, gin.H{"status": "INTERNAL_SERVER_ERROR",
			"message": err.Error()})
		return
	}
	context.JSON(200, scheme)

}

// Delete and return deleted scheme
func (schemeController *SchemeController) DeleteSchemeByID(context *gin.Context) {
	
	id := context.Param("ID")
	// Validate id
	err := uuid.Validate(id)
	if err != nil {
		context.JSON(404, gin.H{"error": "INVALID_ID_FORMAT"})
		return
	}
	scheme, err := schemeController.SchemeService.GetSchemeByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(404, gin.H{"error": "Scheme not found"})
			return
		}
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, scheme)
}

func (schemeController *SchemeController) DeleteAllSchemes(context *gin.Context) {

	schemesDeleted, err := schemeController.SchemeService.DeleteAllSchemes()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{
		"schemesDeleted": schemesDeleted,
	})
}