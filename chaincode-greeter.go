/*
   Mayur learns Chaincode
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Simple Chaincode implementation
type SimpleChaincode struct {
}

// Main - boilerplate code for entry point
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Initialize World State
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 0 {				//Ensure proper usage of 'init' without any arguments
		return nil, errors.New("Incorrect number of arguments for Init invocation - expecting none.")
	}
	err := stub.PutState("greetee", []byte("mysterious person!") )
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Invoke() is running " + function)

	// Handle different functions
	if function == "init" {				//Used as reset
		t.Init(stub, "init", args)
	} else if function == "change" {		//Mandatory first argument represents name of new greetee
		if len(args) != 1 {			//Ensure that the one and only one argument has been passed in
			return nil, errors.New("Incorrect number of arguments for Change invocation - expecting 1.")
		}

		err := stub.PutState("greetee", []byte(args[0])) 	//Write the new greetee name into the chaincode state
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	fmt.Println("Invoke() did not find function: " + function)					//Log error message
	return nil, errors.New("Invoke() called with unknown function name: " + function)
}

// Query is our entry point for read operations
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query() is running function \"" + function + "\"")

	if len(args) != 0 {						//Ensure proper usage of Query functions without any arguments
		return nil, errors.New("Incorrect number of arguments to Query() - expecting none.")
	}

	valAsBytes, err := stub.GetState("greetee")	//Read most current name from World State
	if err != nil {							//Handle error if any
		return nil, errors.New("{\"Error\":\"Failed to get state for key 'greetee'. Details - " + err.Error() + "\"}")
	}

	// Handle different functions
	if function == "read" {						
		fmt.Println(string(valAsBytes))				//Display name currently used in greeting
		return nil, nil;
	} else if function == "greet" {					//Read name from World State
		fmt.Println("Hello " + string(valAsBytes) + "!")	//Display greeting
		return nil, nil;
	}

	fmt.Println("Query() did not find function name: " + function)			//Log error
	return nil, errors.New("Query() received unknown function: " + function)
}
