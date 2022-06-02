package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/rc")
	{
		admin.GET("", ras.PlaceHolderController)      // return all RC
		admin.POST("/new", ras.PlaceHolderController) // new RC
		admin.GET("/:id", ras.PlaceHolderController)  // get RC just overview details

		// NOtice, events, new company must have an all query param
		admin.GET("/:id/notice", ras.PlaceHolderController)               // all notices in details
		admin.POST("/:id/notice/new", ras.PlaceHolderController)          // new notice
		admin.POST("/:id/notice/:id/reminder", ras.PlaceHolderController) // reminder: send mail to all
		admin.DELETE("/:id/notice/:id", ras.PlaceHolderController)        // delete notice

		admin.GET("/:id/new-company", ras.PlaceHolderController) // Company signup request from auth

		admin.GET("/:id/company", ras.PlaceHolderController)              // all registerd compnay
		admin.POST("/:id/company/new", ras.PlaceHolderController)         // add compnay to RC from master
		admin.GET("/:id/company/:id", ras.PlaceHolderController)          // get company
		admin.GET("/:id/company/:id/proforma", ras.PlaceHolderController) // all proforma

		admin.POST("/:id/ppo-pio", ras.PlaceHolderController) // add ppo-pio

		admin.GET("/:id/student", ras.PlaceHolderController) // get all students of rc
		admin.GET("/:id/student/:id", ras.PlaceHolderController)
		admin.PUT("/:id/student/:id", ras.PlaceHolderController)
		admin.POST("/:id/student", ras.PlaceHolderController)       // bulk post/ enroll in RC
		admin.POST("/:id/student/stats", ras.PlaceHolderController) // query branch wise stats

		admin.GET("/:id/student/:id/questions", ras.PlaceHolderController)
		admin.PUT("/:id/student/:id/questions/:id", ras.PlaceHolderController)

		admin.GET("/:id/resume", ras.PlaceHolderController)
		admin.POST("/:id/resume", ras.PlaceHolderController) // bulk accept/reject

	}
}

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student/rc")
	{
		student.GET("", ras.PlaceHolderController)                 // get registered rc
		student.POST("/:id/enrollment", ras.PlaceHolderController) // enrolment question
		student.GET("/:id/notice", ras.PlaceHolderController)      // cache
		student.GET("/:id/resume", ras.PlaceHolderController)
		student.POST("/:id/resume/new", ras.PlaceHolderController)
	}
}

func CompanyRouter(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.POST("/create", ras.PlaceHolderController)
		student.PUT("/:id", ras.PlaceHolderController)
		student.GET("/:id", ras.PlaceHolderController)
		student.GET("/all", ras.PlaceHolderController)
		student.GET("/programs", ras.PlaceHolderController)
		student.GET("/departments", ras.PlaceHolderController)
		student.GET("/program-departments", ras.PlaceHolderController)
	}
}
