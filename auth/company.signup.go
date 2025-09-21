package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/mail"
	"gorm.io/gorm"
)

func companySignUpHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signupReq CompanySignUpRequest

		err := ctx.ShouldBindJSON(&signupReq)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existing User
		if err := db.Where("user_id = ?", signupReq.Email).First(&existing).Error; err == nil {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "User already registered with this email",
			})
			return
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}


		id, err := createCompany(ctx, &signupReq)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logrus.Infof("A Company %s made signUp request with id %d", signupReq.CompanyName, id)
		mail_channel <- mail.GenerateMail(signupReq.Email,
			"Registration requested on RAS",
			"Dear "+signupReq.Name+",\n\nWe got your request for registration on Recruitment Automation System, IIT Kanpur. We will get back to you soon. For any queries, please get in touch with us at spo@iitk.ac.in.")

		mail_channel <- mail.GenerateMail("spo@iitk.ac.in",
			"Registration requested on RAS",
			"Company "+signupReq.CompanyName+" has requested to be registered on RAS. The details are as follows:\n\n"+
				"Name: "+signupReq.Name+"\n"+
				"Designation: "+signupReq.Designation+"\n"+
				"Email: "+signupReq.Email+"\n"+
				"Phone: "+signupReq.Phone+"\n"+
				"Comments: "+signupReq.Comments+"\n")

		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully Requested"})
	}
}
