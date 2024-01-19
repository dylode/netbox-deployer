package netbox

type EventType string

type ModelName string

type ModelID int

const (
	EventCreated EventType = "created"
	EventUpdated EventType = "updated"
	EventDeleted EventType = "deleted"
)

type WebhookEvent struct {
	EventType EventType `json:"event"`
	ModelName ModelName `json:"model"`
	ModelID   ModelID   `json:"id"`
}
