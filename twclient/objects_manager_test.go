package twclient

import(
  "log"
  "strings"
  . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Objects Manager Test", func(){

    const(
      host = "192.168.0.109"
      port = "7443"
      basicauth = false
      uname = "twcadm"
      pass = "zxcv1234"
      cert="/opt/tcpwave/certs/client.crt"
      key="/opt/tcpwave/certs/client.key"
      requestTimeout=100
      poolSize=10
    )

    conf := &HostConfig{Host:host, Port:port}
    objMan := NewObjectManager(*conf, key, cert, poolSize, requestTimeout)

    Describe("Network Operation", func(){
      It("Should create network", func(){
          networkObj := &Network{MaskLen: 24}
          networkObj.Name = "Test Network 2"
          networkObj.Description = "test network"
          networkObj.Organization = "Internal"
          networkObj.Addr1 = 192
          networkObj.Addr2 = 168
          networkObj.Addr3 = 2
          networkObj.Addr4 = 0
          networkObj.DMZVisible = "no"
          _,err := objMan.CreateNetwork(*networkObj)
          if err!=nil{
            if strings.Contains(err.Error(), "TIMS-2007") {
              Expect(err.Error()).To(Equal("TIMS-2007: Network address range clashes with an existing network"))
            }
          }else{
            Expect(err).To(Equal(nil))
          }
      })

      It("Should Fetch network", func(){
          ipAddr := "192.168.2.0"
          org := "Internal"
          rsp,err := objMan.GetNetwork(ipAddr, org)
          if err!=nil{
            log.Print(err.Error())
            Expect(err.Error()).To(Equal("500 Request failed."))
          }else{
            addr := strings.Split(rsp.Address, "/")[0]
            log.Print(addr)
            Expect(addr).To(Equal("192.168.2.0"))
          }
      })
    })

    Describe("Subnet Operation", func(){
      It("Should create Subnet", func(){
          subnetObj := &Subnet{MaskLen: 26}
          subnetObj.Name = "Test Subnet 1"
          subnetObj.Description = "test subnet"
          subnetObj.Organization = "Internal"
          subnetObj.Addr1 = 192
          subnetObj.Addr2 = 168
          subnetObj.Addr3 = 1
          subnetObj.Addr4 = 64
          subnetObj.RouterAddr = "192.168.1.65"
          subnetObj.NetworkAddr = "192.168.1.0"
          subnetObj.PrimaryDomain = "saurabh.tcpwave.com"
          _,err := objMan.CreateSubnet(*subnetObj)
          if err!=nil{
            if strings.Contains(err.Error(), "TIMS-2011") {
              Expect(err.Error()).To(Equal("TIMS-2011: Subnet address range clashes with an existing subnet"))
            }
          }else{
            Expect(err).To(Equal(nil))
          }
      })
      It("Should Fetch subnet", func(){
          ipAddr := "192.168.1.64"
          org := "Internal"
          rsp,err := objMan.GetSubnet(ipAddr, org)
          if err!=nil{
            log.Print(err.Error())
            Expect(err.Error()).To(Equal("500 Request failed."))
          }else{
            addr := strings.Split(rsp.SubnetAddress, "/")[0]
            log.Print(rsp.SubnetAddress)
            Expect(addr).To(Equal("192.168.1.64"))
          }
      })
    })

    Describe("Ip Address Operation", func(){
        It("Should fetch next free IP", func(){
          ipAddr := "192.168.1.64/26"
          org := "Internal"
          rsp,err := objMan.GetNextFreeIP(ipAddr, org)
          if err!=nil{
            log.Print(err.Error())
            Expect(err.Error()).To(Equal("500 Request failed."))
          }else{
            log.Print(rsp)
            Expect(rsp).To(Equal("192.168.1.67"))
          }
        })

        It("Should Create Ip Address", func(){
          ipAddr := "192.168.1.66"
          subnetAddr := "192.168.1.64"
          domain := "saurabh.tcpwave.com"
          org := "Internal"
          name := "k8s_pod_1"
          rsp,err := objMan.CreateIPAddress(ipAddr, "", subnetAddr, domain, org, name)
          if err!=nil{
            log.Print(err.Error())
            Expect(err.Error()).To(Equal("500 Internal Server Error"))
          }else{
            log.Print(rsp)
            Expect(rsp).To(Equal(""))
          }
        })
    })
})
