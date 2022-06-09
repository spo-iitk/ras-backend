package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {

	admin := r.Group("/api/admin/application/rc/:rid")
	{
		admin.GET("/company/:cid/proforma", getProformaByCompanyHandler)
		admin.GET("/events", getAllEventsByRCHandler)
		admin.GET("/student/stats", getStats)
		admin.POST("/pio-ppo", postPPOPIOHandler)

		admin.GET("/resume", ras.PlaceHolderController)
		admin.POST("/resume", ras.PlaceHolderController) // bulk accept/reject

		proforma := admin.Group("/proforma/:pid")
		{
			proforma.GET("", getProformaHandler)
			proforma.PUT("", putProformaHandler)
			proforma.DELETE("", deleteProformaHandler)

			proforma.GET("/question", getQuestionsByProformaHandler)
			proforma.GET("/question/:qid", getQuestionHandler)
			proforma.PUT("/question/:qid", putQuestionHandler)
			proforma.POST("/question", postQuestionHandler)

			proforma.POST("/email", proformaEmailHandler(mail_channel)) // edit proforma
			// excel and resume pending

			proforma.GET("/event", getEventsByProformaHandler)                            // edit proforma
			proforma.POST("/event", postEventHandler)                                     // edit proforma
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
		student.GET("/event/:eid", getEventHandler)

		student.GET("/stats", ras.PlaceHolderController)
		student.GET("/resume", ras.PlaceHolderController)
		student.POST("/resume", ras.PlaceHolderController)
	}
}
func CompanyRouter(r *gin.Engine) {
	company := r.Group("/api/company/application/rc/:rid")
	{
		company.GET("/proforma", getProformaForCompanyHandler)
		company.POST("/proforma", postProformaByCompanyHandler)

		company.PUT("/proforma", putProformaByCompanyHandler)
		company.GET("/proforma/:pid", getProformaHandler)
		company.DELETE("/proforma/:pid", deleteProformaByCompanyHandler)

		company.GET("/proforma/:pid/event", getEventsByProformaHandler)
		company.POST("/event", postEventByCompanyHandler)
		company.GET("/event/:eid", getEventHandler)

		company.PUT("/event/:eid", putEventByCompanyHandler)
		company.DELETE("/event/:eid", deleteEventByCompanyHandler)

		company.GET("/event/:eid/students", getStudentsByEventHandler)
	}
}
