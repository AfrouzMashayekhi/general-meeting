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

ORG=Trader
ORG_NAME=trader
P0PORT=7051
P1PORT=8051
CAPORT=7054
PEERPEM=organizations/peerOrganizations/trader.share.com/tlsca/tlsca.trader.share.com-cert.pem
CAPEM=organizations/peerOrganizations/trader.share.com/ca/ca.trader.share.com-cert.pem

echo "$(json_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/trader.share.com/connection-trader.json
echo "$(yaml_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/trader.share.com/connection-trader.yaml

ORG=Regulator
ORG_NAME=regulator
P0PORT=9051
P1PORT=10051
CAPORT=8054
PEERPEM=organizations/peerOrganizations/regulator.share.com/tlsca/tlsca.regulator.share.com-cert.pem
CAPEM=organizations/peerOrganizations/regulator.share.com/ca/ca.regulator.share.com-cert.pem

echo "$(json_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/regulator.share.com/connection-regulator.json
echo "$(yaml_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/regulator.share.com/connection-regulator.yaml

ORG=Company
ORG_NAME=company
P0PORT=11051
P1PORT=12051
CAPORT=10054
PEERPEM=organizations/peerOrganizations/company.share.com/tlsca/tlsca.company.share.com-cert.pem
CAPEM=organizations/peerOrganizations/company.share.com/ca/ca.company.share.com-cert.pem

echo "$(json_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/company.share.com/connection-company.json
echo "$(yaml_ccp $ORG $ORG_NAME $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/company.share.com/connection-company.yaml
