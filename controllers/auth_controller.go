package controllers

import (
    "net/http"
    "go-hospital-app/config"
    "go-hospital-app/models"
    "go-hospital-app/utils"
    "github.com/gin-gonic/gin"
)

type loginDto struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
    var dto loginDto
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var user models.User
    if err := config.DB.Where("username = ?", dto.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
    if err := utils.CheckPassword(user.Password, dto.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
    token, _ := utils.GenerateToken(user.ID, user.Role)
    c.JSON(http.StatusOK, gin.H{"token": token, "role": user.Role})
}