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

// SupplyClass acts as the class of all supplies
var SupplyClass = iot.AssetClass{
	Name:        "supply",
	Prefix:      "SUP",
	AssetIDPath: "supply.supplyId",
}

var supplyIDs []string

func newSupply() iot.Asset {
	return iot.Asset{
		Class:        SupplyClass,
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

var createAssetSupply iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplyClass.CreateAsset(stub, args, "createAssetSupply", []iot.QPropNV{})
}

var replaceAssetSupply iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplyClass.ReplaceAsset(stub, args, "replaceAssetSupply", []iot.QPropNV{})
}

var updateAssetSupply iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplyClass.UpdateAsset(stub, args, "updateAssetSupply", []iot.QPropNV{})
}

var deleteAssetSupply iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplyClass.DeleteAsset(stub, args)
}

var deleteAssetStateHistorySupply iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplyClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllAssetsSupply iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplyClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromAssetSupply iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplyClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromAssetSupply", []iot.QPropNV{})
}

var readAssetSupply iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplyClass.ReadAsset(stub, args)
}

var readAllAssetsSupply iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Inside readAllAssetsSupply ", args[0])
	return SupplyClass.ReadAllAssets(stub, args)
}

var readAssetStateHistorySupply = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplyClass.ReadAssetStateHistory(stub, args)
}

var supplyShippedAlert iot.AlertName = "SUPPLYSHIP"
var supplyShipRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, supply *iot.Asset) error {
	shipStat, found := iot.GetObjectAsString(supply.State, "supply.supplyStatus")
	if found {
		if shipStat == "shipped" {
			iot.RaiseAlert(supply, supplyShippedAlert)
		} else {
			iot.ClearAlert(supply, supplyShippedAlert)
		}
	}
	return nil
}

func init() {
	iot.AddRule("Supply Shipped Alert", SupplyClass, []iot.AlertName{supplyShippedAlert}, supplyShipRule)
	iot.AddRoute("createAssetSupply", "invoke", SupplyClass, createAssetSupply)
	iot.AddRoute("replaceAssetSupply", "invoke", SupplyClass, replaceAssetSupply)
	iot.AddRoute("updateAssetSupply", "invoke", SupplyClass, updateAssetSupply)
	iot.AddRoute("deleteAssetSupply", "invoke", SupplyClass, deleteAssetSupply)
	iot.AddRoute("deleteAssetStateHistorySupply", "invoke", SupplyClass, deleteAssetStateHistorySupply)
	iot.AddRoute("deleteAllAssetsSupply", "invoke", SupplyClass, deleteAllAssetsSupply)
	iot.AddRoute("deletePropertiesFromAssetSupply", "invoke", SupplyClass, deletePropertiesFromAssetSupply)
	iot.AddRoute("readAssetSupply", "query", SupplyClass, readAssetSupply)
	iot.AddRoute("readAssetStateHistorySupply", "query", SupplyClass, readAssetStateHistorySupply)
	iot.AddRoute("readAllAssetsSupply", "query", SupplyClass, readAllAssetsSupply)
}
