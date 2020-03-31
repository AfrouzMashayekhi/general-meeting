function createCustomer {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/customer.share.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/customer.share.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-customer --tls.certfiles ${PWD}/organizations/fabric-ca/customer/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-customer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-customer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-customer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-customer.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/customer.share.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-customer --id.name peer0 --id.secret peer0pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer"' --tls.certfiles ${PWD}/organizations/fabric-ca/customer/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-customer --id.name user1 --id.secret user1pw --id.type client --id.attrs '"hf.Registrar.Roles=client"' --tls.certfiles ${PWD}/organizations/fabric-ca/customer/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-customer --id.name customeradmin --id.secret customeradminpw --id.type admin --id.attrs '"hf.Registrar.Roles=admin"' --tls.certfiles ${PWD}/organizations/fabric-ca/customer/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/customer.share.com/peers
  mkdir -p organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-customer -M ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/msp --csr.hosts peer0.customer.share.com --tls.certfiles ${PWD}/organizations/fabric-ca/customer/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/customer.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-customer -M ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls --enrollment.profile tls --csr.hosts peer0.customer.share.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/customer/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls/server.key

  mkdir ${PWD}/organizations/peerOrganizations/customer.share.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/customer.share.com/msp/tlscacerts/ca.crt

  mkdir ${PWD}/organizations/peerOrganizations/customer.share.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/customer.share.com/tlsca/tlsca.customer.share.com-cert.pem

  mkdir ${PWD}/organizations/peerOrganizations/customer.share.com/ca
  cp ${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/customer.share.com/ca/ca.customer.share.com-cert.pem

  mkdir -p organizations/peerOrganizations/customer.share.com/users
  mkdir -p organizations/peerOrganizations/customer.share.com/users/User1@customer.share.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca-customer -M ${PWD}/organizations/peerOrganizations/customer.share.com/users/User1@customer.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/customer/tls-cert.pem
  set +x

  mkdir -p organizations/peerOrganizations/customer.share.com/users/Admin@customer.share.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://customeradmin:customeradminpw@localhost:7054 --caname ca-customer -M ${PWD}/organizations/peerOrganizations/customer.share.com/users/Admin@customer.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/customer/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/customer.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/customer.share.com/users/Admin@customer.share.com/msp/config.yaml

}


function createRegulator {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/regulator.share.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/regulator.share.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:8054 --caname ca-regulator --tls.certfiles ${PWD}/organizations/fabric-ca/regulator/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-regulator.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-regulator.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-regulator.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-regulator.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/regulator.share.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-regulator --id.name peer0 --id.secret peer0pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer"' --tls.certfiles ${PWD}/organizations/fabric-ca/regulator/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-regulator --id.name user1 --id.secret user1pw --id.type client --id.attrs '"hf.Registrar.Roles=client"' --tls.certfiles ${PWD}/organizations/fabric-ca/regulator/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-regulator --id.name regulatoradmin --id.secret regulatoradminpw --id.type admin --id.attrs '"hf.Registrar.Roles=admin"' --tls.certfiles ${PWD}/organizations/fabric-ca/regulator/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/regulator.share.com/peers
  mkdir -p organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-regulator -M ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/msp --csr.hosts peer0.regulator.share.com --tls.certfiles ${PWD}/organizations/fabric-ca/regulator/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/regulator.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-regulator -M ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls --enrollment.profile tls --csr.hosts peer0.regulator.share.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/regulator/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls/server.key

  mkdir ${PWD}/organizations/peerOrganizations/regulator.share.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/regulator.share.com/msp/tlscacerts/ca.crt

  mkdir ${PWD}/organizations/peerOrganizations/regulator.share.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/regulator.share.com/tlsca/tlsca.regulator.share.com-cert.pem

  mkdir ${PWD}/organizations/peerOrganizations/regulator.share.com/ca
  cp ${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/regulator.share.com/ca/ca.regulator.share.com-cert.pem

  mkdir -p organizations/peerOrganizations/regulator.share.com/users
  mkdir -p organizations/peerOrganizations/regulator.share.com/users/User1@regulator.share.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:8054 --caname ca-regulator -M ${PWD}/organizations/peerOrganizations/regulator.share.com/users/User1@regulator.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/regulator/tls-cert.pem
  set +x

  mkdir -p organizations/peerOrganizations/regulator.share.com/users/Admin@regulator.share.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://regulatoradmin:regulatoradminpw@localhost:8054 --caname ca-regulator -M ${PWD}/organizations/peerOrganizations/regulator.share.com/users/Admin@regulator.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/regulator/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/regulator.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/regulator.share.com/users/Admin@regulator.share.com/msp/config.yaml

}


function createShareDealer {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/sharedealer.share.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/sharedealer.share.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:10054 --caname ca-sharedealer --tls.certfiles ${PWD}/organizations/fabric-ca/sharedealer/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-sharedealer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-sharedealer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-sharedealer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-sharedealer.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/sharedealer.share.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-sharedealer --id.name peer0 --id.secret peer0pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer"' --tls.certfiles ${PWD}/organizations/fabric-ca/sharedealer/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-sharedealer --id.name user1 --id.secret user1pw --id.type client --id.attrs '"hf.Registrar.Roles=client"' --tls.certfiles ${PWD}/organizations/fabric-ca/sharedealer/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-sharedealer --id.name sharedealeradmin --id.secret sharedealeradminpw --id.type admin --id.attrs '"hf.Registrar.Roles=admin"' --tls.certfiles ${PWD}/organizations/fabric-ca/sharedealer/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/sharedealer.share.com/peers
  mkdir -p organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:10054 --caname ca-sharedealer -M ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/msp --csr.hosts peer0.sharedealer.share.com --tls.certfiles ${PWD}/organizations/fabric-ca/sharedealer/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/sharedealer.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:10054 --caname ca-sharedealer -M ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls --enrollment.profile tls --csr.hosts peer0.sharedealer.share.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/sharedealer/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls/server.key

  mkdir ${PWD}/organizations/peerOrganizations/sharedealer.share.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/sharedealer.share.com/msp/tlscacerts/ca.crt

  mkdir ${PWD}/organizations/peerOrganizations/sharedealer.share.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/sharedealer.share.com/tlsca/tlsca.sharedealer.share.com-cert.pem

  mkdir ${PWD}/organizations/peerOrganizations/sharedealer.share.com/ca
  cp ${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/sharedealer.share.com/ca/ca.sharedealer.share.com-cert.pem

  mkdir -p organizations/peerOrganizations/sharedealer.share.com/users
  mkdir -p organizations/peerOrganizations/sharedealer.share.com/users/User1@sharedealer.share.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:10054 --caname ca-sharedealer -M ${PWD}/organizations/peerOrganizations/sharedealer.share.com/users/User1@sharedealer.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/sharedealer/tls-cert.pem
  set +x

  mkdir -p organizations/peerOrganizations/sharedealer.share.com/users/Admin@sharedealer.share.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://sharedealeradmin:sharedealeradminpw@localhost:10054 --caname ca-sharedealer -M ${PWD}/organizations/peerOrganizations/sharedealer.share.com/users/Admin@sharedealer.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/sharedealer/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/sharedealer.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/sharedealer.share.com/users/Admin@sharedealer.share.com/msp/config.yaml

}
function createOrderer {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/ordererOrganizations/share.com

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/share.com
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:9054 --caname ca-orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/ordererOrganizations/share.com/msp/config.yaml


  echo
	echo "Register orderer"
  echo
  set -x
	fabric-ca-client register --caname ca-orderer --id.name orderer --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
    set +x

  echo
  echo "Register the orderer admin"
  echo
  set -x
  fabric-ca-client register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --id.attrs '"hf.Registrar.Roles=admin"' --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

	mkdir -p organizations/ordererOrganizations/share.com/orderers
  mkdir -p organizations/ordererOrganizations/share.com/orderers/share.com

  mkdir -p organizations/ordererOrganizations/share.com/orderers/orderer.share.com

  echo
  echo "## Generate the orderer msp"
  echo
  set -x
	fabric-ca-client enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/msp --csr.hosts orderer.share.com --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/share.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/msp/config.yaml

  echo
  echo "## Generate the orderer-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/tls --enrollment.profile tls --csr.hosts orderer.share.com --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/tls/ca.crt
  cp ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/tls/signcerts/* ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/tls/server.crt
  cp ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/tls/keystore/* ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/tls/server.key

  mkdir ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/msp/tlscacerts/tlsca.share.com-cert.pem

  mkdir ${PWD}/organizations/ordererOrganizations/share.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/share.com/msp/tlscacerts/tlsca.share.com-cert.pem

  mkdir -p organizations/ordererOrganizations/share.com/users
  mkdir -p organizations/ordererOrganizations/share.com/users/Admin@share.com

  echo
  echo "## Generate the admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/share.com/users/Admin@share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/share.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/share.com/users/Admin@share.com/msp/config.yaml


}
