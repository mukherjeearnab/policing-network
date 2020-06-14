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

// Definition of the Investigation Reports structure
type report struct {
	DateTime int    `json:"DateTime"`
	Content  string `json:"Content"`
}

// Definition of the Investigation Arrests structure
type arrest struct {
	CItizenID string `json:"CitizenID"`
	Cause     string `json:"Cause"`
	Date      int    `json:"Date"`
	Mugshot   string `json:"Mugshot"`
}

// Definition of the Investigation structure
type investigation struct {
	Type       string   `json:"Type"`
	ID         string   `json:"ID"`
	FIRID      string   `json:"FIRID"`
	Officer    string   `json:"Officer"`
	Evidence   []string `json:"Evidence"`
	Reports    []report `json:"Reports"`
	AccusedIDs []string `json:"AccusedIDs"`
	Arrests    []arrest `json:"Arrests"`
	Complete   bool     `json:"Complete"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "newInvestigationFromFIR" {
		return cc.newInvestigationFromFIR(stub, params)
	} else if fcn == "readInvestigation" {
		return cc.readInvestigation(stub, params)
	} else if fcn == "updateInvestigation" {
		return cc.updateInvestigation(stub, params)
	} else if fcn == "addEvidence" {
		return cc.addEvidence(stub, params)
	} else if fcn == "addReport" {
		return cc.addReport(stub, params)
	} else if fcn == "addAccusedID" {
		return cc.addAccusedID(stub, params)
	} else if fcn == "addArrest" {
		return cc.addArrest(stub, params)
	} else if fcn == "setComplete" {
		return cc.setComplete(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to create new Investigation
func (cc *Chaincode) newInvestigationFromFIR(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Check if Params are non-empty
	for a := 0; a < 3; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ID := params[0]
	FIRID := params[1]
	Officer := params[2]
	var Evidence []string
	var Reports []report
	var AccusedIDs []string
	var Arrests []arrest
	Complete := false

	// Check if Investigation exists with Key => params[0]
	investigationAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to check if Asset exists!")
	} else if investigationAsBytes != nil {
		return shim.Error("Asset Already Exists!")
	}

	// Generate Asset from params provided
	investigation := &investigation{"investigation",
		ID, FIRID, Officer,
		Evidence, Reports, AccusedIDs, Arrests, Complete}

	// Convert to JSON bytes
	investigationJSONasBytes, err := json.Marshal(investigation)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Asset with Key => params[0]
	err = stub.PutState(ID, investigationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Add Evidence ID to The Investigation with InvestigationID
	args := util.ToChaincodeArgs("addInvestigationToFIR", FIRID, ID)
	response := stub.InvokeChaincode("fir_cc", args, "mainchannel")
	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	// Returned on successful execution of the function
	return shim.Success(investigationJSONasBytes)
}

// Function to read an Investigation
func (cc *Chaincode) readInvestigation(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Asset with Key => params[0]
	investigationAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if investigationAsBytes == nil {
		jsonResp := "{\"Error\":\"Investigation does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(investigationAsBytes)
}

// Function to update an investigation's owner (U of CRUD)
func (cc *Chaincode) updateInvestigation(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	for a := 0; a < 3; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ID := params[0]
	FIRID := params[1]
	Officer := params[2]

	// Get State of Asset with Key => ID
	investigationAsBytes, err := stub.GetState(ID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if investigationAsBytes == nil {
		jsonResp := "{\"Error\":\"Investigation does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Create new Asset Variable
	investigationToUpdate := investigation{}
	err = json.Unmarshal(investigationAsBytes, &investigationToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if Judgement is Complete or NOT
	if investigationToUpdate.Complete {
		return shim.Error("Error: Judgement is Complete & Locked!")
	}

	// Update Investigation
	investigationToUpdate.FIRID = FIRID
	investigationToUpdate.Officer = Officer

	// Convert to Byte[]
	investigationJSONasBytes, err := json.Marshal(investigationToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put updated State of the Asset with Key => ID
	err = stub.PutState(ID, investigationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(investigationJSONasBytes)
}

// Function to add new Evidence
func (cc *Chaincode) addEvidence(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePoFoCi(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting (14+10 = 24)!")
	}

	// Check if Params are non-empty
	for a := 0; a < 2; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	NewEvidence := params[1]

	// Check if Investigation exists with Key => ID
	investigationAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get Citizen Details!")
	} else if investigationAsBytes == nil {
		return shim.Error("Error: Citizen Does NOT Exist!")
	}

	// Create Update struct var
	investigationToUpdate := investigation{}
	err = json.Unmarshal(investigationAsBytes, &investigationToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if Judgement is Complete or NOT
	if investigationToUpdate.Complete {
		return shim.Error("Error: Judgement is Complete & Locked!")
	}

	// Update Invsedtigation.Evidence to append => NewEvidence
	investigationToUpdate.Evidence = append(investigationToUpdate.Evidence, NewEvidence)

	// Convert to JSON bytes
	investigationJSONasBytes, err := json.Marshal(investigationToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Citizen with Key => ID
	err = stub.PutState(ID, investigationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(investigationJSONasBytes)
}

// Function to add new Report
func (cc *Chaincode) addReport(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting (14+10 = 24)!")
	}

	// Check if Params are non-empty
	for a := 0; a < 3; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	DateTime := params[1]
	Content := params[2]
	DateTimeI, err := strconv.Atoi(DateTime)
	if err != nil {
		return shim.Error("Error: Invalid DateTime!")
	}

	// Create Report Struct var
	NewReport := report{DateTimeI, Content}

	// Check if Investigation exists with Key => ID
	investigationAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get Citizen Details!")
	} else if investigationAsBytes == nil {
		return shim.Error("Error: Citizen Does NOT Exist!")
	}

	// Create Update struct var
	investigationToUpdate := investigation{}
	err = json.Unmarshal(investigationAsBytes, &investigationToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if Judgement is Complete or NOT
	if investigationToUpdate.Complete {
		return shim.Error("Error: Judgement is Complete & Locked!")
	}

	// Update Invsedtigation.Evidence to append => NewEvidence
	investigationToUpdate.Reports = append(investigationToUpdate.Reports, NewReport)

	// Convert to JSON bytes
	investigationJSONasBytes, err := json.Marshal(investigationToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Citizen with Key => ID
	err = stub.PutState(ID, investigationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(investigationJSONasBytes)
}

// Function to add new AccusedID
func (cc *Chaincode) addAccusedID(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting (14+10 = 24)!")
	}

	// Check if Params are non-empty
	for a := 0; a < 2; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	NewAccusedID := params[1]

	// Check if Investigation exists with Key => ID
	investigationAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get Citizen Details!")
	} else if investigationAsBytes == nil {
		return shim.Error("Error: Citizen Does NOT Exist!")
	}

	// Create Update struct var
	investigationToUpdate := investigation{}
	err = json.Unmarshal(investigationAsBytes, &investigationToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if Judgement is Complete or NOT
	if investigationToUpdate.Complete {
		return shim.Error("Error: Judgement is Complete & Locked!")
	}

	// Update Invsedtigation.AccusedIDs to append => NewAccusedID
	investigationToUpdate.AccusedIDs = append(investigationToUpdate.AccusedIDs, NewAccusedID)

	// Convert to JSON bytes
	investigationJSONasBytes, err := json.Marshal(investigationToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of Updated Investigation with Key => ID
	err = stub.PutState(ID, investigationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(investigationJSONasBytes)
}

// Function to add new Arrest
func (cc *Chaincode) addArrest(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting (14+10 = 24)!")
	}

	// Check if Params are non-empty
	for a := 0; a < 5; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	CitizenID := params[1]
	Cause := params[2]
	Date := params[3]
	Mugshot := params[4]
	DateI, err := strconv.Atoi(Date)
	if err != nil {
		return shim.Error("Error: Invalid DateTime!")
	}

	// Create Report Struct var
	NewArrest := arrest{CitizenID, Cause, DateI, Mugshot}

	// Check if Investigation exists with Key => ID
	investigationAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get Citizen Details!")
	} else if investigationAsBytes == nil {
		return shim.Error("Error: Citizen Does NOT Exist!")
	}

	// Create Update struct var
	investigationToUpdate := investigation{}
	err = json.Unmarshal(investigationAsBytes, &investigationToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if Judgement is Complete or NOT
	if investigationToUpdate.Complete {
		return shim.Error("Error: Judgement is Complete & Locked!")
	}

	// Update Invsedtigation.Evidence to append => NewEvidence
	investigationToUpdate.Arrests = append(investigationToUpdate.Arrests, NewArrest)

	// Convert to JSON bytes
	investigationJSONasBytes, err := json.Marshal(investigationToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Citizen with Key => ID
	err = stub.PutState(ID, investigationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(investigationJSONasBytes)
}

// Function to set Complete Signal
func (cc *Chaincode) setComplete(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	ID := params[0]

	// Get State of Asset with Key => ID
	investigationAsBytes, err := stub.GetState(ID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if investigationAsBytes == nil {
		jsonResp := "{\"Error\":\"Asset does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Create new Asset Variable
	investigationToUpdate := investigation{}
	err = json.Unmarshal(investigationAsBytes, &investigationToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if Judgement is Complete or NOT
	if investigationToUpdate.Complete {
		return shim.Error("Error: Judgement is Complete & Locked!")
	}

	// Update investigation.Complete => true
	investigationToUpdate.Complete = true

	// Convert to Byte[]
	investigationJSONasBytes, err := json.Marshal(investigationToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put updated State of the Asset with Key => ID
	err = stub.PutState(ID, investigationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
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

// Authenticate => IdentityProvider
func authenticatePolice(mspID string, certCN string) bool {
	return (mspID == "PoliceMSP") && (certCN == "ca.police.example.com")
}

// Authenticate => Police / Forensics / Citizen
func authenticatePoFoCi(mspID string, certCN string) bool {
	return (mspID == "PoliceMSP") && (certCN == "ca.police.example.com") || (mspID == "ForensicsMSP") && (certCN == "ca.forensics.example.com") || (mspID == "CitizenMSP") && (certCN == "ca.citizen.example.com")
}
