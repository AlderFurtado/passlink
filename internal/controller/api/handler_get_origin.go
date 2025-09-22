package controller

import (
	"net/http"

	"github.com/AlderFurtado/passlink/internal/domain/usecase"
)

func HandlerGetOrigin(w http.ResponseWriter, r *http.Request) {
	destinyRequest := "http://localhost:8080" + r.URL.Path
	origin, err := usecase.FindOriginLinkUseCase(destinyRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, origin, http.StatusFound)
}
