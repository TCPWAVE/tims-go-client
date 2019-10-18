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
In order to use this library in some project, use the below command.
  **go get "github/TCPWAVE/tims-go-client/twclient"**
