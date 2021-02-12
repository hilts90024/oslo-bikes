package models

type SystemInformationResponse struct {
	LastUpdated uint32
	TTL         uint
	Data        *SystemInformation
}

type SystemInformation struct {
	SystemID    string `json:"system_id"`
	Language    string
	Name        string
	Operator    string
	Timezone    string
	PhoneNumber string `json:"phone_number"`
	Email       string
}
