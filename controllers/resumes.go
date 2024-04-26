package controllers

import (
	"fmt"
	"net/http"
	"os"
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
	currentUser := ctx.MustGet("currentUser").(models.User)
	resumeID := ctx.Param("resumeID")
	fmt.Println(resumeID)
	var resume models.Resume
	err := rc.DB.First(&resume, resumeID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Resume not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error retrieving resume from database"})
		return
	}

	if currentUser.Role == "candidate" && currentUser.ID != resume.CandidateID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are Not Allowed to perform this action"})
		return
	}

	filePath := filepath.Join("uploads", resume.Filename)
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Resume not found"})
		return
	}
	ctx.File(filePath)

}
