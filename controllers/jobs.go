package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ouhabmoh/HR/models"
	"gorm.io/gorm"
)

type JobController struct {
	DB *gorm.DB
}

func NewJobController(DB *gorm.DB) JobController {
	return JobController{DB}
}

func (jc *JobController) CreateJob(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateJobRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	format := "2006-01-02"
	deadline, errors := time.Parse(format, payload.Deadline)
	if errors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Deadline format is invalid"})
		return
	}

	now := time.Now()
	if now.Compare(deadline) == 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Deadline can not be older than today"})
		return
	}

	newJob := models.Job{
		Title:          payload.Title,
		Description:    payload.Description,
		Location:       payload.Location,
		EmploymentType: payload.EmploymentType,
		Deadline:       deadline,
		RecruiterID:    currentUser.ID,
	}

	result := jc.DB.Create(&newJob)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Job with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	jobResponse := models.JobResponse{
		ID:             newJob.ID,
		Title:          newJob.Title,
		Description:    newJob.Description,
		Location:       newJob.Location,
		EmploymentType: newJob.EmploymentType,
		Deadline:       newJob.Deadline,
		CreatedAt:      newJob.CreatedAt,
		UpdatedAt:      newJob.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": jobResponse})
}

func (jc *JobController) UpdateJob(ctx *gin.Context) {
	jobID, _ := strconv.Atoi(ctx.Param("jobID"))
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.UpdateJobRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedJob models.Job
	result := jc.DB.First(&updatedJob, "id = ?", jobID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No job with that ID exists"})
		return
	}

	jobToUpdate := models.Job{
		Title:          payload.Title,
		Description:    payload.Description,
		Location:       payload.Location,
		EmploymentType: payload.EmploymentType,
		Deadline:       payload.Deadline,
		RecruiterID:    currentUser.ID,
	}

	jc.DB.Model(&updatedJob).Updates(jobToUpdate)

	jobResponse := models.JobResponse{
		ID:             updatedJob.ID,
		Title:          updatedJob.Title,
		Description:    updatedJob.Description,
		Location:       updatedJob.Location,
		EmploymentType: updatedJob.EmploymentType,
		Deadline:       updatedJob.Deadline,
		CreatedAt:      updatedJob.CreatedAt,
		UpdatedAt:      updatedJob.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": jobResponse})
}

func (jc *JobController) FindJobByID(ctx *gin.Context) {
	jobID, _ := strconv.Atoi(ctx.Param("jobID"))
	var job models.Job
	result := jc.DB.First(&job, "id = ?", jobID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No job with that ID exists"})
		return
	}

	jobResponse := models.JobResponse{
		ID:             job.ID,
		Title:          job.Title,
		Description:    job.Description,
		Location:       job.Location,
		EmploymentType: job.EmploymentType,
		Deadline:       job.Deadline,
		CreatedAt:      job.CreatedAt,
		UpdatedAt:      job.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": jobResponse})
}

func (jc *JobController) FindJobs(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit
	var jobs []models.Job

	results := jc.DB.Limit(intLimit).Offset(offset).Find(&jobs)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	var jobResponses []models.JobResponse
	for _, job := range jobs {
		jobResponse := models.JobResponse{
			ID:             job.ID,
			Title:          job.Title,
			Description:    job.Description,
			Location:       job.Location,
			EmploymentType: job.EmploymentType,
			Deadline:       job.Deadline,
			CreatedAt:      job.CreatedAt,
			UpdatedAt:      job.UpdatedAt,
		}
		jobResponses = append(jobResponses, jobResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(jobResponses), "data": jobResponses})
}

func (jc *JobController) DeleteJob(ctx *gin.Context) {
	jobID, _ := strconv.Atoi(ctx.Param("jobID"))
	result := jc.DB.Delete(&models.Job{}, "id = ?", jobID)
	log.Printf("%+v", result)
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No job with that ID exists"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (jc *JobController) Apply(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	jobID, _ := strconv.Atoi(ctx.Param("jobID"))
	var request *models.CreateApplicationRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	file, err := ctx.FormFile("resume")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error getting resume file: %v", err)
		return
	}

	// Save the uploaded file to the server
	err = ctx.SaveUploadedFile(file, "./uploads/"+file.Filename)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error saving resume file: %v", err)
		return
	}

	var job models.Job
	result := jc.DB.First(&job, "id = ?", jobID)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No job with that ID exists"})
		return
	}

	newApplication := models.Application{
		JobID:       jobID,
		CandidateID: currentUser.ID,
		Status:      "pending",
		Evaluation:  nil,
	}

	result = jc.DB.Create(&newApplication)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "You Already Applied to this job posting"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	applicationResponse := models.ApplicationResponse{
		ID:          newApplication.ID,
		JobID:       newApplication.JobID,
		CandidateID: newApplication.CandidateID,
		Status:      newApplication.Status,

		CreatedAt: newApplication.CreatedAt,
		UpdatedAt: newApplication.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": applicationResponse})
}

func (jc *JobController) GetJobApplications(ctx *gin.Context) {
	jobID, _ := strconv.Atoi(ctx.Param("jobID"))
	var applications []models.Application
	result := jc.DB.Where("job_id = ?", jobID).Find(&applications)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
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
