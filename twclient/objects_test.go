package twclient

import(
  "log"
  "encoding/json"
  . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Objects Type Test", func(){
    Describe("Network Object", func(){
        It("should prepare json object correctly", func(){
            network := &Network{MaskLen: 16}
            network.Name = "Test network"
            network.Addr1=192
            network.Addr2=168
            network.Addr3=0
            network.Addr4=0
            netStr,_ := json.Marshal(network)
            log.Printf("Network json : %s",string(netStr))
            Expect(string(netStr)).To(Equal("{\"name\":\"Test network\",\"addr1\":192,\"addr2\":168,\"addr3\":0,\"addr4\":0,\"mask_length\":16}"))
        })
    })
})
