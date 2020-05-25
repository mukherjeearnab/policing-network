package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Definition of the Asset structure
type asset struct {
	ID    string `json:"objID"`
	Name  string `json:"objName"`
	Owner string `json:"objOwner"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "createAsset" {
		return cc.createAsset(stub, params)
	} else if fcn == "readAsset" {
		return cc.readAsset(stub, params)
	} else if fcn == "updateAsset" {
		return cc.updateAsset(stub, params)
	} else if fcn == "deleteAsset" {
		return cc.deleteAsset(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to create new asset (C of CRUD)
func (cc *Chaincode) createAsset(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(params[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(params[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	// Check if Asset exists with Key => params[0]
	assetAsBytes, err := stub.GetState(params[0])
	if err != nil {
		return shim.Error("Failed to check if Asset exists!")
	} else if assetAsBytes != nil {
		return shim.Error("Asset Already Exists!")
	}

	// Generate Asset from params provided
	asset := &asset{params[0], params[1], params[2]}
	assetJSONasBytes, err := json.Marshal(asset)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Asset with Key => params[0]
	err = stub.PutState(params[0], assetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an asset (R of CRUD)
func (cc *Chaincode) readAsset(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Asset with Key => params[0]
	assetAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if assetAsBytes == nil {
		jsonResp := "{\"Error\":\"Asset does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(assetAsBytes)
}

// Function to update an asset's owner (U of CRUD)
func (cc *Chaincode) updateAsset(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(params[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	// Get State of Asset with Key => params[0]
	assetAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if assetAsBytes == nil {
		jsonResp := "{\"Error\":\"Asset does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Create new Asset Variable
	assetToTransfer := asset{}
	err = json.Unmarshal(assetAsBytes, &assetToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update asset.Owner => params[1]
	assetToTransfer.Owner = params[1]

	// Convert to Byte[]
	assetJSONasBytes, err := json.Marshal(assetToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put updated State of the Asset with Key => params[0]
	err = stub.PutState(params[0], assetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Delete an asset (D of CRUD)
func (cc *Chaincode) deleteAsset(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Delete the State with Key => params[0]
	err := stub.DelState(params[0])
	if err != nil {
		return shim.Error("Failed to delete Asset: " + err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}
