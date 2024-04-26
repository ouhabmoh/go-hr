package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/ouhabmoh/HR/models"
	"gorm.io/gorm"
)

type ResumeController struct {
	DB *gorm.DB
}

func NewResumeController(DB *gorm.DB) ResumeController {
	return ResumeController{DB}
}

func (rc *ResumeController) GetResumeByID(ctx *gin.Context) {
	resumeID := ctx.Param("resumeID")
	fmt.Println(resumeID)
	var resume models.Resume
	err := rc.DB.First(&resume, resumeID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.String(http.StatusNotFound, "Resume not found")
			return
		}
		ctx.String(http.StatusInternalServerError, "Error retrieving resume from database")
		return
	}

	filePath := filepath.Join("uploads", resume.Filename)
	ctx.File(filePath)

}
