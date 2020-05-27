package main

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Definition of the offence structure
type offence struct {
	Nature      string `json:"Nature"`
	Particulars string `json:"Particulars"`
}

// Definition of the FIR structure
type fir struct {
	Type                 string  `json:"Type"`
	ID                   string  `json:"ID"`
	CitizenID            string  `json:"CitizenID"`
	PoliceStation        string  `json:"PoliceStation"`
	District             string  `json:"District"`
	PlaceOfOccurence     string  `json:"PlaceOfOccurence"`
	DateHour             int     `json:"DateHour"`
	Offence              offence `json:"Offence"`
	DescriptionOfAccused string  `json:"DescriptionOfAccused"`
	DetailsOfWitness     string  `json:"DetailsOfWitness"`
	Complaint            string  `json:"Complaint"`
	InvestigationID      string  `json:"InvestigationID"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "createNewFIR" {
		return cc.createNewFIR(stub, params)
	} else if fcn == "readFIR" {
		return cc.readFIR(stub, params)
	} else if fcn == "queryFIR" {
		return cc.queryFIR(stub, params)
	} else if fcn == "addInvestigationToFIR" {
		return cc.addInvestigationToFIR(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to create new FIR
func (cc *Chaincode) createNewFIR(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCitizen(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	for a := 0; a < 11; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ID := params[0]
	CitizenID := params[1]
	PoliceStation := params[2]
	District := params[3]
	PlaceOfOccurence := params[4]
	DateHour := params[5]
	Nature := params[6]
	Particulars := params[7]
	DescriptionOfAccused := params[8]
	DetailsOfWitness := params[9]
	Complaint := params[10]
	InvestigationID := ""
	DateHourI, err := strconv.Atoi(DateHour)
	if err != nil {
		return shim.Error("Error: Invalid DateHour!")
	}

	// Check if FIR exists with Key => ID
	FIRAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to check if FIR exists!")
	} else if FIRAsBytes != nil {
		return shim.Error("FIR Already Exists!")
	}

	// Generate Offence from Params
	offence := offence{Nature, Particulars}

	// Generate FIR from params provided
	FIR := &fir{"FIR",
		ID, CitizenID, PoliceStation, District,
		PlaceOfOccurence, DateHourI, offence,
		DescriptionOfAccused, DetailsOfWitness,
		Complaint, InvestigationID}

	FIRJSONasBytes, err := json.Marshal(FIR)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated FIR with Key => ID
	err = stub.PutState(ID, FIRJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an FIR
func (cc *Chaincode) readFIR(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of FIR with Key => params[0]
	FIRAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if FIRAsBytes == nil {
		jsonResp := "{\"Error\":\"FIR does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(FIRAsBytes)
}

// Function to query FIRs
func (cc *Chaincode) queryFIR(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	regex := "(?i:.*%s.*)"
	search := "{\"selector\": {\"$and\": [{\"CitizenID\": { \"$regex\": \"%s\" }},{\"PoliceStation\": { \"$regex\": \"%s\" }}]}}"

	if len(params[0]) <= 0 && len(params[1]) <= 0 {
		// 0 0
		search = fmt.Sprintf(search, ".*", ".*")
	} else if len(params[0]) <= 0 && len(params[1]) > 0 {
		// 0 1
		search = fmt.Sprintf(search, ".*", fmt.Sprintf(regex, params[1]))
	} else if len(params[0]) > 0 && len(params[1]) <= 0 {
		// 1 0
		search = fmt.Sprintf(search, fmt.Sprintf(regex, params[0]), ".*")
	} else {
		// 1 1
		search = fmt.Sprintf(search, fmt.Sprintf(regex, params[0]), fmt.Sprintf(regex, params[1]))
	}

	queryResults, err := getQueryResultForQueryString(stub, search)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

// Function to Add an InvestigationID to FIR
func (cc *Chaincode) addInvestigationToFIR(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
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
	InvestigationID := params[1]

	// Check if FIR exists with Key => ID
	FIRAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get FIR Details!")
	} else if FIRAsBytes == nil {
		return shim.Error("Error: FIR Does NOT Exist!")
	}

	// Create Update struct var
	FIRToUpdate := fir{}
	err = json.Unmarshal(FIRAsBytes, &FIRToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update FIR
	FIRToUpdate.InvestigationID = InvestigationID

	// Convert to JSON bytes
	FIRJSONasBytes, err := json.Marshal(FIRToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated FIR with Key => ID
	err = stub.PutState(ID, FIRJSONasBytes)
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

// Authenticate => Citizen
func authenticateCitizen(mspID string, certCN string) bool {
	return (mspID == "CitizenMSP") && (certCN == "ca.FIR.example.com")
}

func authenticatePolice(mspID string, certCN string) bool {
	return (mspID == "PoliceMSP") && (certCN == "ca.police.example.com")
}

// Query Helpers
// +++++++++++++

// Construct Query Response from Iterator
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	return &buffer, nil
}

// Get Query Result for Query String
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
