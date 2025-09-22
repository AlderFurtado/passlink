package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AlderFurtado/passlink/internal/controller/dto"
	"github.com/AlderFurtado/passlink/internal/domain/usecase"
)

func HandlerNewLink(w http.ResponseWriter, r *http.Request) {
	var req dto.LinkRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	link, err := usecase.GenerateLinkUseCase(req.Link, req.IsPaid)
	if err != nil {
		http.Error(w, "Erro ao gerar novo link", http.StatusBadRequest)
		return
	}
	newLink := dto.LinkResponse{Link: link}
	json.NewEncoder(w).Encode(newLink)

}
