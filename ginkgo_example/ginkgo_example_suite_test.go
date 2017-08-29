package ginkgo_example_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGinkgoExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoExample Suite")
}
