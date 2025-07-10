package controllers

import (
    "net/http"
    "strconv"
    "go-hospital-app/config"
    "go-hospital-app/models"
    "github.com/gin-gonic/gin"
)

func CreatePatient(c *gin.Context) {
    var p models.Patient
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&p)
    c.JSON(http.StatusCreated, p)
}

func GetPatients(c *gin.Context) {
    var list []models.Patient
    config.DB.Find(&list)
    c.JSON(http.StatusOK, list)
}

func GetPatient(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var p models.Patient
    if err := config.DB.First(&p, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, p)
}

func UpdatePatient(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var p models.Patient
    if err := config.DB.First(&p, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&p)
    c.JSON(http.StatusOK, p)
}

func DeletePatient(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := config.DB.Unscoped().Delete(&models.Patient{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
        return
    }
    c.Status(http.StatusNoContent)
}