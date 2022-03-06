package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	camp "github.com/samalgundi/api/lib/campsite"
)

var _ = Describe("Storing information about a campsite", Label("campsite"), func() {
	var campsite camp.Campsite

	When("information about a new campsite is received", func() {
		It("store information", func() {
			Expect(camp.AddCampSite).To(Succeed())
		})
	})
})
