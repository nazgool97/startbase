package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/nazgool97/startbase/internal/db"
	"github.com/nazgool97/startbase/internal/mail"
	"github.com/nazgool97/startbase/internal/models"

	"github.com/gin-gonic/gin"
)

// POST /forgot-password
func ForgotPassword(c *gin.Context) {
	email := c.PostForm("email")
	var user models.User
	if db.DB.Where("email = ?", email).First(&user).Error != nil {
		// Не показываем, существует ли email
		c.JSON(http.StatusOK, gin.H{"message": "If email exists, reset link sent"})
		return
	}

	token := randomToken(32)
	exp := time.Now().Add(1 * time.Hour)
	db.DB.Model(&user).Updates(map[string]interface{}{
		"reset_token":   token,
		"reset_expires": exp,
	})

	mail.SendReset(email, token)
	c.JSON(http.StatusOK, gin.H{"message": "If email exists, reset link sent"})
}

// POST /reset-password
func ResetPassword(c *gin.Context) {
	token := c.PostForm("token")
	newPass := c.PostForm("password")
	var user models.User
	if db.DB.Where("reset_token = ? AND reset_expires > ?", token, time.Now()).First(&user).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired token"})
		return
	}
	db.DB.Model(&user).Updates(map[string]interface{}{
		"password":      newPass, // TODO bcrypt
		"reset_token":   nil,
		"reset_expires": nil,
	})
	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
}

func randomToken(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}