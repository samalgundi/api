package cos_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	camp "github.com/samalgundi/api/lib/campsite"
	cos "github.com/samalgundi/api/lib/cos"
)

var _ = Describe("Cos", func() {

	var myTestCampsite camp.Campsite

	BeforeEach(func() {
		myTestCampsite = camp.Campsite{
			Name:    "Villa BB",
			Country: "Germany",
			City:    "Boeblingen",
			Zip:     "71032",
		}
	})

	Describe("Adding campsites", func() {
		It("should not throw and error", func() {
			cos.PutCos(myTestCampsite)
			Expect(cos.PutCos(myTestCampsite)).To(BeNil())
		})
	})

})
