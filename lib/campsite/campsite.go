package camp

//campsite json request payload is as follows,
//{
//  "name":    "thename",
//  "country": "thecountry",
//  "city":    "thecity",
//  "zip":     "zipcode"
//}
type Campsite struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
}

func AddCampSite() {

}
