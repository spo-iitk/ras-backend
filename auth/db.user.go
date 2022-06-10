package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
)

func createUser(ctx *gin.Context, user *User) (uint, error) {
	tx := db.WithContext(ctx).Create(user)
	return user.ID, tx.Error
}

func getPasswordAndRole(ctx *gin.Context, userID string) (string, constants.Role, error) {
	var user User
	tx := db.WithContext(ctx).Where("user_id = ?", userID).First(&user)
	return user.Password, user.RoleID, tx.Error
}

func updatePassword(ctx *gin.Context, userID string, password string) error {
	tx := db.WithContext(ctx).Model(&User{}).Where("user_id = ?", userID).Update("password", password)
	return tx.Error
}
