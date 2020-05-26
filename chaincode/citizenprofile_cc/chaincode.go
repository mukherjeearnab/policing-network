package main

import (
	"bytes"
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

// Definition of the Asset structure
type citizen struct {
	_type         string   `json:"_type"`
	ID            string   `json:"ID"`
	Name          string   `json:"Name"`
	Email         string   `json:"Email"`
	Phone         string   `json:"Phone"`
	DOB           int      `json:"DOB"`
	Gender        string   `json:"Gender"`
	BloodGroup    string   `json:"BloodGroup"`
	EyeColor      string   `json:"EyeColor"`
	Nationality   string   `json:"Nationality"`
	Address       string   `json:"Address"`
	FathersName   string   `json:"FathersName"`
	MothersName   string   `json:"MothersName"`
	Religion      string   `json:"Religion"`
	Occupation    string   `json:"Occupation"`
	Fingerprint   []string `json:"Fingerprint"`
	VerdictRecord []string `json:"VerdictRecord"`
}

// Init function.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke function.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "createNewCitizenProfile" {
		return cc.createNewCitizenProfile(stub, params)
	} else if fcn == "updateCitizenProfile" {
		return cc.updateCitizenProfile(stub, params)
	} else if fcn == "readCitizenProfile" {
		return cc.readCitizenProfile(stub, params)
	} else if fcn == "queryCitizenProfile" {
		return cc.queryCitizenProfile(stub, params)
	} else if fcn == "addVerdictRecord" {
		return cc.addVerdictRecord(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to Create a New Citizen Profile.
func (cc *Chaincode) createNewCitizenProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateIdentityProvider(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) < 24 {
		return shim.Error("Incorrect number of arguments. Expecting (14+10 = 24)!")
	}

	// Check if Params are non-empty
	for a := 0; a < 14; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	Name := params[1]
	Email := params[2]
	Phone := params[3]
	DOB := params[4]
	Gender := params[5]
	BloodGroup := params[6]
	EyeColor := params[7]
	Nationality := params[8]
	Address := params[9]
	FathersName := params[10]
	MothersName := params[11]
	Religion := params[12]
	Occupation := params[13]
	var Fingerprint []string
	var VerdictRecord []string

	for a := 14; a < 24; a++ {
		Fingerprint = append(Fingerprint, params[a])
	}
	DOBI, err := strconv.Atoi(DOB)
	if err != nil {
		return shim.Error("Error: Invalid DOB!")
	}

	// Check if Citizen exists with Key => ID
	citizenAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to check if Citizen exists!")
	} else if citizenAsBytes != nil {
		return shim.Error("Citizen Already Exists!")
	}

	// Generate Citizen from params provided
	citizen := &citizen{"citizenProfile",
		ID, Name, Email, Phone,
		DOBI, Gender, BloodGroup, EyeColor,
		Nationality, Address, FathersName,
		MothersName, Religion, Occupation,
		Fingerprint, VerdictRecord}

	// Convert to JSON bytes
	citizenJSONasBytes, err := json.Marshal(citizen)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Citizen with Key => ID
	err = stub.PutState(ID, citizenJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(citizenJSONasBytes)
}

// Function to Update Citizen Profile.
func (cc *Chaincode) updateCitizenProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateIdentityProvider(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) < 24 {
		return shim.Error("Incorrect number of arguments. Expecting (14+10 = 24)!")
	}

	// Check if Params are non-empty
	for a := 0; a < 14; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	Name := params[1]
	Email := params[2]
	Phone := params[3]
	DOB := params[4]
	Gender := params[5]
	BloodGroup := params[6]
	EyeColor := params[7]
	Nationality := params[8]
	Address := params[9]
	FathersName := params[10]
	MothersName := params[11]
	Religion := params[12]
	Occupation := params[13]
	var Fingerprint []string

	for a := 14; a < 24; a++ {
		Fingerprint = append(Fingerprint, params[a])
	}
	DOBI, err := strconv.Atoi(DOB)
	if err != nil {
		return shim.Error("Error: Invalid DOB!")
	}

	// Check if Citizen exists with Key => ID
	citizenAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get Citizen Details!")
	} else if citizenAsBytes == nil {
		return shim.Error("Error: Citizen Does NOT Exist!")
	}

	// Create Update struct var
	citizenToUpdate := citizen{}
	err = json.Unmarshal(citizenAsBytes, &citizenToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Citizen
	citizenToUpdate.Name = Name
	citizenToUpdate.Email = Email
	citizenToUpdate.Phone = Phone
	citizenToUpdate.DOB = DOBI
	citizenToUpdate.Gender = Gender
	citizenToUpdate.BloodGroup = BloodGroup
	citizenToUpdate.EyeColor = EyeColor
	citizenToUpdate.Nationality = Nationality
	citizenToUpdate.Address = Address
	citizenToUpdate.FathersName = FathersName
	citizenToUpdate.MothersName = MothersName
	citizenToUpdate.Religion = Religion
	citizenToUpdate.Occupation = Occupation
	citizenToUpdate.Fingerprint = Fingerprint

	// Convert to JSON bytes
	citizenJSONasBytes, err := json.Marshal(citizenToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Citizen with Key => ID
	err = stub.PutState(ID, citizenJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(citizenJSONasBytes)
}

// Function to Create a New Citizen Profile.
func (cc *Chaincode) readCitizenProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Asset with Key => params[0]
	citizenAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if citizenAsBytes == nil {
		jsonResp := "{\"Error\":\"Citizen does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(citizenAsBytes)
}

// Function to Create a New Citizen Profile.
func (cc *Chaincode) queryCitizenProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting (14+10 = 24)!")
	}

	regex := "(?i:.*%s.*)"
	search := "{\"selector\": {\"$and\": [{\"Name\": { \"$regex\": \"%s\" }},{\"Phone\": \"%s\"}, {\"Gender\": \"%s\"}]}}"

	if len(params[0]) <= 0 && len(params[1]) <= 0 && len(params[2]) <= 0 {
		// 0 0 0
		search = fmt.Sprintf(search, ".*", ".*", ".*")
	} else if len(params[0]) <= 0 && len(params[1]) <= 0 && len(params[2]) > 0 {
		// 0 0 1
		search = fmt.Sprintf(search, ".*", ".*", fmt.Sprintf(regex, params[2]))
	} else if len(params[0]) <= 0 && len(params[1]) > 0 && len(params[2]) <= 0 {
		// 0 1 0
		search = fmt.Sprintf(search, ".*", fmt.Sprintf(regex, params[1]), ".*")
	} else if len(params[0]) <= 0 && len(params[1]) > 0 && len(params[2]) > 0 {
		// 0 1 1
		search = fmt.Sprintf(search, ".*", fmt.Sprintf(regex, params[1]), fmt.Sprintf(regex, params[2]))
	} else if len(params[0]) > 0 && len(params[1]) <= 0 && len(params[2]) <= 0 {
		// 1 0 0
		search = fmt.Sprintf(search, fmt.Sprintf(regex, params[0]), ".*", ".*")
	} else if len(params[0]) > 0 && len(params[1]) <= 0 && len(params[2]) > 0 {
		// 1 0 1
		search = fmt.Sprintf(search, fmt.Sprintf(regex, params[0]), ".*", fmt.Sprintf(regex, params[2]))
	} else if len(params[0]) > 0 && len(params[1]) > 0 && len(params[2]) <= 0 {
		// 1 1 0
		search = fmt.Sprintf(search, fmt.Sprintf(regex, params[0]), fmt.Sprintf(regex, params[1]), ".*")
	} else {
		// 1 1 1
		search = fmt.Sprintf(search, fmt.Sprintf(regex, params[0]), fmt.Sprintf(regex, params[1]), fmt.Sprintf(regex, params[2]))
	}

	queryResults, err := getQueryResultForQueryString(stub, search)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

// Function to Create a New Citizen Profile.
func (cc *Chaincode) addVerdictRecord(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCourt(creatorOrg, creatorCertIssuer) {
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
	NewVerdict := params[1]

	// Check if Citizen exists with Key => ID
	citizenAsBytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get Citizen Details!")
	} else if citizenAsBytes == nil {
		return shim.Error("Error: Citizen Does NOT Exist!")
	}

	// Create Update struct var
	citizenToUpdate := citizen{}
	err = json.Unmarshal(citizenAsBytes, &citizenToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Citizen.VerdictRecord to append => NewVerdict
	citizenToUpdate.VerdictRecord = append(citizenToUpdate.VerdictRecord, NewVerdict)

	// Convert to JSON bytes
	citizenJSONasBytes, err := json.Marshal(citizenToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Citizen with Key => ID
	err = stub.PutState(ID, citizenJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(citizenJSONasBytes)
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
func authenticateIdentityProvider(mspID string, certCN string) bool {
	return (mspID == "IdentityProviderMSP") && (certCN == "ca.identityprovider.example.com")
}

func authenticateCourt(mspID string, certCN string) bool {
	return (mspID == "CourtMSP") && (certCN == "ca.court.example.com")
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
