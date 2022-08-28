package api

import (
	//"html/template"
	//"path"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelpBook(c *gin.Context) {
    books := fetch_helpbooks()
    c.HTML(http.StatusOK, "index.html", gin.H{
      "books": books,
  })
}

func HelpBookIndex(c *gin.Context) {
    books := fetch_helpbook_name(c.Param("name"))
    c.HTML(http.StatusOK, "book.html", gin.H{
      "books": books,
  })
}
