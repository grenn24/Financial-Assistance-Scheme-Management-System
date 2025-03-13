package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/grenn24/financial-assistance-scheme-management-system/services"
)

type SchemeController struct {
	Db *gorm.DB
	SchemeService *services.SchemeService
}

func (this *SchemeController) GetAllSchemes(context *gin.Context) {
	schemes, err := this.SchemeService.GetAllSchemes()
	if (err != nil) {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, schemes);
}

func (this *SchemeController) GetSchemeByID(context *gin.Context) {
	ID := context.Param("ID")
	scheme, err := this.SchemeService.GetSchemeByID(ID)
	if (err != nil) {
		if (err == gorm.ErrRecordNotFound) {
			context.JSON(404, gin.H{"error": "Scheme not found"})
			return
		}
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, scheme);
}

func (this *SchemeController) CreateScheme(context *gin.Context) {

}

// Delete and return deleted scheme
func (this *SchemeController) DeleteSchemeByID(context *gin.Context) {
	ID := context.Param("ID")
	scheme, err := this.SchemeService.GetSchemeByID(ID)
	if (err != nil) {
		if (err == gorm.ErrRecordNotFound) {
			context.JSON(404, gin.H{"error": "Scheme not found"})
			return
		}
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, scheme);
}
