package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// SupplyChainChaincode is a chaincode that implements a supply chain tracking system
type SupplyChainChaincode struct {
}

// Asset represents an asset in the supply chain
type Asset struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Owner    string  `json:"owner"`
}

// Transaction represents a transaction in the supply chain
type Transaction struct {
	ID        string  `json:"id"`
	AssetID   string  `json:"assetID"`
	FromOwner string  `json:"fromOwner"`
	ToOwner   string  `json:"toOwner"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

// Init initializes the chaincode
func (c *SupplyChainChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is the entry point for chaincode invocations
func (c *SupplyChainChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "createAsset" {
		return c.createAsset(stub, args)
	} else if function == "getAsset" {
		return c.getAsset(stub, args)
	} else if function == "updateAsset" {
		return c.updateAsset(stub, args)
	} else if function == "deleteAsset" {
		return c.deleteAsset(stub, args)
	} else if function == "transferAsset" {
		return c.transferAsset(stub, args)
	} else if function == "createTransaction" {
		return c.createTransaction(stub, args)
	} else if function == "getTransaction" {
		return c.getTransaction(stub, args)
	} else if function == "getAssetHistory" {
		return c.getAssetHistory(stub, args)
	} else if function == "getTransactionHistory" {
		return c.getTransactionHistory(stub, args)
	} else {
		return shim.Error("Invalid function name.")
	}
}

// createAsset creates a new asset in the supply chain
func (c *SupplyChainChaincode) createAsset(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	assetID := args[0]
	name := args[1]
	quantity, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid quantity. Expecting an integer")
	}
	price, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return shim.Error("Invalid price. Expecting a float")
	}
	owner := stub.GetCreator()

	asset := Asset{
		ID:       assetID,
		Name:     name,
		Quantity: quantity,
		Price:    price,
		Owner:    string(owner),
	}

	assetBytes, err := json.Marshal(asset)
	if err != nil {
		return
