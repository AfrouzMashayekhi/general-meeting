function createTrader {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/trader.share.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/trader.share.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-trader --tls.certfiles ${PWD}/organizations/fabric-ca/trader/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-trader.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-trader.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-trader.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-trader.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/trader.share.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-trader --id.name peer0 --id.secret peer0pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer"' --tls.certfiles ${PWD}/organizations/fabric-ca/trader/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-trader --id.name user1 --id.secret user1pw --id.type client --id.attrs '"hf.Registrar.Roles=client"' --tls.certfiles ${PWD}/organizations/fabric-ca/trader/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-trader --id.name traderadmin --id.secret traderadminpw --id.type admin --id.attrs '"hf.Registrar.Roles=admin"' --tls.certfiles ${PWD}/organizations/fabric-ca/trader/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/trader.share.com/peers
  mkdir -p organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-trader -M ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/msp --csr.hosts peer0.trader.share.com --tls.certfiles ${PWD}/organizations/fabric-ca/trader/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/trader.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-trader -M ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/tls --enrollment.profile tls --csr.hosts peer0.trader.share.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/trader/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/tls/server.key

  mkdir ${PWD}/organizations/peerOrganizations/trader.share.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/trader.share.com/msp/tlscacerts/ca.crt

  mkdir ${PWD}/organizations/peerOrganizations/trader.share.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/trader.share.com/tlsca/tlsca.trader.share.com-cert.pem

  mkdir ${PWD}/organizations/peerOrganizations/trader.share.com/ca
  cp ${PWD}/organizations/peerOrganizations/trader.share.com/peers/peer0.trader.share.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/trader.share.com/ca/ca.trader.share.com-cert.pem

  mkdir -p organizations/peerOrganizations/trader.share.com/users
  mkdir -p organizations/peerOrganizations/trader.share.com/users/User1@trader.share.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca-trader -M ${PWD}/organizations/peerOrganizations/trader.share.com/users/User1@trader.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/trader/tls-cert.pem
  set +x

  mkdir -p organizations/peerOrganizations/trader.share.com/users/Admin@trader.share.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://traderadmin:traderadminpw@localhost:7054 --caname ca-trader -M ${PWD}/organizations/peerOrganizations/trader.share.com/users/Admin@trader.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/trader/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/trader.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/trader.share.com/users/Admin@trader.share.com/msp/config.yaml

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


function createCompany {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/company.share.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/company.share.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:10054 --caname ca-company --tls.certfiles ${PWD}/organizations/fabric-ca/company/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-company.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-company.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-company.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-company.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/company.share.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-company --id.name peer0 --id.secret peer0pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer"' --tls.certfiles ${PWD}/organizations/fabric-ca/company/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-company --id.name user1 --id.secret user1pw --id.type client --id.attrs '"hf.Registrar.Roles=client"' --tls.certfiles ${PWD}/organizations/fabric-ca/company/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-company --id.name companyadmin --id.secret companyadminpw --id.type admin --id.attrs '"hf.Registrar.Roles=admin"' --tls.certfiles ${PWD}/organizations/fabric-ca/company/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/company.share.com/peers
  mkdir -p organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:10054 --caname ca-company -M ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/msp --csr.hosts peer0.company.share.com --tls.certfiles ${PWD}/organizations/fabric-ca/company/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/company.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:10054 --caname ca-company -M ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/tls --enrollment.profile tls --csr.hosts peer0.company.share.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/company/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/tls/server.key

  mkdir ${PWD}/organizations/peerOrganizations/company.share.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/company.share.com/msp/tlscacerts/ca.crt

  mkdir ${PWD}/organizations/peerOrganizations/company.share.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/company.share.com/tlsca/tlsca.company.share.com-cert.pem

  mkdir ${PWD}/organizations/peerOrganizations/company.share.com/ca
  cp ${PWD}/organizations/peerOrganizations/company.share.com/peers/peer0.company.share.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/company.share.com/ca/ca.company.share.com-cert.pem

  mkdir -p organizations/peerOrganizations/company.share.com/users
  mkdir -p organizations/peerOrganizations/company.share.com/users/User1@company.share.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:10054 --caname ca-company -M ${PWD}/organizations/peerOrganizations/company.share.com/users/User1@company.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/company/tls-cert.pem
  set +x

  mkdir -p organizations/peerOrganizations/company.share.com/users/Admin@company.share.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://companyadmin:companyadminpw@localhost:10054 --caname ca-company -M ${PWD}/organizations/peerOrganizations/company.share.com/users/Admin@company.share.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/company/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/company.share.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/company.share.com/users/Admin@company.share.com/msp/config.yaml

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
