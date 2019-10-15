package twclient

import(
  "crypto/tls"
  "crypto/x509"
  . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("test config module", func(){
    const(
      certFile = "/opt/tcpwave/certs/client.crt"
      keyFile = "/opt/tcpwave/certs/client.key"
    )
    Context("with valid certificates", func(){
        var certPool *x509.CertPool
        var keypair  tls.Certificate
        It("should correctly parse certificates", func(){
            certPool, keypair = GetCertificates(certFile, keyFile)
            Expect(certPool).NotTo(Equal(nil))
            Expect(keypair).NotTo(Equal(nil))
        })

    })

})
