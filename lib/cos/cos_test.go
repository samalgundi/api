package cos_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/google/uuid"
	definitions "github.com/samalgundi/api/lib/campsite"
	cos "github.com/samalgundi/api/lib/cos"
)

var _ = Describe("Cos", func() {

	var myTestCampsite definitions.Location

	BeforeEach(func() {
		myTestCampsite = definitions.Location{
			Name:    "Villa BB",
			Country: "Germany",
			City:    "Boeblingen",
			Zip:     "71032",
			Type:    "campsite",
		}

	})

	Describe("Adding campsites", func() {
		It("should not throw an error", func() {
			myTestCampsite.UUID = uuid.New().String()
			// create the object
			Expect(cos.PutObjectIntoCos(myTestCampsite)).To(BeNil())
			// delete the object
			Expect(cos.DeleteObjectFromCos(myTestCampsite)).To(BeNil())
		})
	})

})
