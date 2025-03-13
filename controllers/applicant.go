package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApplicantController struct {
			Db *gorm.DB
}

func (applicantController *ApplicantController) GetAllApplicants(context *gin.Context, db *gorm.DB) {

}

func (applicantController *ApplicantController) CreateApplicant(context *gin.Context, db *gorm.DB) {

}

func (applicantController *ApplicantController) DeleteApplicantByID(context *gin.Context, db *gorm.DB) {
	//ID := context.Param("ID")
}
