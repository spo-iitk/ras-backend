package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {
	r.GET("/api/admin/rc", getAllRCHandler)
	r.POST("/api/admin/rc", postRCHandler)
	r.PUT("/api/admin/rc", editRCHandler)

	admin := r.Group("/api/admin/rc/:rid")
	admin.Use(checkAdminAccessToRC())
	{
		admin.GET("", getRCHandler)
		admin.GET("/count", getRCCountHandler)

		admin.GET("/notice", getAllNoticesHandler)
		admin.POST("/notice", postNoticeHandler(mail_channel))
		admin.PUT("/notice", putNoticeHandler)
		admin.POST("/notice/:nid/reminder", postReminderHandler(mail_channel))
		admin.DELETE("/notice/:nid", deleteNoticeHandler)

		admin.GET("/company", getAllCompaniesHandler)
		admin.POST("/company", postNewCompanyHandler)
		admin.PUT("/company", putCompanyHandler)
		admin.GET("/company/:cid", getCompanyHandler)
		admin.DELETE("/company/:cid", deleteCompanyHandler)
		admin.GET("/company/:cid/history", getCompanyHistoryHandler)

		admin.GET("/student", getAllStudentsHandler)

		admin.GET("/student/:sid", getStudentHandler)
		admin.POST("/student/:sid/clarification", postClarificationHandler(mail_channel))
		admin.DELETE("/student/:sid", deleteStudentHandler)

		admin.POST("/student", postStudentsHandler(mail_channel))
		admin.PUT("/student", putStudentHandler)

		admin.PUT("/student/freeze", bulkFreezeStudentsHandler(mail_channel))

		admin.PUT("/student/sync", syncStudentsHandler)
		// route not in use
		// admin.PUT("/student/deregister", deregisterAllStudentsHandler)

		admin.GET("/student/questions", getStudentQuestionsHandler)
		admin.POST("/student/question", postStudentQuestionHandler)
		admin.PUT("/student/question", putStudentQuestionHandler)
		admin.DELETE("/student/question/:qid", deleteStudentQuestionHandler)

		admin.GET("/student/:sid/question/answers", getStudentAnswersHandler)
		admin.GET("/student/:sid/resume", getResumesHandler)

		admin.GET("/resume", getAllResumesHandler)
		admin.PUT("/resume/:rsid/verify", putResumeVerifyHandler(mail_channel))
	}
}

func StudentRouter(r *gin.Engine) {
	r.GET("/api/student/rc", getStudentRCHandler)
	r.GET("/api/student/rc/:rid", studentWhoamiHandler)
	student := r.Group("/api/student/rc/:rid")
	student.Use(ensureActiveStudent())
	{
		student.GET("/notice", getAllNoticesForStudentHandler)

		student.GET("/enrollment", getStudentEnrollmentHandler)
		student.POST("/enrollment/:qid/answer", postEnrollmentAnswerHandler)

		student.POST("/resume", postStudentResumeHandler)
		student.GET("/resume", getStudentResumeHandler)
	}
}

func CompanyRouter(r *gin.Engine) {
	r.GET("/api/company/whoami", companyWhoamiHandler)
	company := r.Group("/api/company/rc")
	{
		company.GET("", getCompanyRCHandler)                    // get registered rc
		company.GET("/all", getAllRCHandlerForCompany)          // get all rc
		company.POST("/:rid/enrollment", enrollCompanyHandler) // enroll a company to a rc
		company.GET("/:rid/hr", getCompanyRCHRHandler)
	}
}
