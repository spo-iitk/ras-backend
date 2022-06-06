package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/mail"
)

func CompanySignUpHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signupReq CompanySignUpRequest

		if err := ctx.ShouldBindJSON(&signupReq); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := createCompany(ctx, &signupReq)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logrus.Infof("A Company %s made signUp request with id %d", signupReq.CompanyName, id)
		mail_channel <- mail.GenerateMail(signupReq.Email, signupReq.CompanyName+"registered on RAS", "Dear "+signupReq.CompanyName+",\n\nYou have been registered on RAS. Wait, you'll hear from us soon.")

		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully Requested"})
	}
}
