package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
)

var apiconf ApiConfiguration

const API_CONFIG = "/config/api_config.json"

type ApiConfiguration struct {
	ApiKey            string
	AuthEndpoint      string
	ServiceEndpoint   string
	ServiceInstanceID string
	BucketName        string
}

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

// Builds a configuration object
func buildConfigurationObject(api_configuration ApiConfiguration) ApiConfiguration {

	log.Println("Building configuration file.")

	apiconf := &ApiConfiguration{
		ApiKey:            api_configuration.ApiKey,
		AuthEndpoint:      api_configuration.AuthEndpoint,
		ServiceEndpoint:   api_configuration.ServiceEndpoint,
		ServiceInstanceID: api_configuration.ServiceInstanceID,
		BucketName:        api_configuration.BucketName,
	}

	return *apiconf
}

// Loads a configuration file, found in /config/api_config.json
func loadConfigurationFile() (ApiConfiguration, error) {

	log.Println("Loading configuration file.")

	api_configuration := ApiConfiguration{}

	// Using runtime.Caller, to make sure we get the path where the program is being executed
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return api_configuration, errors.New("error calling runtime caller")
	}

	// Reading configuration file
	api_configuration_error := gonfig.GetConf(path.Dir(filename)+string(os.PathSeparator)+API_CONFIG, &api_configuration)

	if api_configuration_error != nil {
		return api_configuration, api_configuration_error
	}

	return api_configuration, nil
}

func addCampSite(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting addCampSite execution.")

	//get input from body
	var newCampsite Campsite
	json.NewDecoder(r.Body).Decode(&newCampsite)
	log.Print("Successfully received new campsite info: ", newCampsite)

	conf := aws.NewConfig().
		WithEndpoint(apiconf.ServiceEndpoint).
		WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(),
			apiconf.AuthEndpoint, apiconf.ApiKey, apiconf.ServiceInstanceID)).
		WithS3ForcePathStyle(true)

	sess := session.Must(session.NewSession())
	client := s3.New(sess, conf)

	// Variables and random content to sample, replace when appropriate
	bucketName := apiconf.BucketName
	key := "campsite.json"
	out, _ := json.Marshal(newCampsite)

	content := bytes.NewReader([]byte(string(out)))

	input := s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   content,
	}

	// Call Function to upload (Put) an object
	result, _ := client.PutObject(&input)
	fmt.Println(result)
}

func main() {

	log.Println("Starting api execution.")

	// Loading App Id configuration file
	api_configuration, api_configuration_error := loadConfigurationFile()
	if api_configuration_error != nil {
		log.Println("Could not load configuration file.")
	}

	// Building global apiconf object, using api configuration
	apiconf = buildConfigurationObject(api_configuration)

	//register URL paths and handlers
	r := mux.NewRouter()

	r.HandleFunc("/campsite", addCampSite).Methods("POST")
	http.Handle("/", r)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
