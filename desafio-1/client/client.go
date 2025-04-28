package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Erro ao criar requisição:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Erro ao fazer requisição:", err)
	}
	defer resp.Body.Close()

	var cotacao Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		log.Fatal("Erro ao decodificar resposta:", err)
	}

	err = ioutil.WriteFile("cotacao.txt", []byte("Dólar: "+cotacao.Bid), 0644)
	if err != nil {
		log.Fatal("Erro ao escrever no arquivo:", err)
	}

	log.Println("Cotação salva em cotacao.txt:", cotacao.Bid)
}
