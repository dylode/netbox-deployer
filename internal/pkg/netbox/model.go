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

type VirtualMachineComponent[T any] struct {
	ID   ModelID
	Data T
}

type InterfaceIPAddress struct {
	Address string `model:"ipaddress"`
}

type VirtualMachineInterface struct {
	IPAddresses []VirtualMachineComponent[InterfaceIPAddress] `model:"ipaddress"`
	VID         int32                                         `model:"vlan"`
	MacAddress  string                                        `model:"vminterface"`
}

type VirtualMachineDisk struct {
	Name string
	Size uint64
}

type VirtualMachine struct {
	Status     string                                             `model:"virtualmachine"`
	CPUs       uint                                               `model:"virtualmachine"`
	Memory     uint64                                             `model:"virtualmachine"`
	Tags       []VirtualMachineComponent[string]                  `model:"tag"`
	Interfaces []VirtualMachineComponent[VirtualMachineInterface] `model:"vminterface"`
	Disks      []VirtualMachineComponent[VirtualMachineDisk]      `model:"vmdisk"`
}
