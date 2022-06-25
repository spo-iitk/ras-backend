package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {
	r.GET("/api/admin/rc", getAllRC)
	r.POST("/api/admin/rc", postRC)
	r.PUT("/api/admin/rc", editRCHandler)

	admin := r.Group("/api/admin/rc/:rid")
	{
		admin.GET("", getRC)
		admin.GET("/count", getRCCount)

		admin.GET("/notice", getAllNotices)
		admin.POST("/notice", postNotice)
		admin.POST("/notice/:nid/reminder", postReminder(mail_channel))
		admin.DELETE("/notice/:nid", deleteNotice)

		admin.GET("/company", getAllCompanies) // all registerd compnay
		admin.POST("/company", postNewCompany) // add compnay to RC from master
		admin.PUT("/company", putCompany)      // add compnay to RC from master
		admin.GET("/company/:cid", getCompany) // get company
		admin.DELETE("/company/:cid", deleteCompanybyID)

		admin.GET("/student", getAllStudents)

		admin.GET("/student/:sid", getStudentByID)
		admin.POST("/student/:sid/clarification", postClarificationHandler(mail_channel))
		admin.DELETE("/student/:sid", deleteStudentByID)

		admin.POST("/student", postStudents)
		admin.PUT("/student", putStudent)

		admin.PUT("/student/freeze", bulkFreezeStudents)

		admin.GET("/student/questions", getStudentQuestions)
		admin.POST("/student/question", postStudentQuestion)
		admin.PUT("/student/question", putStudentQuestion)
		admin.DELETE("/student/question/:qid", deleteStudentQuestionHandler)

		admin.GET("/student/:sid/question/answers", getStudentAnswers)
		// admin.PUT("/student/:sid/question", putStudentAnswer)
		// admin.DELETE("/student/:sid/question/:qid", deleteStudentAnswerHandler)

		admin.GET("/resume", getAllResumes)
		admin.GET("/resume/:rsid", getResume)
		admin.PUT("/resume/:rsid/verify", putResumeVerify)
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

		student.POST("/resume", postStudentResume) // add resume
		student.GET("/resume", getStudentResume)   // get all resume
	}
}

func CompanyRouter(r *gin.Engine) {
	r.GET("/api/company/whoami", companyWhoamiHandler)
	company := r.Group("/api/company/rc")
	{
		company.GET("", getCompanyRecruitmentCycle) // get registered rc
		company.GET("/:rid/hr", getCompanyRCHRHandler)
	}
}
