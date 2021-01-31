#!/bin/bash

function one_line_pem() {
    echo "$(awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1)"
}

function json_ccp() {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
    -e "s/\${ORGMSP}/$2/" \
    -e "s/\${P0PORT}/$3/" \
    -e "s/\${CAPORT}/$4/" \
    -e "s#\${PEERPEM}#$PP#" \
    -e "s#\${CAPEM}#$CP#" \
    ../../connections/ccp-template.json
}

ORG=citizen
ORGMSP=Citizen
P0PORT=7051
CAPORT=7054
PEERPEM=../crypto-config/peerOrganizations/citizen.lean.com/tlsca/tlsca.citizen.lean.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/citizen.lean.com/ca/ca.citizen.lean.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-citizen.json

ORG=police
ORGMSP=Police
P0PORT=8051
CAPORT=8054
PEERPEM=../crypto-config/peerOrganizations/police.lean.com/tlsca/tlsca.police.lean.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/police.lean.com/ca/ca.police.lean.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-police.json
ORG=forensics
ORGMSP=Forensics
P0PORT=9051
CAPORT=9054
PEERPEM=../crypto-config/peerOrganizations/forensics.lean.com/tlsca/tlsca.forensics.lean.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/forensics.lean.com/ca/ca.forensics.lean.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-forensics.json

ORG=court
ORGMSP=Court
P0PORT=10051
CAPORT=10054
PEERPEM=../crypto-config/peerOrganizations/court.lean.com/tlsca/tlsca.court.lean.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/court.lean.com/ca/ca.court.lean.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-court.json

ORG=identityprovider
ORGMSP=IdentityProvider
P0PORT=11051
CAPORT=11054
PEERPEM=../crypto-config/peerOrganizations/identityprovider.lean.com/tlsca/tlsca.identityprovider.lean.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/identityprovider.lean.com/ca/ca.identityprovider.lean.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-identityprovider.json
