package definitions

import "github.com/google/uuid"

//campsite json request payload is as follows,
//{
//  "name":    "thename",
//  "country": "thecountry",
//  "city":    "thecity",
//  "zip":     "zipcode",
//  "type":    "type"
//}
type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
	UUID    string `json:"uuid"`
	Type    string `json:"type"`
}

// returns a new Location with a uuid
func NewLocation(t string) *Location {
	l := Location{Type: t}

	l.UUID = uuid.New().String()

	return &l
}
