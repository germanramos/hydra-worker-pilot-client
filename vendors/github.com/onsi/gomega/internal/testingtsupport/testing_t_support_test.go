package testingtsupport_test

import (
	. "github.com/innotech/hydra-worker-pilot-client/vendors/github.com/onsi/gomega"

	"testing"
)

func TestTestingT(t *testing.T) {
	RegisterTestingT(t)
	Ω(true).Should(BeTrue())
}
