package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)

var samples = `

{
    "API": {
        "createAssetOrder": {
            "args": [
                {
                    "order": {
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "goodsOrServicesType": "Goods / Services",
                        "industryType": "Industry type - to start with Auto / Pharma / Food",
                        "orderContentType": "Raw Material/ Component/Subassembly / Finished order",
                        "orderDeliveryLocation": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        },
                        "orderDescription": "Description of the order",
                        "orderDistributor": {
                            "bolId": "Bill of Lading number",
                            "distId": "Participant Id of order distributor",
                            "distName": "Participant name of order distributor"
                        },
                        "orderFulfiller": {
                            "Id": "Participant Id of order fulfiller",
                            "Type": "Participant type of order fulfiller",
                            "actualFulfilmentDate": "Fulfilment date",
                            "batchIds": "batch ids used to fill order",
                            "batchSizeDelivered": 123.456,
                            "committedDeliveryDate": "Order due date",
                            "orderReceivedDate": "Date the order was received"
                        },
                        "orderFulfilmentLocation": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        },
                        "orderId": "An order's ID",
                        "orderIssuer": {
                            "batchSizeRequired": 123.456,
                            "oderIssuerId": "Participant Id of order issuer",
                            "oderIssuerType": "Participant type of order issuer",
                            "orderDueDate": "Order due date",
                            "orderIssueDate": "Year of Manufacturing"
                        },
                        "orderStatus": "The current status of the order",
                        "orderType": "Manufacturing / Distribution/MRO"
                    }
                }
            ],
            "function": "createAssetOrder"
        },
        "createAssetProduct": {
            "args": [
                {
                    "product": {
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "industryType": "Industry type - to start with Auto / Pharma / Food",
                        "madeOf": [
                            {
                                "assetBatchId": "Batch Id of supply / productModelId of product",
                                "assetCount": 123.456,
                                "assetDescription": "Description of the supply/supplies",
                                "assetIds": "Comma separated list of actual supplies",
                                "assetType": "Is this a product(subassembly) or a supply asset"
                            }
                        ],
                        "makerId": "A maker's ID",
                        "manufacturingDate": "Date of Manufacturing",
                        "orderId": "An order's ID",
                        "productBatchId": "Batch id to which the product belongs",
                        "productDescription": "Description of the finished product",
                        "productId": "A product's ID",
                        "productInternalId": "A product's Internal ID",
                        "productModel": "Product Model",
                        "productRecalls": [
                            {
                                "assetType": "Is this a product(subassembly) or a supply asset",
                                "newassetIds": "Asset Ids of supply / product in recall",
                                "originalassetIds": "Asset Ids of supply / product in original schema"
                            }
                        ],
                        "productSchemaId": "A product schema ID",
                        "productStatus": "The current status of the product",
                        "productType": "FP(FinishedProduct/Subassembly",
                        "productionLocation": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        },
                        "sellerId": "A seller's ID"
                    }
                }
            ],
            "function": "createAssetProduct"
        },
        "createAssetSupply": {
            "args": [
                {
                    "supply": {
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "industryType": "Industry type - to start with Auto / Pharma / Food",
                        "makerId": "A maker's ID",
                        "makerproductId": "A product's ID",
                        "orderId": "An order's ID",
                        "shippedDate": "Date of shipment to maker",
                        "supplierId": "A supplier's ID",
                        "supplyAggregationLevel": "Supply type - RawMaterial / Component / Operational",
                        "supplyAvailableDate": "Date of manufacture",
                        "supplyBatchId": "Supplier's production batch under which this supply falls",
                        "supplyCount": 123.456,
                        "supplyDescription": "Description about the supply",
                        "supplyId": "A supply's ID",
                        "supplyStatus": "The current status of the supply",
                        "supplyType": "eg.LightingCircuits, EngineSubasembly etc "
                    }
                }
            ],
            "function": "createAssetSupply"
        },
        "createParticipantMaker": {
            "args": [
                {
                    "maker": {
                        "industryType": "Industry, to start off with is auto, pharma or food",
                        "makerAddress": {
                            "city": "carpe noctem",
                            "country": "carpe noctem",
                            "name": "carpe noctem",
                            "phone": "carpe noctem",
                            "street": "carpe noctem",
                            "website": "carpe noctem",
                            "zipcode": "carpe noctem"
                        },
                        "makerId": "A maker's ID",
                        "sellers": [
                            {
                                "seller": {
                                    "sellerId": "A seller's ID"
                                }
                            }
                        ],
                        "suppliers": [
                            {
                                "supplier": {
                                    "supplierId": "A supplier's ID"
                                }
                            }
                        ]
                    }
                }
            ],
            "function": "createParticipantMaker"
        },
        "createParticipantSeller": {
            "args": [
                {
                    "seller": {
                        "industryType": "Industry, to start off with is auto, pharma or food",
                        "makers": [
                            {
                                "maker": {
                                    "makerId": "A maker's ID"
                                }
                            }
                        ],
                        "sellerAddress": {
                            "city": "carpe noctem",
                            "country": "carpe noctem",
                            "name": "carpe noctem",
                            "phone": "carpe noctem",
                            "street": "carpe noctem",
                            "website": "carpe noctem",
                            "zipcode": "carpe noctem"
                        },
                        "sellerId": "A seller's ID"
                    }
                }
            ],
            "function": "createParticipantSeller"
        },
        "createParticipantSupplier": {
            "args": [
                {
                    "supplier": {
                        "industryType": "Industry, to start off with is auto, pharma or food",
                        "makers": [
                            {
                                "maker": {
                                    "makerId": "A maker's ID"
                                }
                            }
                        ],
                        "supplierAddress": {
                            "city": "carpe noctem",
                            "country": "carpe noctem",
                            "name": "carpe noctem",
                            "phone": "carpe noctem",
                            "street": "carpe noctem",
                            "website": "carpe noctem",
                            "zipcode": "carpe noctem"
                        },
                        "supplierId": "A supplier's ID"
                    }
                }
            ],
            "function": "createParticipantSupplier"
        },
        "updateAssetOrder": {
            "args": [
                {
                    "order": {
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "goodsOrServicesType": "Goods / Services",
                        "industryType": "Industry type - to start with Auto / Pharma / Food",
                        "orderContentType": "Raw Material/ Component/Subassembly / Finished order",
                        "orderDeliveryLocation": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        },
                        "orderDescription": "Description of the order",
                        "orderDistributor": {
                            "bolId": "Bill of Lading number",
                            "distId": "Participant Id of order distributor",
                            "distName": "Participant name of order distributor"
                        },
                        "orderFulfiller": {
                            "Id": "Participant Id of order fulfiller",
                            "Type": "Participant type of order fulfiller",
                            "actualFulfilmentDate": "Fulfilment date",
                            "batchIds": "batch ids used to fill order",
                            "batchSizeDelivered": 123.456,
                            "committedDeliveryDate": "Order due date",
                            "orderReceivedDate": "Date the order was received"
                        },
                        "orderFulfilmentLocation": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        },
                        "orderId": "An order's ID",
                        "orderIssuer": {
                            "batchSizeRequired": 123.456,
                            "oderIssuerId": "Participant Id of order issuer",
                            "oderIssuerType": "Participant type of order issuer",
                            "orderDueDate": "Order due date",
                            "orderIssueDate": "Year of Manufacturing"
                        },
                        "orderStatus": "The current status of the order",
                        "orderType": "Manufacturing / Distribution/MRO"
                    }
                }
            ],
            "function": "updateAssetOrder"
        },
        "updateAssetProduct": {
            "args": [
                {
                    "product": {
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "industryType": "Industry type - to start with Auto / Pharma / Food",
                        "madeOf": [
                            {
                                "assetBatchId": "Batch Id of supply / productModelId of product",
                                "assetCount": 123.456,
                                "assetDescription": "Description of the supply/supplies",
                                "assetIds": "Comma separated list of actual supplies",
                                "assetType": "Is this a product(subassembly) or a supply asset"
                            }
                        ],
                        "makerId": "A maker's ID",
                        "manufacturingDate": "Date of Manufacturing",
                        "orderId": "An order's ID",
                        "productBatchId": "Batch id to which the product belongs",
                        "productDescription": "Description of the finished product",
                        "productId": "A product's ID",
                        "productInternalId": "A product's Internal ID",
                        "productModel": "Product Model",
                        "productRecalls": [
                            {
                                "assetType": "Is this a product(subassembly) or a supply asset",
                                "newassetIds": "Asset Ids of supply / product in recall",
                                "originalassetIds": "Asset Ids of supply / product in original schema"
                            }
                        ],
                        "productSchemaId": "A product schema ID",
                        "productStatus": "The current status of the product",
                        "productType": "FP(FinishedProduct/Subassembly",
                        "productionLocation": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        },
                        "sellerId": "A seller's ID"
                    }
                }
            ],
            "function": "updateAssetProduct"
        },
        "updateAssetSupply": {
            "args": [
                {
                    "supply": {
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "industryType": "Industry type - to start with Auto / Pharma / Food",
                        "makerId": "A maker's ID",
                        "makerproductId": "A product's ID",
                        "orderId": "An order's ID",
                        "shippedDate": "Date of shipment to maker",
                        "supplierId": "A supplier's ID",
                        "supplyAggregationLevel": "Supply type - RawMaterial / Component / Operational",
                        "supplyAvailableDate": "Date of manufacture",
                        "supplyBatchId": "Supplier's production batch under which this supply falls",
                        "supplyCount": 123.456,
                        "supplyDescription": "Description about the supply",
                        "supplyId": "A supply's ID",
                        "supplyStatus": "The current status of the supply",
                        "supplyType": "eg.LightingCircuits, EngineSubasembly etc "
                    }
                }
            ],
            "function": "updateAssetSupply"
        },
        "updateParticipantMaker": {
            "args": [
                {
                    "maker": {
                        "industryType": "Industry, to start off with is auto, pharma or food",
                        "makerAddress": {
                            "city": "carpe noctem",
                            "country": "carpe noctem",
                            "name": "carpe noctem",
                            "phone": "carpe noctem",
                            "street": "carpe noctem",
                            "website": "carpe noctem",
                            "zipcode": "carpe noctem"
                        },
                        "makerId": "A maker's ID",
                        "sellers": [
                            {
                                "seller": {
                                    "sellerId": "A seller's ID"
                                }
                            }
                        ],
                        "suppliers": [
                            {
                                "supplier": {
                                    "supplierId": "A supplier's ID"
                                }
                            }
                        ]
                    }
                }
            ],
            "function": "updateParticipantMaker"
        },
        "updateParticipantSeller": {
            "args": [
                {
                    "seller": {
                        "industryType": "Industry, to start off with is auto, pharma or food",
                        "makers": [
                            {
                                "maker": {
                                    "makerId": "A maker's ID"
                                }
                            }
                        ],
                        "sellerAddress": {
                            "city": "carpe noctem",
                            "country": "carpe noctem",
                            "name": "carpe noctem",
                            "phone": "carpe noctem",
                            "street": "carpe noctem",
                            "website": "carpe noctem",
                            "zipcode": "carpe noctem"
                        },
                        "sellerId": "A seller's ID"
                    }
                }
            ],
            "function": "updateParticipantSeller"
        },
        "updateParticipantSupplier": {
            "args": [
                {
                    "supplier": {
                        "industryType": "Industry, to start off with is auto, pharma or food",
                        "makers": [
                            {
                                "maker": {
                                    "makerId": "A maker's ID"
                                }
                            }
                        ],
                        "supplierAddress": {
                            "city": "carpe noctem",
                            "country": "carpe noctem",
                            "name": "carpe noctem",
                            "phone": "carpe noctem",
                            "street": "carpe noctem",
                            "website": "carpe noctem",
                            "zipcode": "carpe noctem"
                        },
                        "supplierId": "A supplier's ID"
                    }
                }
            ],
            "function": "updateParticipantSupplier"
        }
    },
    "Model": {
        "eventIOTContractPlatformInvokeResult": {
            "name": "EVT.IOTCP.INVOKE.RESULT",
            "payload": {
                "properties": "NO TYPE PROPERTY"
            }
        },
        "maker": {
            "industryType": "Industry, to start off with is auto, pharma or food",
            "makerAddress": {
                "city": "carpe noctem",
                "country": "carpe noctem",
                "name": "carpe noctem",
                "phone": "carpe noctem",
                "street": "carpe noctem",
                "website": "carpe noctem",
                "zipcode": "carpe noctem"
            },
            "makerId": "A maker's ID",
            "sellers": [
                {
                    "seller": {
                        "sellerId": "A seller's ID"
                    }
                }
            ],
            "suppliers": [
                {
                    "supplier": {
                        "supplierId": "A supplier's ID"
                    }
                }
            ]
        },
        "order": {
            "common": {
                "appdata": [
                    {
                        "K": "carpe noctem",
                        "V": "carpe noctem"
                    }
                ],
                "deviceID": "A unique identifier for the device that sent the current event",
                "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                "location": {
                    "latitude": 123.456,
                    "longitude": 123.456
                }
            },
            "goodsOrServicesType": "Goods / Services",
            "industryType": "Industry type - to start with Auto / Pharma / Food",
            "orderContentType": "Raw Material/ Component/Subassembly / Finished order",
            "orderDeliveryLocation": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "orderDescription": "Description of the order",
            "orderDistributor": {
                "bolId": "Bill of Lading number",
                "distId": "Participant Id of order distributor",
                "distName": "Participant name of order distributor"
            },
            "orderFulfiller": {
                "Id": "Participant Id of order fulfiller",
                "Type": "Participant type of order fulfiller",
                "actualFulfilmentDate": "Fulfilment date",
                "batchIds": "batch ids used to fill order",
                "batchSizeDelivered": 123.456,
                "committedDeliveryDate": "Order due date",
                "orderReceivedDate": "Date the order was received"
            },
            "orderFulfilmentLocation": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "orderId": "An order's ID",
            "orderIssuer": {
                "batchSizeRequired": 123.456,
                "oderIssuerId": "Participant Id of order issuer",
                "oderIssuerType": "Participant type of order issuer",
                "orderDueDate": "Order due date",
                "orderIssueDate": "Year of Manufacturing"
            },
            "orderStatus": "The current status of the order",
            "orderType": "Manufacturing / Distribution/MRO"
        },
        "product": {
            "common": {
                "appdata": [
                    {
                        "K": "carpe noctem",
                        "V": "carpe noctem"
                    }
                ],
                "deviceID": "A unique identifier for the device that sent the current event",
                "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                "location": {
                    "latitude": 123.456,
                    "longitude": 123.456
                }
            },
            "industryType": "Industry type - to start with Auto / Pharma / Food",
            "madeOf": [
                {
                    "assetBatchId": "Batch Id of supply / productModelId of product",
                    "assetCount": 123.456,
                    "assetDescription": "Description of the supply/supplies",
                    "assetIds": "Comma separated list of actual supplies",
                    "assetType": "Is this a product(subassembly) or a supply asset"
                }
            ],
            "makerId": "A maker's ID",
            "manufacturingDate": "Date of Manufacturing",
            "orderId": "An order's ID",
            "productBatchId": "Batch id to which the product belongs",
            "productDescription": "Description of the finished product",
            "productId": "A product's ID",
            "productInternalId": "A product's Internal ID",
            "productModel": "Product Model",
            "productRecalls": [
                {
                    "assetType": "Is this a product(subassembly) or a supply asset",
                    "newassetIds": "Asset Ids of supply / product in recall",
                    "originalassetIds": "Asset Ids of supply / product in original schema"
                }
            ],
            "productSchemaId": "A product schema ID",
            "productStatus": "The current status of the product",
            "productType": "FP(FinishedProduct/Subassembly",
            "productionLocation": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "sellerId": "A seller's ID"
        },
        "seller": {
            "industryType": "Industry, to start off with is auto, pharma or food",
            "makers": [
                {
                    "maker": {
                        "makerId": "A maker's ID"
                    }
                }
            ],
            "sellerAddress": {
                "city": "carpe noctem",
                "country": "carpe noctem",
                "name": "carpe noctem",
                "phone": "carpe noctem",
                "street": "carpe noctem",
                "website": "carpe noctem",
                "zipcode": "carpe noctem"
            },
            "sellerId": "A seller's ID"
        },
        "supplier": {
            "industryType": "Industry, to start off with is auto, pharma or food",
            "makers": [
                {
                    "maker": {
                        "makerId": "A maker's ID"
                    }
                }
            ],
            "supplierAddress": {
                "city": "carpe noctem",
                "country": "carpe noctem",
                "name": "carpe noctem",
                "phone": "carpe noctem",
                "street": "carpe noctem",
                "website": "carpe noctem",
                "zipcode": "carpe noctem"
            },
            "supplierId": "A supplier's ID"
        },
        "supply": {
            "common": {
                "appdata": [
                    {
                        "K": "carpe noctem",
                        "V": "carpe noctem"
                    }
                ],
                "deviceID": "A unique identifier for the device that sent the current event",
                "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                "location": {
                    "latitude": 123.456,
                    "longitude": 123.456
                }
            },
            "industryType": "Industry type - to start with Auto / Pharma / Food",
            "makerId": "A maker's ID",
            "makerproductId": "A product's ID",
            "orderId": "An order's ID",
            "shippedDate": "Date of shipment to maker",
            "supplierId": "A supplier's ID",
            "supplyAggregationLevel": "Supply type - RawMaterial / Component / Operational",
            "supplyAvailableDate": "Date of manufacture",
            "supplyBatchId": "Supplier's production batch under which this supply falls",
            "supplyCount": 123.456,
            "supplyDescription": "Description about the supply",
            "supplyId": "A supply's ID",
            "supplyStatus": "The current status of the supply",
            "supplyType": "eg.LightingCircuits, EngineSubasembly etc "
        }
    }
}`


	var readAssetSamples iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(samples), nil
	}

	func init() {
		iot.AddRoute("readAssetSamples", "query", iot.SystemClass, readAssetSamples)
	}
	