package services

import (
	"github.com/grenn24/financial-assistance-scheme-management-system/models"
	"gorm.io/gorm"
)

type ApplicationService struct {
	Db *gorm.DB
}

func (applicationSer *ApplicationService) GetAllApplications() ([]models.Application, error) {
	var applications []models.Application
	result := applicationSer.Db.Find(&applications)

	if result.Error != nil {
		return nil, result.Error
	}
	return applications, nil
}

func (applicationService *ApplicationService) GetApplicationByID(id string) (*models.Application, error) {
	var applications models.Application
	result := applicationService.Db.First(&applications,  "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &applications, nil
}

func (applicationService *ApplicationService) CreateApplication(application *models.Application) (*models.Application, error) {

	result := applicationService.Db.Create(&application)

	if result.Error != nil {
		return nil, result.Error
	}

	return application, nil
}

func (applicationService *ApplicationService) DeleteApplicationByID(id string) (*models.Application, error) {

	scheme, err := applicationService.GetApplicationByID(id)
	if err != nil {
		return nil, err
	}
	result := applicationService.Db.Delete(&models.Application{},  "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	return scheme, nil
}


func (applicationService *ApplicationService) DeleteAllApplications() (int, error) {

	result := applicationService.Db.Where("true").Delete(&models.Application{})

	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

