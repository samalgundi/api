package definitions

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

func AddCampSite() {

}
