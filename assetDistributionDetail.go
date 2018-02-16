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

// DistributionDetailClass acts as the class of all distributionDetails
var DistributionDetailClass = iot.AssetClass{
	Name:        "distributionDetail",
	Prefix:      "DISTDET",
	AssetIDPath: "distributionDetail.distributionDetailId",
}

func newDistributionDetail() iot.Asset {
	return iot.Asset{
		Class:        DistributionDetailClass,
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

var createAssetDistributionDetail iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// Ideally check the supply ids against the supplies 'all assets' 
	fmt.Println(args)
	res, err:= DistributionDetailClass.CreateAsset(stub, args, "createAssetDistributionDetail", []iot.QPropNV{})
	if err !=nil {
		return res, err
	} 
	//else {
		//res, err:= SupplyClass.UpdateAsset(stub, args, "updateAssetSupply", []iot.QPropNV{})
		
	//}
	return res, nil
}

var replaceAssetDistributionDetail iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return DistributionDetailClass.ReplaceAsset(stub, args, "replaceAssetDistributionDetail", []iot.QPropNV{})
}

var updateAssetDistributionDetail iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return DistributionDetailClass.UpdateAsset(stub, args, "updateAssetDistributionDetail", []iot.QPropNV{})
}

var deleteAssetDistributionDetail iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return DistributionDetailClass.DeleteAsset(stub, args)
}

var deleteAssetStateHistoryDistributionDetail iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return DistributionDetailClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllAssetsDistributionDetail iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return DistributionDetailClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromAssetDistributionDetail iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return DistributionDetailClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromAssetDistributionDetail", []iot.QPropNV{})
}

var readAssetDistributionDetail iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return DistributionDetailClass.ReadAsset(stub, args)
}

var readAllAssetsDistributionDetail iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return DistributionDetailClass.ReadAllAssets(stub, args)
}

var readAssetStateHistoryDistributionDetail = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return DistributionDetailClass.ReadAssetStateHistory(stub, args)
}

var distributionDetailShippedAlert iot.AlertName = "PRODSHIP"
var distributionDetailShipRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, distributionDetail *iot.Asset) error {
	shipStat, found := iot.GetObjectAsString(distributionDetail.State, "distributionDetail.shipmentStatus")
	if found {
		if shipStat == "shipped" {
			iot.RaiseAlert(distributionDetail, distributionDetailShippedAlert)
		} else {
			iot.ClearAlert(distributionDetail, distributionDetailShippedAlert)
		}
	}
	return nil
}

func init() {
	iot.AddRule("DistributionDetail Shipped Alert", DistributionDetailClass, []iot.AlertName{distributionDetailShippedAlert}, distributionDetailShipRule)
	iot.AddRoute("createAssetDistributionDetail", "invoke", DistributionDetailClass, createAssetDistributionDetail)
	iot.AddRoute("replaceAssetDistributionDetail", "invoke", DistributionDetailClass, replaceAssetDistributionDetail)
	iot.AddRoute("updateAssetDistributionDetail", "invoke", DistributionDetailClass, updateAssetDistributionDetail)
	iot.AddRoute("deleteAssetDistributionDetail", "invoke", DistributionDetailClass, deleteAssetDistributionDetail)
	iot.AddRoute("deleteAssetStateHistoryDistributionDetail", "invoke", DistributionDetailClass, deleteAssetStateHistoryDistributionDetail)
	iot.AddRoute("deleteAllAssetsDistributionDetail", "invoke", DistributionDetailClass, deleteAllAssetsDistributionDetail)
	iot.AddRoute("deletePropertiesFromAssetDistributionDetail", "invoke", DistributionDetailClass, deletePropertiesFromAssetDistributionDetail)
	iot.AddRoute("readAssetDistributionDetail", "query", DistributionDetailClass, readAssetDistributionDetail)
	iot.AddRoute("readAssetStateHistoryDistributionDetail", "query", DistributionDetailClass, readAssetStateHistoryDistributionDetail)
	iot.AddRoute("readAllAssetsDistributionDetail", "query", DistributionDetailClass, readAllAssetsDistributionDetail)
}
