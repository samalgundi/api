package cos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/tkanos/gonfig"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"

	camp "github.com/samalgundi/api/lib/campsite"
)

// expects a configuration file at this location
const cosConfig = "/config/config.json"

// configuration contains all configuration that is necessary for this package
type configuration struct {
	APIKey            string
	AuthEndpoint      string
	ServiceEndpoint   string
	ServiceInstanceID string
	BucketName        string
	COSClientConf     *aws.Config
}

// stores the configuration information
var conf configuration

// Init initializes the COS package
func init() {

	log.Println("Starting cos.Init execution.")

	// Loading App Id configuration file
	confError := loadConfigurationFile(&conf)
	if confError != nil {
		log.Println("Could not load configuration file.")
	}

	// configuration for the COS client
	conf.COSClientConf = aws.NewConfig().
		WithEndpoint(conf.ServiceEndpoint).
		WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(),
			conf.AuthEndpoint, conf.APIKey, conf.ServiceInstanceID)).
		WithS3ForcePathStyle(true)

}

// Loads a configuration file, found in /config/api_config.json
func loadConfigurationFile(c *configuration) error {

	log.Println("Loading configuration file.")

	// Using runtime.Caller, to make sure we get the path where the program is being executed
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return errors.New("error calling runtime caller")
	}

	// Reading configuration file
	cosConfigurationError := gonfig.GetConf(path.Dir(filename)+string(os.PathSeparator)+cosConfig, c)

	if cosConfigurationError != nil {
		return cosConfigurationError
	}

	return nil
}

// PutCos writes a file to COS
func PutCos(c camp.Campsite) error {

	log.Println("Starting cos.PutCos execution.")

	sess := session.Must(session.NewSession())
	client := s3.New(sess, conf.COSClientConf)

	// Variables and random content to sample, replace when appropriate
	bucketName := conf.BucketName
	key := "campsite.json"
	out, _ := json.Marshal(c)

	content := bytes.NewReader([]byte(string(out)))

	input := s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   content,
	}

	// Call Function to upload (Put) an object
	result, _ := client.PutObject(&input)
	fmt.Println(result)

	return nil

}
