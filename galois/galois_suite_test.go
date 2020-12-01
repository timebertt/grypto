package galois_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGalois(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Galois Field Suite")
}
