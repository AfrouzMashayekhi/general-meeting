#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/msp/tlscacerts/tlsca.share.com-cert.pem
export PEER0_CUSTOMER_CA=${PWD}/organizations/peerOrganizations/customer.share.com/peers/peer0.customer.share.com/tls/ca.crt
export PEER0_REGULATOR_CA=${PWD}/organizations/peerOrganizations/regulator.share.com/peers/peer0.regulator.share.com/tls/ca.crt
export PEER0_SHAREDEALER_CA=${PWD}/organizations/peerOrganizations/sharedealer.share.com/peers/peer0.sharedealer.share.com/tls/ca.crt

# Set OrdererOrg.Admin globals
setOrdererGlobals() {
  export CORE_PEER_LOCALMSPID="OrdererMSP"
  export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/ordererOrganizations/share.com/orderers/orderer.share.com/msp/tlscacerts/tlsca.share.com-cert.pem
  export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/ordererOrganizations/share.com/users/Admin@share.com/msp
}

# Set environment variables for the peer org
setGlobals() {
  ORG=$1
  if [ $ORG -eq 1 ]; then
    export CORE_PEER_LOCALMSPID="CustomerMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_CUSTOMER_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/customer.share.com/users/Admin@customer.share.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
  elif [ $ORG -eq 2 ]; then
    export CORE_PEER_LOCALMSPID="RegulatorMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_REGULATOR_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/regulator.share.com/users/Admin@regulator.share.com/msp
    export CORE_PEER_ADDRESS=localhost:9051

  elif [ $ORG -eq 3 ]; then
    export CORE_PEER_LOCALMSPID="SharedealerMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_SHAREDEALER_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/sharedealer.share.com/users/Admin@sharedealer.share.com/msp
    export CORE_PEER_ADDRESS=localhost:11051
  else
    echo "================== ERROR !!! ORG Unknown =================="
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

# parsePeerConnectionParameters $@
# Helper function that takes the parameters from a chaincode operation
# (e.g. invoke, query, instantiate) and checks for an even number of
# peers and associated org, then sets $PEER_CONN_PARMS and $PEERS
parsePeerConnectionParameters() {
  # check for uneven number of peer and org parameters

  PEER_CONN_PARMS=""
  PEERS=""
  # while [ "$#" -gt 0 ]; do
  #   setGlobals $1
  #   PEER="peer0.org$1"
  #   PEERS="$PEERS $PEER"
  #   PEER_CONN_PARMS="$PEER_CONN_PARMS --peerAddresses $CORE_PEER_ADDRESS"
  #   if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "true" ]; then
  #     TLSINFO=$(eval echo "--tlsRootCertFiles \$PEER0_ORG$1_CA")
  #     PEER_CONN_PARMS="$PEER_CONN_PARMS $TLSINFO"
  #   fi
  #   # shift by two to get the next pair of peer/org parameters
  #   shift
  # done
  # remove leading space for output
  while [ "$#" -gt 0 ]; do
    setGlobals $1
    if [ $1 -eq 1 ]; then
      PEER="peer0.customer"
      PEERS="$PEERS $PEER"
      PEER_CONN_PARMS="$PEER_CONN_PARMS --peerAddresses $CORE_PEER_ADDRESS"
      if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "true" ]; then
        TLSINFO=$(eval echo "--tlsRootCertFiles \$PEER0_CUSTOMER_CA")
        PEER_CONN_PARMS="$PEER_CONN_PARMS $TLSINFO"
      fi
    elif [ $1 -eq 2 ]; then
      PEER="peer0.regulator"
      PEERS="$PEERS $PEER"
      PEER_CONN_PARMS="$PEER_CONN_PARMS --peerAddresses $CORE_PEER_ADDRESS"
      if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "true" ]; then
        TLSINFO=$(eval echo "--tlsRootCertFiles \$PEER0_REGULATOR_CA")
        PEER_CONN_PARMS="$PEER_CONN_PARMS $TLSINFO"
      fi
    elif [ $1 -eq 3 ]; then
      PEER="peer0.sharedealer"
      PEERS="$PEERS $PEER"
      PEER_CONN_PARMS="$PEER_CONN_PARMS --peerAddresses $CORE_PEER_ADDRESS"
      if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "true" ]; then
        TLSINFO=$(eval echo "--tlsRootCertFiles \$PEER0_SHAREDEALER_CA")
        PEER_CONN_PARMS="$PEER_CONN_PARMS $TLSINFO"
      fi
    else
      echo "================== ERROR !!! ORG Unknown =================="
    fi
    shift
  done
  PEERS="$(echo -e "$PEERS" | sed -e 's/^[[:space:]]*//')"
}

verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo
    exit 1
  fi
}
