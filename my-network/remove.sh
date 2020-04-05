 #!/bin/bash
 sudo   rm -rf organizations/fabric-ca/customer/msp organizations/fabric-ca/customer/tls-cert.pem organizations/fabric-ca/customer/ca-cert.pem organizations/fabric-ca/customer/IssuerPublicKey organizations/fabric-ca/customer/IssuerRevocationPublicKey organizations/fabric-ca/customer/fabric-ca-server.db
 sudo   rm -rf organizations/fabric-ca/regulator/msp organizations/fabric-ca/regulator/tls-cert.pem organizations/fabric-ca/regulator/ca-cert.pem organizations/fabric-ca/regulator/IssuerPublicKey organizations/fabric-ca/regulator/IssuerRevocationPublicKey organizations/fabric-ca/regulator/fabric-ca-server.db
 sudo   rm -rf organizations/fabric-ca/sharedealer/msp organizations/fabric-ca/sharedealer/tls-cert.pem organizations/fabric-ca/sharedealer/ca-cert.pem organizations/fabric-ca/sharedealer/IssuerPublicKey organizations/fabric-ca/sharedealer/IssuerRevocationPublicKey organizations/fabric-ca/sharedealer/fabric-ca-server.db
  sudo  rm -rf organizations/fabric-ca/ordererOrg/msp organizations/fabric-ca/ordererOrg/tls-cert.pem organizations/fabric-ca/ordererOrg/ca-cert.pem organizations/fabric-ca/ordererOrg/IssuerPublicKey organizations/fabric-ca/ordererOrg/IssuerRevocationPublicKey organizations/fabric-ca/ordererOrg/fabric-ca-server.db
    # remove channel and script artifacts
 sudo   rm -rf channel-artifacts log.txt fabcar.tar.gz fabcar

