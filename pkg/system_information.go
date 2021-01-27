package pkg

// {
// 	"last_updated": 1553592653,
// 	"ttl": 10,
// 	"data": {
// 		"system_id": "oslobysykkel",
// 		"language": "nb",
// 		"name": "Oslo Bysykkel",
// 		"operator": "UIP Oslo Bysykkel AS",
// 		"timezone": "Europe/Oslo",
// 		"phone_number": "+4791589700",
// 		"email": "post@oslobysykkel.no"
// 	}
// }
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
