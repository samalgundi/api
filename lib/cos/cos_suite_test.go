package cos_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	cos "github.com/samalgundi/api/lib/cos"
)

func TestCos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cos Suite")
}

var myTestBucket string

var _ = BeforeSuite(func() {

	myTestBucket = "bucket" + fmt.Sprint(time.Now().Unix())
	log.Println("Creating bucket:", myTestBucket)

	cos.SetBucket(myTestBucket)
	Expect(cos.CreateBucket(myTestBucket)).To(BeNil())
})

var _ = AfterSuite(func() {

	log.Println("Deleting bucket:", myTestBucket)
	Expect(cos.DeleteBucket(myTestBucket)).To(BeNil())
})
