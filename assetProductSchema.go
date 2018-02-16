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

// ProductSchemaClass acts as the class of all productschemas
var ProductSchemaClass = iot.AssetClass{
	Name:        "productschema",
	Prefix:      "PRODSCH",
	AssetIDPath: "productschema.productschemaId",
}

var productschemaIDs []string

func newProductSchema() iot.Asset {
	return iot.Asset{
		Class:        ProductSchemaClass,
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

var createAssetProductSchema iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductSchemaClass.CreateAsset(stub, args, "createAssetProductSchema", []iot.QPropNV{})
}

var replaceAssetProductSchema iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductSchemaClass.ReplaceAsset(stub, args, "replaceAssetProductSchema", []iot.QPropNV{})
}

var updateAssetProductSchema iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductSchemaClass.UpdateAsset(stub, args, "updateAssetProductSchema", []iot.QPropNV{})
}

var deleteAssetProductSchema iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductSchemaClass.DeleteAsset(stub, args)
}

var deleteAssetStateHistoryProductSchema iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductSchemaClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllAssetsProductSchema iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductSchemaClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromAssetProductSchema iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductSchemaClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromAssetProductSchema", []iot.QPropNV{})
}

var readAssetProductSchema iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductSchemaClass.ReadAsset(stub, args)
}

var readAllAssetsProductSchema iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Inside readAllAssetsProductSchema ", args[0])
	return ProductSchemaClass.ReadAllAssets(stub, args)
}

var readAssetStateHistoryProductSchema = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductSchemaClass.ReadAssetStateHistory(stub, args)
}

var productschemaUsedAlert iot.AlertName = "SUPPLYSHIP"
var productschemaUseRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, productschema *iot.Asset) error {
	schemaStat, found := iot.GetObjectAsString(productschema.State, "productschema.productschemaStatus")
	if found {
		if schemaStat == "inuse" {
			iot.RaiseAlert(productschema, productschemaUsedAlert)
		} else {
			iot.ClearAlert(productschema, productschemaUsedAlert)
		}
	}
	return nil
}

func init() {
	iot.AddRule("ProductSchema Used Alert", ProductSchemaClass, []iot.AlertName{productschemaUsedAlert}, productschemaUseRule)
	iot.AddRoute("createAssetProductSchema", "invoke", ProductSchemaClass, createAssetProductSchema)
	iot.AddRoute("replaceAssetProductSchema", "invoke", ProductSchemaClass, replaceAssetProductSchema)
	iot.AddRoute("updateAssetProductSchema", "invoke", ProductSchemaClass, updateAssetProductSchema)
	iot.AddRoute("deleteAssetProductSchema", "invoke", ProductSchemaClass, deleteAssetProductSchema)
	iot.AddRoute("deleteAssetStateHistoryProductSchema", "invoke", ProductSchemaClass, deleteAssetStateHistoryProductSchema)
	iot.AddRoute("deleteAllAssetsProductSchema", "invoke", ProductSchemaClass, deleteAllAssetsProductSchema)
	iot.AddRoute("deletePropertiesFromAssetProductSchema", "invoke", ProductSchemaClass, deletePropertiesFromAssetProductSchema)
	iot.AddRoute("readAssetProductSchema", "query", ProductSchemaClass, readAssetProductSchema)
	iot.AddRoute("readAssetStateHistoryProductSchema", "query", ProductSchemaClass, readAssetStateHistoryProductSchema)
	iot.AddRoute("readAllAssetsProductSchema", "query", ProductSchemaClass, readAllAssetsProductSchema)
}
