package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"xm-companies/pkg/kafka"
	"xm-companies/pkg/models"
)

// DeleteCompany is the API handler for deleting a Company
func (h *CompaniesHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err := h.db.DeleteCompany(id)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		log.Printf("Error deleting company: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(map[string]string{"status": "deleted"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

	go func() {
		event := models.CompanyEvent{
			CompanyID: id,
			EventType: models.EventTypeDeleted,
		}
		b, _ := json.Marshal(event)
		kafkaMessage := kafka.Message{
			Key:   id,
			Value: b,
		}
		go h.producer.Produce(&kafkaMessage)
	}()
}
