package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlderFurtado/passlink/internal/controller"
)

func main() {
	http.HandleFunc("/", controller.HandlerApi)
	fmt.Printf("Running in port %v", "8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error to running application")
	}
}
