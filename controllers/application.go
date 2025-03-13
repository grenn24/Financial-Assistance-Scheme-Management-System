package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApplicationController struct {
		Db *gorm.DB
		
}

func (applicationController *ApplicationController) GetAllApplications(context *gin.Context, db *gorm.DB) {

}

func (applicationController *ApplicationController) CreateApplication(context *gin.Context, db *gorm.DB) {

}

func (applicationController *ApplicationController) DeleteApplicationByID(context *gin.Context, db *gorm.DB) {
	//ID := context.Param("ID")
}
