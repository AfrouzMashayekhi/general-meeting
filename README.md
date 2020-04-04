[//]: # (SPDX-License-Identifier: CC-BY-4.0)

## Hyperledger Fabric For General Meeting

Please visit the [installation instructions](https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html)
to ensure you have the correct prerequisites installed. Please use the
version of the documentation that matches the version of the software you
intend to use to ensure alignment.

## Prerequisties 

Make sure you have installed 
1. git
2. curl
3. docker
4. docker-compose
5. go
6. python
7. node-js

**Note:** for setting you working space
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
#working directory $GOPATH/src
cd $GOPATH/src
```
## Download Binaries and Docker Images

The installation instructions will utilize `prepare.sh` (available in the fabric repository)
script to download all of the requisite Hyperledger Fabric binaries and docker
images, and tag the images with the 'latest' tag. Optionally,
specify a version for fabric, fabric-ca and thirdparty images. If versions
are not passed, the latest available versions will be downloaded.

The script will also clone fabric-samples repository using the version tag that
is aligned with the Fabric version.

Note: you need to use roxy for docker hub and git installtion

```bash
#your working space should be GOPATH/src
curl -sSL https://gist.githubusercontent.com/afrouzMashaykhi/026ab4f4aa825915c2c9d30001da43d6/raw/7b8746e11f8c59cc25580f6776d1da1665068925/prepare.sh | bash -- 2.0.0 1.4.6 0.4.18

```
After successfuly installtion you should have list of images that have been installed<br/>
**Note** for running bash you should have proxy for Docker Hub within Iran access Internet.<br/>
```bash
#copy bins that generated in fabric-samples/bin to go/bin
cp -r $GOPATH/src/fabric-samples/bin  $GOPATH/bin

```

## Setup infrastructure of share organizantion

This project setup a cluster of 3 Organizations with Fabric-ca and an orderer with which stores data in couch db.
 - Share Dealer
 - Customers
 - Regulator

 **Working Directory** `$GOPATH/src/fabric-samples/my-network/` 

 ```bash

 ```


