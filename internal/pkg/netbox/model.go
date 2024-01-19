package netbox

type EventType string

type ModelName string

type ModelID int

const (
	Created EventType = "created"
	Updated EventType = "updated"
	Deleted EventType = "deleted"
)

type WebhookEvent struct {
	EventType EventType `json:"event"`
	ModelName ModelName `json:"model"`
	ModelID   ModelID   `json:"id"`
}
