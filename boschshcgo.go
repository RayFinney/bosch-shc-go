package bosch_shc_go

type BoschShcGo interface {
	// GetDevices returns all devices.
	GetDevices() ([]Device, error)
	// GetDevice returns a device by id.
	GetDevice(id string) (Device, error)

	// GetRooms returns all rooms.
	GetRooms() ([]Room, error)
	// GetRoom returns a room by id.
	GetRoom(id string) (Room, error)

	// GetScenarios returns all scenarios.
	GetScenarios() ([]Scenario, error)
	// GetScenario returns a scenario by id.
	GetScenario(id string) (Scenario, error)
	// TriggerScenario triggers a scenario by id.
	TriggerScenario(id string) error

	// GetMessages returns all messages.
	GetMessages() ([]Message, error)
}
