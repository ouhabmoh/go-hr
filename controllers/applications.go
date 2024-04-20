package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouhabmoh/HR/models"
	"gorm.io/gorm"
)

type ApplicationController struct {
	DB *gorm.DB
}

func NewApplicationController(DB *gorm.DB) ApplicationController {
	return ApplicationController{DB}
}

func (ac *ApplicationController) UpdateApplication(ctx *gin.Context) {
	applicationID, _ := strconv.Atoi(ctx.Param("applicationID"))
	var payload *models.UpdateApplicationRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedApplication models.Application
	result := ac.DB.First(&updatedApplication, "id = ?", applicationID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No application with that ID exists"})
		return
	}

	if payload.Status != nil {
		updatedApplication.Status = *payload.Status
	}

	if payload.Evaluation != nil {
		updatedApplication.Evaluation = new(int)
		*updatedApplication.Evaluation = *payload.Evaluation
	} else {
		updatedApplication.Evaluation = nil
	}

	ac.DB.Model(&updatedApplication).Updates(updatedApplication)

	applicationResponse := models.ApplicationResponse{
		ID:          updatedApplication.ID,
		JobID:       updatedApplication.JobID,
		CandidateID: updatedApplication.CandidateID,
		Status:      updatedApplication.Status,

		CreatedAt: updatedApplication.CreatedAt,
		UpdatedAt: updatedApplication.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": applicationResponse})
}

func (ac *ApplicationController) GetApplicationByID(ctx *gin.Context) {
	applicationID, _ := strconv.Atoi(ctx.Param("applicationID"))
	currentUser := ctx.MustGet("currentUser").(models.User)
	var application models.Application
	result := ac.DB.First(&application, "id = ?", applicationID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No application with that ID exists"})
		return
	}

	if currentUser.Role == "candidate" && currentUser.ID != application.CandidateID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are Not Allowed to perform this action"})
		return
	}

	applicationResponse := models.ApplicationResponse{
		ID:          application.ID,
		JobID:       application.JobID,
		CandidateID: application.CandidateID,
		Status:      application.Status,

		CreatedAt: application.CreatedAt,
		UpdatedAt: application.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": applicationResponse})
}

func (ac *ApplicationController) FindApplications(ctx *gin.Context) {
	jobID, _ := strconv.Atoi(ctx.Query("jobId"))
	var applications []models.Application
	result := ac.DB.Where("job_id = ?", jobID).Find(&applications)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	var applicationResponses []models.ApplicationResponse
	for _, application := range applications {
		applicationResponse := models.ApplicationResponse{
			ID:          application.ID,
			JobID:       application.JobID,
			CandidateID: application.CandidateID,
			Status:      application.Status,

			CreatedAt: application.CreatedAt,
			UpdatedAt: application.UpdatedAt,
		}
		applicationResponses = append(applicationResponses, applicationResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": applicationResponses})
}

func (ac *ApplicationController) GetApplicationsByCandidate(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var applications []models.Application
	result := ac.DB.Where("candidate_id = ?", currentUser.ID).Find(&applications)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	var applicationResponses []models.ApplicationResponse
	for _, application := range applications {
		applicationResponse := models.ApplicationResponse{
			ID:          application.ID,
			JobID:       application.JobID,
			CandidateID: application.CandidateID,
			Status:      application.Status,
			CreatedAt:   application.CreatedAt,
			UpdatedAt:   application.UpdatedAt,
		}
		applicationResponses = append(applicationResponses, applicationResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": applicationResponses})
}
