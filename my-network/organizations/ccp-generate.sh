#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $6)
    local CP=$(one_line_pem $7)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${ORG_NAME}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${P1PORT}/$4/" \
        -e "s/\${CAPORT}/$5/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $6)
    local CP=$(one_line_pem $7)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${ORG_NAME}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${P1PORT}/$4/" \
        -e "s/\${CAPORT}/$5/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n        /g'
}

ORG=Customer
ORG_NAME=customer
P0PORT=7051
P1PORT=8051
CAPORT=7054
PEERPEM=organizations/peerOrganizations/customer.share.com/tlsca/tlsca.customer.share.com-cert.pem
CAPEM=organizations/peerOrganizations/customer.share.com/ca/ca.customer.share.com-cert.pem

echo "$(json_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/customer.share.com/connection-customer.json
echo "$(yaml_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/customer.share.com/connection-customer.yaml

ORG=Regulator
ORG_NAME=regulator
P0PORT=9051
P1PORT=10051
CAPORT=8054
PEERPEM=organizations/peerOrganizations/regulator.share.com/tlsca/tlsca.regulator.share.com-cert.pem
CAPEM=organizations/peerOrganizations/regulator.share.com/ca/ca.regulator.share.com-cert.pem

echo "$(json_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/regulator.share.com/connection-regulator.json
echo "$(yaml_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/regulator.share.com/connection-regulator.yaml

ORG=Sharedealer
ORG_NAME=sharedealer
P0PORT=11051
P1PORT=12051
CAPORT=10054
PEERPEM=organizations/peerOrganizations/sharedealer.share.com/tlsca/tlsca.sharedealer.share.com-cert.pem
CAPEM=organizations/peerOrganizations/sharedealer.share.com/ca/ca.sharedealer.share.com-cert.pem

echo "$(json_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/sharedealer.share.com/connection-sharedealer.json
echo "$(yaml_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/sharedealer.share.com/connection-sharedealer.yaml
