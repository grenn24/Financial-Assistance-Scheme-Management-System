package services

import (
	"github.com/google/uuid"
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

func (schemeService *SchemeService) GetEligibleSchemes(id string) ([]models.Scheme, error) {
	var applicant models.Applicant
	result := schemeService.Db.First(&applicant, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	var schemes []models.Scheme
	result = schemeService.Db.Joins("JOIN scheme_criteria AS criteria ON criteria.scheme_id = schemes.id").
		Where("criteria.employment_status = ? OR criteria.marital_status = ? OR criteria.has_children = ?", applicant.EmploymentStatus, applicant.MaritalStatus).
		Preload("Benefits").
		Preload("Criteria").
		Find(&schemes)

	if result.Error != nil {
		return nil, result.Error
	}
	return schemes, nil
}

func (schemeService *SchemeService) CreateScheme(scheme *models.Scheme) (*models.Scheme, error) {
	// Start Transaction
	tx := schemeService.Db.Begin()

	result := tx.Create(&scheme)
	if result.Error != nil {
		return nil, result.Error
	}
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

func (schemeService *SchemeService) UpdateScheme(scheme *models.UpdateSchemeRequest, id string) (*models.Scheme, error) {
	tx := schemeService.Db.Begin()
	result := tx.Model(&models.Scheme{}).Where("id = ?", id).Updates(&scheme)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(scheme.Benefits) > 0 {
		// Delete existing scheme benefits
		result := tx.Where("scheme_id = ?", id).Delete(&models.SchemeBenefit{})
		if result.Error != nil {
			return nil, result.Error
		}

		for i, benefit := range scheme.Benefits {
			benefit.SchemeID = uuid.MustParse(id)
			scheme.Benefits[i] = benefit
		}
		result = tx.Create(scheme.Benefits)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	if scheme.Criteria != nil {
		// Delete existing scheme criteria
		result := tx.Where("scheme_id = ?", id).Delete(&models.SchemeCriteria{})
		if result.Error != nil {
			return nil, result.Error
		}
		scheme.Criteria.SchemeID = uuid.MustParse(id)
		result = tx.Create(&scheme.Criteria)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	tx.Commit()

	return schemeService.GetSchemeByID(id)
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
