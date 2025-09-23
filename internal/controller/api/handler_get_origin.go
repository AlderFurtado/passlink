package controller

import (
	"fmt"
	"net/http"

	"github.com/AlderFurtado/passlink/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

func HandlerGetOrigin(c *gin.Context) {
	destinyRequest := "http://localhost:8080/get?term=" + c.Query("term")
	fmt.Print(destinyRequest)
	origin, err := usecase.FindOriginLinkUseCase(destinyRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.Redirect(http.StatusMovedPermanently, origin)
}
