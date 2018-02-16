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
	"encoding/json"
	"strconv"
	"strings"
	//"errors"
)

// ProductClass acts as the class of all products
var ProductClass = iot.AssetClass{
	Name:        "product",
	Prefix:      "PROD",
	AssetIDPath: "product.productId",
}

func newProduct() iot.Asset {
	return iot.Asset{
		Class:        ProductClass,
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
// Structure of Assets that comprise the product

var createAssetProduct iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var assetMapsTemplate interface{}
	var args1 []string
	var args2 []string
	var s = strings.Split(args[0],"\"madeOf\":")
	sFinalargsfirstpart:= s[0]+"\"madeOf\": {";
	var s1 =strings.Split(s[1],",\"makerId\":")
	sFinalargslastpart:= "},\"makerId\":"+ s1[1]
	
	//var assetIsMadeOf IsMadeOf
	
	err := json.Unmarshal([]byte(args[0]), &assetMapsTemplate)
	if err != nil {
		// Product data has syntax error
		log.Criticalf("Product data failed to unmarshal: %s\n", err)
		return nil,err
	}
	//fmt.Println("Before assetMapsTemplate")
	if assetMapsTemplate!=nil {
		assetMaps := assetMapsTemplate.(map[string]interface{})
			//fmt.Println("After assetMapsTemplate")
		assetProduct:= assetMaps["product"].(map[string]interface{})
		//fmt.Println("After assetMaps")
		if assetProduct["madeOf"]!=nil {
			assetMadeOf:= assetProduct["madeOf"].(map[string]interface{})
			//fmt.Println("After assetProduct")
			//fmt.Println(assetMadeOf)
			iCount:=0;
			for _, v := range assetMadeOf {
				assetMadeOfInstance:=v.(map[string]interface{})
				//fmt.Println("After assetMadeOfInstance")
				//fmt.Println("assetMadeOfInstance is : ", assetMadeOfInstance)
				//iMadeOfIter:=0
				//fmt.Println(k, v)
				sAssetBatchId:=""
				sAssetType:=""
				iSubAssetCount:=0
				sSubAssetIds:= ""
				sAssetDescription:=""
				for k1, v1 := range assetMadeOfInstance {
					//fmt.Println(k1, " is ", v1)
					if k1=="assetBatchId" {sAssetBatchId =v1.(string)}
					if k1=="assetType" {sAssetType =v1.(string)}
					if k1=="assetCount" { 
						fSupplyCount:= v1.(float64) 
						iSubAssetCount = int(fSupplyCount)
					}
				}
				
				///////////////////////////////////////
				// Update the supply record
				///////////////////////////////////////
				if sAssetType=="Supply"  {
					sQueryString:="{\"filter\":{\"match\":\"all\",\"select\":{\"0\":{\"qprop\":\"assetstate.supply.supplyBatchId\",\"value\":\""+sAssetBatchId+"\"},\"1\":{\"qprop\":\"assetstate.supply.supplyStatus\",\"value\":\"New\"}}}}"
					//fmt.Println("sQueryString is ", sQueryString)
					args1 = append(args1, sQueryString)
					//fmt.Println("args1 for fetching a supply record is ", args1)
					//fmt.Println("args1[0] ", args1[0])
					supplyRecords, err:=readAllAssetsSupply(stub, args1)
					//fmt.Println("Here!")
					if err != nil {
						return nil, err
					}
					//fmt.Println("successfully fetched the required supply records")
					//fmt.Println(string(supplyRecords))
					var supplyData interface{}
					json.Unmarshal(supplyRecords, &supplyData)
					supplyMap := supplyData.([]interface{})
					//fmt.Println("supplyMap ", supplyMap)
					for key, val := range supplyMap {
						//fmt.Println("int(key ", int(key), ", iSubAssetCount ", iSubAssetCount)
						if(int(key)<iSubAssetCount){
							supplyRecord:=val.(map[string]interface{})
							//fmt.Println("supplyRecord is ", supplyRecord)
							for key1, val1 := range supplyRecord {
								if key1=="assetstate" {
									assetState := val1.(map[string]interface{})
									//fmt.Println("assetState is : ", assetState)
										for _, val2 := range assetState {
											assetStateSupply := val2.(map[string]interface{})
											//fmt.Println("assetStateSupply ",assetStateSupply)
											for key3, val3 := range assetStateSupply {
												//fmt.Println("key3 is ", key3)
											if key3 =="supplyId" {
												//fmt.Println("supply id value is ", val3.(string), "sSubAssetIds is ", sSubAssetIds)
												if sSubAssetIds =="" {
													sSubAssetIds = val3.(string)
												} else {
													//fmt.Println("Inside else of sSubAssetIds")
													sSubAssetIds = sSubAssetIds+","+val3.(string)
												}
												sQueryString2:="{\"supply\":{\"industryType\":\"AUTO\",\"supplyId\":\""+val3.(string)+"\",\"supplyStatus\":\"InUse\"}}"
												args2 = append(args2, sQueryString2)
												//fmt.Println("args2 is ",args2)
												_,err:=updateAssetSupply(stub, args2)
												if err !=nil {
													//fmt.Println(err)
													return nil, err
												}
											}
											if sAssetDescription =="" && key3 =="supplyDescription" {
												sAssetDescription=val3.(string)
											}
											args2= []string{}
										}
									}

								}
							}
						} else {
							break;
						}
					}
					
					// Append the updated supply ids and description to the record
					//fmt.Println("sAssetDescription is ",sAssetDescription)
					//fmt.Println("Supply ids are ",sSubAssetIds)
					assetMadeOfInstanceString:= "\""+strconv.Itoa(iCount)+"\":{\"assetBatchId\":\""+sAssetBatchId+"\",\"assetType\":\""+sAssetType+"\", \"assetCount\":\""+strconv.Itoa(iSubAssetCount) +"\"";
					assetMadeOfInstanceString+= ", \"assetIds\":\""+sSubAssetIds+"\", \"assetDescription\":\""+sAssetDescription+"\"}"
					//fmt.Println("assetMadeOfInstanceString is "+assetMadeOfInstanceString)
					if iCount>0 {
						sFinalargsfirstpart = sFinalargsfirstpart+", "+assetMadeOfInstanceString
					} else {
						sFinalargsfirstpart = sFinalargsfirstpart+assetMadeOfInstanceString
					}
					iCount++
				}// end of one supply batch id
				if sAssetType=="Product"  {
					sQueryString:="{\"filter\":{\"match\":\"all\",\"select\":{\"0\":{\"qprop\":\"assetstate.product.productBatchId\",\"value\":\""+sAssetBatchId+"\"},\"1\":{\"qprop\":\"assetstate.product.productStatus\",\"value\":\"New\"}}}}"
					fmt.Println("sQueryString is ", sQueryString)
					args1 = append(args1, sQueryString)
					//fmt.Println("args1[0] ", args1[0])
					//fmt.Println("supply count ",iSubAssetCount)
					productRecords, err:=readAllAssetsProduct(stub, args1)
					fmt.Println("Here!")
					if err != nil {
						return nil, err
					}
					fmt.Println("successfully fetched the required product records")
					//fmt.Println(string(productRecords))
					var productData interface{}
					json.Unmarshal(productRecords, &productData)
					productMap := productData.([]interface{})
					for key, val := range productMap {
						fmt.Println("product count ",iSubAssetCount)
						if(int(key)<iSubAssetCount){
							productRecord:=val.(map[string]interface{})
							for key1, val1 := range productRecord {
								if key1=="assetstate" {
									assetState := val1.(map[string]interface{})
									fmt.Println("assetState is : ", assetState)
										for _, val2 := range assetState {
											assetStateProduct := val2.(map[string]interface{})
											for key3, val3 := range assetStateProduct {
											if key3 =="productId" {
												//fmt.Println("supply id value is ", val3.(string))
												sQueryString2:="{\"product\":{\"industryType\":\"AUTO\",\"productId\":\""+val3.(string)+"\",\"productStatus\":\"InUse\"}}"
												args2 = append(args2, sQueryString2)
												fmt.Println("args2 is ",args2)
												_,err:=updateAssetProduct(stub, args2)
												if err !=nil {
													fmt.Println(err)
													return nil, err
												}
												if sSubAssetIds =="" {
													sSubAssetIds = val3.(string)
												} else {
													sSubAssetIds = sSubAssetIds+","+val3.(string)
												}

											}
											if sAssetDescription =="" && key3 =="productDescription" {
												sAssetDescription=val3.(string)
											}
											args2= []string{}
										}
									}

								}
							}
						} else {
							break;
						}
					}
					//fmt.Println("sAssetDescription is ",sAssetDescription)
					//fmt.Println("Product ids are ",sSubAssetIds)
					assetMadeOfInstanceString:= "\""+strconv.Itoa(iCount)+"\":{\"assetBatchId\":\""+sAssetBatchId+"\",\"assetType\":\""+sAssetType+"\", \"assetCount\":\""+strconv.Itoa(iSubAssetCount) +"\"";
					assetMadeOfInstanceString+= ", \"assetIds\":\""+sSubAssetIds+"\", \"assetDescription\":\""+sAssetDescription+"\"}"
					//fmt.Println("assetMadeOfInstanceString is "+assetMadeOfInstanceString)
					if iCount>0 {
						sFinalargsfirstpart = sFinalargsfirstpart+", "+assetMadeOfInstanceString
					} else {
						sFinalargsfirstpart = sFinalargsfirstpart+assetMadeOfInstanceString
					}
					iCount++
				}// end of one product batch id
				
				// Add supply list and description
				args1= []string{}
			}
		}
	}
	//fmt.Println("Here again!")
	//fmt.Println("createAssetProduct command is ", args[0])
	/*s := strings.Split(args[0],"\"madeOf\":")
	fmt.Println("first part of split ", s[0])
	sFinalargsfirstpart:= s[0]+"\"madeOf\":";
	//fmt.Println("second part of split ", s[1])
	s1 :=strings.Split(s[1],",\"makerId\":")
	fmt.Println("first part of second split ", s1[0])
	//fmt.Println("second part of second split ", s1[1])
	sFinalargslastpart:= ",\"makerId\":"+ s1[1]
	fmt.Println("Final args is ", sFinalargsfirstpart )
	fmt.Println("Final args is ", sFinalargslastpart )*/
	sUpdArgs:=sFinalargsfirstpart+sFinalargslastpart
	//fmt.Println("sUpdArgs is ",sUpdArgs)
	//fmt.Println("args[0] is ",args[0])
	args[0] = sUpdArgs
	res, err:= ProductClass.CreateAsset(stub, args, "createAssetProduct", []iot.QPropNV{})
	//fmt.Println(iot.QPropNV{})
	
	if err !=nil {
		return res, err
	} 
	return res, nil
	
	//return nil, nil
}

var replaceAssetProduct iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductClass.ReplaceAsset(stub, args, "replaceAssetProduct", []iot.QPropNV{})
}

