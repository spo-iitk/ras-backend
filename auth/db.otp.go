package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func saveOTP(ctx *gin.Context, otp *OTP) error {
	tx := db.WithContext(ctx).Create(&otp)
	return tx.Error
}

func verifyOTP(ctx *gin.Context, userID string, otp string) (bool, error) {
	var otpObj OTP
	tx := db.WithContext(ctx).Where("user_id = ? AND otp = ? AND expires > ?", userID, otp, time.Now().UnixMilli()).First(&otpObj)
	switch tx.Error {
	case nil:
		db.WithContext(ctx).Delete(&otpObj)
		return true, nil
	case gorm.ErrRecordNotFound:
		return false, nil
	default:
		return false, tx.Error
	}
}

func cleanupOTP() {
	for {
		db.Unscoped().Delete(OTP{}, "expires < ?", time.Now().Add(-24*time.Hour).UnixMilli())
		time.Sleep(time.Hour * 24)
	}
}
