package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"xm-companies/pkg/kafka"
	"xm-companies/pkg/models"
)

// CreateCompany is the API handler for creating a new Company
func (h *CompaniesHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := company.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	company.ID = uuid.New().String()
	err := h.db.InsertCompany(company)
	if err != nil {
		// Check if the error is due to a duplicate name
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			http.Error(w, "Company with the same name already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error inserting company: %v", err)
		return
	}

	response, _ := json.Marshal(company)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

	go func() {
		event := models.CompanyEvent{
			CompanyID: company.ID,
			EventType: models.EventTypeCreated,
		}
		b, _ := json.Marshal(event)
		kafkaMessage := kafka.Message{
			Key:   company.ID,
			Value: b,
		}
		go h.producer.Produce(&kafkaMessage)
	}()
}
