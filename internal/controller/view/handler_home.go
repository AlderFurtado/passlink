package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}
