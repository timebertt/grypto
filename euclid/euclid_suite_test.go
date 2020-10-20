package euclid_test

import (
  "testing"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

func TestEuclid(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Euclidean Algorithm Suite")
}
