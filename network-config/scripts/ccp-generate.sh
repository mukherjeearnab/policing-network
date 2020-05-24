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
PEERPEM=../crypto-config/peerOrganizations/citizen.example.com/tlsca/tlsca.citizen.example.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/citizen.example.com/ca/ca.citizen.example.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-citizen.json

ORG=police
ORGMSP=Police
P0PORT=8051
CAPORT=8054
PEERPEM=../crypto-config/peerOrganizations/police.example.com/tlsca/tlsca.police.example.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/police.example.com/ca/ca.police.example.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-police.json

ORG=forensics
ORGMSP=Forensics
P0PORT=9051
CAPORT=9054
PEERPEM=../crypto-config/peerOrganizations/forensics.example.com/tlsca/tlsca.forensics.example.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/forensics.example.com/ca/ca.forensics.example.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-forensics.json

ORG=court
ORGMSP=Court
P0PORT=10051
CAPORT=10054
PEERPEM=../crypto-config/peerOrganizations/court.example.com/tlsca/tlsca.court.example.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/court.example.com/ca/ca.court.example.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-court.json

ORG=identityprovider
ORGMSP=IdentityProvider
P0PORT=11051
CAPORT=11054
PEERPEM=../crypto-config/peerOrganizations/identityprovider.example.com/tlsca/tlsca.identityprovider.example.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/identityprovider.example.com/ca/ca.identityprovider.example.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-identityprovider.json
