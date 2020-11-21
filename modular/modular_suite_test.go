package modular_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestModular(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Modular Suite")
}
