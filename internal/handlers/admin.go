package handlers

import (
	"net/http"
	"github.com/nazgool97/startbase/internal/db"
	"github.com/nazgool97/startbase/internal/models"

	"github.com/gin-gonic/gin"
)

type userDTO struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// GET /admin/users  – список всех пользователей
func AdminListUsers(c *gin.Context) {
	var users []models.User
	db.DB.Select("id, email, role").Find(&users)

	list := make([]userDTO, len(users))
	for i, u := range users {
		list[i] = userDTO{ID: u.ID, Email: u.Email, Role: u.Role}
	}
	c.HTML(http.StatusOK, "users.html", gin.H{"Users": list})
}

// PUT /admin/users/:id/role  – изменить роль
func AdminSetRole(c *gin.Context) {
	type body struct {
		Role string `json:"role" binding:"required,oneof=admin user"`
	}
	var b body
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := db.DB.Model(&models.User{}).Where("id = ?", id).Update("role", b.Role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}