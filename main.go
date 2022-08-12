package main

import (
  "github.com/gin-gonic/gin"
  "linuxbook/models"
  "linuxbook/views"
)

func main() {
  r := gin.Default()

  models.ConnectDatabase()

  r.GET("/help", controllers.GetHelpBook)
  r.POST("/create", controllers.CreateHelpBook)

  r.Run()
}
