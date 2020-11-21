package rsa32_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRSA32(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RSA32 Suite")
}
