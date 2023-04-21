package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type AssetTransferChaincode struct {
}

func (c *AssetTransferChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (c *AssetTransferChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "transferAsset" {
		return c.transferAsset(stub, args)
	}

	return shim.Error("Invalid function name.")
}

func (c *AssetTransferChaincode) transferAsset(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	assetID := args[0]
	oldOwner := args[1]
	newOwner := args[2]

	// Get the asset state from the ledger
	assetBytes, err := stub.GetState(assetID)
	if err != nil {
		return shim.Error("Failed to get asset: " + err.Error())
	} else if assetBytes == nil {
		return shim.Error("Asset does not exist")
	}

	// Transfer ownership of the asset
	err = stub.PutState(assetID, []byte(newOwner))
	if err != nil {
		return shim.Error("Failed to transfer asset: " + err.Error())
	}

	// Log the asset transfer
	fmt.Printf("Asset %s transferred from %s to %s", assetID, oldOwner, newOwner)

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(AssetTransferChaincode))
	if err != nil {
		fmt.Printf("Error starting asset transfer chaincode: %s", err)
	}
}
