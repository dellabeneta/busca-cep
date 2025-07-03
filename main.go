// main.go (versão final que lida com a inconsistência da API)
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	// ALTERAÇÃO 1: Mudamos de bool para string para aceitar a resposta da API
	Erro string `json:"erro"`
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/consultar", consultaCepHandler)
	log.Println("Servidor iniciado em http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Não foi possível iniciar o servidor: %s\n", err)
	}
}

func consultaCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if len(cep) != 8 {
		http.Error(w, "CEP com formato inválido", http.StatusBadRequest)
		return
	}

	url := "https://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Erro ao contatar a API ViaCEP", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "CEP com formato inválido", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erro ao ler a resposta da API", http.StatusInternalServerError)
		return
	}

	var viaCEPData ViaCEPResponse
	if err := json.Unmarshal(body, &viaCEPData); err != nil {
		http.Error(w, "Erro ao decodificar a resposta JSON", http.StatusInternalServerError)
		return
	}

	// ALTERAÇÃO 2: Comparamos o texto para ver se o CEP não foi encontrado
	if viaCEPData.Erro == "true" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}