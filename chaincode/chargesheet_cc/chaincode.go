package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Definition of the Accused Persons structure
type accusedPersons struct {
	CitizenID string `json:"CitizenID"`
	Status    string `json:"Status"`
}

// Definition of the Brief Report structure
type briefReport struct {
	Content   string   `json:"Content"`
	Documents []string `json:"Documents"`
}

// Definition of the Charged Persons structure
type chargedPersons struct {
	CitizenID     string   `json:"CitizenID"`
	SectionOfLaws []string `json:"SectionOfLaws"`
}

// Definition of the ChargeSheet structure
type chargeSheet struct {
	Type                  string           `json:"Type"`
	ID                    string           `json:"ID"`
	Name                  string           `json:"Name"`
	FIRIDs                []string         `json:"FIRIDs"`
	DateTime              int              `json:"DateTime"`
	SectionOfLaws         []string         `json:"SectionOfLaws"`
	InvestigatingOfficers []string         `json:"InvestigatingOfficers"`
	InvestigationIDs      []string         `json:"InvestigationIDs"`
	AccusedPersons        []accusedPersons `json:"AccusedPersons"`
	BriefReport           []briefReport    `json:"BriefReport"`
	ChargedPersons        []chargedPersons `json:"ChargedPersons"`
	DespatchDate          int              `json:"DespatchDate"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "createNewChargeSheet" {
		return cc.createNewChargeSheet(stub, params)
	} else if fcn == "readChargeSheet" {
		return cc.readChargeSheet(stub, params)
	} else if fcn == "addFIRIDs" {
		return cc.addFIRIDs(stub, params)
	} else if fcn == "addSectionOfLaw" {
		return cc.addSectionOfLaw(stub, params)
	} else if fcn == "addInvestigatingOfficer" {
		return cc.addInvestigatingOfficer(stub, params)
	} else if fcn == "addInvestigatingID" {
		return cc.addInvestigatingID(stub, params)
	} else if fcn == "addAccusedPerson" {
		return cc.addAccusedPerson(stub, params)
	} else if fcn == "addBriefReport" {
		return cc.addBriefReport(stub, params)
	} else if fcn == "addChargedPerson" {
		return cc.addChargedPerson(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to create new ChargeSheet
func (cc *Chaincode) createNewChargeSheet(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Check if Params are non-empty
	for a := 0; a < 4; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ID := params[0]
	Name := params[1]
	var FIRIDs []string
	DateTime := params[2]
	var SectionOfLaws []string
	var InvestigatingOfficers []string
	var InvestigationIDs []string
	var AccusedPersons []accusedPersons
	var BriefReport []briefReport
	var ChargedPersons []chargedPersons
	DespatchDate := params[3]
	DespatchDateI, err := strconv.Atoi(DespatchDate)
	if err != nil {
		return shim.Error("Error: Invalid DespatchDate!")
	}
	DateTimeI, err := strconv.Atoi(DateTime)
	if err != nil {
		return shim.Error("Error: Invalid DateTime!")
	}

	// Check if Asset exists with Key => ID
	chargeSheetAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to check if Asset exists!")
	} else if chargeSheetAsBytes != nil {
		return shim.Error("Asset Already Exists!")
	}

	// Generate ChargeSheet from params provided
	chargeSheet := &chargeSheet{"chargesheet",
		ID, Name, FIRIDs, DateTimeI,
		SectionOfLaws, InvestigatingOfficers, InvestigationIDs,
		AccusedPersons, BriefReport, ChargedPersons, DespatchDateI}

	chargeSheetJSONasBytes, err := json.Marshal(chargeSheet)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Asset with Key => ID
	err = stub.PutState(ID, chargeSheetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read a ChargeSheet
func (cc *Chaincode) readChargeSheet(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Asset with Key => params[0]
	chargeSheetAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if chargeSheetAsBytes == nil {
		jsonResp := "{\"Error\":\"Asset does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(chargeSheetAsBytes)
}

// Function to Add FIR's to the ChargeSheet
func (cc *Chaincode) addFIRIDs(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	// Check if Params are non-empty
	for a := 0; a < 2; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	NewFIR := params[1]

	// Check if ChargeSheet exists with Key => ID
	chargeSheetAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get ChargeSheet Details!")
	} else if chargeSheetAsBytes == nil {
		return shim.Error("Error: ChargeSheet Does NOT Exist!")
	}

	// Create Update struct var
	chargeSheetToUpdate := chargeSheet{}
	err = json.Unmarshal(chargeSheetAsBytes, &chargeSheetToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update ChargeSheet.FIRIDs to append => NewFIR
	chargeSheetToUpdate.FIRIDs = append(chargeSheetToUpdate.FIRIDs, NewFIR)

	// Convert to JSON bytes
	chargeSheetJSONasBytes, err := json.Marshal(chargeSheetToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated ChargeSheet with Key => ID
	err = stub.PutState(ID, chargeSheetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(chargeSheetJSONasBytes)
}

// Function to Add new Section of Law violated
func (cc *Chaincode) addSectionOfLaw(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	// Check if Params are non-empty
	for a := 0; a < 2; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	NewSectionOfLaw := params[1]

	// Check if ChargeSheet exists with Key => ID
	chargeSheetAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get ChargeSheet Details!")
	} else if chargeSheetAsBytes == nil {
		return shim.Error("Error: ChargeSheet Does NOT Exist!")
	}

	// Create Update struct var
	chargeSheetToUpdate := chargeSheet{}
	err = json.Unmarshal(chargeSheetAsBytes, &chargeSheetToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update ChargeSheet.SectionOfLaws to append => NewSectionOfLaw
	chargeSheetToUpdate.SectionOfLaws = append(chargeSheetToUpdate.SectionOfLaws, NewSectionOfLaw)

	// Convert to JSON bytes
	chargeSheetJSONasBytes, err := json.Marshal(chargeSheetToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated ChargeSheet with Key => ID
	err = stub.PutState(ID, chargeSheetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(chargeSheetJSONasBytes)
}

// Function to Add Officers, who Investigated
func (cc *Chaincode) addInvestigatingOfficer(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	// Check if Params are non-empty
	for a := 0; a < 2; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	NewOfficer := params[1]

	// Check if ChargeSheet exists with Key => ID
	chargeSheetAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get ChargeSheet Details!")
	} else if chargeSheetAsBytes == nil {
		return shim.Error("Error: ChargeSheet Does NOT Exist!")
	}

	// Create Update struct var
	chargeSheetToUpdate := chargeSheet{}
	err = json.Unmarshal(chargeSheetAsBytes, &chargeSheetToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update ChargeSheet.InvestigatingOfficers to append => NewOfficer
	chargeSheetToUpdate.InvestigatingOfficers = append(chargeSheetToUpdate.InvestigatingOfficers, NewOfficer)

	// Convert to JSON bytes
	chargeSheetJSONasBytes, err := json.Marshal(chargeSheetToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated ChargeSheet with Key => ID
	err = stub.PutState(ID, chargeSheetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(chargeSheetJSONasBytes)
}

// Function to Add ID of Investigation's conducted
func (cc *Chaincode) addInvestigatingID(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2!")
	}

	// Check if Params are non-empty
	for a := 0; a < 2; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	NewInvestigationID := params[1]

	// Check if ChargeSheet exists with Key => ID
	chargeSheetAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get ChargeSheet Details!")
	} else if chargeSheetAsBytes == nil {
		return shim.Error("Error: ChargeSheet Does NOT Exist!")
	}

	// Create Update struct var
	chargeSheetToUpdate := chargeSheet{}
	err = json.Unmarshal(chargeSheetAsBytes, &chargeSheetToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update ChargeSheet.InvestigationIDs to append => NewInvestigationID
	chargeSheetToUpdate.InvestigationIDs = append(chargeSheetToUpdate.InvestigationIDs, NewInvestigationID)

	// Convert to JSON bytes
	chargeSheetJSONasBytes, err := json.Marshal(chargeSheetToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated ChargeSheet with Key => ID
	err = stub.PutState(ID, chargeSheetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(chargeSheetJSONasBytes)
}

// Function to Add new Accused Person
func (cc *Chaincode) addAccusedPerson(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3!")
	}

	// Check if Params are non-empty
	for a := 0; a < 3; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	CitizenID := params[1]
	Status := params[2]
	NewAccused := accusedPersons{CitizenID, Status}

	// Check if ChargeSheet exists with Key => ID
	chargeSheetAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get ChargeSheet Details!")
	} else if chargeSheetAsBytes == nil {
		return shim.Error("Error: ChargeSheet Does NOT Exist!")
	}

	// Create Update struct var
	chargeSheetToUpdate := chargeSheet{}
	err = json.Unmarshal(chargeSheetAsBytes, &chargeSheetToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update ChargeSheet.AccusedPersons to append => NewAccused
	chargeSheetToUpdate.AccusedPersons = append(chargeSheetToUpdate.AccusedPersons, NewAccused)

	// Convert to JSON bytes
	chargeSheetJSONasBytes, err := json.Marshal(chargeSheetToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated ChargeSheet with Key => ID
	err = stub.PutState(ID, chargeSheetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(chargeSheetJSONasBytes)
}

// Function to Add new Section to Brief-Report
func (cc *Chaincode) addBriefReport(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2+!")
	}

	// Check if Params are non-empty
	for a := 0; a < len(params); a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	Content := params[1]
	var Documents []string
	for a := 2; a < len(params); a++ {
		Documents = append(Documents, params[a])
	}

	NewReport := briefReport{Content, Documents}

	// Check if ChargeSheet exists with Key => ID
	chargeSheetAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get ChargeSheet Details!")
	} else if chargeSheetAsBytes == nil {
		return shim.Error("Error: ChargeSheet Does NOT Exist!")
	}

	// Create Update struct var
	chargeSheetToUpdate := chargeSheet{}
	err = json.Unmarshal(chargeSheetAsBytes, &chargeSheetToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update ChargeSheet.BriefReport to append => NewReport
	chargeSheetToUpdate.BriefReport = append(chargeSheetToUpdate.BriefReport, NewReport)

	// Convert to JSON bytes
	chargeSheetJSONasBytes, err := json.Marshal(chargeSheetToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated ChargeSheet with Key => ID
	err = stub.PutState(ID, chargeSheetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(chargeSheetJSONasBytes)
}

// Function to Add new Charged Person
func (cc *Chaincode) addChargedPerson(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting (2+)!")
	}

	// Check if Params are non-empty
	for a := 0; a < len(params); a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	CitizenID := params[1]
	var SectionOfLaws []string
	for a := 2; a < len(params); a++ {
		SectionOfLaws = append(SectionOfLaws, params[a])
	}

	NewChargedPerson := chargedPersons{CitizenID, SectionOfLaws}

	// Check if ChargeSheet exists with Key => ID
	chargeSheetAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get ChargeSheet Details!")
	} else if chargeSheetAsBytes == nil {
		return shim.Error("Error: ChargeSheet Does NOT Exist!")
	}

	// Create Update struct var
	chargeSheetToUpdate := chargeSheet{}
	err = json.Unmarshal(chargeSheetAsBytes, &chargeSheetToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update ChargeSheet.ChargedPersons to append => NewChargedPerson
	chargeSheetToUpdate.ChargedPersons = append(chargeSheetToUpdate.ChargedPersons, NewChargedPerson)

	// Convert to JSON bytes
	chargeSheetJSONasBytes, err := json.Marshal(chargeSheetToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated ChargeSheet with Key => ID
	err = stub.PutState(ID, chargeSheetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(chargeSheetJSONasBytes)
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

// Authenticate => Police
func authenticatePolice(mspID string, certCN string) bool {
	return (mspID == "PoliceMSP") && (certCN == "ca.police.example.com")
}
