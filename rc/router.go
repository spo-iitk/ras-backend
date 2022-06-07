package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/auth"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(r *gin.Engine) {
	r.GET("/api/admin/rc", getAllRC)
	r.POST("/api/admin/rc/new", postRC)

	admin := r.Group("/api/admin/rc/:rid")
	{
		admin.GET("", getRC)

		// NOtice, events, new company must have an all query param
		admin.GET("/notice", getAllNotices)
		admin.POST("/notice/new", postNotice)
		admin.POST("/notice/:nid/reminder", ) // reminder: send mail to all
		admin.DELETE("/notice/:nid", deleteNotice)        // delete notice

		admin.GET("/new-company", auth.CompanySignUpHandler(mail_channel)) // Company signup request from auth, need some middleware?

		admin.GET("/company", getAllCompanies)     // all registerd compnay
		admin.POST("/company/new", postNewCompany) // add compnay to RC from master
		admin.GET("/company/:cid", getCompany)     // get company

		admin.PUT("/pio-ppo/:sid", putPPOPIO) // add ppo-pio

		admin.GET("/student", getAllStudents) // get all students of rc
		admin.GET("/student/:sid", getStudent)
		admin.PUT("/student/:sid", putStudent)
		admin.POST("/student", postStudents) // bulk post/ enroll in RC

		admin.POST("/student/stats", ras.PlaceHolderController) // query branch wise stats, clarity needed.

		admin.GET("/student/:sid/questions", getStudentQuestions)
		admin.POST("/student/:sid/question", postStudentQuestion)
		admin.PUT("/student/:sid/questions/:qid", putStudentQuestion)
		admin.DELETE("/student/:sid/questions/:qid", deleteStudentQuestionHandler)

		admin.GET("/student/:sid/questions/:qid/answer", getStudentAnswers)      //get answer
		admin.PUT("/student/:sid/questions/:qid", putStudentAnswer)              // edit answer
		admin.DELETE("/student/:sid/questions/:qid", deleteStudentAnswerHandler) // delete answer

		admin.GET("/resume", ras.PlaceHolderController)
		admin.POST("/resume", ras.PlaceHolderController) // bulk accept/reject

	}
}

func StudentRouter(r *gin.Engine) {
	r.GET("/api/rc/:rid/student/notice", ras.PlaceHolderController) // cache
	student := r.Group("/api/rc/:rid/student/:sid")
	{
		student.GET("", getStudent) // get registered rc

		student.GET("/enrollment", getStudentEnrollment)              // enrolment question
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
