package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(r *gin.Engine) {
	r.GET("/api/admin/rc", ras.PlaceHolderController)      // return all RC
	r.POST("/api/admin/rc/new", ras.PlaceHolderController) // new RC
	admin := r.Group("/api/admin/rc/:rid")
	{
		admin.GET("", ras.PlaceHolderController) // get RC just overview details

		// NOtice, events, new company must have an all query param
		admin.GET("/notice", ras.PlaceHolderController)                // all notices in details
		admin.POST("/notice/new", ras.PlaceHolderController)           // new notice
		admin.POST("/notice/:nid/reminder", ras.PlaceHolderController) // reminder: send mail to all
		admin.DELETE("/notice/:nid", ras.PlaceHolderController)        // delete notice

		admin.GET("/new-company", ras.PlaceHolderController) // Company signup request from auth

		admin.GET("/company", ras.PlaceHolderController)               // all registerd compnay
		admin.POST("/company/new", ras.PlaceHolderController)          // add compnay to RC from master
		admin.GET("/company/:cid", ras.PlaceHolderController)          // get company
		admin.GET("/company/:cid/proforma", ras.PlaceHolderController) // all proforma

		admin.POST("/ppo-pio", ras.PlaceHolderController) // add ppo-pio

		admin.GET("/student", ras.PlaceHolderController) // get all students of rc
		admin.GET("/student/:sid", ras.PlaceHolderController)
		admin.PUT("/student/:sid", ras.PlaceHolderController)
		admin.POST("/student", ras.PlaceHolderController)       // bulk post/ enroll in RC
		admin.POST("/student/stats", ras.PlaceHolderController) // query branch wise stats

		admin.GET("/student/:sid/questions", ras.PlaceHolderController)
		admin.PUT("/student/:sid/questions/:qid", ras.PlaceHolderController)

		admin.GET("/resume", ras.PlaceHolderController)
		admin.POST("/resume", ras.PlaceHolderController) // bulk accept/reject

	}
}

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student/rc")
	{
		student.GET("", ras.PlaceHolderController)                  // get registered rc
		student.POST("/:rid/enrollment", ras.PlaceHolderController) // enrolment question
		student.GET("/:rid/notice", ras.PlaceHolderController)      // cache
		student.GET("/:rid/resume", ras.PlaceHolderController)
		student.POST("/:rid/resume/new", ras.PlaceHolderController)
	}
}

func CompanyRouter(r *gin.Engine) {
	company := r.Group("/api/company/rc")
	{
		company.GET("", ras.PlaceHolderController) // get registered rc
	}
}
