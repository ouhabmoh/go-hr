package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ouhabmoh/HR/controllers"
	"github.com/ouhabmoh/HR/middleware"
)

type ResumeRouteController struct {
	resumeController controllers.ResumeController
}

func NewRouteResumeController(resumeController controllers.ResumeController) ResumeRouteController {
	return ResumeRouteController{resumeController}
}

func (rc *ResumeRouteController) ResumeRoute(rg *gin.RouterGroup) {
	router := rg.Group("resumes")

	router.Use(middleware.DeserializeUser())

	router.GET("/:resumeID", rc.resumeController.GetResumeByID)
}
