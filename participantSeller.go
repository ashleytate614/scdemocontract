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

// SellerClass acts as the class of all supplies
var SellerClass = iot.AssetClass{
	Name:        "seller",
	Prefix:      "SLR",
	AssetIDPath: "seller.sellerId",
}

func newSeller() iot.Asset {
	return iot.Asset{
		Class:        SellerClass,
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

var createParticipantSeller iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//Ideally create a seller list. For now in maker, check all participant sellers

	return SellerClass.CreateAsset(stub, args, "createParticipantSeller", []iot.QPropNV{})
}

var replaceParticipantSeller iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SellerClass.ReplaceAsset(stub, args, "replaceParticipantSeller", []iot.QPropNV{})
}

var updateParticipantSeller iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SellerClass.UpdateAsset(stub, args, "updateParticipantSeller", []iot.QPropNV{})
}

var deleteParticipantSeller iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SellerClass.DeleteAsset(stub, args)
}

var deleteParticipantStateHistorySeller iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SellerClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllParticipantsSeller iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SellerClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromParticipantSeller iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SellerClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromParticipantSeller", []iot.QPropNV{})
}

var readParticipantSeller iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SellerClass.ReadAsset(stub, args)
}

var readAllParticipantsSeller iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SellerClass.ReadAllAssets(stub, args)
}

var readParticipantStateHistorySeller = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SellerClass.ReadAssetStateHistory(stub, args)
}

var sellerShippedAlert iot.AlertName = "SUPPLYSHIP"
var sellerShipRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, seller *iot.Asset) error {
	shipStat, found := iot.GetObjectAsString(seller.State, "seller.shipmentStatus")
	if found {
		if shipStat == "shipped" {
			iot.RaiseAlert(seller, sellerShippedAlert)
		} else {
			iot.ClearAlert(seller, sellerShippedAlert)
		}
	}
	return nil
}

func init() {
	iot.AddRule("Seller Shipped Alert", SellerClass, []iot.AlertName{sellerShippedAlert}, sellerShipRule)
	iot.AddRoute("createParticipantSeller", "invoke", SellerClass, createParticipantSeller)
	iot.AddRoute("replaceParticipantSeller", "invoke", SellerClass, replaceParticipantSeller)
	iot.AddRoute("updateParticipantSeller", "invoke", SellerClass, updateParticipantSeller)
	iot.AddRoute("deleteParticipantSeller", "invoke", SellerClass, deleteParticipantSeller)
	iot.AddRoute("deleteParticipantStateHistorySeller", "invoke", SellerClass, deleteParticipantStateHistorySeller)
	iot.AddRoute("deleteAllParticipantsSeller", "invoke", SellerClass, deleteAllParticipantsSeller)
	iot.AddRoute("deletePropertiesFromParticipantSeller", "invoke", SellerClass, deletePropertiesFromParticipantSeller)
	iot.AddRoute("readParticipantSeller", "query", SellerClass, readParticipantSeller)
	iot.AddRoute("readParticipantStateHistorySeller", "query", SellerClass, readParticipantStateHistorySeller)
	iot.AddRoute("readAllParticipantsSeller", "query", SellerClass, readAllParticipantsSeller)
}
