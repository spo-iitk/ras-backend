package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(r *gin.Engine) {

	admin := r.Group("/api/admin/application/rc/:rid")
	{
		admin.GET("/company/:cid/proforma", ras.PlaceHolderController) // all proforma
		admin.GET("/events", ras.PlaceHolderController)                // all events by date by schedule/not schedule
		admin.POST("/student/stats", ras.PlaceHolderController)        // query branch wise stats
		admin.POST("/pio-ppo", ras.PlaceHolderController)              // add ppo-pio, to events

		performa := admin.Group("/proforma/:pid")
		{
			performa.GET("", ras.PlaceHolderController) // 1 proforma
			performa.PUT("", ras.PlaceHolderController) // edit proforma

			performa.GET("/question", ras.PlaceHolderController)      // all proforma
			performa.GET("/question/:qid", ras.PlaceHolderController) // all proforma
			performa.PUT("/question/:qid", ras.PlaceHolderController) // all proforma
			performa.POST("/question/new", ras.PlaceHolderController) // all proforma

			performa.POST("/email", ras.PlaceHolderController) // edit proforma
			// excel and resume pending

			performa.GET("/event", ras.PlaceHolderController)                // edit proforma
			performa.POST("/event/new", ras.PlaceHolderController)           // edit proforma
			performa.POST("/event/:eid/reminder", ras.PlaceHolderController) // edit proforma
			performa.PUT("/event/:eid", ras.PlaceHolderController)           // edit proforma
			performa.DELETE("/event/:eid", ras.PlaceHolderController)        // edit proforma

			performa.GET("/event/:eid/student", ras.PlaceHolderController)      // 1 proforma add students to event i.e. pass to next stage
			performa.POST("/event/:eid/student", ras.PlaceHolderController)     // 1 proforma add students to event i.e. pass to next stage
			performa.GET("/event/:eid/coordinator", ras.PlaceHolderController)  // 1 proforma add students to event i.e. pass to next stage
			performa.POST("/event/:eid/coordinator", ras.PlaceHolderController) // 1 proforma add students to event i.e. pass to next stage
		}
	}
}

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/application/student/rc/:rid") // abhishek will sort this
	{
		student.GET("/proforma", ras.PlaceHolderController)
		student.GET("/proforma/:pid", ras.PlaceHolderController)
		student.POST("/proforma/:pid/new", ras.PlaceHolderController) // question post isme hi honge
		student.DELETE("/:aid", ras.PlaceHolderController)
		student.GET("", ras.PlaceHolderController)
		student.GET("/events", ras.PlaceHolderController)     // all events by date
		student.GET("/event/:eid", ras.PlaceHolderController) // all events by date
		student.GET("/stats", ras.PlaceHolderController)      // all events by date

	}
}
func CompanyRouter(r *gin.Engine) {
	company := r.Group("/api/application/company/rc/:rid/performa")
	{
		company.GET("", ras.PlaceHolderController)         // enrolment question
		company.GET("/:pid", ras.PlaceHolderController)    // enrolment question
		company.PUT("/:pid", ras.PlaceHolderController)    // if ownwr
		company.DELETE("/:pid", ras.PlaceHolderController) // if ownwr

		company.GET("/:pid/event", ras.PlaceHolderController)         // all envents
		company.GET("/:pid/event/:eid", ras.PlaceHolderController)    // 1 envents
		company.PUT("/:pid/event/:eid", ras.PlaceHolderController)    // 1 envents
		company.DELETE("/:pid/event/:eid", ras.PlaceHolderController) // 1 envents

		company.GET("/:pid/event/:eid/students", ras.PlaceHolderController) // students of event
	}
}
