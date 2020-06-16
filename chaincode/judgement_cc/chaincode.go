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
	Evidence string `json:"Evidence"`
	Content  string `json:"Content"`
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
	Hearings       []hearing      `json:"Hearings"`
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
	} else if fcn == "createHearing" {
		return cc.createHearing(stub, params)
	} else if fcn == "concludeHearing" {
		return cc.concludeHearing(stub, params)
	} else if fcn == "initFinalJudgement" {
		return cc.initFinalJudgement(stub, params)
	} else if fcn == "addEvidenceToFinalJudgement" {
		return cc.addEvidenceToFinalJudgement(stub, params)
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
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	// Check if Params are non-empty
	for a := 0; a < 3; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ID := params[0]
	ChargeSheetID := params[1]
	HearingDate := params[2]
	var Hearings []hearing
	var FinalJudgement finalJudgement
	Complete := false
	HearingDateI, err := strconv.Atoi(HearingDate)
	if err != nil {
		return shim.Error("Error: Invalid HearingDate!")
	}
	var Conclusion conclusion

	FirstHearing := hearing{HearingDateI, Conclusion}
	Hearings = append(Hearings, FirstHearing)

	// Check if Asset exists with Key => ID
	judgementReportAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to check if JudgementReport exists!")
	} else if judgementReportAsBytes != nil {
		return shim.Error("JudgementReport Already Exists!")
	}

	// Generate Asset from params provided
	judgementReport := &judgementReport{"judgementReport",
		ID, ChargeSheetID, Hearings, FinalJudgement, Complete}

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

// Function to Add new Hearing Session
func (cc *Chaincode) createHearing(stub shim.ChaincodeStubInterface, params []string) sc.Response {
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
	HearingDate := params[1]
	HearingDateI, err := strconv.Atoi(HearingDate)
	if err != nil {
		return shim.Error("Error: Invalid HearingDate!")
	}
	var Conclusion conclusion
	NewHearing := hearing{HearingDateI, Conclusion}

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

	if judgementReportToUpdate.Hearings[len(judgementReportToUpdate.Hearings)-1].Conclusion.Content == "" {
		return shim.Error("Error: Previous Hearing NOT Concluded!")
	}

	// Append judgementReport.Hearings => NewHearing
	judgementReportToUpdate.Hearings = append(judgementReportToUpdate.Hearings, NewHearing)

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

// Function to Add new Hearing Session
func (cc *Chaincode) concludeHearing(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCourt(creatorOrg, creatorCertIssuer) {
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
	Evidence := params[1]
	Content := params[2]

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

	//Update Latest Hearing (Conclude)
	judgementReportToUpdate.Hearings[len(judgementReportToUpdate.Hearings)-1].Conclusion.Evidence = Evidence
	judgementReportToUpdate.Hearings[len(judgementReportToUpdate.Hearings)-1].Conclusion.Content = Content

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

// Function to Init. Final Judgement
func (cc *Chaincode) initFinalJudgement(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCourt(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Check if Params are non-empty
	for a := 0; a < 10; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ID := params[0]
	Date := params[1]
	PreliminaryIssues := params[2]
	SummaryOfProsecutionsCase := params[3]
	SummaryOfDefendantsCase := params[4]
	IssuesToBeDetermined := params[5]
	var Evidence []string
	StatutoryLaws := params[6]
	CaseLaws := params[7]
	Guilt := params[8]
	AggravatingMitigatingCircumstances := params[9]
	var Sentence []sentence
	DateI, err := strconv.Atoi(Date)
	if err != nil {
		return shim.Error("Error: Invalid FinalJudgement Date!")
	}

	Introduction := introduction{PreliminaryIssues, SummaryOfProsecutionsCase,
		SummaryOfDefendantsCase, IssuesToBeDetermined}
	ApplicableLaw := applicableLaw{StatutoryLaws, CaseLaws}
	Deliberations := deliberations{Guilt, AggravatingMitigatingCircumstances, Sentence}

	FinalJudgement := finalJudgement{DateI,
		Introduction, Evidence,
		ApplicableLaw, Deliberations}

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
	judgementReportToUpdate.FinalJudgement = FinalJudgement

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

// Add Evidence to Final Judgement
func (cc *Chaincode) addEvidenceToFinalJudgement(stub shim.ChaincodeStubInterface, params []string) sc.Response {
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

	// Check if FinalJudgement is Initialized
	if judgementReportToUpdate.FinalJudgement.Date == 0 {
		return shim.Error("Error: Final Judgement NOT Initialized!")
	}

	// Update judgementReport.Evidence => NewEvidence
	judgementReportToUpdate.FinalJudgement.Evidence = append(judgementReportToUpdate.FinalJudgement.Evidence, NewEvidence)

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

// Function to Add Sentence to Citizen
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

	// Check if FinalJudgement is Initialized
	if judgementReportToUpdate.FinalJudgement.Date == 0 {
		return shim.Error("Error: Final Judgement NOT Initialized!")
	}

	// Update judgementReport.Deliberations.Sentence => NewSentence
	judgementReportToUpdate.FinalJudgement.Deliberations.Sentence = append(judgementReportToUpdate.FinalJudgement.Deliberations.Sentence, NewSentence)

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
