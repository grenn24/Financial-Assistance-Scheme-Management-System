package services

import (
	"fmt"
	"strconv"

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

	result := schemeService.Db.Preload("Household").First(&applicant, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	schoolLevelFilter := ""

	// add SchoolLevel filter conditions
	var hasPrimary bool
	var hasSecondary bool
	var hasTertiary bool
	var hasChildrenInSchool bool
	for i := 0; i < len(applicant.Household) && (!hasPrimary || !hasSecondary || !hasTertiary); i++ {
		if applicant.Household[i].Relation == "son" || applicant.Household[i].Relation == "daughter" {
			if applicant.Household[i].SchoolLevel != nil {
				switch *applicant.Household[i].SchoolLevel {
				case "primary":
					if !hasPrimary {
						schoolLevelFilter += fmt.Sprintf("OR (criteria.has_children->>'school_level' = '%s') ", "primary")
						hasPrimary = true
					}
				case "secondary":
					if !hasSecondary {
						schoolLevelFilter += fmt.Sprintf("OR (criteria.has_children->>'school_level' = '%s') ", "secondary")
						hasSecondary = true
					}
				case "tertiary":
					if !hasTertiary {
						schoolLevelFilter += fmt.Sprintf("OR (criteria.has_children->>'school_level' = '%s') ", "tertiary")
						hasTertiary = true
					}
				}
			}
			// include schemes that allows children from all school levels
			if !hasChildrenInSchool {
				schoolLevelFilter += fmt.Sprintf("OR (criteria.has_children->>'school_level' = '%s') ", "all")
				hasChildrenInSchool = true
			}
		}
	}

	childrenCount := 0
	for _, member := range applicant.Household {
		if member.Relation == "son" || member.Relation == "daughter" {
			childrenCount++
		}
	}

	var schemes []models.Scheme
	result = schemeService.Db.Joins("JOIN scheme_criteria AS criteria ON criteria.scheme_id = schemes.id").
		Where(
			fmt.Sprintf("(criteria.employment_status IS NULL OR criteria.employment_status = ?) AND (criteria.marital_status IS NULL OR criteria.marital_status = ?) AND (criteria.has_children->>'school_level' IS NULL %s) AND (criteria.has_children->>'number' IS NULL OR criteria.has_children->>'number' = ?)", schoolLevelFilter),
			applicant.EmploymentStatus, applicant.MaritalStatus, strconv.Itoa(childrenCount),
		).
		Preload("Benefits").
		Preload("Criteria").
		Find(&schemes)

	if result.Error != nil {
		return nil, result.Error
	}
	return schemes, nil
}

func (schemeService *SchemeService) SearchSchemes(query string) ([]models.Scheme, error) {
	var schemes []models.Scheme
	result := schemeService.Db.Preload("Benefits").Preload("Criteria").Where("name ILIKE ?", "%"+query+"%").Find(&schemes)

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
		tx.Rollback()
		return nil, result.Error
	}

	// Commit Transaction
	tx.Commit()

	return scheme, nil
}

func (schemeService *SchemeService) UpdateScheme(scheme *models.UpdateSchemeRequest, id string) (*models.Scheme, error) {
	_, err := schemeService.GetSchemeByID(id)
	if err != nil {
		return nil, err
	}
	tx := schemeService.Db.Begin()
	result := tx.Model(&models.Scheme{}).Where("id = ?", id).Updates(&scheme)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	if len(scheme.Benefits) > 0 {
		// Delete existing scheme benefits
		result := tx.Where("scheme_id = ?", id).Delete(&models.SchemeBenefit{})
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}

		for i, benefit := range scheme.Benefits {
			benefit.SchemeID = uuid.MustParse(id)
			scheme.Benefits[i] = benefit
		}
		result = tx.Create(scheme.Benefits)
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	if scheme.Criteria != nil {
		// Delete existing scheme criteria
		result := tx.Where("scheme_id = ?", id).Delete(&models.SchemeCriteria{})
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
		scheme.Criteria.SchemeID = uuid.MustParse(id)
		result = tx.Create(&scheme.Criteria)
		if result.Error != nil {
			tx.Rollback()
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
