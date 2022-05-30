package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func companySignUpHandler(ctx *gin.Context) {
	var signupReq CompanySignUpRequest

	if err := ctx.ShouldBindJSON(&signupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := createCompany(ctx, &signupReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A Company %s made signUp request with id %d", signupReq.CompanyName, id)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully Requested"})
}
