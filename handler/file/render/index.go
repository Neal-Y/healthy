package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FrontendPage(r *gin.Engine) {
	r.GET("/home", ShowIndex)
}

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
