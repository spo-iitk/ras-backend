package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createUser(ctx *gin.Context, user *User) (uint, error) {
	tx := db.WithContext(ctx).Create(user)
	return user.ID, tx.Error
}

func getPasswordAndRole(ctx *gin.Context, userID string) (string, Role, error) {
	var user User
	tx := db.WithContext(ctx).Where("user_id = ?", userID).First(&user)
	return user.Password, user.RoleID, tx.Error
}

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

func updatePassword(ctx *gin.Context, userID string, password string) error {
	tx := db.WithContext(ctx).Model(&User{}).Where("user_id = ?", userID).Update("password", password)
	return tx.Error
}

func cleanupOTP() {
	for {
		time.Sleep(time.Hour * 24)
		db.Unscoped().Delete(OTP{}, "expires > ?", time.Now().Add(-24*time.Hour).UnixMilli())
	}
}

func createCompany(ctx *gin.Context, company *CompanySignUpRequest) (uint, error) {
	tx := db.WithContext(ctx).Create(company)
	return company.ID, tx.Error
}
