package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func file() *os.File {
	f, err := os.OpenFile("credits.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		log.Fatal(err)
	}
	return f
}

func creditsHandler(ctx *gin.Context) {
	middleware.Authenticator()(ctx)
	user_id := middleware.GetUserID(ctx)
	role_id := middleware.GetRoleID(ctx)

	log.SetOutput(logf)
	log.Println("User:", user_id, "Role:", role_id)

	if user_id == "" {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"role_id": role_id, "user_id": user_id})
}
