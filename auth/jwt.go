package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

var jwtExpiration = viper.GetInt("JWT.EXPIRATION")
var signingKey = []byte(viper.GetString("JWT.PRIVATE_KEY"))

type CustomClaims struct {
	UserID string `json:"user_id"`
	RoleID Role   `json:"role_id"`
	jwt.StandardClaims
}

func generateToken(userID string, roleID Role) (string, error) {
	claims := CustomClaims{
		userID,
		roleID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(jwtExpiration) * time.Minute).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
			Issuer:    "auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}
