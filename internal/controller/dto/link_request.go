package dto

type LinkRequest struct {
	Link   string `json:"link"` // exportado para JSON funcionar
	IsPaid bool   `json:"bool"` // exportado para JSON funcionar
}
