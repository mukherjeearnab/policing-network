package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Definition of the Asset structure
type introduction struct {
	PreliminaryIssues         string `json:"PreliminaryIssues"`
	SummaryOfProsecutionsCase string `json:"SummaryOfProsecutionsCase"`
	SummaryOfDefendantsCase   string `json:"SummaryOfDefendantsCase"`
	IssuesToBeDetermined      string `json:"IssuesToBeDetermined"`
}

type applicableLaw struct {
	StatutoryLaws string `json:"StatutoryLaws"`
	CaseLaws      string `json:"CaseLaws"`
}

type sentence struct {
	CitizenID string `json:"CitizenID"`
	Statement string `json:"Statement"`
}

type deliberations struct {
	Guilt                              string     `json:"Guilt"`
	AggravatingMitigatingCircumstances string     `json:"AggravatingMitigatingCircumstances"`
	Sentence                           []sentence `json:"Sentence"`
}

type finalJudgement struct {
	Date          int           `json:"Date"`
	Introduction  introduction  `json:"Introduction"`
	Evidence      []string      `json:"Evidence"`
	ApplicableLaw applicableLaw `json:"ApplicableLaw"`
	Deliberations deliberations `json:"Deliberations"`
}

type conclusion struct {
	Evidence []string `json:"Evidence"`
	Content  string   `json:"Content"`
}

type hearing struct {
	Date       int        `json:"Date"`
	Conclusion conclusion `json:"Conclusion"`
}

// Definition of the Asset structure
type judgementReport struct {
	Type           string         `json:"Type"`
	ID             string         `json:"ID"`
	ChargeSheetID  string         `json:"ChargeSheetID"`
	Hearings       hearing        `json:"Hearings"`
	FinalJudgement finalJudgement `json:"FinalJudgement"`
	Complete       bool           `json:"Complete"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "createNewJudgementReport" {
		return cc.createNewJudgementReport(stub, params)
	} else if fcn == "readJudgementReport" {
		return cc.readJudgementReport(stub, params)
	} else if fcn == "addEvidence" {
		return cc.addEvidence(stub, params)
	} else if fcn == "addSentence" {
		return cc.addSentence(stub, params)
	} else if fcn == "setComplete" {
		return cc.setComplete(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to create new judgementReport (C of CRUD)
func (cc *Chaincode) createNewJudgementReport(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCourt(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	// Check if Params are non-empty
	for a := 0; a < 9; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ID := params[0]
	PreliminaryIssues := params[1]
	SummaryOfProsecutionsCase := params[2]
	SummaryOfDefendantsCase := params[3]
	IssuesToBeDetermined := params[4]
	var Evidence []string
	StatutoryLaws := params[5]
	CaseLaws := params[6]
	Guilt := params[7]
	AggravatingMitigatingCircumstances := params[8]
	var Sentence []sentence
	Complete := false

	Introduction := introduction{PreliminaryIssues, SummaryOfProsecutionsCase,
		SummaryOfDefendantsCase, IssuesToBeDetermined}
	ApplicableLaw := applicableLaw{StatutoryLaws, CaseLaws}
	Deliberations := deliberations{Guilt, AggravatingMitigatingCircumstances, Sentence}

	// Check if Asset exists with Key => ID
	judgementReportAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to check if JudgementReport exists!")
	} else if judgementReportAsBytes != nil {
		return shim.Error("JudgementReport Already Exists!")
	}

	// Generate Asset from params provided
	judgementReport := &judgementReport{"judgementReport",
		ID, Introduction, Evidence, ApplicableLaw,
		Deliberations, Complete}

	// Convert to JSON bytes
	judgementReportJSONasBytes, err := json.Marshal(judgementReport)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Asset with Key => ID
	err = stub.PutState(ID, judgementReportJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an judgementReport (R of CRUD)
func (cc *Chaincode) readJudgementReport(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Asset with Key => params[0]
	judgementReportAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if judgementReportAsBytes == nil {
		jsonResp := "{\"Error\":\"Asset does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(judgementReportAsBytes)
}

// Function to update an judgementReport's owner (U of CRUD)
func (cc *Chaincode) addEvidence(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCourt(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	for a := 0; a < 2; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ID := params[0]
	NewEvidence := params[1]

	// Get State of Asset with Key => ID
	judgementReportAsBytes, err := stub.GetState(ID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if judgementReportAsBytes == nil {
		jsonResp := "{\"Error\":\"Asset does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Create new Asset Variable
	judgementReportToUpdate := judgementReport{}
	err = json.Unmarshal(judgementReportAsBytes, &judgementReportToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if Judgement is Complete or NOT
	if judgementReportToUpdate.Complete {
		return shim.Error("Error: Judgement is Complete & Locked!")
	}

	// Update judgementReport.Evidence => NewEvidence
	judgementReportToUpdate.Evidence = append(judgementReportToUpdate.Evidence, NewEvidence)

	// Convert to Byte[]
	judgementReportJSONasBytes, err := json.Marshal(judgementReportToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put updated State of the Asset with Key => ID
	err = stub.PutState(ID, judgementReportJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Delete an judgementReport (D of CRUD)
func (cc *Chaincode) addSentence(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCourt(creatorOrg, creatorCertIssuer) {
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
	CitizenID := params[1]
	Statement := params[2]
	NewSentence := sentence{CitizenID, Statement}

	// Get State of Asset with Key => ID
	judgementReportAsBytes, err := stub.GetState(ID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if judgementReportAsBytes == nil {
		jsonResp := "{\"Error\":\"Asset does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Create new Asset Variable
	judgementReportToUpdate := judgementReport{}
	err = json.Unmarshal(judgementReportAsBytes, &judgementReportToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if Judgement is Complete or NOT
	if judgementReportToUpdate.Complete {
		return shim.Error("Error: Judgement is Complete & Locked!")
	}

	// Update judgementReport.Deliberations.Sentence => NewSentence
	judgementReportToUpdate.Deliberations.Sentence = append(judgementReportToUpdate.Deliberations.Sentence, NewSentence)

	// Convert to Byte[]
	judgementReportJSONasBytes, err := json.Marshal(judgementReportToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put updated State of the Asset with Key => ID
	err = stub.PutState(ID, judgementReportJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Add Judgement ID to The Citizen Profile with CitizenID
	args := util.ToChaincodeArgs("addVerdictRecord", CitizenID, ID)
	response := stub.InvokeChaincode("citizenprofile_cc", args, "mainchannel")
	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Func
func (cc *Chaincode) setComplete(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCourt(creatorOrg, creatorCertIssuer) {
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
	judgementReportAsBytes, err := stub.GetState(ID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if judgementReportAsBytes == nil {
		jsonResp := "{\"Error\":\"Asset does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Create new Asset Variable
	judgementReportToUpdate := judgementReport{}
	err = json.Unmarshal(judgementReportAsBytes, &judgementReportToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if Judgement is Complete or NOT
	if judgementReportToUpdate.Complete {
		return shim.Error("Error: Judgement is Complete & Locked!")
	}

	// Update judgementReport.Complete => true
	judgementReportToUpdate.Complete = true

	// Convert to Byte[]
	judgementReportJSONasBytes, err := json.Marshal(judgementReportToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put updated State of the Asset with Key => ID
	err = stub.PutState(ID, judgementReportJSONasBytes)
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

// Authenticate => Court
func authenticateCourt(mspID string, certCN string) bool {
	return (mspID == "CourtMSP") && (certCN == "ca.court.example.com")
}