var updateAssetProduct iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("update product args is ", args[0])
	return ProductClass.UpdateAsset(stub, args, "updateAssetProduct", []iot.QPropNV{})
}

var deleteAssetProduct iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductClass.DeleteAsset(stub, args)
}

var deleteAssetStateHistoryProduct iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllAssetsProduct iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromAssetProduct iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromAssetProduct", []iot.QPropNV{})
}

var readAssetProduct iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductClass.ReadAsset(stub, args)
}

var readAllAssetsProduct iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductClass.ReadAllAssets(stub, args)
}

var readAssetStateHistoryProduct = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ProductClass.ReadAssetStateHistory(stub, args)
}

var productShippedAlert iot.AlertName = "PRODSHIP"
var productShipRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, product *iot.Asset) error {
	shipStat, found := iot.GetObjectAsString(product.State, "product.shipmentStatus")
	if found {
		if shipStat == "shipped" {
			iot.RaiseAlert(product, productShippedAlert)
		} else {
			iot.ClearAlert(product, productShippedAlert)
		}
	}
	return nil
}

func init() {
	iot.AddRule("Product Shipped Alert", ProductClass, []iot.AlertName{productShippedAlert}, productShipRule)
	iot.AddRoute("createAssetProduct", "invoke", ProductClass, createAssetProduct)
	iot.AddRoute("replaceAssetProduct", "invoke", ProductClass, replaceAssetProduct)
	iot.AddRoute("updateAssetProduct", "invoke", ProductClass, updateAssetProduct)
	iot.AddRoute("deleteAssetProduct", "invoke", ProductClass, deleteAssetProduct)
	iot.AddRoute("deleteAssetStateHistoryProduct", "invoke", ProductClass, deleteAssetStateHistoryProduct)
	iot.AddRoute("deleteAllAssetsProduct", "invoke", ProductClass, deleteAllAssetsProduct)
	iot.AddRoute("deletePropertiesFromAssetProduct", "invoke", ProductClass, deletePropertiesFromAssetProduct)
	iot.AddRoute("readAssetProduct", "query", ProductClass, readAssetProduct)
	iot.AddRoute("readAssetStateHistoryProduct", "query", ProductClass, readAssetStateHistoryProduct)
	iot.AddRoute("readAllAssetsProduct", "query", ProductClass, readAllAssetsProduct)
}
