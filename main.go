package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"

	camp "github.com/samalgundi/api/lib/campsite"
	cos "github.com/samalgundi/api/lib/cos"
	
)

const apiConfig = "/config/api_config.json"

type configuration struct {
	  CosConfig cos.Configuration
}

// Builds a configuration object
func buildConfigurationObject(c configuration) configuration {

	log.Println("Building configuration file.")

	apiconf := &configuration{
		CosConfig: c.CosConfig,
	}

	return *apiconf
}

// Loads a configuration file, found in /config/api_config.json
func loadConfigurationFile() (configuration, error) {

	log.Println("Loading configuration file.")

	apiConfiguration := configuration{}

	// Using runtime.Caller, to make sure we get the path where the program is being executed
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return apiConfiguration, errors.New("error calling runtime caller")
	}

	// Reading configuration file
	apiConfigurationError := gonfig.GetConf(path.Dir(filename)+string(os.PathSeparator)+apiConfig, &apiConfiguration)

	if apiConfigurationError != nil {
		return apiConfiguration, apiConfigurationError
	}

	return apiConfiguration, nil
}

func addCampSite(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting addCampSite execution.")

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

	//	conf := aws.NewConfig().
	//		WithEndpoint(apiconf.ServiceEndpoint).
	//		WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(),
	//			apiconf.AuthEndpoint, apiconf.ApiKey, apiconf.ServiceInstanceID)).
	//		WithS3ForcePathStyle(true)
	//
	//	sess := session.Must(session.NewSession())
	//	client := s3.New(sess, conf)
	//
	//	// Variables and random content to sample, replace when appropriate
	//	bucketName := apiconf.BucketName
	//	key := "campsite.json"
	//	out, _ := json.Marshal(newCampsite)
	//
	//	log.Println(out)
	//
	//	content := bytes.NewReader([]byte(string(out)))
	//
	//	input := s3.PutObjectInput{
	//		Bucket: aws.String(bucketName),
	//		Key:    aws.String(key),
	//		Body:   content,
	//	}
	//
	//	// Call Function to upload (Put) an object
	//	result, _ := client.PutObject(&input)
	//	fmt.Println(result)
}

func init() {

	// Loading App Id configuration file
	apiConfiguration, apiConfigurationError := loadConfigurationFile()
	if apiConfigurationError != nil {
		log.Println("Could not load configuration file.")
	}

	// Building global apiconf object, using api configuration
	apiconf := buildConfigurationObject(apiConfiguration)

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
