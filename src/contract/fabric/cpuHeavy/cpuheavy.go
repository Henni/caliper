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

type CpuHeavyChaincode struct {

}

func (t *CpuHeavyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *CpuHeavyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	for i := 0; i < 10000000000; i++ {
	}
	return shim.Success(nil)
}

func (t *CpuHeavyChaincode) Query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func  main()  {
	err := shim.Start(new(CpuHeavyChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %v \n", err)
	}
}
