package des_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DES Suite")
}
