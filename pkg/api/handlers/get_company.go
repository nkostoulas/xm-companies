package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetCompany is the API handler for fetching a Company by id
func (h *CompaniesHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	company, err := h.db.SelectCompany(id)
	if err != nil {
		log.Printf("Error fetching company: %v", err)
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(company)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
