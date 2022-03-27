package cos_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/google/uuid"
	cos "github.com/samalgundi/api/lib/cos"
	defs "github.com/samalgundi/api/lib/definitions"
)

var _ = Describe("Cos", func() {

	myTestCampsite := defs.NewLocation("campsite")

	BeforeEach(func() {

		myTestCampsite.Name = "Villa BB"
		myTestCampsite.Country = "Germany"
		myTestCampsite.City = "Boeblingen"
		myTestCampsite.Zip = "71032"

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
