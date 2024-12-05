package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"xm-companies/pkg/kafka"
	"xm-companies/pkg/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// PatchCompany is the API handler for updating a Company
func (h *CompaniesHandler) PatchCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updates models.UpdateCompany
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Ensure at least one field is provided for update
	if updates.Name == nil && updates.Description == nil && updates.NumEmployees == nil &&
		updates.IsRegistered == nil && updates.Type == nil {
		http.Error(w, "No valid fields provided for update", http.StatusBadRequest)
		return
	}

	if err := updates.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.db.UpdateCompany(id, updates)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			http.Error(w, "Company with the same name already exists", http.StatusConflict)
			return
		}
		log.Printf("Error updating company: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(map[string]string{"status": "updated"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

	go func() {
		event := models.CompanyEvent{
			CompanyID: id,
			EventType: models.EventTypePatched,
		}
		b, _ := json.Marshal(event)
		kafkaMessage := kafka.Message{
			Key:   id,
			Value: b,
		}
		go h.producer.Produce(&kafkaMessage)
	}()
}
