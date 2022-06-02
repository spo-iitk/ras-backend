package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/rc/:id") // abhishek will sort this
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

		admin.GET("/company/:id/proforma/:id/event", ras.PlaceHolderController)               // edit proforma
		admin.POST("/company/:id/proforma/:id/event/new", ras.PlaceHolderController)          // edit proforma
		admin.POST("/company/:id/proforma/:id/event/:id/reminder", ras.PlaceHolderController) // edit proforma
		admin.PUT("/company/:id/proforma/:id/event/:id", ras.PlaceHolderController)           // edit proforma
		admin.DELETE("/company/:id/proforma/:id/event/:id", ras.PlaceHolderController)        // edit proforma

		admin.GET("/company/:id/proforma/:id/event/:id/student", ras.PlaceHolderController)      // 1 proforma add students to event i.e. pass to next stage
		admin.POST("/company/:id/proforma/:id/event/:id/student", ras.PlaceHolderController)     // 1 proforma add students to event i.e. pass to next stage
		admin.GET("/company/:id/proforma/:id/event/:id/coordinator", ras.PlaceHolderController)  // 1 proforma add students to event i.e. pass to next stage
		admin.POST("/company/:id/proforma/:id/event/:id/coordinator", ras.PlaceHolderController) // 1 proforma add students to event i.e. pass to next stage

		admin.GET("/events", ras.PlaceHolderController) // all events by date by schedule/not schedule
	}
}

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student/rc/:id") // abhishek will sort this
	{
		student.GET("/proforma", ras.PlaceHolderController)
		student.GET("/proforma/:id", ras.PlaceHolderController)
		student.POST("/application/proforma/:id/new", ras.PlaceHolderController) // question post isme hi honge
		student.DELETE("/application/:id", ras.PlaceHolderController)
		student.GET("/application", ras.PlaceHolderController)
		student.GET("/events", ras.PlaceHolderController)    // all events by date
		student.GET("/event/:id", ras.PlaceHolderController) // all events by date
		student.GET("/stats", ras.PlaceHolderController)     // all events by date

	}
}
