package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {
	r.GET("/api/admin/rc", getAllRC)
	r.POST("/api/admin/rc/new", postRC)

	admin := r.Group("/api/admin/rc/:rid")
	{
		admin.GET("", getRC)

		// NOtice, events, new company must have an all query param
		admin.GET("/notice", getAllNotices)
		admin.POST("/notice/new", postNotice)
		admin.POST("/notice/:nid/reminder", postReminder(mail_channel))
		admin.DELETE("/notice/:nid", deleteNotice)

		admin.GET("/company", getAllCompanies)     // all registerd compnay
		admin.POST("/company/new", postNewCompany) // add compnay to RC from master
		admin.GET("/company/:cid", getCompany)     // get company

		admin.POST("/pio-ppo", postPPOPIO) // add ppo-pio

		admin.GET("/student", getAllStudents) // get all students of rc
		admin.GET("/student/:sid", getStudent)
		admin.PUT("/student", putStudent)
		admin.POST("/student", postStudents) // bulk post/ enroll in RC

		admin.GET("/student/questions", getStudentQuestions)
		admin.POST("/student/question", postStudentQuestion)
		admin.PUT("/student/questions", putStudentQuestion)
		admin.DELETE("/student/questions/:qid", deleteStudentQuestionHandler)

		admin.GET("/student/:sid/questions/answer", getStudentAnswers)           //get answer
		admin.PUT("/student/:sid/questions", putStudentAnswer)                   // edit answer
		admin.DELETE("/student/:sid/questions/:qid", deleteStudentAnswerHandler) // delete answer

	}
}

func StudentRouter(r *gin.Engine) {
	r.GET("/api/rc/:rid/student/notice", ras.PlaceHolderController) // cache
	student := r.Group("/api/rc/:rid/student/:sid")
	{
		student.GET("", getStudent) // get registered rc

		student.GET("/enrollment", getStudentEnrollment)              // enrolment question + answers
		student.POST("/enrollment/:qid/answer", postEnrollmentAnswer) // enrolment answer

		student.GET("/resume", ras.PlaceHolderController)
		student.POST("/resume/new", ras.PlaceHolderController)
	}
}

func CompanyRouter(r *gin.Engine) {
	company := r.Group("/api/company/rc")
	{
		company.GET("", ras.PlaceHolderController) // get registered rc
	}
}
