package services

import (
	_ "fmt"

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
		return nil, result.Error
	}

	for _, householdMember := range applicant.Household {
		householdMember.HouseholdOwnerID = applicant.ID
	}

	result = tx.Create(&applicant.Household)

	if result.Error != nil {
		return nil, result.Error
	}

	// Commit Transaction
	tx.Commit()

	return applicant, nil
}

func (applicantService *ApplicantService) UpdateApplicant(applicant *models.Applicant, id string) (*models.Applicant, error) {
	result := applicantService.Db.Model(&models.Applicant{}).Updates(&applicant)
	if result.Error != nil {
		return nil, result.Error
	}

	return applicant, nil
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
