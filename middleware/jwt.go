package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	_ "github.com/spo-iitk/ras-backend/config"
)

var (
	jwtExpirationLong  int
	jwtExpirationShort int
	signingKey         []byte
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	RoleID uint   `json:"role_id"`
	jwt.StandardClaims
}

func init() {
	jwtExpirationLong = viper.GetInt("JWT.EXPIRATION.LONG")
	jwtExpirationShort = viper.GetInt("JWT.EXPIRATION.SHORT")
	signingKey = []byte(viper.GetString("JWT.PRIVATE_KEY"))
}

func GenerateToken(userID string, roleID uint, long bool) (string, error) {
	var jwtExpiration int
	if long {
		jwtExpiration = jwtExpirationLong
	} else {
		jwtExpiration = jwtExpirationShort
	}

	claims := CustomClaims{
		userID,
		roleID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(jwtExpiration) * time.Minute).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
			Issuer:    "ras",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func validateToken(encodedToken string) (string, uint, error) {

	claims := &CustomClaims{}
	_, err := jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", 20, err
	}

	return claims.UserID, claims.RoleID, nil
}
