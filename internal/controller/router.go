package controller

import (
	"net/http"
	"strings"

	api "github.com/AlderFurtado/passlink/internal/controller/api"
	view "github.com/AlderFurtado/passlink/internal/controller/view"
)

func HandlerApi(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/get") && r.Method == "GET" {
		api.HandlerGetOrigin(w, r)
	} else if strings.Contains(r.URL.Path, "/newLink") && r.Method == "POST" {
		api.HandlerNewLink(w, r)
	} else {
		view.HandlerHome(w, r)
	}
}
