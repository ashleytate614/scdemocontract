/*
Copyright (c) 2016 IBM Corporation and other Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

Contributors:
Kim Letkeman - Initial Contribution
*/

// v0.1 KL -- new IOT sample with Trade Lane properties and behaviors

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
	"fmt"
)

// OrderClass acts as the class of all orders
var OrderClass = iot.AssetClass{
	Name:        "order",
	Prefix:      "ORDR",
	AssetIDPath: "order.orderId",
}

func newOrder() iot.Asset {
	return iot.Asset{
		Class:        OrderClass,
		AssetKey:     "",
		State:        nil,
		EventIn:      nil,
		TXNID:        "",
		TXNTS:        nil,
		EventOut:     nil,
		AlertsActive: nil,
		Compliant:    true,
	}
}

var createAssetOrder iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// Ideally check the supply ids against the supplies 'all assets' 
	fmt.Println(args)
	res, err:= OrderClass.CreateAsset(stub, args, "createAssetOrder", []iot.QPropNV{})
	if err !=nil {
		return res, err
	} 
	//else {
		//res, err:= SupplyClass.UpdateAsset(stub, args, "updateAssetSupply", []iot.QPropNV{})
		
	//}
	return res, nil
}

var replaceAssetOrder iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return OrderClass.ReplaceAsset(stub, args, "replaceAssetOrder", []iot.QPropNV{})
}

var updateAssetOrder iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return OrderClass.UpdateAsset(stub, args, "updateAssetOrder", []iot.QPropNV{})
}

var deleteAssetOrder iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return OrderClass.DeleteAsset(stub, args)
}

var deleteAssetStateHistoryOrder iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return OrderClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllAssetsOrder iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return OrderClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromAssetOrder iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return OrderClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromAssetOrder", []iot.QPropNV{})
}

var readAssetOrder iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return OrderClass.ReadAsset(stub, args)
}

var readAllAssetsOrder iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return OrderClass.ReadAllAssets(stub, args)
}

var readAssetStateHistoryOrder = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return OrderClass.ReadAssetStateHistory(stub, args)
}

var orderShippedAlert iot.AlertName = "PRODSHIP"
var orderShipRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, order *iot.Asset) error {
	shipStat, found := iot.GetObjectAsString(order.State, "order.shipmentStatus")
	if found {
		if shipStat == "shipped" {
			iot.RaiseAlert(order, orderShippedAlert)
		} else {
			iot.ClearAlert(order, orderShippedAlert)
		}
	}
	return nil
}

func init() {
	iot.AddRule("Order Shipped Alert", OrderClass, []iot.AlertName{orderShippedAlert}, orderShipRule)
	iot.AddRoute("createAssetOrder", "invoke", OrderClass, createAssetOrder)
	iot.AddRoute("replaceAssetOrder", "invoke", OrderClass, replaceAssetOrder)
	iot.AddRoute("updateAssetOrder", "invoke", OrderClass, updateAssetOrder)
	iot.AddRoute("deleteAssetOrder", "invoke", OrderClass, deleteAssetOrder)
	iot.AddRoute("deleteAssetStateHistoryOrder", "invoke", OrderClass, deleteAssetStateHistoryOrder)
	iot.AddRoute("deleteAllAssetsOrder", "invoke", OrderClass, deleteAllAssetsOrder)
	iot.AddRoute("deletePropertiesFromAssetOrder", "invoke", OrderClass, deletePropertiesFromAssetOrder)
	iot.AddRoute("readAssetOrder", "query", OrderClass, readAssetOrder)
	iot.AddRoute("readAssetStateHistoryOrder", "query", OrderClass, readAssetStateHistoryOrder)
	iot.AddRoute("readAllAssetsOrder", "query", OrderClass, readAllAssetsOrder)
}
