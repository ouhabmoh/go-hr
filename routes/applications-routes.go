package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ouhabmoh/HR/controllers"
	"github.com/ouhabmoh/HR/middleware"
)

type ApplicationRouteController struct {
	applicationController controllers.ApplicationController
}

func NewRouteApplicationController(applicationController controllers.ApplicationController) ApplicationRouteController {
	return ApplicationRouteController{applicationController}
}

func (ac *ApplicationRouteController) ApplicationRoute(rg *gin.RouterGroup) {
	router := rg.Group("applications")
	router.Use(middleware.DeserializeUser())

	// candidate apply for a job, we can write this endpoint in a different manner, like jobs/:jobId/applications,
	// this will have more meaning and clarify the realtion between the job and application and also the current logged in candidate will be the candidate of the applications
	router.POST("/", ac.applicationController.CreateApplication)

	router.PATCH("/:applicationID", ac.applicationController.UpdateApplication)
	router.GET("/:applicationID", ac.applicationController.GetApplicationByID)
	router.GET("/", ac.applicationController.FindApplications)
	router.GET("/me", ac.applicationController.GetApplicationsByCandidate)
}
