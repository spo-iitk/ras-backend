package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {

	admin := r.Group("/api/admin/application/rc/:rid")
	{
		admin.GET("/count", getApplicationCountHandler)
		admin.GET("/student/stats", getStatsHandler)
		admin.POST("/pio-ppo", postPPOPIOHandler)

		admin.GET("/event", getAllEventsByRCHandler)
		admin.GET("/event/:eid", getEventHandler)

		admin.GET("/company/:cid/proforma", getProformaByCompanyHandler)

		admin.GET("/proforma", getAllProformasHandler)
		admin.POST("/proforma", postProformaHandler)
		admin.PUT("/proforma", putProformaHandler)
		admin.PUT("/proforma/hide", hideProformaHandler)

		proforma := admin.Group("/proforma/:pid")
		{
			proforma.GET("", getProformaHandler)
			proforma.DELETE("", deleteProformaHandler)

			proforma.GET("/question", getQuestionsByProformaHandler)
			proforma.POST("/quxestion", postQuestionHandler)
			proforma.PUT("/question/:qid", putQuestionHandler)

			proforma.POST("/email", proformaEmailHandler(mail_channel))

			proforma.GET("/event", getEventsByProformaHandler)
			proforma.POST("/event", postEventHandler)
			proforma.PUT("/event", putEventHandler)
			proforma.DELETE("/event/:eid", deleteEventHandler)

			proforma.GET("/event/:eid/student", getStudentsByEventHandler)
			proforma.POST("/event/:eid/student", postStudentsByEventHandler)
			// proforma.GET("/event/:eid/coordinator", getCoordinatorsByEventHandler)
			// proforma.POST("/event/:eid/coordinator", postCoordinatorByEventHandler)
		}
	}
}

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student/application/rc/:rid")
	{
		student.GET("/proforma", getProformasForStudentHandler)
		student.GET("/proforma/:pid", getProformaForStudentHandler)
		student.GET("/proforma/:pid/event", getEventsByProformaForStudentHandler)

		student.GET("/opening/:pid", getApplicationHandler)
		student.POST("/opening/:pid", postApplicationHandler)
		student.DELETE("/opening/:pid", deleteApplicationHandler)

		student.GET("/event", getEventsByStudentHandler)
		student.GET("/event/:eid", getEventHandler)

		student.GET("/stats", getStatsHandler)
	}
}
func CompanyRouter(r *gin.Engine) {
	company := r.Group("/api/company/application/rc/:rid")
	{
		company.GET("/proforma", getProformaForCompanyHandler)
		company.POST("/proforma", postProformaByCompanyHandler)

		company.PUT("/proforma", putProformaByCompanyHandler)
		company.GET("/proforma/:pid", getProformaHandlerForCompany)
		company.DELETE("/proforma/:pid", deleteProformaByCompanyHandler)

		company.GET("/proforma/:pid/event", getEventsByProformaForCompanyHandler)
		company.POST("/event", postEventByCompanyHandler)
		company.GET("/event/:eid", getEventHandler)

		company.PUT("/event", putEventByCompanyHandler)
		company.DELETE("/event/:eid", deleteEventByCompanyHandler)

		company.GET("/event/:eid/student", getStudentsByEventForCompanyHandler)
	}
}
