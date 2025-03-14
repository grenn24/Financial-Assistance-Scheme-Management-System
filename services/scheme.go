package services

import (
	 "github.com/grenn24/financial-assistance-scheme-management-system/models"
	"gorm.io/gorm"
)

type SchemeService struct {
	Db *gorm.DB
}

func (schemeService *SchemeService) GetAllSchemes() ([]models.Scheme, error) {
	var schemes []models.Scheme
	result := schemeService.Db.Find(&schemes)

	if result.Error != nil {
		return nil, result.Error
	}
	return schemes, nil
}

func (schemeService *SchemeService) GetSchemeByID(id string) (*models.Scheme, error) {
	var scheme models.Scheme
	result := schemeService.Db.First(&scheme,  "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &scheme, nil
}

func (schemeService *SchemeService) CreateScheme(scheme *models.Scheme) (*models.Scheme, error) {

	result := schemeService.Db.Create(&scheme)

	if result.Error != nil {
		return nil, result.Error
	}

	return scheme, nil
}

func (schemeService *SchemeService) DeleteSchemeByID(id string) (*models.Scheme, error) {

	scheme, err := schemeService.GetSchemeByID(id)
	if err != nil {
		return nil, err
	}
	result := schemeService.Db.Delete(&models.Scheme{},  "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	return scheme, nil
}


func (schemeService *SchemeService) DeleteAllSchemes() (int, error) {

	result := schemeService.Db.Where("true").Delete(&models.Scheme{})

	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}