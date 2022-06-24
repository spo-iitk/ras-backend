package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {

	admin := r.Group("/api/admin/application/rc/:rid")
	{
		admin.GET("/count", getApplicationCount)

		admin.GET("/company/:cid/proforma", getProformaByCompanyHandler)
		admin.GET("/event", getAllEventsByRCHandler)
		admin.GET("/event/:eid", getEventHandler)
		admin.GET("/student/stat", getStats)
		admin.POST("/pio-ppo", postPPOPIOHandler)

		admin.GET("/resume", ras.PlaceHolderController)
		admin.POST("/resume", ras.PlaceHolderController)

		admin.PUT("/proforma", putProformaHandler)
		admin.PUT("/proforma/hide", hideProformaHandler)
		admin.POST("/proforma", postProformaHandler)

		proforma := admin.Group("/proforma/:pid")
		{
			proforma.GET("", getProformaHandler)
			proforma.DELETE("", deleteProformaHandler)

			proforma.GET("/question", getQuestionsByProformaHandler)
			proforma.GET("/question/:qid", getQuestionHandler)
			proforma.PUT("/question/:qid", putQuestionHandler)
			proforma.POST("/question", postQuestionHandler)

			proforma.POST("/email", proformaEmailHandler(mail_channel))

			proforma.GET("/event", getEventsByProformaHandler)
			proforma.POST("/event", postEventHandler)
			proforma.PUT("/event", putEventHandler)
			proforma.DELETE("/event/:eid", deleteEventHandler)

			proforma.GET("/event/:eid/student", getStudentsByEventHandler)
			proforma.POST("/event/:eid/student", postStudentsByEventHandler)
			proforma.GET("/event/:eid/coordinator", getCoordinatorsByEventHandler)
			proforma.POST("/event/:eid/coordinator", postCoordinatorByEventHandler)
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

		student.GET("/event", getEventsByStudentHandler)
		student.GET("/event/:eid", getEventHandler)

		student.GET("/stat", ras.PlaceHolderController)
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

		company.PUT("/event", putEventByCompanyHandler)
		company.DELETE("/event/:eid", deleteEventByCompanyHandler)

		company.GET("/event/:eid/student", getStudentsByEventHandler)
	}
}
