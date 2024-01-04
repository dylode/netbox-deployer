package worker

type UpdateEvent string

const (
	Created UpdateEvent = "created"
	Updated UpdateEvent = "updated"
	Deleted UpdateEvent = "deleted"
)

type UpdateModel string

const (
	Virtualization UpdateModel = "virtualization"
	IPAM           UpdateModel = "ipam"
)

type UpdateModelID int

type Update struct {
	Event UpdateEvent   `json:"event"`
	Model UpdateModel   `json:"model"`
	ID    UpdateModelID `json:"id"`
}
