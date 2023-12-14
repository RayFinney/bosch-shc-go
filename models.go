package bosch_shc_go

const StatusAvailable = "AVAILABLE"
const StatusUnavailable = "UNAVAILABLE"

type Options struct {
	ShcIp               string
	ShcPort             int
	ValidateCertificate bool
	ApiVersion          string
}

type Error struct {
	Type       string `json:"@type"`
	ErrorCode  string `json:"errorCode"`
	StatusCode int    `json:"statusCode"`
}

type Device struct {
	Type             string   `json:"@type"`
	RootDeviceId     string   `json:"rootDeviceId"`
	Id               string   `json:"id"`
	DeviceServiceIds []string `json:"deviceServiceIds"`
	Manufacturer     string   `json:"manufacturer"`
	RoomId           string   `json:"roomId"`
	DeviceModel      string   `json:"deviceModel"`
	Serial           string   `json:"serial"`
	Profile          string   `json:"profile"`
	Name             string   `json:"name"`
	Status           string   `json:"status"`
}

type Room struct {
	Type   string `json:"@type"`
	Id     string `json:"id"`
	IconId string `json:"iconId"`
	Name   string `json:"name"`
}

type Scenario struct {
	Type    string        `json:"@type"`
	Id      string        `json:"id"`
	IconId  string        `json:"iconId"`
	Name    string        `json:"name"`
	Actions []interface{} `json:"actions"`
}
