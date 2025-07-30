package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/nazgool97/startbase/internal/models"
	"github.com/nazgool97/startbase/internal/db"
	"github.com/nazgool97/startbase/internal/middleware"
)

var secret = []byte("your-secret-key")

func Signup(c *gin.Context) {
    email := c.PostForm("email")
    password := c.PostForm("password")

    if email == "" || password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email and password required"})
        return
    }

    user := models.User{Email: email, Password: password}
    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "user created"})
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := db.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString(secret)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func AdminLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
		return
	}

	if user.Password != password {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := middleware.GenerateJWT(user.Email)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Error generating token"})
		return
	}

	// Сохраняем токен в cookie
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/admin/dashboard")
}

func AdminDashboard(c *gin.Context) {
	email := c.MustGet("email").(string)
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"Email": email})
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/admin/login")
}