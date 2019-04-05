package ginkgo_example_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoExample Suite")
}
