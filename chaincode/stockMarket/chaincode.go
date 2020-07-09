/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
//Package stockMarket define Stock market Objects
package stockMarket

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
)

// SmartContract provides functions for managing a stock
type SmartContract struct {
	contractapi.Contract
}

// stock describes basic details of what makes up a stock

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Stock
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create generalmeeting chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting generalmeeting chaincode: %s", err.Error())
	}
}
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// Get the args from the transaction proposal
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}
	// Set up any variables or assets here by calling stub.PutState()
	// We store the key and the value on the ledger
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The 'set'
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()
	var result string
	var err error
	if fn == "set" {
		result, err = set(stub, args)
	} else {
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	// Return the result as success payload
	return shim.Success([]byte(result))
}


// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
} /
/ Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}
	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0],

			err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}