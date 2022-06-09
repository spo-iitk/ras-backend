package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {

	admin := r.Group("/api/admin/application/rc/:rid")
	{
		admin.GET("/company/:cid/proforma", getProformaByCompanyID) // all proforma
		admin.GET("/events", getAllEventsByRCHandler)               // all events by date by schedule/not schedule
		admin.GET("/student/stats", getStats)                       // query branch wise stats
		admin.POST("/pio-ppo", postPPOPIOHandler)                   // add ppo-pio, to events

		admin.GET("/resume", ras.PlaceHolderController)
		admin.POST("/resume", ras.PlaceHolderController) // bulk accept/reject

		proforma := admin.Group("/proforma/:pid")
		{
			proforma.GET("", getProformaHandler) // 1 proforma
			proforma.PUT("", putProforma)        // edit proforma

			proforma.GET("/question", getQuestionsByPID)      // all proforma
			proforma.GET("/question/:qid", getQuestionsByQID) // all proforma
			proforma.PUT("/question/:qid", putQuestion)       // all proforma
			proforma.POST("/question/new", postQuestion)      // all proforma

			proforma.POST("/email", proformaEmailHandler(mail_channel)) // edit proforma
			// excel and resume pending

			proforma.GET("/event", getEventsByPIDHandler)                                 // edit proforma
			proforma.POST("/event/new", postEventHandler)                                 // edit proforma
			proforma.POST("/event/:eid/reminder", postEventReminderHandler(mail_channel)) // edit proforma
			proforma.PUT("/event", putEventHandler)                                       // edit proforma
			proforma.DELETE("/event/:eid", deleteEventHandler)                            // edit proforma

			proforma.GET("/event/:eid/student", getStudentsByEventHandler)          // 1 proforma add students to event i.e. pass to next stage
			proforma.POST("/event/:eid/student", postStudentsByEventHandler)        // 1 proforma add students to event i.e. pass to next stage
			proforma.GET("/event/:eid/coordinator", getCoordinatorsByEventHandler)  // 1 proforma add students to event i.e. pass to next stage
			proforma.POST("/event/:eid/coordinator", postCoordinatorByEventHandler) // 1 proforma add students to event i.e. pass to next stage

		}
	}
}

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student/application/rc/:rid")
	{
		student.GET("/proforma", getProformaByRIDHandler)
		student.GET("/proforma/:pid", getProformaHandler)

		student.POST("/proforma/:pid", postApplicationHandler)
		student.DELETE("/proforma/:pid", deleteApplicationHandler)

		student.GET("/events", getEventsByStudentHandler)
		student.GET("/event/:eid", getEventsByIDHandler)

		student.GET("/stats", ras.PlaceHolderController)
		student.GET("/resume", ras.PlaceHolderController)
		student.POST("/resume/new", ras.PlaceHolderController)

	}
}
func CompanyRouter(r *gin.Engine) {
	company := r.Group("/api/company/application/rc/:rid/proforma")
	{
		company.GET("", getProformaByCompanyID)            // all perfroma by company id
		company.POST("/new", postProformaByCompanyID)      // add new proforma
		company.GET("/:pid", getProformaHandler)           // 1 proforma by id
		company.PUT("", putProformaByCompanyID)            // if ownwr
		company.DELETE("/:pid", deleteProformaByCompanyID) // if ownwr

		company.GET("/:pid/event", ras.PlaceHolderController)         // all envents
		company.GET("/:pid/event/:eid", ras.PlaceHolderController)    // 1 envents
		company.PUT("/:pid/event/:eid", ras.PlaceHolderController)    // 1 envents
		company.DELETE("/:pid/event/:eid", ras.PlaceHolderController) // 1 envents

		company.GET("/:pid/event/:eid/students", ras.PlaceHolderController) // students of event
	}
}
