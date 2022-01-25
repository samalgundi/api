package cos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"

	camp "github.com/samalgundi/api/lib/campsite"
)

// Configuration contains all configuration that is necessary for this package
type Configuration struct {
	APIKey            string
	AuthEndpoint      string	
	ServiceEndpoint   string
	ServiceInstanceID string
	BucketName        string
	COSClientConf     *aws.Config
}

// stores the configuration information
var conf Configuration

// Init initializes the COS package
func Init(c Configuration) {

	conf.BucketName = c.BucketName

	// configuration for the COS client
	conf.COSClientConf = aws.NewConfig().
		WithEndpoint(c.ServiceEndpoint).
		WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(),
			c.AuthEndpoint, c.APIKey, c.ServiceInstanceID)).
		WithS3ForcePathStyle(true)
}

// PutCos writes a file to COS
func PutCos(c camp.Campsite) {

	sess := session.Must(session.NewSession())
	client := s3.New(sess, conf.COSClientConf)

	// Variables and random content to sample, replace when appropriate
	bucketName := conf.BucketName
	key := "campsite.json"
	out, _ := json.Marshal(c)

	log.Println(out)

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
