package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("secret_key")

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		email, ok := claims["email"].(string)
		if !ok {
			return "", errors.New("invalid claims")
		}
		return email, nil
	}

	return "", errors.New("invalid token")
}

func AuthRequired(c *gin.Context) {
	tokenString := ""

	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		cookie, err := c.Cookie("token")
		if err == nil {
			tokenString = cookie
		}
	}

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	email, err := ValidateJWT(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.Set("email", email)
	c.Next()
}
