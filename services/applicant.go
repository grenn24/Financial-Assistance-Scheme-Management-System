package services

import (
	"fmt"
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
	result := applicantService.Db.Preload("Parent").Preload("Spouse").Find(&applicants)

	if result.Error != nil {
		return nil, result.Error
	}
	return applicants, nil
}

func (applicantService *ApplicantService) GetApplicantByID(id string) (*models.Applicant, error) {
	var applicant models.Applicant
	var household []models.Applicant

	// Retrieve applicant
	result := applicantService.Db.Preload("Parent").Preload("Spouse").First(&applicant, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Retrieve household members separately
	result = applicantService.Db.
		Select("*, "+`CASE 
			WHEN parent_id IS NOT NULL AND sex = 'male' THEN 'son'
			WHEN parent_id IS NOT NULL AND sex != 'male' THEN 'daughter'
			WHEN spouse_id IS NOT NULL AND sex = 'male' THEN 'husband'
			WHEN spouse_id IS NOT NULL AND sex != 'male' THEN 'wife'
			ELSE 'Unknown'
		END AS relation`,
		).Where("parent_id = ? OR spouse_id = ?", id, id).
		Find(&household)

		fmt.Println(household)

	if result.Error != nil {
		return nil, result.Error
	}
	applicant.Household = household

	return &applicant, nil
}

func (applicantService *ApplicantService) CreateApplicant(applicant *models.Applicant) (*models.Applicant, error) {

	result := applicantService.Db.Create(&applicant)

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

