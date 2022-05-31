package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var jwtExpiration = viper.GetInt("JWT.EXPIRATION")
var signingKey = []byte(viper.GetString("JWT.PRIVATE_KEY"))

type CustomClaims struct {
	UserID string `json:"user_id"`
	RoleID uint   `json:"role_id"`
	jwt.StandardClaims
}

func GenerateToken(userID string, roleID uint) (string, error) {
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
		return "", 12, err
	}

	return claims.UserID, claims.RoleID, nil
}

func handleExpiredJWT(ctx *gin.Context, err error, cookieValue string) {
	jwtError, ok := err.(*jwt.ValidationError)
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	logrus.Infoln("ddsfsafasfdgfsafsgdsf")

	if jwtError.Errors&jwt.ValidationErrorMalformed != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	} else if jwtError.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) == 0 {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	logrus.Infoln("dsf")
	token, _, err := new(jwt.Parser).ParseUnverified(cookieValue, CustomClaims{})
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	logrus.Infoln("dssdff")

	claims, ok := token.Claims.(CustomClaims)
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	logrus.Infoln("dsfasdfdsaf")

	tokenValue, err := GenerateToken(claims.UserID, claims.RoleID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", tokenValue, 0, "", "", true, true)

	ctx.Next()
}
