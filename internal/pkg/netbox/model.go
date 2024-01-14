package netbox

type UpdateEvent string

type Model string

type ModelID int

const (
	Created UpdateEvent = "created"
	Updated UpdateEvent = "updated"
	Deleted UpdateEvent = "deleted"
)

type Update struct {
	Event UpdateEvent `json:"event"`
	Model Model       `json:"model"`
	ID    ModelID     `json:"id"`
}
