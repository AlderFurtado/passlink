package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AlderFurtado/passlink/internal/domain/usecase"
)

type LinkRequest struct {
	Link   string `json:"link"` // exportado para JSON funcionar
	IsPaid bool   `json:"bool"` // exportado para JSON funcionar
}

type LinkResponse struct {
	Link string `json:"link"`
}

func handlerGetOrigin(w http.ResponseWriter, r *http.Request) {
	destinyRequest := "http://localhost:8080" + r.URL.Path
	origin, err := usecase.FindOriginLinkUseCase(destinyRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, origin, http.StatusFound)
}

func handlerNewLink(w http.ResponseWriter, r *http.Request) {
	var req LinkRequest
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
	newLink := LinkResponse{Link: link}
	json.NewEncoder(w).Encode(newLink)

}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html lang="pt-BR">
	<head>
		<meta charset="UTF-8">
		<title>Gerador de Link</title>
		<style>
			body { font-family: Arial, sans-serif; background: #f0f4f8; display:flex; justify-content:center; align-items:center; height:100vh; margin:0; }
			.container { background:#fff; padding:30px 40px; border-radius:12px; box-shadow:0 8px 20px rgba(0,0,0,0.1); text-align:center; width:400px; }
			input[type="text"] { width:80%; padding:10px; margin-bottom:20px; border-radius:6px; border:1px solid #ccc; font-size:16px; }
			button { padding:10px 20px; background:#4CAF50; color:#fff; border:none; border-radius:6px; cursor:pointer; font-size:16px; }
			button:hover { background:#45a049; }
			.result { margin-top:20px; font-size:18px; color:#333; word-break: break-all; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Gerador de Link</h2>
			<input type="text" id="inputLink" placeholder="Digite o link aqui" />
			<br>
			<button onclick="gerarLink()">Gerar</button>
			<div class="result" id="result"></div>
		</div>

		<script>
			async function gerarLink() {
				const input = document.getElementById("inputLink").value;
				if (!input) { alert("Digite um link válido!"); return; }

				try {
					const response = await fetch("/newLink", {
						method: "POST",
						headers: { "Content-Type": "application/json" },
						body: JSON.stringify({ link: input, isPaid: false })
					});

					if (!response.ok) throw new Error("Erro ao gerar link");

					const data = await response.json();
					const resultDiv = document.getElementById("result");
				
					document.getElementById("result").innerText = "Novo link: " + data.link;
				} catch (err) {
					document.getElementById("result").innerText = "Erro: " + err.message;
				}
			}
		</script>
	</body>
	</html>
	`)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/get") && r.Method == "GET" {
		handlerGetOrigin(w, r)
	} else if strings.Contains(r.URL.Path, "/newLink") && r.Method == "POST" {
		handlerNewLink(w, r)
	} else {
		homeHandler(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
