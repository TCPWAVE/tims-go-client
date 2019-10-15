package twclient

import(
	"time"
	"errors"
	"strings"
	"encoding/json"
	glog "github.com/golang/glog"
)
// ObjectManager manages all object types
type ObjectManager struct {
	Client		*RestClient
}

// NewObjectManager create new instance
func NewObjectManager(conf HostConfig, key string, cert string, poolSize int, timeout time.Duration) (*ObjectManager) {
	clnt := NewRestClient(conf, key, cert, poolSize, timeout)
	objMgr := &ObjectManager{Client:clnt}
	return objMgr
}

// IObjectManager defines methods for managing objects
type IObjectManager interface {
  CreateNetwork(network Network) (string,error)
	GetNetwork(ipAddr string, orgName string) (*Network,error)
	CreateSubnet(subnet Subnet) (string,error)
	GetSubnet(subnetAddr string, orgName string) (*Subnet,error)
	GetNextFreeIp(subnetAddr string, orgName string) (string,error)
	CreateIPAddress(ipAddr string, macAddr string, subnetAddr string, domain string, org string, name string)(string,error)
	DeleteIPAddress(ip string, subnetAddr string, organization string) error
}

// CreateNetwork creates new Network
func (objMgr *ObjectManager) CreateNetwork(network Network) (string, error){
	glog.Infof("Creating Network : %s", network.Name)
	payload := &Payload{Method:"POST", Path:"network/add", Body: network}
	res,err := objMgr.Client.MakeRequest(*payload)
	if err!=nil{
		glog.Errorf("Error creating network :: Msg :'%s', Status : %s", err.ErrorMsg, err.Status)
		return "", errors.New(err.ErrorMsg)
	}
	resStr := string(res)
	return resStr,nil
}

// GetNetwork Fetches existing network details
func (objMgr *ObjectManager) GetNetwork(ipAddr string, orgName string) (*Network,error){
	glog.Infof("Fetching details for Network : %s", ipAddr)
	payload :=&Payload{Method:"GET", Path:"network/detailsByIP"}
	params := make(map[string]string)
	ip := strings.Split(ipAddr, "/")[0]
	adds := strings.Split(ip, ".")
	params["addr1"] = adds[0]
	params["addr2"] = adds[1]
	params["addr3"] = adds[2]
	params["addr4"] = adds[3]
	params["organizationName"] = orgName
	payload.Params = params
	res,err := objMgr.Client.MakeRequest(*payload)
	if err!=nil{
		glog.Errorf("Error Fetching network :: Status : %s", err.Status)
		return nil, errors.New(err.Status)
	}
	var network Network
	err1 := json.Unmarshal(res, &network)
	if err1 != nil {
		glog.Errorf("Error Parsing Network Response : '%s'", err1)
		return nil,err1
	}
	return &network, nil
}

// CreateSubnet creates a new Subnet
func (objMgr *ObjectManager) CreateSubnet(subnet Subnet) (string,error){
	glog.Infof("Creating subnet : %s", subnet.Name)
	payload := &Payload{Method:"POST", Path:"subnet/add", Body: subnet}
	res,err := objMgr.Client.MakeRequest(*payload)
	if err!=nil{
		glog.Errorf("Error creating subnet :: Msg :'%s', Status : %s", err.ErrorMsg, err.Status)
		return "", errors.New(err.ErrorMsg)
	}
	resStr := string(res)
	return resStr,nil
}

// GetSubnet fetches esisting subnet
func (objMgr *ObjectManager) GetSubnet(subnetAddr string, orgName string) (*Subnet,error){
	glog.Infof("Fetching details for subnet : %s", subnetAddr)
	payload :=&Payload{Method:"GET", Path:"subnet/getSubnetData"}
	params := make(map[string]string)
	ip := strings.Split(subnetAddr, "/")[0]
	params["subnet_address"] = ip
	params["org_name"] = orgName
	payload.Params = params
	res,err := objMgr.Client.MakeRequest(*payload)
	if err!=nil{
		glog.Errorf("Error Fetching subnet :: Msg :'%s', Status : %s", err.ErrorMsg, err.Status)
		return nil, errors.New(err.Status)
	}
	var subnet Subnet
	err1 := json.Unmarshal(res, &subnet)
	if err1 != nil {
		glog.Errorf("Error Parsing Subnet Response : '%s'", err1)
		return nil,err1
	}
	return &subnet, nil
}

// GetNextFreeIP creates a new Subnet
func (objMgr *ObjectManager) GetNextFreeIP(subnetAddr string, orgName string) (string,error){
	glog.Infof("Fetching Next Free Ip for subnet : %s", subnetAddr)
	//payload :=&Payload{Method:"GET", Path:"object/nextfreeip"}
	payload :=&Payload{Method:"GET", Path:"object/getNextFreeIP"}
	params := make(map[string]string)
	ip := strings.Split(subnetAddr, "/")[0]
	params["subnet_addr"] = ip
	params["org_name"] = orgName
	payload.Params = params
	res,err := objMgr.Client.MakeRequest(*payload)
	if err!=nil{
		glog.Errorf("Error Fetching Next Free Ip :: Msg :'%s', Status : %s", err.ErrorMsg, err.Status)
		return "", errors.New(err.Status)
	}
	return string(res), nil
}

// CreateIPAddress creates Ip object
func (objMgr *ObjectManager) CreateIPAddress(ipAddr string, macAddr string, subnetAddr string, domain string, org string, name string)(string,error){
	glog.Infof("Creating Ip %s in subnet : %s", ipAddr, subnetAddr)
	ipObj := NewIPObject(ipAddr, subnetAddr, domain, org, name)
	payload :=&Payload{Method:"POST", Path:"object/add", Body:ipObj}
	res,err := objMgr.Client.MakeRequest(*payload)
	if err!=nil{
		glog.Errorf("Error Creating Ip :: Msg :'%s', Status : %s", err.ErrorMsg, err.Status)
		return "", errors.New(err.Status)
	}
	return string(res), nil
}

// DeleteIPAddress deletes Ip object
func (objMgr *ObjectManager) DeleteIPAddress(ip string, subnetAddr string, organization string) error {
	glog.Infof("Deleting Ip %s in subnet : %s", ip, subnetAddr)
	addresses := []string{ip}
	ipDelObj := &IPObjDel{
		Addresses: 	addresses,
		Org : 			organization,
		DelRRSChked : 0,
	}
	payload := &Payload{Method:"POST", Path:"object/reclaimObjects", Body: *ipDelObj}
	_,err := objMgr.Client.MakeRequest(*payload)
	if err!=nil{
		glog.Errorf("Error Deleting Ip :: Msg :'%s', Status : %s", err.ErrorMsg, err.Status)
		return errors.New(err.Status)
	}
	return nil
}
