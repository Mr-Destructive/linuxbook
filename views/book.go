package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "linuxbook/models"
)

func GetHelpBook(c *gin.Context) {
  var helpbooks []models.HelpBook
  models.DB.Find(&helpbooks)

  c.JSON(http.StatusOK, gin.H{"data": helpbooks})
}

func CreateHelpBook(c *gin.Context) {
  var input models.CreateHelpBookInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  helpbook := models.HelpBook{Name: input.Name, Description: input.Description, Parameters: input.Parameters}
  models.DB.Create(&helpbook)

  c.JSON(http.StatusOK, gin.H{"data": helpbook})
}

func IndexHelpBook(c *gin.Context) {
 var helpbook models.HelpBook

  if err := models.DB.Where("Name= ?", c.Param("name")).First(&helpbook).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": helpbook})
}

func UpdateHelpBook(c *gin.Context) {
  var helpbook models.HelpBook
  if err := models.DB.Where("id = ?", c.Param("id")).First(&helpbook).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
  var input models.CreateHelpBookInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&helpbook).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": helpbook})
}

func DeleteHelpBook(c *gin.Context) {
  var helpbook models.HelpBook
  if err := models.DB.Where("id = ?", c.Param("id")).First(&helpbook).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&helpbook)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
