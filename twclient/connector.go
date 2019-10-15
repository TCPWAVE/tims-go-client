package twclient

import(
  "bytes"
  "log"
  "time"
  "strings"
  "net/http"
  "net/url"
  "io/ioutil"
  "crypto/tls"
  "crypto/x509"
  "encoding/json"

  glog "github.com/golang/glog"
)

// Credentials houses client keys
type Credentials struct{
  keyFile       string
  certFile      string
  keyPair       tls.Certificate
  certPool      *x509.CertPool
}

// RestClient is a HTTP Client
type RestClient struct{
  clientConf        ClientConfig
  client          http.Client
}

// Init initialize rest client
func (rc *RestClient) Init(){
  glog.Info("Initialising rest client")
  tlsConfig := tls.Config{
    RootCAs: rc.clientConf.CertPool,
    InsecureSkipVerify: true,
    Certificates: []tls.Certificate{rc.clientConf.CertKeyPair},
  }
  tr := &http.Transport{
		TLSClientConfig: &tlsConfig,
		MaxIdleConnsPerHost: rc.clientConf.HTTPPoolConnections,
	}
  rc.client = http.Client{Transport: tr, Timeout: rc.clientConf.HTTPRequestTimeout * time.Second}
}

func getReqBody(payload Payload) ([]byte){
  var objJSON []byte
	var err error
  if payload.Body!=nil{
	   objJSON, err = json.Marshal(payload.Body)
  }
	if err != nil {
		glog.Errorf("Cannot marshal object '%s': %s", payload.Body, err)
		return nil
	}
	return objJSON
}

func prepareURL(hostConf HostConfig, payload Payload)(urlStr string){
  pathArr  := []string{"tims", "rest", payload.Path}
  qry := ""
	vals := url.Values{}
  if len(payload.Params) > 0 {
    for k,v := range payload.Params{
      vals.Set(k,v)
    }
		qry = vals.Encode()
  }
  u := url.URL{
		Scheme:   "https",
		Host:     hostConf.Host + ":" + hostConf.Port,
		Path:     strings.Join(pathArr, "/"),
		RawQuery: qry,
	}
  urlStr = u.String()
  return
}

// MakeRequest performs rest call
func (rc *RestClient) MakeRequest(payload Payload)([]byte, *ErrorResp){
  // build url
  urlStr := prepareURL(rc.clientConf.HostConfig, payload)
  log.Print(urlStr)
  // prepares req body
  var bodyStr []byte
  if payload.Body !=nil && payload.Method != "GET"{
    bodyStr = getReqBody(payload)
    log.Print("Body : " + string(bodyStr))
    glog.Infof("Object json : '%s'", string(bodyStr))
  }
  // creating request
  var req *http.Request
  method :=payload.Method
  req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(bodyStr))
  if err != nil {
		glog.Errorf("Error preparing HTTP request : '%s'", err)
		return nil, &ErrorResp{StatusCode:800, ErrorMsg:"Error Building Request"}
	}
  if method !="GET"{
    req.Header.Add("Content-Type", "application/json")
  }
  
  // Send Request
  var resp *http.Response
  resp, err = rc.client.Do(req)
  if err != nil {
    glog.Errorf("Error executing http request : %s", err)
	} else if !(resp.StatusCode == http.StatusOK || (resp.StatusCode == http.StatusCreated && req.Method == "POST")) {
		errRsp := getHTTPResponseError(resp)
		return nil, &errRsp
	}

	defer resp.Body.Close()
	res, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		glog.Errorf("Http Reponse ioutil.ReadAll() Error: '%s'", err)
		return nil, &ErrorResp{StatusCode:801, ErrorMsg:"Error Parsing Response"}
	}
	return res, nil
}

func getHTTPResponseError(resp *http.Response) ErrorResp {
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
  errObj :=&ErrorResp{Status:resp.Status, StatusCode:resp.StatusCode, ErrorMsg:string(content)}
	return *errObj
}

// NewRestClient returns new Rest Client
func NewRestClient(conf HostConfig, key string, cert string, connPool int, connTimeout time.Duration) ( *RestClient){
  cl := &RestClient{}
  cl.clientConf.HostConfig = conf
  caPool, certKeyPair := GetCertificates(cert, key)
  cl.clientConf.TransportConfig.CertPool = caPool
  cl.clientConf.TransportConfig.CertKeyPair = certKeyPair
  cl.clientConf.TransportConfig.KeyFile = key
  cl.clientConf.TransportConfig.CertFile = cert
  cl.clientConf.TransportConfig.HTTPPoolConnections = connPool
  cl.clientConf.TransportConfig.HTTPRequestTimeout = connTimeout
  cl.Init()
  return cl
}
