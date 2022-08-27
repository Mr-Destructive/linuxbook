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

  r.GET("/help", api.GetHelpBook)
  r.POST("/create", api.CreateHelpBook)
  r.GET("/get/:name", api.IndexHelpBook)
  r.PATCH("/update/:id", api.UpdateHelpBook)
  r.DELETE("/delete/:id", api.DeleteHelpBook)
  r.GET("/helpbooks",  api.HelpBook)
  //r.GET("/helpbooks", func(c *gin.Context) {
  //  var books []models.HelpBook
  //  log.Print(books)
  //  models.DB.Find(&books)
  //  c.HTML(http.StatusOK, "index.html", gin.H{
  //    "books": books,
  //  })
  //})

  r.Run()
}
