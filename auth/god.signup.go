package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/constants"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
)

func godSignUpHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware.Authenticator()(ctx)
		if middleware.GetRoleID(ctx) != constants.GOD {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Only God can sign up for GOD"})
			return
		}
		
		var req User

		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.UserID == "" || req.Password == "" || req.Name == "" || req.RoleID == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing fields"})
			return
		}

		if req.RoleID == constants.STUDENT || req.RoleID == constants.COMPANY || req.RoleID == constants.GOD {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
			return
		}

		pass := req.Password
		req.Password = hashAndSalt(req.Password)

		id, err := firstOrCreateUser(ctx, &req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		mail_channel <- mail.GenerateMail(req.UserID, "Registered on RAS", "Dear "+req.Name+",\n\nYou have been registered as an Admin.\n"+"Your new credentials are: \n\nUser ID: "+req.UserID+"\nPassword: "+pass)

		logrus.Info("Admin registered: ", req.UserID, " by ", middleware.GetUserID(ctx), " with id ", id)
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	}
}
