package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
)

func createUser(ctx *gin.Context, user *User) (uint, error) {
	tx := db.WithContext(ctx).Create(user)
	return user.ID, tx.Error
}

func firstOrCreateUser(ctx *gin.Context, user *User) (uint, error) {
	tx := db.WithContext(ctx).Where("user_id = ? ", user.UserID).FirstOrCreate(user)
	return user.ID, tx.Error
}

func getAllCompaniesAdded(ctx *gin.Context) ([]User, error) {
	var companies []User
	tx := db.WithContext(ctx).Where("role_id = 2").Order("created_at DESC").Limit(50).Find(&companies)
	return companies, tx.Error
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

func setLastLogin(ctx *gin.Context, userID string) error {
	tx := db.WithContext(ctx).Model(&User{}).Where("user_id = ?", userID).Update("last_login", time.Now().Unix())
	return tx.Error
}
