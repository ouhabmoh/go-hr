package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ouhabmoh/HR/controllers"
	"github.com/ouhabmoh/HR/middleware"
)

type JobRouteController struct {
	jobController controllers.JobController
}

func NewRouteJobController(jobController controllers.JobController) JobRouteController {
	return JobRouteController{jobController}
}

func (jc *JobRouteController) JobRoute(rg *gin.RouterGroup) {

	router := rg.Group("jobs")
	router.Use(middleware.DeserializeUser())
	router.POST("/", middleware.AuthorizeRoles(middleware.RoleRecruiter), jc.jobController.CreateJob)
	router.GET("/", jc.jobController.FindJobs)
	router.PUT("/:jobID", middleware.AuthorizeRoles(middleware.RoleRecruiter), jc.jobController.UpdateJob)
	router.GET("/:jobID", jc.jobController.FindJobByID)
	router.DELETE("/:jobID", middleware.AuthorizeRoles(middleware.RoleRecruiter), jc.jobController.DeleteJob)
}
