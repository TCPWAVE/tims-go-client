# TIMS Rest Client written in GO language

## 1. Overview
TCPWave provides an extremely powerful Rest APIs to simplify any and every trivial task in IP Address management. This project aims to provide a client library in GO language. Currently this supports only limited functionalities which is required for Kubernetes and IPAM integration, and hopefully will be extended to support most of the APIs.
Supported Functionalities are:
  * Create Network
  * Fetch Network details
  * Create Subnet
  * Fetch Subnet details
  * Get Next Free IP in the subnet
  * Create IP Object in IPAM

## 2. Authentication
This client does not support **Basic Authorization** and uses a more secure **Certificates based Authorization**. One has to provide path to client certificate and key files to successfully create the client and use it. The certificates must be authorized to communicate with TIMS IPAM.
If a fresh cretificate is created, then these need to be imported into TIMS IPAM before use.

## 3. Usage
In order to import this library in your project, use the below command.
  **go get "github/TCPWAVE/tims-go-client/twclient"**

### Sample code snippet
    ```GO
    package main

    import(
        twc "github.com/TCPWAVE/tims-go-client/twclient"
    )

    func main(){
      // Create config object
      hostConfig = new(twc.HostConfig)
      config.Host="192.168.0.10"
      config.Port="80"

      CERT_FILE="/path/to/cert/file"
      KEY_FILE="/path/to/key/file"
      HTTP_CONNS_POOL=10
      HTTP_REQ_TIMEOUT=30

      // Create Object Manager Instance
      objMgr := twc.NewObjectManager(hostConfig, KEY_FILE, CERT_FILE, HTTP_CONNS_POOL, HTTP_REQ_TIMEOUT)

      // Once Object Managet is created this object can be used to access all supported
      // functionalities. Below I have shown for Fetching Network detail
      // Fetch Network details for organization = Tcpwave and Address = 192.168.0.0/16
      objMgr.GetNetwork("192.168.0.0/16", "Tcpwave")

    }

    ```
