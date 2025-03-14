package services

import (
	"fmt"

	"github.com/grenn24/financial-assistance-scheme-management-system/models"
	"gorm.io/gorm"
)

type SchemeService struct {
	Db *gorm.DB
}

func (schemeService *SchemeService) GetAllSchemes() ([]models.Scheme, error) {
	var schemes []models.Scheme
	result := schemeService.Db.Preload("Benefits").Preload("Criteria").Find(&schemes)

	if result.Error != nil {
		return nil, result.Error
	}
	return schemes, nil
}

func (schemeService *SchemeService) GetSchemeByID(id string) (*models.Scheme, error) {
	var scheme models.Scheme
	result := schemeService.Db.Preload("Benefits").Preload("Criteria").First(&scheme, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &scheme, nil
}

func (schemeService *SchemeService) CreateScheme(scheme *models.Scheme) (*models.Scheme, error) {
	// Start Transaction
	tx := schemeService.Db.Begin()

	result := tx.Create(&scheme)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(scheme)
	scheme.Criteria.SchemeID = scheme.ID
	result = tx.Create(&scheme.Criteria)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, benefit := range scheme.Benefits {
		benefit.SchemeID = scheme.ID
	}

	result = tx.Create(&scheme.Benefits)
	if result.Error != nil {
		return nil, result.Error
	}

	// Commit Transaction
	tx.Commit()

	return scheme, nil
}

func (schemeService *SchemeService) UpdateScheme(scheme *models.Scheme, id string) (*models.Scheme, error) {
	result := schemeService.Db.Model(&models.Scheme{}).Updates(&scheme)
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
	result := schemeService.Db.Delete(&models.Scheme{}, "id = ?", id)

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
