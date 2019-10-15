package twclient

import (
	"log"
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTCPwaveGoClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TcpWaveGoClient Test Suite")
}

var _ = BeforeSuite(func(){
	log.Print("Before Suite Execution")
})

var _ = AfterSuite(func() {
  log.Print("After Suite Execution")
})
