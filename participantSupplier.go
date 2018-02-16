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
)

// SupplierClass acts as the class of all supplies
var SupplierClass = iot.AssetClass{
	Name:        "supplier",
	Prefix:      "SUPR",
	AssetIDPath: "supplier.supplierId",
}

func newSupplier() iot.Asset {
	return iot.Asset{
		Class:        SupplierClass,
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

var createParticipantSupplier iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.CreateAsset(stub, args, "createParticipantSupplier", []iot.QPropNV{})
}

var replaceParticipantSupplier iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.ReplaceAsset(stub, args, "replaceParticipantSupplier", []iot.QPropNV{})
}

var updateParticipantSupplier iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.UpdateAsset(stub, args, "updateParticipantSupplier", []iot.QPropNV{})
}

var deleteParticipantSupplier iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.DeleteAsset(stub, args)
}

var deleteParticipantStateHistorySupplier iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllParticipantsSupplier iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromParticipantSupplier iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromParticipantSupplier", []iot.QPropNV{})
}

var readParticipantSupplier iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.ReadAsset(stub, args)
}

var readAllParticipantsSupplier iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.ReadAllAssets(stub, args)
}

var readParticipantStateHistorySupplier = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SupplierClass.ReadAssetStateHistory(stub, args)
}

var supplierShippedAlert iot.AlertName = "SUPPLYSHIP"
var supplierShipRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, supplier *iot.Asset) error {
	shipStat, found := iot.GetObjectAsString(supplier.State, "supplier.shipmentStatus")
	if found {
		if shipStat == "shipped" {
			iot.RaiseAlert(supplier, supplierShippedAlert)
		} else {
			iot.ClearAlert(supplier, supplierShippedAlert)
		}
	}
	return nil
}

func init() {
	iot.AddRule("Supplier Shipped Alert", SupplierClass, []iot.AlertName{supplierShippedAlert}, supplierShipRule)
	iot.AddRoute("createParticipantSupplier", "invoke", SupplierClass, createParticipantSupplier)
	iot.AddRoute("replaceParticipantSupplier", "invoke", SupplierClass, replaceParticipantSupplier)
	iot.AddRoute("updateParticipantSupplier", "invoke", SupplierClass, updateParticipantSupplier)
	iot.AddRoute("deleteParticipantSupplier", "invoke", SupplierClass, deleteParticipantSupplier)
	iot.AddRoute("deleteParticipantStateHistorySupplier", "invoke", SupplierClass, deleteParticipantStateHistorySupplier)
	iot.AddRoute("deleteAllParticipantsSupplier", "invoke", SupplierClass, deleteAllParticipantsSupplier)
	iot.AddRoute("deletePropertiesFromParticipantSupplier", "invoke", SupplierClass, deletePropertiesFromParticipantSupplier)
	iot.AddRoute("readParticipantSupplier", "query", SupplierClass, readParticipantSupplier)
	iot.AddRoute("readParticipantStateHistorySupplier", "query", SupplierClass, readParticipantStateHistorySupplier)
	iot.AddRoute("readAllParticipantsSupplier", "query", SupplierClass, readAllParticipantsSupplier)
}
