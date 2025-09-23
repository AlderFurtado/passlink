package controller

import (
	"net/http"

	"github.com/AlderFurtado/passlink/internal/controller/dto"
	"github.com/AlderFurtado/passlink/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

func HandlerNewLink(c *gin.Context) {
	var req dto.LinkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	link, err := usecase.GenerateLinkUseCase(req.Link, req.IsPaid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"link": link,
	})

}
