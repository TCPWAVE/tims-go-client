package twclient

import (
  "strings"
  "strconv"
)
// Payload hold request payload
type Payload struct {
  Method      string
  Path        string
  Params      map[string]string
  Body        interface{}
}

// Base base object
type Base struct {
  Name              string    `json:"name,omitempty"`
  Description       string    `json:"description,omitempty"`
  Organization      string    `json:"organization_name,omitempty"`
  OrganizationID    int       `json:"organization_id,omitempty"`
  Addr1             int       `json:"addr1"`
  Addr2             int       `json:"addr2"`
  Addr3             int       `json:"addr3"`
  Addr4             int       `json:"addr4"`
  Address           string    `json:"address,omitempty"`
}

// Network Represents IPV4 network
type Network struct {
  Base
  MaskLen       int       `json:"mask_length,omitempty"`
  DMZVisible    string    `json:"dmzVisible,omitempty"`
  DNSSecEnable  string    `json:"dnssec_enable,omitempty"`
  Discovery     string    `json:"enable_discovery,omitempty"`
}

// Subnet Represents IPV4 network
type Subnet struct {
  Base
  MaskLen           int       `json:"mask_length,omitempty"`
  RouterAddr        string    `json:"routerAddress,omitempty"`
  NetworkAddr       string    `json:"network_address,omitempty"`
  PrimaryDomain     string    `json:"primary_domain,omitempty"`
  SubnetAddress     string    `json:"fullAddress,omitempty"`
  NetworkMask       int       `json:"network_mask,omitempty"`
}

// IPObject to create IP Address
type IPObject struct {
  Base
  AllocType       int         `json:"alloc_type"`
  Class           string      `json:"class_code,omitempty"`
  Domain          string      `json:"domain_name,omitempty"`
  SubnetAddr      string      `json:"subnet_address,omitempty"`
  UpdateNaA       bool        `json:"update_ns_a,omitempty"`
  UpdateNsPtr     bool        `json:"update_ns_ptr,omitempty"`
  DynUpdRrsA      bool        `json:"dyn_update_rrs_a,omitempty"`
  DynUpdRrsPtr    bool        `json:"dyn_update_rrs_ptr,omitempty"`
  DynUpdRrsCName  bool        `json:"dyn_update_rrs_cname,omitempty"`
  DynUpdRrsMx     bool        `json:"dyn_update_rrs_mx,omitempty"`
}

// NewIPObject creates new instance of Ip object
func NewIPObject(ipAddr string, subnet string, domain string, org string, name string)(*IPObject){
  ipSlice := strings.Split(ipAddr, ".")
  ipObj := &IPObject{AllocType:1, Class:"Others", UpdateNaA:true,
    UpdateNsPtr:true, DynUpdRrsA:true, DynUpdRrsPtr:true,
    DynUpdRrsCName:true, DynUpdRrsMx:true}
  ipObj.Name = name
  ipObj.Domain = domain
  ipObj.Organization = org
  ipObj.SubnetAddr = subnet
  adr, _ := strconv.Atoi(ipSlice[0])
  ipObj.Addr1 = adr
  adr1, _ := strconv.Atoi(ipSlice[1])
  ipObj.Addr2 = adr1
  adr2, _ := strconv.Atoi(ipSlice[2])
  ipObj.Addr3 = adr2
  adr3, _ := strconv.Atoi(ipSlice[3])
  ipObj.Addr4 = adr3
  return ipObj
}

// ErrorResp custom response
type ErrorResp struct {
  Status      string
  ErrorMsg    string
  StatusCode  int
}

// IPObjDel for deleting IP Address
type IPObjDel struct {
  Addresses     []string      `json:"addressArray,omitempty"`
  Org           string        `json:"organization_name,omitempty"`
  DelRRSChked   uint          `json:"isDeleterrsChecked"`
}
