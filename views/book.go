package api

import (
	//"html/template"
	//"path"
  "net/http"
	"github.com/gin-gonic/gin"
)

func HelpBook(c *gin.Context) {
    books := fetch_helpbooks()
    //fp := path.Join("templates", "index.html")
    //tmpl, _ := template.ParseFiles(fp)
    c.HTML(http.StatusOK, "index.html", gin.H{
      "books": books,
  })
}
