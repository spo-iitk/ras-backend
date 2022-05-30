package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type signUp struct {
	UserID     string `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func signUpHandler(ctx *gin.Context) {
	var signupReq signUp
	if err := ctx.ShouldBindJSON(&signupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd := HashAndSalt(signupReq.Password)

	createUser(ctx, &User{
		UserID: signupReq.UserID,
		Name: signupReq.Name,
		Password: hashedPwd,
	})

	ctx.JSON(http.StatusOK, gin.H{"status": "succesfully signed up"})
}
