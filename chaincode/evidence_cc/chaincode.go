package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Definition of the Evidence structure
type evidence struct {
	Type            string `json:"Type"`
	ID              string `json:"ID"`
	MimeType        string `json:"MimeType"`
	Extention       string `json:"Extention"`
	Description     string `json:"Description"`
	DateTime        int    `json:"DateTime"`
	InvestigationID string `json:"InvestigationID"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "addEvidence" {
		return cc.addEvidence(stub, params)
	} else if fcn == "readEvidence" {
		return cc.readEvidence(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to add new Evidence
func (cc *Chaincode) addEvidence(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePoFoCi(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	for a := 0; a < 6; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ID := params[0]
	MimeType := params[1]
	Extention := params[2]
	Description := params[3]
	DateTime := params[4]
	InvestigationID := params[5]
	DateTimeI, err := strconv.Atoi(DateTime)
	if err != nil {
		return shim.Error("Error: Invalid DateTime!")
	}

	// Check if Evidence exists with Key => params[0]
	evidenceAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to check if Evidence exists!")
	} else if evidenceAsBytes != nil {
		return shim.Error("Evidence Already Exists!")
	}

	// Generate Evidence from params provided
	evidence := &evidence{"evidence",
		ID, MimeType, Extention, Description, DateTimeI, InvestigationID}

	// Get JSON bytes of Evidence struct
	evidenceJSONasBytes, err := json.Marshal(evidence)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Evidence with Key => params[0]
	err = stub.PutState(ID, evidenceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Add Evidence ID to The Investigation with InvestigationID
	args := util.ToChaincodeArgs("addEvidence", InvestigationID, ID)
	response := stub.InvokeChaincode("investigation_cc", args, "mainchannel")
	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an Evidence
func (cc *Chaincode) readEvidence(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Evidence with Key => params[0]
	evidenceAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if evidenceAsBytes == nil {
		jsonResp := "{\"Error\":\"Evidence does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(evidenceAsBytes)
}

// ---------------------------------------------
// Helper Functions
// ---------------------------------------------

// Authentication
// ++++++++++++++

// Get Tx Creator Info
func getTxCreatorInfo(stub shim.ChaincodeStubInterface) (string, string, error) {
	var mspid string
	var err error
	var cert *x509.Certificate
	mspid, err = cid.GetMSPID(stub)

	if err != nil {
		fmt.Printf("Error getting MSP identity: %sn", err.Error())
		return "", "", err
	}

	cert, err = cid.GetX509Certificate(stub)
	if err != nil {
		fmt.Printf("Error getting client certificate: %sn", err.Error())
		return "", "", err
	}

	return mspid, cert.Issuer.CommonName, nil
}

// Authenticate => Police / Forensics / Citizen
func authenticatePoFoCi(mspID string, certCN string) bool {
	return (mspID == "PoliceMSP") && (certCN == "ca.police.example.com") || (mspID == "ForensicsMSP") && (certCN == "ca.forensics.example.com") || (mspID == "CitizenMSP") && (certCN == "ca.citizen.example.com")
}
