package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/rc/:id")
	{
		admin.GET("/company/:id/proforma", ras.PlaceHolderController)     // all proforma
		admin.GET("/company/:id/proforma/:id", ras.PlaceHolderController) // 1 proforma
		admin.PUT("/company/:id/proforma/:id", ras.PlaceHolderController) // edit proforma

		admin.GET("/company/:id/proforma/:id/question", ras.PlaceHolderController)      // all proforma
		admin.GET("/company/:id/proforma/:id/question/:id", ras.PlaceHolderController)  // all proforma
		admin.PUT("/company/:id/proforma/:id/question/:id", ras.PlaceHolderController)  // all proforma
		admin.POST("/company/:id/proforma/:id/question/new", ras.PlaceHolderController) // all proforma

		admin.POST("/company/:id/proforma/:id/email", ras.PlaceHolderController) // edit proforma
		// excel and resume pending

		admin.GET("/company/:id/proforma/:id/event", ras.PlaceHolderController)        // edit proforma
		admin.POST("/company/:id/proforma/:id/event/new", ras.PlaceHolderController)   // edit proforma
		admin.POST("/company/:id/proforma/:id/event/:id/reminder", ras.PlaceHolderController)   // edit proforma
		admin.PUT("/company/:id/proforma/:id/event/:id", ras.PlaceHolderController)    // edit proforma
		admin.DELETE("/company/:id/proforma/:id/event/:id", ras.PlaceHolderController) // edit proforma

		admin.GET("/company/:id/proforma/:id/event/:id/student", ras.PlaceHolderController)      // 1 proforma add students to event i.e. pass to next stage
		admin.POST("/company/:id/proforma/:id/event/:id/student", ras.PlaceHolderController)     // 1 proforma add students to event i.e. pass to next stage
		admin.GET("/company/:id/proforma/:id/event/:id/coordinator", ras.PlaceHolderController)  // 1 proforma add students to event i.e. pass to next stage
		admin.POST("/company/:id/proforma/:id/event/:id/coordinator", ras.PlaceHolderController) // 1 proforma add students to event i.e. pass to next stage

		admin.GET("/events/", ras.PlaceHolderController) // all events by date by schedule/not schedule
	}
}

func Router(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.POST("/create", ras.PlaceHolderController)
		student.PUT("/:id", ras.PlaceHolderController)
		student.GET("/:id", ras.PlaceHolderController)
		student.GET("/all", ras.PlaceHolderController)
		student.GET("/programs", ras.PlaceHolderController)
		student.GET("/departments", ras.PlaceHolderController)
		student.GET("/program-departments", ras.PlaceHolderController)
	}
}
