package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {

	admin := r.Group("/api/admin/application/rc/:rid")
	{
		admin.GET("/count", getApplicationCountHandler)
		admin.GET("/stats", getStatsHandler)
		admin.POST("/pio-ppo", postPPOPIOHandler)

		admin.GET("/event", getAllEventsByRCHandler)
		admin.GET("/event/:eid", getEventHandler)
		admin.DELETE("event/:eid/student", deleteAllStudentsFromEventHandler)
		admin.DELETE("event/:eid/student/:sid", deleteStudentFromEventHandler)

		admin.GET("/pvf", getAllPvfForAdminHandler)
		admin.PUT("/pvf", putPVFHandlerForAdmin)
		admin.GET("/pvf/:pid", getPvfForAdminHandler)
		admin.GET("pvf/:pid/verification/send", sendVerificationLinkForPvfHandler(mail_channel))
		admin.GET("pvf/student/:sid/verification/send", sendVerificationLinkForStudentAllPvfHandler(mail_channel))
		admin.DELETE("pvf/:pid", deletePVFHandler)
		admin.GET("pvf/student/:sid", getAllStudentPvfHandler)

		admin.GET("/company/:cid/proforma", getProformaByCompanyHandler)
		admin.GET("/company/:cid/stats", getCompanyRecruitStatsHandler)
		admin.POST("/company/count", fetchCompanyRecruitCountHandler)

		admin.GET("/proforma", getAllProformasHandler)
		admin.POST("/proforma", postProformaHandler)
		admin.PUT("/proforma", putProformaHandler)
		admin.PUT("/proforma/hide", hideProformaHandler)
		admin.GET("/view/:sid", viewApplicationsAdminHandler)

		proforma := admin.Group("/proforma/:pid")
		{
			proforma.GET("", getProformaHandler)
			proforma.DELETE("", deleteProformaHandler)

			proforma.GET("/question", getQuestionsByProformaHandler)
			proforma.POST("/question", postQuestionHandler)
			proforma.PUT("/question/:qid", putQuestionHandler)
			proforma.DELETE("/question/:qid", deleteQuestionHandler)
			// proforma.GET("/answers", getAnswersForProforma)
			proforma.POST("/email", proformaEmailHandler(mail_channel))

			proforma.GET("/event", getEventsByProformaHandler)
			proforma.POST("/event", postEventHandler)
			proforma.PUT("/event", putEventHandler)
			proforma.DELETE("/event/:eid", deleteEventHandler)

			proforma.GET("/event/:eid/student", getStudentsByEventHandler)
			proforma.POST("/event/:eid/student", postStudentsByEventHandler(mail_channel))
			// proforma.GET("/event/:eid/coordinator", getCoordinatorsByEventHandler)
			// proforma.POST("/event/:eid/coordinator", postCoordinatorByEventHandler)

			proforma.GET("/students", getStudentsByRole)
		}
	}
}

func StudentRouter(mail_channel chan mail.Mail, r *gin.Engine) {
	student := r.Group("/api/student/application/rc/:rid")
	student.Use(ensureActiveStudent())
	{
		student.GET("/proforma", getProformasForStudentHandler)
		student.GET("/proforma/:pid", getProformaForStudentHandler)
		student.GET("/proforma/:pid/event", getEventsByProformaForStudentHandler)

		student.POST("/pvf", postPvfForStudentHandler)
		student.PUT("/pvf/:pid", putPVFForStudentHandler)
		student.GET("/pvf", getAllPvfForStudentHandler)
		student.GET("/pvf/:pid", getPvfForStudentHandler)
		student.DELETE("pvf/:pid", deletePVFHandler)

		student.GET("/opening", getProformasForEligibleStudentHandler)
		student.GET("/opening/:pid", getApplicationHandler)
		student.POST("/opening/:pid", postApplicationHandler(mail_channel))
		student.DELETE("/opening/:pid", deleteApplicationHandler)

		student.GET("/event", getEventsByStudentHandler)
		student.GET("/event/:eid", getEventHandler)
		student.GET("/event/:eid/students", getStudentsByEventForStudentHandler)

		student.GET("/view", viewApplicationsHandler)
		student.GET("/stats", getStatsHandler)
	}
}
func CompanyRouter(r *gin.Engine) {
	company := r.Group("/api/company/application/rc/:rid")
	company.Use(ensureCompany())
	{
		company.GET("/proforma", getProformaForCompanyHandler)
		company.POST("/proforma", postProformaByCompanyHandler)

		company.PUT("/proforma", putProformaByCompanyHandler)
		company.GET("/proforma/:pid", getProformaHandlerForCompany)
		company.GET("/proforma/:pid/students", getStudentsForCompanyByRole)
		company.DELETE("/proforma/:pid", deleteProformaByCompanyHandler)

		company.GET("/proforma/:pid/event", getEventsByProformaForCompanyHandler)
		company.POST("/event", postEventByCompanyHandler)
		company.GET("/event/:eid", getEventHandler)

		company.PUT("/event", putEventByCompanyHandler)
		company.DELETE("/event/:eid", deleteEventByCompanyHandler)

		company.GET("/event/:eid/student", getStudentsByEventForCompanyHandler)
	}
}

func PvfVerificationRouter(r *gin.Engine) {
	pvf := r.Group("/api/verification")
	pvf.Use()
	{
		pvf.GET("/pvf", getPvfForVerificationHandler)
		// pvf.PUT("pvf/:pid/verify", verifyPvfHandler)
		pvf.PUT("/pvf", putPVFHandler)
	}
}
