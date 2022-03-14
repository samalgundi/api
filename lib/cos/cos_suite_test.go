package cos_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cos Suite")
}
