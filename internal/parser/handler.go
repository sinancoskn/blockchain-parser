package parser

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ParserHandler struct {
	parser Parser
}

func NewParserHandler(parser Parser) *ParserHandler {
	return &ParserHandler{parser}
}

func (s *ParserHandler) GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	block := s.parser.GetCurrentBlock()
	response := map[string]int{"current_block": block}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *ParserHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Address string `json:"address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.Address = strings.ToLower(req.Address)

	success := s.parser.Subscribe(req.Address)
	response := map[string]bool{"success": success}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *ParserHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}

	address = strings.ToLower(address)

	success := s.parser.Unsubscribe(address)
	response := map[string]bool{"success": success}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *ParserHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}

	address = strings.ToLower(address)

	transactions := s.parser.GetTransactions(address)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
