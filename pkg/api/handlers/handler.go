package handlers

import (
	companiesdb "xm-companies/pkg/db"
	"xm-companies/pkg/kafka"
)

// CompaniesHandler is the companies HTTP API handler
type CompaniesHandler struct {
	db       *companiesdb.DB
	producer *kafka.Producer
}

// NewCompaniesHandler returns a new CompaniesHandler
func NewCompaniesHandler(db *companiesdb.DB, producer *kafka.Producer) *CompaniesHandler {
	return &CompaniesHandler{db: db, producer: producer}
}
