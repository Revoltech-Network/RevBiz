package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// MarketplaceChaincode is a chaincode that implements a Marketplace and team work management system
type MarketplaceChaincode struct {
}

// User represents a user in the system
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Product represents a product in the Marketplace
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Owner       string  `json:"owner"`
}

// Order represents an order in the Marketplace
type Order struct {
	ID         string  `json:"id"`
	ProductID  string  `json:"productID"`
	Buyer      string  `json:"buyer"`
	Seller     string  `json:"seller"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`
}

// Project represents a team work project in the system
type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Leader      string `json:"leader"`
}

// Task represents a task in a team work project
type Task struct {
	ID          string `json:"id"`
	ProjectID   string `json:"projectID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
	Status      string `json:"status"`
}

// Init initializes the chaincode
func (c *MarketplaceChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is the entry point for chaincode invocations
func (c *MarketplaceChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "createUser" {
		return c.createUser(stub, args)
	} else if function == "getUser" {
		return c.getUser(stub, args)
	} else if function == "updateUser" {
		return c.updateUser(stub, args)
	} else if function == "deleteUser" {
		return c.deleteUser(stub, args)
	} else if function == "createProduct" {
		return c.createProduct(stub, args)
	} else if function == "getProduct" {
		return c.getProduct(stub, args)
	} else if function == "updateProduct" {
		return c.updateProduct(stub, args)
	} else if function == "deleteProduct" {
		return c.deleteProduct(stub, args)
	} else if function == "createOrder" {
		return c.createOrder(stub, args)
	} else if function == "getOrder" {
		return c.getOrder(stub, args)
	} else if function == "updateOrderStatus" {
		return c.updateOrderStatus(stub, args)
	} else if function == "getProject" {
		return c.getProject(stub, args)
	} else if function == "createTask" {
