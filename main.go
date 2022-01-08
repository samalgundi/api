package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/tkanos/gonfig"
)

var conf ApiConfiguration

const API_CONFIG = "/config/api_config.json"

type ApiConfiguration struct {
	ApiKey            string
	ServiceInstanceID string
	AuthEndpoint      string
	ServiceEndpoint   string
}

// Builds a configuration object
func buildConfigurationObject(api_configuration ApiConfiguration) ApiConfiguration {

	log.Println("Building configuration file.")

	conf := &ApiConfiguration{
		ApiKey:            api_configuration.ApiKey,
		ServiceInstanceID: api_configuration.ServiceInstanceID,
		AuthEndpoint:      api_configuration.AuthEndpoint,
		ServiceEndpoint:   api_configuration.ServiceEndpoint,
	}

	return *conf
}

// Loads a configuration file, found in /config/api_config.json
func loadConfigurationFile() (ApiConfiguration, error) {

	log.Println("Loading configuration file.")

	api_configuration := ApiConfiguration{}

	// Using runtime.Caller, to make sure we get the path where the program is being executed
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return api_configuration, errors.New("Error calling runtime caller.")
	}

	// Reading configuration file
	api_configuration_error := gonfig.GetConf(path.Dir(filename)+string(os.PathSeparator)+API_CONFIG, &api_configuration)

	if api_configuration_error != nil {
		return api_configuration, api_configuration_error
	}

	return api_configuration, nil
}

func main() {

	log.Println("Starting api execution.")

	// Loading App Id configuration file
	api_configuration, api_configuration_error := loadConfigurationFile()
	if api_configuration_error != nil {
		log.Println("Could not load configuration file.")
	}

	// Building global conf object, using api configuration
	conf = buildConfigurationObject(api_configuration)

	newBucket := "new-bucketee-campsitelist"

	conf := aws.NewConfig().
		WithEndpoint(conf.ServiceEndpoint).
		WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(),
			conf.AuthEndpoint, conf.ApiKey, conf.ServiceInstanceID)).
		WithS3ForcePathStyle(true)

	sess := session.Must(session.NewSession())
	client := s3.New(sess, conf)

	input := &s3.CreateBucketInput{
		Bucket: aws.String(newBucket),
	}
	client.CreateBucket(input)

	d, _ := client.ListBuckets(&s3.ListBucketsInput{})
	fmt.Println(d)
}
