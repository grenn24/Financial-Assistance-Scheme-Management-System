package services

import (
	_ "fmt"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/grenn24/financial-assistance-scheme-management-system/models"
	"gorm.io/gorm"
)

type ApplicantService struct {
	Db *gorm.DB
}

func (applicantService *ApplicantService) GetAllApplicants() ([]models.Applicant, error) {
	var applicants []models.Applicant
	result := applicantService.Db.Preload("Household").Find(&applicants)

	if result.Error != nil {
		return nil, result.Error
	}
	return applicants, nil
}

func (applicantService *ApplicantService) GetApplicantByID(id string) (*models.Applicant, error) {
	var applicant models.Applicant

	// Retrieve applicant
	result := applicantService.Db.Preload("Household").First(&applicant, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}

func (applicantService *ApplicantService) CreateApplicant(applicant *models.Applicant) (*models.Applicant, error) {
	// Start Transaction
	tx := applicantService.Db.Begin()
	result := tx.Create(&applicant)

	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	for _, householdMember := range applicant.Household {

		householdMember.HouseholdOwnerID = applicant.ID
	}
	if len(applicant.Household) > 0 {
		tx.Rollback()
		result = tx.Create(&applicant.Household)
	} else {
		applicant.Household = []models.HouseholdMember{}
	}

	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// Commit Transaction
	tx.Commit()

	return applicant, nil
}

func (applicantService *ApplicantService) UpdateApplicant(applicant *models.UpdateApplicantRequest, id string) (*models.Applicant, error) {
	_, err := applicantService.GetApplicantByID(id)
	if err != nil {
		return nil, err
	}
	tx := applicantService.Db.Begin()
	
	result := tx.Model(&models.Applicant{}).Where("id = ?", id).Updates(&applicant)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	if len(applicant.Household) > 0 {
		// Delete existing household members
		result := applicantService.Db.Where("household_owner_id = ?", id).Delete(&models.HouseholdMember{})
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
		for index, householdMember := range applicant.Household {
			householdMember.HouseholdOwnerID = uuid.MustParse(id)
			applicant.Household[index] = householdMember
		}
		result = tx.Create(applicant.Household)
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}

	}

	tx.Commit()

	return applicantService.GetApplicantByID(id)
}

func (applicantService *ApplicantService) DeleteApplicantByID(id string) (*models.Applicant, error) {

	scheme, err := applicantService.GetApplicantByID(id)
	if err != nil {
		return nil, err
	}
	result := applicantService.Db.Delete(&models.Applicant{}, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	return scheme, nil
}

func (applicantService *ApplicantService) DeleteAllApplicants() (int, error) {

	result := applicantService.Db.Where("true").Delete(&models.Applicant{})

	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
