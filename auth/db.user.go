package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
)

func firstOrCreateUser(ctx *gin.Context, user *User) (uint, error) {
	tx := db.WithContext(ctx).Create(user)
	if tx.Error != nil {
		tx = db.WithContext(ctx).Where("user_id = ?", user.UserID).Updates(user)
	}
	return user.ID, tx.Error
}

func getPasswordAndRole(ctx *gin.Context, userID string) (string, constants.Role, error) {
	var user User
	tx := db.WithContext(ctx).Where("user_id = ?", userID).First(&user)
	return user.Password, user.RoleID, tx.Error
}

func updatePassword(ctx *gin.Context, userID string, password string) (bool, error) {
	tx := db.WithContext(ctx).Model(&User{}).Where("user_id = ? AND role_id = ?", userID, constants.STUDENT).Update("password", password)
	return tx.RowsAffected > 0, tx.Error
}

func setLastLogin(userID string) error {
	tx := db.Model(&User{}).Where("user_id = ?", userID).Update("last_login", time.Now().UnixMilli())
	return tx.Error
}
