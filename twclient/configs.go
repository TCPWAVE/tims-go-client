package twclient

import(
  "time"
  "io/ioutil"
  "crypto/tls"
  "crypto/x509"

  log "github.com/golang/glog"
)

// HostConfig holds config for connecting to Ipam Server
type HostConfig struct{
  Host                 string
  Port                 string
}

// TransportConfig contains HTTP transport config
type TransportConfig struct {
  CertFile               string
  KeyFile                string
	CertPool               *x509.CertPool
  CertKeyPair            tls.Certificate
	HTTPRequestTimeout     time.Duration // in seconds
	HTTPPoolConnections    int
}

// ClientConfig hold all config required for creating client
type ClientConfig struct{
  HostConfig
  TransportConfig
}

// GetCertificates reads and prepares certificate files
func GetCertificates(certFile string, keyFile string) (caPool *x509.CertPool, keyPair tls.Certificate) {
  cert, err := ioutil.ReadFile(certFile)
  if err != nil {
			log.Errorf("Cannot load certificate file '%s'", certFile)
      return
	}

  caPool = x509.NewCertPool()
  if !caPool.AppendCertsFromPEM(cert){
    log.Errorf("Cannot append certificate from file '%s'", certFile)
    return
  }

  crt, err1 := tls.LoadX509KeyPair(certFile, keyFile)
  if err1 != nil {
    log.Errorf("Cannot load certificate pair '%s'", err1)
    return
  }
  keyPair = crt
  return
}

// NewClientConfig creates new client config
func NewClientConfig(hostConf *HostConfig, certFile string, keyFile string, httpRequestTimeout int, httpPoolConnections int) (ClientConfig){
  cfg := &ClientConfig{HostConfig: *hostConf}
  certPool, certKeypair := GetCertificates(certFile, keyFile)
  cfg.TransportConfig.CertKeyPair = certKeypair
  cfg.TransportConfig.CertPool = certPool
  cfg.TransportConfig.KeyFile = keyFile
  cfg.TransportConfig.CertFile = certFile
  cfg.TransportConfig.HTTPPoolConnections = httpPoolConnections
  cfg.TransportConfig.HTTPRequestTimeout = time.Duration(httpRequestTimeout)
  return *cfg
}

// NewTransportConfig creates the transport configuration
func NewTransportConfig(certFile string, keyFile string, httpRequestTimeout int, httpPoolConnections int) (cfg TransportConfig) {
  certPool, certKeypair := GetCertificates(certFile, keyFile)
  cfg.CertPool = certPool
  cfg.HTTPPoolConnections = httpPoolConnections
  cfg.HTTPRequestTimeout = time.Duration(httpRequestTimeout)
  cfg.CertKeyPair = certKeypair
  return
}
