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
	{
		admin.GET("", getRCHandler)
		admin.GET("/count", getRCCountHandler)

		admin.GET("/notice", getAllNoticesHandler)
		admin.POST("/notice", postNoticeHandler)
		admin.POST("/notice/:nid/reminder", postReminderHandler(mail_channel))
		admin.DELETE("/notice/:nid", deleteNoticeHandler)

		admin.GET("/company", getAllCompaniesHandler)
		admin.POST("/company", postNewCompanyHandler)
		admin.PUT("/company", putCompanyHandler)
		admin.GET("/company/:cid", getCompanyHandler)
		admin.DELETE("/company/:cid", deleteCompanyHandler)

		admin.GET("/student", getAllStudentsHandler)

		admin.GET("/student/:sid", getStudentHandler)
		admin.POST("/student/:sid/clarification", postClarificationHandler(mail_channel))
		admin.DELETE("/student/:sid", deleteStudentHandler)

		admin.POST("/student", postStudentsHandler)
		admin.PUT("/student", putStudentHandler)

		admin.PUT("/student/freeze", bulkFreezeStudentsHandler)

		admin.GET("/student/questions", getStudentQuestionsHandler)
		admin.POST("/student/question", postStudentQuestionHandler)
		admin.PUT("/student/question", putStudentQuestionHandler)
		admin.DELETE("/student/question/:qid", deleteStudentQuestionHandler)

		admin.GET("/student/:sid/question/answers", getStudentAnswersHandler)

		admin.GET("/resume", getAllResumesHandler)
		admin.GET("/resume/:rsid", getResumeHandler)
		admin.PUT("/resume/:rsid/verify", putResumeVerifyHandler)
	}
}

func StudentRouter(r *gin.Engine) {
	r.GET("/api/student/rc", getStudentRCHandler)
	student := r.Group("/api/student/rc/:rid")
	{
		student.GET("/notice", getAllNoticesForStudentHandler)
		student.GET("", studentWhoamiHandler)

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
		company.GET("", getCompanyRCHandler) // get registered rc
		company.GET("/:rid/hr", getCompanyRCHRHandler)
	}
}
