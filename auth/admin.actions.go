package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
)

func hrSignUpHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware.Authenticator()(ctx)
		if middleware.GetUserID(ctx) == "" {
			return
		}

		if middleware.GetRoleID(ctx) != constants.GOD && middleware.GetRoleID(ctx) != constants.OPC {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Only God and OPC can sign up for HR"})
			return
		}

		var request User
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if request.Name == "" || request.Password == "" || request.UserID == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		pass := request.Password
		request.Password = hashAndSalt(request.Password)
		request.RoleID = constants.COMPANY

		id, err := firstOrCreateUser(ctx, &request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mail_channel <- mail.GenerateMail(request.UserID, "New Credentials generated", "Your new credentials are: \n\nUser ID: "+request.UserID+"\nPassword: "+pass+"\n\nYou can reset the password from <a href= \"https://placement.iitk.ac.in/reset-password\">here</a>")
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	}
}
