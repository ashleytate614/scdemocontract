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
import "fmt"

// MakerClass acts as the class of all supplies
var MakerClass = iot.AssetClass{
	Name:        "maker",
	Prefix:      "MKR",
	AssetIDPath: "maker.makerId",
}

func newMaker() iot.Asset {
	return iot.Asset{
		Class:        MakerClass,
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

var createParticipantMaker iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//Check 
	return MakerClass.CreateAsset(stub, args, "createParticipantMaker", []iot.QPropNV{})
}

var replaceParticipantMaker iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return MakerClass.ReplaceAsset(stub, args, "replaceParticipantMaker", []iot.QPropNV{})
}

var updateParticipantMaker iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return MakerClass.UpdateAsset(stub, args, "updateParticipantMaker", []iot.QPropNV{})
}

var deleteParticipantMaker iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return MakerClass.DeleteAsset(stub, args)
}

var deleteParticipantStateHistoryMaker iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return MakerClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllParticipantsMaker iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return MakerClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromParticipantMaker iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return MakerClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromParticipantMaker", []iot.QPropNV{})
}

var readParticipantMaker iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return MakerClass.ReadAsset(stub, args)
}

var readAllParticipantsMaker iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return MakerClass.ReadAllAssets(stub, args)
}

var readParticipantStateHistoryMaker = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return MakerClass.ReadAssetStateHistory(stub, args)
}

var makerInvalidSupplierAlert iot.AlertName = "INV-SUPP"
var makerInvalidSupplierRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, maker *iot.Asset) error {
	suplStat, found := iot.GetObjectAsString(maker.State, "maker.suppliers")
	if found {
		fmt.Println("Supplier details found")
		fmt.Println(string(suplStat))
	}
	return nil
}

var makerShippedAlert iot.AlertName = "PRODSHIP"
var makerShipRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, maker *iot.Asset) error {
	shipStat, found := iot.GetObjectAsString(maker.State, "maker.shipmentStatus")
	if found {
		if shipStat == "shipped" {
			iot.RaiseAlert(maker, makerShippedAlert)
		} else {
			iot.ClearAlert(maker, makerShippedAlert)
		}
	}
	return nil
}

func init() {
	iot.AddRule("Maker Shipped Alert", MakerClass, []iot.AlertName{makerShippedAlert}, makerShipRule)
	iot.AddRoute("createParticipantMaker", "invoke", MakerClass, createParticipantMaker)
	iot.AddRoute("replaceParticipantMaker", "invoke", MakerClass, replaceParticipantMaker)
	iot.AddRoute("updateParticipantMaker", "invoke", MakerClass, updateParticipantMaker)
	iot.AddRoute("deleteParticipantMaker", "invoke", MakerClass, deleteParticipantMaker)
	iot.AddRoute("deleteParticipantStateHistoryMaker", "invoke", MakerClass, deleteParticipantStateHistoryMaker)
	iot.AddRoute("deleteAllParticipantsMaker", "invoke", MakerClass, deleteAllParticipantsMaker)
	iot.AddRoute("deletePropertiesFromParticipantMaker", "invoke", MakerClass, deletePropertiesFromParticipantMaker)
	iot.AddRoute("readParticipantMaker", "query", MakerClass, readParticipantMaker)
	iot.AddRoute("readParticipantStateHistoryMaker", "query", MakerClass, readParticipantStateHistoryMaker)
	iot.AddRoute("readAllParticipantsMaker", "query", MakerClass, readAllParticipantsMaker)
}
