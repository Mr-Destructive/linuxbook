package main

import (
	"linuxbook/models"
	"linuxbook/views"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.LoadHTMLGlob("templates/**")
  models.ConnectDatabase()

  // API
  r.GET("/help", api.GetHelpBook)
  r.POST("/create", api.CreateHelpBook)
  r.GET("/get/:name", api.IndexHelpBook)
  r.PATCH("/update/:id", api.UpdateHelpBook)
  r.DELETE("/delete/:id", api.DeleteHelpBook)

  // Templates
  r.GET("/helpbooks",  api.HelpBook)
  r.GET("/helpbooks/:name",  api.HelpBookIndex)

  r.Run()
}
