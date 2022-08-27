package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "linuxbook/models"
)

func GetHelpBook(c *gin.Context) {
  var books []models.HelpBook
  models.DB.Find(&books)

  c.JSON(http.StatusOK, gin.H{"data": books})
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
 var book models.HelpBook

  if err := models.DB.Where("Name= ?", c.Param("name")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": book})
}
