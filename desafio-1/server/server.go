package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Cotacao struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func main() {
	http.HandleFunc("/cotacao", CotacaoHandler)
	log.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, "Erro ao criar request", http.StatusInternalServerError)
		log.Println("Erro ao criar request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Erro ao buscar cotação", http.StatusInternalServerError)
		log.Println("Erro ao buscar cotação:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, "Erro ao decodificar cotação", http.StatusInternalServerError)
		log.Println("Erro ao decodificar cotação:", err)
		return
	}

	cotacao := result["USDBRL"]

	if err := salvarCotacao(cotacao); err != nil {
		log.Println("Erro ao salvar cotação:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"bid": cotacao.Bid,
	})
}

func salvarCotacao(cotacao Cotacao) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS cotacoes (
			id INTEGER PRIMARY KEY,
			code TEXT,
			codein TEXT,
			name TEXT,
			high TEXT,
			low TEXT,
			varBid TEXT,
			pctChange TEXT,
			bid TEXT,
			ask TEXT,
			timestamp TEXT,
			create_date TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`); err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, `
		INSERT INTO cotacoes (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		cotacao.Code, cotacao.Codein, cotacao.Name, cotacao.High, cotacao.Low, cotacao.VarBid, cotacao.PctChange, cotacao.Bid, cotacao.Ask, cotacao.Timestamp, cotacao.CreateDate); err != nil {
		return err
	}

	return nil
}
