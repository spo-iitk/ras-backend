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

func fetchUser(ctx *gin.Context, user *User, userID string) error {
	tx := db.WithContext(ctx).Where("user_id = ?", userID).First(&user)
	return tx.Error
}
func fetchAdmin(ctx *gin.Context, user *User, ID string) error {
	tx := db.WithContext(ctx).Where("id = ?", ID).First(&user)
	return tx.Error
}
func fetchAdmins(ctx *gin.Context, users *[]User) error {
	tx := db.WithContext(ctx).Where("role_id >= 100").Find(&users)
	return tx.Error
}

func getPasswordAndRole(ctx *gin.Context, userID string) (string, constants.Role, bool, error) {
	var user User
	tx := db.WithContext(ctx).Where("user_id = ? AND is_active = ?", userID, true).First(&user)
	return user.Password, user.RoleID, user.IsActive, tx.Error
}

func getUserRole(ctx *gin.Context, ID uint) (constants.Role, error) {
	var user User
	tx := db.WithContext(ctx).Where("ID = ?", ID).First(&user)
	return user.RoleID, tx.Error
}

func updatePassword(ctx *gin.Context, userID string, password string) (bool, error) {
	tx := db.WithContext(ctx).Model(&User{}).Where("user_id = ?", userID).Update("password", password)
	return tx.RowsAffected > 0, tx.Error
}

func updatePasswordbyGod(ctx *gin.Context, userID string, password string) (bool, error) {
	tx := db.WithContext(ctx).Model(&User{}).Where("user_id = ?", userID).Update("password", password)
	return tx.RowsAffected > 0, tx.Error
}

func updateRole(ctx *gin.Context, ID uint, roleID constants.Role) error {
	tx := db.WithContext(ctx).Model(&User{}).Where("ID = ?", ID).Update("role_id", roleID)
	return tx.Error
}

func setLastLogin(userID string) error {
	tx := db.Model(&User{}).Where("user_id = ?", userID).Update("last_login", time.Now().UnixMilli())
	return tx.Error
}

func toggleActive(ctx *gin.Context, ID uint) (bool, error) {
	var currStatus bool
	tx := db.WithContext(ctx).Model(&User{}).Where("ID = ?", ID).Select("is_active").First(&currStatus)
	if tx.Error != nil {
		return false, tx.Error
	}
	// fmt.Printf(currentStatus.First())
	tx = db.WithContext(ctx).Model(&User{}).Where("ID = ?", ID).Update("is_active", !currStatus)
	return !currStatus, tx.Error
}
