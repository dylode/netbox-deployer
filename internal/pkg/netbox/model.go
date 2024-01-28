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

type virtualMachineComponent[T any] struct {
	ID   ModelID
	Data T
}

type interfaceIPAddress struct {
	Address string `model:"ipaddress"`
}

type virtualMachineInterface struct {
	IPAddresses []virtualMachineComponent[interfaceIPAddress] `model:"ipaddress"`
	VID         int32                                         `model:"vlan"`
	MacAddress  string                                        `model:"vminterface"`
}

type virtualMachineDisk struct {
	Name string
	Size uint64
}

type VirtualMachine struct {
	Status     string                                             `model:"virtualmachine"`
	CPUs       uint                                               `model:"virtualmachine"`
	Memory     uint64                                             `model:"virtualmachine"`
	Tags       []virtualMachineComponent[string]                  `model:"tag"`
	Interfaces []virtualMachineComponent[virtualMachineInterface] `model:"vminterface"`
	Disks      []virtualMachineComponent[virtualMachineDisk]      `model:"vmdisk"`
}
