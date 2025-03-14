package services

import (
	"github.com/grenn24/financial-assistance-scheme-management-system/models"
	"gorm.io/gorm"
)

type ApplicantService struct {
	Db *gorm.DB
}

func (applicantService *ApplicantService) GetAllApplicants() ([]models.Applicant, error) {
	var applicants []models.Applicant
	result := applicantService.Db.Find(&applicants)

	if result.Error != nil {
		return nil, result.Error
	}
	return applicants, nil
}

func (applicantService *ApplicantService) GetApplicantByID(id string) (*models.Applicant, error) {
	var applicants models.Applicant
	result := applicantService.Db.First(&applicants, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &applicants, nil
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
	result := applicantService.Db.Delete(&models.Applicant{}, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return scheme, nil
}
