package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {
	r.GET("/api/admin/rc", getAllRC)
	r.POST("/api/admin/rc", postRC)

	admin := r.Group("/api/admin/rc/:rid")
	{
		admin.GET("", getRC)
		admin.GET("/count", getRCCount)

		// NOtice, events, new company must have an all query param

		admin.GET("/notice", getAllNotices)
		admin.POST("/notice", postNotice)
		admin.POST("/notice/:nid/reminder", postReminder(mail_channel))
		admin.DELETE("/notice/:nid", deleteNotice)

		admin.GET("/company", getAllCompanies) // all registerd compnay
		admin.POST("/company", postNewCompany) // add compnay to RC from master
		admin.GET("/company/:cid", getCompany) // get company

		admin.GET("/student", getAllStudents) // get all students of rc
		admin.GET("/student/:sid", getStudentByID)
		admin.PUT("/student", putStudent)
		admin.POST("/student", postStudents) // bulk post/ enroll in RC

		admin.GET("/student/questions", getStudentQuestions)
		admin.POST("/student/question", postStudentQuestion)
		admin.PUT("/student/question", putStudentQuestion)
		admin.DELETE("/student/question/:qid", deleteStudentQuestionHandler)

		admin.GET("/student/:sid/question/answers", getStudentAnswers)          // get answer
		admin.PUT("/student/:sid/question", putStudentAnswer)                   // edit answer
		admin.DELETE("/student/:sid/question/:qid", deleteStudentAnswerHandler) // delete answer
	}
}

func StudentRouter(r *gin.Engine) {
	r.GET("/api/student/rc", getStudentRC)
	student := r.Group("/api/student/rc/:rid")
	{
		student.GET("/notice", getAllNotices) // cache
		student.GET("", getStudent)           // get registered rc

		student.GET("/enrollment", getStudentEnrollment)              // enrolment question + answers
		student.POST("/enrollment/:qid/answer", postEnrollmentAnswer) // enrolment answer
	}
}

func CompanyRouter(r *gin.Engine) {
	company := r.Group("/api/company/rc")
	{
		company.GET("", getCompanyRecruitmentCycle) // get registered rc
	}
}
