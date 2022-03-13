package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	camp "github.com/samalgundi/api/lib/campsite"
	cos "github.com/samalgundi/api/lib/cos"

)


func addCampSite(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting addCampSite execution:", r.FormValue("name"))

	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		log.Println("ParseForm() err: ", err)
		return
	}

	newCampsite := camp.Campsite{}

	newCampsite.Name = r.FormValue("name")
	newCampsite.Country = r.FormValue("country")
	newCampsite.City = r.FormValue("city")
	newCampsite.Zip = r.FormValue("zip")

	log.Println(newCampsite)
	cos.PutCos(newCampsite)

}

func main() {

	log.Println("Starting api execution.")

	//register URL paths and handlers
	r := mux.NewRouter()

	r.HandleFunc("/campsite", addCampSite).Methods("POST")
	http.Handle("/", r)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
