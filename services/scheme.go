package services

import (
	models "github.com/grenn24/financial-assistance-scheme-management-system/models/Scheme"
	"gorm.io/gorm"
)

type SchemeService struct {
	Db *gorm.DB
}

func (this *SchemeService) GetAllSchemes() ([]models.Scheme, error) {
	var schemes []models.Scheme
	result := this.Db.Find(&schemes)

	if result.Error != nil {
		return nil, result.Error
	}
	return schemes, nil
}

func (this *SchemeService) GetSchemeByID(id string) (*models.Scheme, error) {
	var scheme models.Scheme
	result := this.Db.First(&scheme, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &scheme, nil
}

func (this *SchemeService) DeleteSchemeByID(id string) (*models.Scheme, error) {

	scheme, err := this.GetSchemeByID(id)
	if err != nil {
		return nil, err
	}
	result := this.Db.Delete(&models.Scheme{}, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return scheme, nil
}
