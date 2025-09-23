package main

import (
	api "github.com/AlderFurtado/passlink/internal/controller/api"
	view "github.com/AlderFurtado/passlink/internal/controller/view"
	"github.com/gin-gonic/gin"
)

func main() {
	// http.HandleFunc("/", controller.HandlerApi)
	// fmt.Printf("Running in port %v", "8080")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal("Error to running application")
	// }

	router := gin.Default()

	router.LoadHTMLGlob("../../../internal/controller/view/templates/*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("newLink", api.HandlerNewLink)
	router.GET("/get", api.HandlerGetOrigin)
	router.GET("/", view.HandlerHome)

	router.Run(":8080") // listens on 0.0.0.0:8080 by default
}
