# XM Companies Microservice

## Overview
This microservice handles CRUD operations for company data and integrates with Kafka for event messaging.
It provides JWT-secured REST APIs for creating, updating, deleting, and retrieving company records.

### Features
- **CRUD operations**: Create, Read, Update, Delete companies.
- **Secure APIs**: JWT authentication for restricted access.
- **Database**: Postgres SQL for storing companies data.
- **Kafka Integration**: Produces Kafka events for every mutating operation.
- **Containerized**: Docker and Docker Compose for seamless setup.

## Prerequisites
- Docker & Docker Compose
- Go (for local development)

## Setup

### 1. Clone the Repository
```bash
git clone https://github.com/nkostoulas/xm-companies
cd xm-companies
```

### 2. Start Services
```bash
docker-compose up --build
```
This command will:
- Build the application image
- Setup Postgres and migrations
- Start Kafka and Zookeeper
- Start the application

### 3. Access the Application
The application runs on `http://localhost:8080`.

## API Endpoints

### Authentication
Include the JWT token in the `Authorization` header for all endpoints:
```
Authorization: Bearer <token>
```

A JWT token signed by the application secret must be used, e.g. using jwt.io.

### CRUD Endpoints
- **Create a Company**
  ```bash
  POST /companies
  ```
  Request Body:
  ```json
  {
    "name": "Example Company",
    "description": "Sample description",
    "num_employees": 50,
    "is_registered": true,
    "type": "Corporation"
  }
  ```

- **Get a Company**
  ```bash
  GET /companies/{id}
  ```

- **Update a Company**
  ```bash
  PATCH /companies/{id}
  ```
  Request Body:
  ```json
  {
    "name": "Updated Name"
  }
  ```

- **Delete a Company**
  ```bash
  DELETE /companies/{id}
  ```

## Kafka Integration
Connect via the Kafka API:
- Topic: `company.events.v1`
- Broker: `localhost:9092`

## Development

### Run Locally
Ensure DB and Kafka are running locally or in Docker. Then:
```bash
go run cmd/main.go
```

### Testing
Tests can be run using Go's `testing` package. Example:
```bash
go test ./...
```

## Deployment
The application is containerized. Deploy it using Docker Compose in production.

## Future Improvements
- unit tests
- integration tests
- proper secrets management for jwt authentication
- proper secrets management for database access
- add authentication to Kafka
