package models

// EventType defines the event type
type EventType string

const (
	EventTypeCreated EventType = "created"
	EventTypePatched EventType = "patched"
	EventTypeDeleted EventType = "deleted"
)

// CompanyEvent defines a company event
type CompanyEvent struct {
	CompanyID string    `json:"company_id"`
	EventType EventType `json:"event_type"`
}
