package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {

	admin := r.Group("/api/admin/application/rc/:rid")
	{
		admin.GET("/company/:cid/proforma", getPerformaByCompanyID) // all proforma
		admin.GET("/events", getAllEventsByRCHandler)             // all events by date by schedule/not schedule
		admin.POST("/student/stats", ras.PlaceHolderController)     // query branch wise stats
		admin.POST("/pio-ppo", postPPOPIOHandler)           // add ppo-pio, to events

		admin.GET("/resume", ras.PlaceHolderController)
		admin.POST("/resume", ras.PlaceHolderController) // bulk accept/reject

		performa := admin.Group("/proforma/:pid")
		{
			performa.GET("", getPerformaByPID) // 1 proforma
			performa.PUT("", putPerforma)      // edit proforma

			performa.GET("/question", getQuestionsByPID)      // all proforma
			performa.GET("/question/:qid", getQuestionsByQID) // all proforma
			performa.PUT("/question/:qid", putQuestion)       // all proforma
			performa.POST("/question/new", postQuestion)      // all proforma

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
	student := r.Group("/api/student/application/rc/:rid") // abhishek will sort this
	{
		student.GET("/proforma", getPerformaByRID)
		student.GET("/proforma/:pid", getPerformaByPID)

		student.POST("/proforma/new", ras.PlaceHolderController) // question post isme hi honge
		student.DELETE("/:aid", ras.PlaceHolderController)
		student.GET("", ras.PlaceHolderController)
		student.GET("/events", ras.PlaceHolderController)     // all events by date
		student.GET("/event/:eid", ras.PlaceHolderController) // all events by date
		student.GET("/stats", ras.PlaceHolderController)      // all events by date
		student.GET("/resume", ras.PlaceHolderController)
		student.POST("/resume/new", ras.PlaceHolderController)

	}
}
func CompanyRouter(r *gin.Engine) {
	company := r.Group("/api/application/company/:cid/rc/:rid/performa")
	{
		company.GET("", getPerformaByCompanyID)            // all perfroma by company id
		company.POST("/new", postPerformaByCompanyID)      // add new proforma
		company.GET("/:pid", getPerformaByPID)             // 1 performa by id
		company.PUT("", putPerformaByCompanyID)            // if ownwr
		company.DELETE("/:pid", deletePerformaByCompanyID) // if ownwr

		company.GET("/:pid/event", ras.PlaceHolderController)         // all envents
		company.GET("/:pid/event/:eid", ras.PlaceHolderController)    // 1 envents
		company.PUT("/:pid/event/:eid", ras.PlaceHolderController)    // 1 envents
		company.DELETE("/:pid/event/:eid", ras.PlaceHolderController) // 1 envents

		company.GET("/:pid/event/:eid/students", ras.PlaceHolderController) // students of event
	}
}
