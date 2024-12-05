package server

import (
	"xm-companies/internal/api/handlers"
	"xm-companies/internal/api/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// NewServer returns a new API server
func NewServer(
	handler *handlers.CompaniesHandler,
	jwtMiddleware *middleware.JWTMiddleware,
) *mux.Router {
	r := mux.NewRouter()

	// Public route
	r.HandleFunc("/companies/{id}", handler.GetCompany).Methods("GET")

	// Routes that require auth
	secured := r.PathPrefix("/").Subrouter()
	secured.Use(jwtMiddleware.AuthHandler)
	secured.HandleFunc("/companies", handler.CreateCompany).Methods("POST")
	secured.HandleFunc("/companies/{id}", handler.PatchCompany).Methods("PATCH")
	secured.HandleFunc("/companies/{id}", handler.DeleteCompany).Methods("DELETE")

	return r
}
