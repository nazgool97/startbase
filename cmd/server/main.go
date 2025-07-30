package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"

	"github.com/nazgool97/startbase/internal/db"
	"github.com/nazgool97/startbase/internal/handlers"
	"github.com/nazgool97/startbase/internal/middleware"
	"github.com/nazgool97/startbase/internal/otel"
)

func main() {
	// 🔍 Инициализация трассировки
	shutdown := otel.InitTracer("startbase")
	defer shutdown(context.Background())

	db.Init()

	r := gin.Default()

	// 📦 Включаем OpenTelemetry middleware
	r.Use(otel.Middleware("startbase"))

	// --- HTML templates
	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("templates/*.html")

	// --- Public
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	r.GET("/admin/login", handlers.ShowLogin)
	r.POST("/admin/login", handlers.AdminLogin)
	r.GET("/admin/logout", handlers.Logout)
	r.POST("/forgot-password", handlers.ForgotPassword)

	r.GET("/reset", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
			<form method="post" action="/reset-password">
			  <input type="hidden" name="token" value="` + c.Query("token") + `">
			  New password: <input name="password" type="password" required><br>
			  <button type="submit">Reset</button>
			</form>`))
	})

	r.POST("/reset-password", handlers.ResetPassword)

	// --- Protected admin group
	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired)
	{
		admin.GET("/dashboard", handlers.AdminDashboard)
		admin.GET("/users", handlers.AdminListUsers)
		admin.PUT("/users/:id/role", handlers.AdminSetRole)
	}

	// --- Auth API
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	api := r.Group("/")
	api.Use(middleware.AuthRequired)
	api.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You are authorized!"})
	})

	// 🚀 Запускаем сервер
	r.Run(":8080")
}