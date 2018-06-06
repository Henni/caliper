/**
* Copyright 2017 HUAWEI. All Rights Reserved.
*
* SPDX-License-Identifier: Apache-2.0
*
*/

package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const ERROR_WRONG_FORMAT = "{\"code\":301, \"reason\": \"command format is wrong\"}"
const ERROR_NOT_IMPLEMENTED = "{\"code\":302, \"reason\": \"method is not implemented\"}"


type KeyValueChaincode struct {

}

func (t *KeyValueChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}
	for key := range args[0] {
		err := shim.PutState(key, []byte(args[1]))
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
}

func (t *KeyValueChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "query" {
		err := shim.GetState(args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(money)
	}

	if function == "update" {
		err := shim.PutState(args[0], []byte(args[1]))
		if err != nil {
			return shim.Error(err.Error())
		}
		return (stub, args)
	}

	if function == "scan" {
		return shim.Error(ERROR_NOT_IMPLEMENTED)
	}

	if function == "insert" {
		return shim.Error(ERROR_NOT_IMPLEMENTED)
	}

	return shim.Error(ERROR_WRONG_FORMAT)
}

func  main()  {
	err := shim.Start(new(KeyValueChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %v \n", err)
	}

}
