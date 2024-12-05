package main

import (
	"database/sql"
	"log"
	"net/http"

	"xm-companies/internal/api/handlers"
	"xm-companies/internal/api/middleware"
	"xm-companies/internal/api/server"
	"xm-companies/internal/config"
	companiesdb "xm-companies/internal/db"
	"xm-companies/pkg/kafka"

	_ "github.com/lib/pq"
)

func mustDBConnect(dbHost string) *sql.DB {
	dbConn, err := sql.Open("postgres", dbHost)
	if err != nil {
		log.Panicf("Failed to establish database connection: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Panicf("Failed to ping database: %v", err)
	}

	log.Println("Database connected successfully.")
	return dbConn
}

func main() {
	appConfig, err := config.InitConfig()
	if err != nil {
		log.Panicf("Error initializing config: %v", err)
	}
	dbConn := mustDBConnect(appConfig.DBConnStr)
	defer dbConn.Close()

	producer, err := kafka.NewProducer(
		&kafka.ProducerConfig{
			Brokers:   appConfig.Brokers,
			TopicName: appConfig.EventTopic,
		},
	)
	if err != nil {
		log.Panicf("Error creating producer: %v", err)
	}
	defer producer.Close()

	db := companiesdb.NewDB(dbConn)
	handler := handlers.NewCompaniesHandler(db, producer)
	jwtMiddleware := middleware.NewJWTMiddleware(appConfig.Secret)

	r := server.NewServer(handler, jwtMiddleware)
	http.Handle("/", r)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
