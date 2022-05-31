package auth

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const charset = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var otpExpiration = viper.GetInt("OTP.EXPIRATION")
var size = viper.GetInt("OTP.SIZE")

type otpRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

func generateOTP() string {
	b := make([]byte, size)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func otpHandler(ctx *gin.Context) {
	var otpReq otpRequest
	if err := ctx.ShouldBindJSON(&otpReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	otp := generateOTP()

	err := saveOTP(ctx, &OTP{
		UserID:  otpReq.UserID,
		OTP:     otp,
		Expires: uint(time.Now().Add(time.Duration(otpExpiration) * time.Minute).UnixMilli()),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "OTP sent"})
}
