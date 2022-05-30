package auth

import (
	"time"

	"github.com/gin-gonic/gin"
)

func createUser(ctx gin.Context, user *User) (uint, error) {
	tx := db.WithContext(&ctx).Create(user)
	return user.ID, tx.Error
}

func getPassword(ctx gin.Context, userID string) (string, error) {
	var user User
	tx := db.WithContext(&ctx).Where("user_id = ?", userID).First(&user)
	return user.Password, tx.Error
}

func saveOTP(ctx gin.Context, otp *OTP) error {
	tx := db.WithContext(&ctx).Create(&otp)
	return tx.Error
}

func verifyOTP(ctx gin.Context, userID string, otp string) (bool, error) {
	var otpObj OTP
	tx := db.WithContext(&ctx).Where("user_id = ? AND otp = ? AND expires <", userID, otp, time.Now().UnixMilli()).First(&otpObj)
	return tx.RowsAffected > 0, tx.Error
}

func updatePassword(ctx gin.Context, userID string, password string) error {
	tx := db.WithContext(&ctx).Model(&User{}).Where("user_id = ?", userID).Update("password", password)
	return tx.Error
}
