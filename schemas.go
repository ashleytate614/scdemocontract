package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)

var schemas = `

{
    "API": {
        "createAssetDistributionDetail": {
            "description": "Creates a new DistributionDetail (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "distributionDetail": {
                                "description": "Distribution Checkpoints",
                                "properties": {
                                    "checkPointEvents": {
                                        "description": "CheckPoint Events",
                                        "type": "string"
                                    },
                                    "checkPointShock": {
                                        "description": "Shock measured at the CheckPoint",
                                        "type": "string"
                                    },
                                    "checkPointTemperature": {
                                        "description": "Temperature measured at the checkpoint",
                                        "type": "string"
                                    },
                                    "checkpointDate": {
                                        "description": "Date-time at which the shipment is expected at the checkpoint",
                                        "type": "string"
                                    },
                                    "checkpointLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "checkpointName": {
                                        "description": "Checkpoint lat longs of distributor",
                                        "type": "string"
                                    },
                                    "distributionDetailId": {
                                        "description": "Unique id of a dist. record",
                                        "type": "string"
                                    },
                                    "distributionId": {
                                        "description": "Participant Id of order distributor",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "distributionDetailId"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createAssetDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createAssetOrder": {
            "description": "Creates a new Order (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "order": {
                                "description": "The changeable properties for a order, also considered its 'event' as a partial state",
                                "properties": {
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "goodsOrServicesType": {
                                        "description": "Goods / Services",
                                        "type": "string"
                                    },
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "orderContentType": {
                                        "description": "Raw Material/ Component/Subassembly / Finished order",
                                        "type": "string"
                                    },
                                    "orderDeliveryLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderDescription": {
                                        "description": "Description of the order",
                                        "type": "string"
                                    },
                                    "orderDistributor": {
                                        "description": "Order distributor details",
                                        "minItems": 1,
                                        "properties": {
                                            "bolId": {
                                                "description": "Bill of Lading number",
                                                "type": "string"
                                            },
                                            "distId": {
                                                "description": "Participant Id of order distributor",
                                                "type": "string"
                                            },
                                            "distName": {
                                                "description": "Participant name of order distributor",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderFulfiller": {
                                        "description": "Order fulfiller details",
                                        "minItems": 1,
                                        "properties": {
                                            "Id": {
                                                "description": "Participant Id of order fulfiller",
                                                "type": "string"
                                            },
                                            "Type": {
                                                "description": "Participant type of order fulfiller",
                                                "type": "string"
                                            },
                                            "actualFulfilmentDate": {
                                                "description": "Fulfilment date",
                                                "type": "string"
                                            },
                                            "batchIds": {
                                                "description": "batch ids used to fill order",
                                                "type": "string"
                                            },
                                            "batchSizeDelivered": {
                                                "description": "Number of assets",
                                                "type": "number"
                                            },
                                            "committedDeliveryDate": {
                                                "description": "Order due date",
                                                "type": "string"
                                            },
                                            "orderReceivedDate": {
                                                "description": "Date the order was received",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderFulfilmentLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    },
                                    "orderIssuer": {
                                        "description": "Order issuer details",
                                        "minItems": 1,
                                        "properties": {
                                            "batchSizeRequired": {
                                                "description": "Number of assets",
                                                "type": "number"
                                            },
                                            "oderIssuerId": {
                                                "description": "Participant Id of order issuer",
                                                "type": "string"
                                            },
                                            "oderIssuerType": {
                                                "description": "Participant type of order issuer",
                                                "type": "string"
                                            },
                                            "orderDueDate": {
                                                "description": "Order due date",
                                                "type": "string"
                                            },
                                            "orderIssueDate": {
                                                "description": "Year of Manufacturing",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderStatus": {
                                        "description": "The current status of the order",
                                        "type": "string"
                                    },
                                    "orderType": {
                                        "description": "Manufacturing / Distribution/MRO",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "orderId",
                                    "industryType",
                                    "orderType",
                                    "orderContentType",
                                    "orderStatus",
                                    "orderDescription",
                                    "orderFulfilmentLocation",
                                    "orderDeliveryLocation"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createAssetOrder"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createAssetProduct": {
            "description": "Creates a new Product (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "product": {
                                "description": "The changeable properties for a product, also considered its 'event' as a partial state",
                                "properties": {
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "madeOf": {
                                        "items": {
                                            "properties": {
                                                "assetBatchId": {
                                                    "description": "Batch Id of supply / productModelId of product",
                                                    "type": "string"
                                                },
                                                "assetCount": {
                                                    "description": "Number of items in the product",
                                                    "type": "number"
                                                },
                                                "assetDescription": {
                                                    "description": "Description of the supply/supplies",
                                                    "type": "string"
                                                },
                                                "assetIds": {
                                                    "description": "Comma separated list of actual supplies",
                                                    "type": "string"
                                                },
                                                "assetType": {
                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "manufacturingDate": {
                                        "description": "Date of Manufacturing",
                                        "type": "string"
                                    },
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    },
                                    "productBatchId": {
                                        "description": "Batch id to which the product belongs",
                                        "type": "string"
                                    },
                                    "productDescription": {
                                        "description": "Description of the finished product",
                                        "type": "string"
                                    },
                                    "productId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    },
                                    "productInternalId": {
                                        "description": "A product's Internal ID",
                                        "type": "string"
                                    },
                                    "productModel": {
                                        "description": "Product Model",
                                        "type": "string"
                                    },
                                    "productRecalls": {
                                        "items": {
                                            "properties": {
                                                "assetType": {
                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                    "type": "string"
                                                },
                                                "newassetIds": {
                                                    "description": "Asset Ids of supply / product in recall",
                                                    "type": "string"
                                                },
                                                "originalassetIds": {
                                                    "description": "Asset Ids of supply / product in original schema",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "productSchemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    },
                                    "productStatus": {
                                        "description": "The current status of the product",
                                        "type": "string"
                                    },
                                    "productType": {
                                        "description": "FP(FinishedProduct/Subassembly",
                                        "type": "string"
                                    },
                                    "productionLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "productId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createAssetProduct"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createAssetProductSchema": {
            "description": "Creates a new ProductSchema (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "productschema": {
                                "description": "The changeable properties for a productschema, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "productSchemaChange": {
                                        "items": {
                                            "properties": {
                                                "assetType": {
                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                    "type": "string"
                                                },
                                                "newassetBatchId": {
                                                    "description": "Batch Id of supply / productModelId of product in recall",
                                                    "type": "string"
                                                },
                                                "originalassetBatchId": {
                                                    "description": "Batch Id of supply / productModelId of product in original schema",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "productSchemaContent": {
                                        "description": "Schema content",
                                        "type": "string"
                                    },
                                    "productcode": {
                                        "description": "Maker's code for this product type",
                                        "type": "string"
                                    },
                                    "productschemaCreationDate": {
                                        "description": "Date of schema definition",
                                        "type": "string"
                                    },
                                    "productschemaDescription": {
                                        "description": "Description about the productschema",
                                        "type": "string"
                                    },
                                    "productschemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    },
                                    "productschemaStatus": {
                                        "description": "New / OnHold / InUse / Archived",
                                        "type": "string"
                                    },
                                    "productschemaType": {
                                        "description": "0 for subassembly, 1 for finished product",
                                        "type": "number"
                                    }
                                },
                                "required": [
                                    "productschemaId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createAssetProductSchema"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createAssetSupply": {
            "description": "Creates a new Supply (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supply": {
                                "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
                                "properties": {
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "makerproductId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    },
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    },
                                    "shippedDate": {
                                        "description": "Date of shipment to maker",
                                        "type": "string"
                                    },
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    },
                                    "supplyAggregationLevel": {
                                        "description": "Supply type - RawMaterial / Component / Operational",
                                        "type": "string"
                                    },
                                    "supplyAvailableDate": {
                                        "description": "Date of manufacture",
                                        "type": "string"
                                    },
                                    "supplyBatchId": {
                                        "description": "Supplier's production batch under which this supply falls",
                                        "type": "string"
                                    },
                                    "supplyCount": {
                                        "description": "Number of assets",
                                        "type": "number"
                                    },
                                    "supplyDescription": {
                                        "description": "Description about the supply",
                                        "type": "string"
                                    },
                                    "supplyId": {
                                        "description": "A supply's ID",
                                        "type": "string"
                                    },
                                    "supplyStatus": {
                                        "description": "The current status of the supply",
                                        "type": "string"
                                    },
                                    "supplyType": {
                                        "description": "eg.LightingCircuits, EngineSubasembly etc ",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "supplyId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createAssetSupply"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createParticipantMaker": {
            "description": "Creates a new maker (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "maker": {
                                "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry, to start off with is auto, pharma or food",
                                        "type": "string"
                                    },
                                    "makerAddress": {
                                        "description": "The makers' physical address and other means of contacting",
                                        "properties": {
                                            "city": {
                                                "type": "string"
                                            },
                                            "country": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            },
                                            "phone": {
                                                "type": "string"
                                            },
                                            "street": {
                                                "type": "string"
                                            },
                                            "website": {
                                                "type": "string"
                                            },
                                            "zipcode": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "sellers": {
                                        "items": {
                                            "properties": {
                                                "seller": {
                                                    "properties": {
                                                        "sellerId": {
                                                            "description": "A seller's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "suppliers": {
                                        "items": {
                                            "properties": {
                                                "supplier": {
                                                    "properties": {
                                                        "supplierId": {
                                                            "description": "A supplier's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    }
                                },
                                "required": [
                                    "makerId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createParticipantMaker"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createParticipantSeller": {
            "description": "Creates a new seller (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "seller": {
                                "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry, to start off with is auto, pharma or food",
                                        "type": "string"
                                    },
                                    "makers": {
                                        "items": {
                                            "properties": {
                                                "maker": {
                                                    "properties": {
                                                        "makerId": {
                                                            "description": "A maker's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "sellerAddress": {
                                        "description": "The sellers' physical address and other means of contacting",
                                        "properties": {
                                            "city": {
                                                "type": "string"
                                            },
                                            "country": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            },
                                            "phone": {
                                                "type": "string"
                                            },
                                            "street": {
                                                "type": "string"
                                            },
                                            "website": {
                                                "type": "string"
                                            },
                                            "zipcode": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "sellerId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createParticipantSeller"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createParticipantSupplier": {
            "description": "Creates a new supplier (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supplier": {
                                "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry, to start off with is auto, pharma or food",
                                        "type": "string"
                                    },
                                    "makers": {
                                        "items": {
                                            "properties": {
                                                "maker": {
                                                    "properties": {
                                                        "makerId": {
                                                            "description": "A maker's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "supplierAddress": {
                                        "description": "The suppliers' physical address and other means of contacting",
                                        "properties": {
                                            "city": {
                                                "type": "string"
                                            },
                                            "country": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            },
                                            "phone": {
                                                "type": "string"
                                            },
                                            "street": {
                                                "type": "string"
                                            },
                                            "website": {
                                                "type": "string"
                                            },
                                            "zipcode": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "supplierId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createParticipantSupplier"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsDistributionDetail": {
            "description": "Delete all distributionDetails from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllAssetsDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsOrder": {
            "description": "Delete all orders from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllAssetsOrder"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsProduct": {
            "description": "Delete all products from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllAssetsProduct"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsProductSchema": {
            "description": "Delete all productschemas from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllAssetsProductSchema"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsSupply": {
            "description": "Delete all supplys from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllAssetsSupply"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllParticipantsMaker": {
            "description": "Delete all makers from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllParticipantsMaker"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllParticipantsSeller": {
            "description": "Delete all sellers from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllParticipantsSeller"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllParticipantsSupplier": {
            "description": "Delete all suppliers from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllParticipantsSupplier"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetDistributionDetail": {
            "description": "Delete a distributionDetail from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "distributionDetail": {
                                "properties": {
                                    "distributionDetailId": {
                                        "description": "Unique id of a dist. record",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetOrder": {
            "description": "Delete a order from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "order": {
                                "properties": {
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetOrder"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetProduct": {
            "description": "Delete a product from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "product": {
                                "properties": {
                                    "productId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetProduct"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetProductSchema": {
            "description": "Delete a productschema from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "productschema": {
                                "properties": {
                                    "productschemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetProductSchema"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetStateHistoryDistributionDetail": {
            "description": "Delete a distributionDetail's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "distributionDetail": {
                                "properties": {
                                    "distributionDetailId": {
                                        "description": "Unique id of a dist. record",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetStateHistoryDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetStateHistoryOrder": {
            "description": "Delete a order's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "order": {
                                "properties": {
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetStateHistoryOrder"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetStateHistoryProduct": {
            "description": "Delete a product's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "product": {
                                "properties": {
                                    "productId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetStateHistoryProduct"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetStateHistoryProductSchema": {
            "description": "Delete a productschema's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "productschema": {
                                "properties": {
                                    "productschemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetStateHistoryProductSchema"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetStateHistorySupply": {
            "description": "Delete a supply's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supply": {
                                "properties": {
                                    "supplyId": {
                                        "description": "A supply's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetStateHistorySupply"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetSupply": {
            "description": "Delete a supply from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supply": {
                                "properties": {
                                    "supplyId": {
                                        "description": "A supply's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetSupply"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteParticipantMaker": {
            "description": "Delete a maker from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "maker": {
                                "properties": {
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteParticipantMaker"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteParticipantSeller": {
            "description": "Delete a seller from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "seller": {
                                "properties": {
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteParticipantSeller"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteParticipantStateHistoryMaker": {
            "description": "Delete a maker's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "maker": {
                                "properties": {
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteParticipantStateHistoryMaker"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteParticipantStateHistorySeller": {
            "description": "Delete a seller's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "seller": {
                                "properties": {
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteParticipantStateHistorySeller"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteParticipantStateHistorySupplier": {
            "description": "Delete a supplier's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supplier": {
                                "properties": {
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteParticipantStateHistorySupplier"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteParticipantSupplier": {
            "description": "Delete a supplier from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supplier": {
                                "properties": {
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteParticipantSupplier"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetDistributionDetail": {
            "description": "Delete one or more properties from a distributionDetail's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "distributionDetail": {
                                "properties": {
                                    "distributionDetailId": {
                                        "description": "Unique id of a dist. record",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "qprops": {
                                "description": "Qualified property names, e.g. distributionDetail.distributionDetailId",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromAssetDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetOrder": {
            "description": "Delete one or more properties from a order's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "order": {
                                "properties": {
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "qprops": {
                                "description": "Qualified property names, e.g. order.orderId",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromAssetOrder"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetProduct": {
            "description": "Delete one or more properties from a product's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "product": {
                                "properties": {
                                    "productId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "qprops": {
                                "description": "Qualified property names, e.g. product.productId",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromAssetProduct"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetProductSchema": {
            "description": "Delete one or more properties from a productschema's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "productschema": {
                                "properties": {
                                    "productschemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "qprops": {
                                "description": "Qualified property names, e.g. productschema.productschemaId",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromAssetProductSchema"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetSupply": {
            "description": "Delete one or more properties from a supply's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "qprops": {
                                "description": "Qualified property names, e.g. supply.supplyId",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "supply": {
                                "properties": {
                                    "supplyId": {
                                        "description": "A supply's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromAssetSupply"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromParticipantMaker": {
            "description": "Delete one or more properties from a maker's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "maker": {
                                "properties": {
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "qprops": {
                                "description": "Qualified property names, e.g. maker.barcode",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromParticipantMaker"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromParticipantSeller": {
            "description": "Delete one or more properties from a seller's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "qprops": {
                                "description": "Qualified property names, e.g. seller.barcode",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "seller": {
                                "properties": {
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromParticipantSeller"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromParticipantSupplier": {
            "description": "Delete one or more properties from a supplier's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "qprops": {
                                "description": "Qualified property names, e.g. supplier.barcode",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "supplier": {
                                "properties": {
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromParticipantSupplier"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteWorldState": {
            "description": "**** WARNING *** Clears the entire contents of world state, redeploy the contract after using this, in debugging mode, will require a restart",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteWorldState"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "initContract": {
            "description": "Sets contract version and nickname",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "nickname": {
                                "default": "IOT Contract Platform",
                                "description": "The nickname of the current contract instance",
                                "type": "string"
                            },
                            "version": {
                                "description": "The version number of the current contract instance",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "initContract"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAllAssetsDistributionDetail": {
            "description": "Returns the state of all distributionDetails, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllAssetsDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of distributionDetail states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^DISTDET": {
                                "description": "A distribution record's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This distribution record's world state distribution record ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the distribution record's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The distribution record's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This distribution record has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "distributionDetail": {
                                                "description": "Distribution Checkpoints",
                                                "properties": {
                                                    "checkPointEvents": {
                                                        "description": "CheckPoint Events",
                                                        "type": "string"
                                                    },
                                                    "checkPointShock": {
                                                        "description": "Shock measured at the CheckPoint",
                                                        "type": "string"
                                                    },
                                                    "checkPointTemperature": {
                                                        "description": "Temperature measured at the checkpoint",
                                                        "type": "string"
                                                    },
                                                    "checkpointDate": {
                                                        "description": "Date-time at which the shipment is expected at the checkpoint",
                                                        "type": "string"
                                                    },
                                                    "checkpointLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "checkpointName": {
                                                        "description": "Checkpoint lat longs of distributor",
                                                        "type": "string"
                                                    },
                                                    "distributionDetailId": {
                                                        "description": "Unique id of a dist. record",
                                                        "type": "string"
                                                    },
                                                    "distributionId": {
                                                        "description": "Participant Id of order distributor",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "distributionDetailId"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "distributionDetail": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The distribution record's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this distribution record",
                                        "properties": {
                                            "distribution record": {
                                                "description": "Distribution Checkpoints",
                                                "properties": {
                                                    "checkPointEvents": {
                                                        "description": "CheckPoint Events",
                                                        "type": "string"
                                                    },
                                                    "checkPointShock": {
                                                        "description": "Shock measured at the CheckPoint",
                                                        "type": "string"
                                                    },
                                                    "checkPointTemperature": {
                                                        "description": "Temperature measured at the checkpoint",
                                                        "type": "string"
                                                    },
                                                    "checkpointDate": {
                                                        "description": "Date-time at which the shipment is expected at the checkpoint",
                                                        "type": "string"
                                                    },
                                                    "checkpointLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "checkpointName": {
                                                        "description": "Checkpoint lat longs of distributor",
                                                        "type": "string"
                                                    },
                                                    "distributionDetailId": {
                                                        "description": "Unique id of a dist. record",
                                                        "type": "string"
                                                    },
                                                    "distributionId": {
                                                        "description": "Participant Id of order distributor",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "distributionDetailId"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllAssetsOrder": {
            "description": "Returns the state of all orders, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllAssetsOrder"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of order states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^PROD": {
                                "description": "A order's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This order's world state order ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the order's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The order's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This order has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "order": {
                                                "description": "The changeable properties for a order, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "goodsOrServicesType": {
                                                        "description": "Goods / Services",
                                                        "type": "string"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "orderContentType": {
                                                        "description": "Raw Material/ Component/Subassembly / Finished order",
                                                        "type": "string"
                                                    },
                                                    "orderDeliveryLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderDescription": {
                                                        "description": "Description of the order",
                                                        "type": "string"
                                                    },
                                                    "orderDistributor": {
                                                        "description": "Order distributor details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "bolId": {
                                                                "description": "Bill of Lading number",
                                                                "type": "string"
                                                            },
                                                            "distId": {
                                                                "description": "Participant Id of order distributor",
                                                                "type": "string"
                                                            },
                                                            "distName": {
                                                                "description": "Participant name of order distributor",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderFulfiller": {
                                                        "description": "Order fulfiller details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "Id": {
                                                                "description": "Participant Id of order fulfiller",
                                                                "type": "string"
                                                            },
                                                            "Type": {
                                                                "description": "Participant type of order fulfiller",
                                                                "type": "string"
                                                            },
                                                            "actualFulfilmentDate": {
                                                                "description": "Fulfilment date",
                                                                "type": "string"
                                                            },
                                                            "batchIds": {
                                                                "description": "batch ids used to fill order",
                                                                "type": "string"
                                                            },
                                                            "batchSizeDelivered": {
                                                                "description": "Number of assets",
                                                                "type": "number"
                                                            },
                                                            "committedDeliveryDate": {
                                                                "description": "Order due date",
                                                                "type": "string"
                                                            },
                                                            "orderReceivedDate": {
                                                                "description": "Date the order was received",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderFulfilmentLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "orderIssuer": {
                                                        "description": "Order issuer details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "batchSizeRequired": {
                                                                "description": "Number of assets",
                                                                "type": "number"
                                                            },
                                                            "oderIssuerId": {
                                                                "description": "Participant Id of order issuer",
                                                                "type": "string"
                                                            },
                                                            "oderIssuerType": {
                                                                "description": "Participant type of order issuer",
                                                                "type": "string"
                                                            },
                                                            "orderDueDate": {
                                                                "description": "Order due date",
                                                                "type": "string"
                                                            },
                                                            "orderIssueDate": {
                                                                "description": "Year of Manufacturing",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderStatus": {
                                                        "description": "The current status of the order",
                                                        "type": "string"
                                                    },
                                                    "orderType": {
                                                        "description": "Manufacturing / Distribution/MRO",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "orderId",
                                                    "industryType",
                                                    "orderType",
                                                    "orderContentType",
                                                    "orderStatus",
                                                    "orderDescription",
                                                    "orderFulfilmentLocation",
                                                    "orderDeliveryLocation"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "order": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The order's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this order",
                                        "properties": {
                                            "order": {
                                                "description": "The changeable properties for a order, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "goodsOrServicesType": {
                                                        "description": "Goods / Services",
                                                        "type": "string"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "orderContentType": {
                                                        "description": "Raw Material/ Component/Subassembly / Finished order",
                                                        "type": "string"
                                                    },
                                                    "orderDeliveryLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderDescription": {
                                                        "description": "Description of the order",
                                                        "type": "string"
                                                    },
                                                    "orderDistributor": {
                                                        "description": "Order distributor details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "bolId": {
                                                                "description": "Bill of Lading number",
                                                                "type": "string"
                                                            },
                                                            "distId": {
                                                                "description": "Participant Id of order distributor",
                                                                "type": "string"
                                                            },
                                                            "distName": {
                                                                "description": "Participant name of order distributor",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderFulfiller": {
                                                        "description": "Order fulfiller details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "Id": {
                                                                "description": "Participant Id of order fulfiller",
                                                                "type": "string"
                                                            },
                                                            "Type": {
                                                                "description": "Participant type of order fulfiller",
                                                                "type": "string"
                                                            },
                                                            "actualFulfilmentDate": {
                                                                "description": "Fulfilment date",
                                                                "type": "string"
                                                            },
                                                            "batchIds": {
                                                                "description": "batch ids used to fill order",
                                                                "type": "string"
                                                            },
                                                            "batchSizeDelivered": {
                                                                "description": "Number of assets",
                                                                "type": "number"
                                                            },
                                                            "committedDeliveryDate": {
                                                                "description": "Order due date",
                                                                "type": "string"
                                                            },
                                                            "orderReceivedDate": {
                                                                "description": "Date the order was received",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderFulfilmentLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "orderIssuer": {
                                                        "description": "Order issuer details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "batchSizeRequired": {
                                                                "description": "Number of assets",
                                                                "type": "number"
                                                            },
                                                            "oderIssuerId": {
                                                                "description": "Participant Id of order issuer",
                                                                "type": "string"
                                                            },
                                                            "oderIssuerType": {
                                                                "description": "Participant type of order issuer",
                                                                "type": "string"
                                                            },
                                                            "orderDueDate": {
                                                                "description": "Order due date",
                                                                "type": "string"
                                                            },
                                                            "orderIssueDate": {
                                                                "description": "Year of Manufacturing",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderStatus": {
                                                        "description": "The current status of the order",
                                                        "type": "string"
                                                    },
                                                    "orderType": {
                                                        "description": "Manufacturing / Distribution/MRO",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "orderId",
                                                    "industryType",
                                                    "orderType",
                                                    "orderContentType",
                                                    "orderStatus",
                                                    "orderDescription",
                                                    "orderFulfilmentLocation",
                                                    "orderDeliveryLocation"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllAssetsProduct": {
            "description": "Returns the state of all products, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllAssetsProduct"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of product states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^PROD": {
                                "description": "A product's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This product's world state product ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the product's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The product's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This product has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "product": {
                                                "description": "The changeable properties for a product, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "madeOf": {
                                                        "items": {
                                                            "properties": {
                                                                "assetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product",
                                                                    "type": "string"
                                                                },
                                                                "assetCount": {
                                                                    "description": "Number of items in the product",
                                                                    "type": "number"
                                                                },
                                                                "assetDescription": {
                                                                    "description": "Description of the supply/supplies",
                                                                    "type": "string"
                                                                },
                                                                "assetIds": {
                                                                    "description": "Comma separated list of actual supplies",
                                                                    "type": "string"
                                                                },
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "manufacturingDate": {
                                                        "description": "Date of Manufacturing",
                                                        "type": "string"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "productBatchId": {
                                                        "description": "Batch id to which the product belongs",
                                                        "type": "string"
                                                    },
                                                    "productDescription": {
                                                        "description": "Description of the finished product",
                                                        "type": "string"
                                                    },
                                                    "productId": {
                                                        "description": "A product's ID",
                                                        "type": "string"
                                                    },
                                                    "productInternalId": {
                                                        "description": "A product's Internal ID",
                                                        "type": "string"
                                                    },
                                                    "productModel": {
                                                        "description": "Product Model",
                                                        "type": "string"
                                                    },
                                                    "productRecalls": {
                                                        "items": {
                                                            "properties": {
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                },
                                                                "newassetIds": {
                                                                    "description": "Asset Ids of supply / product in recall",
                                                                    "type": "string"
                                                                },
                                                                "originalassetIds": {
                                                                    "description": "Asset Ids of supply / product in original schema",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "productSchemaId": {
                                                        "description": "A product schema ID",
                                                        "type": "string"
                                                    },
                                                    "productStatus": {
                                                        "description": "The current status of the product",
                                                        "type": "string"
                                                    },
                                                    "productType": {
                                                        "description": "FP(FinishedProduct/Subassembly",
                                                        "type": "string"
                                                    },
                                                    "productionLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "sellerId": {
                                                        "description": "A seller's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "productId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "product": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The product's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this product",
                                        "properties": {
                                            "product": {
                                                "description": "The changeable properties for a product, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "madeOf": {
                                                        "items": {
                                                            "properties": {
                                                                "assetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product",
                                                                    "type": "string"
                                                                },
                                                                "assetCount": {
                                                                    "description": "Number of items in the product",
                                                                    "type": "number"
                                                                },
                                                                "assetDescription": {
                                                                    "description": "Description of the supply/supplies",
                                                                    "type": "string"
                                                                },
                                                                "assetIds": {
                                                                    "description": "Comma separated list of actual supplies",
                                                                    "type": "string"
                                                                },
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "manufacturingDate": {
                                                        "description": "Date of Manufacturing",
                                                        "type": "string"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "productBatchId": {
                                                        "description": "Batch id to which the product belongs",
                                                        "type": "string"
                                                    },
                                                    "productDescription": {
                                                        "description": "Description of the finished product",
                                                        "type": "string"
                                                    },
                                                    "productId": {
                                                        "description": "A product's ID",
                                                        "type": "string"
                                                    },
                                                    "productInternalId": {
                                                        "description": "A product's Internal ID",
                                                        "type": "string"
                                                    },
                                                    "productModel": {
                                                        "description": "Product Model",
                                                        "type": "string"
                                                    },
                                                    "productRecalls": {
                                                        "items": {
                                                            "properties": {
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                },
                                                                "newassetIds": {
                                                                    "description": "Asset Ids of supply / product in recall",
                                                                    "type": "string"
                                                                },
                                                                "originalassetIds": {
                                                                    "description": "Asset Ids of supply / product in original schema",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "productSchemaId": {
                                                        "description": "A product schema ID",
                                                        "type": "string"
                                                    },
                                                    "productStatus": {
                                                        "description": "The current status of the product",
                                                        "type": "string"
                                                    },
                                                    "productType": {
                                                        "description": "FP(FinishedProduct/Subassembly",
                                                        "type": "string"
                                                    },
                                                    "productionLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "sellerId": {
                                                        "description": "A seller's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "productId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllAssetsProductSchema": {
            "description": "Returns the state of all productschemas, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllAssetsProductSchema"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of productschema states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^SUP": {
                                "description": "A productschema's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This productschema's world state productschema ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the productschema's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The productschema's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This productschema has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "productschema": {
                                                "description": "The changeable properties for a productschema, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "productSchemaChange": {
                                                        "items": {
                                                            "properties": {
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                },
                                                                "newassetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product in recall",
                                                                    "type": "string"
                                                                },
                                                                "originalassetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product in original schema",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "productSchemaContent": {
                                                        "description": "Schema content",
                                                        "type": "string"
                                                    },
                                                    "productcode": {
                                                        "description": "Maker's code for this product type",
                                                        "type": "string"
                                                    },
                                                    "productschemaCreationDate": {
                                                        "description": "Date of schema definition",
                                                        "type": "string"
                                                    },
                                                    "productschemaDescription": {
                                                        "description": "Description about the productschema",
                                                        "type": "string"
                                                    },
                                                    "productschemaId": {
                                                        "description": "A product schema ID",
                                                        "type": "string"
                                                    },
                                                    "productschemaStatus": {
                                                        "description": "New / OnHold / InUse / Archived",
                                                        "type": "string"
                                                    },
                                                    "productschemaType": {
                                                        "description": "0 for subassembly, 1 for finished product",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "productschemaId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "productschema": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The productschema's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this productschema",
                                        "properties": {
                                            "productschema": {
                                                "description": "The changeable properties for a productschema, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "productSchemaChange": {
                                                        "items": {
                                                            "properties": {
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                },
                                                                "newassetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product in recall",
                                                                    "type": "string"
                                                                },
                                                                "originalassetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product in original schema",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "productSchemaContent": {
                                                        "description": "Schema content",
                                                        "type": "string"
                                                    },
                                                    "productcode": {
                                                        "description": "Maker's code for this product type",
                                                        "type": "string"
                                                    },
                                                    "productschemaCreationDate": {
                                                        "description": "Date of schema definition",
                                                        "type": "string"
                                                    },
                                                    "productschemaDescription": {
                                                        "description": "Description about the productschema",
                                                        "type": "string"
                                                    },
                                                    "productschemaId": {
                                                        "description": "A product schema ID",
                                                        "type": "string"
                                                    },
                                                    "productschemaStatus": {
                                                        "description": "New / OnHold / InUse / Archived",
                                                        "type": "string"
                                                    },
                                                    "productschemaType": {
                                                        "description": "0 for subassembly, 1 for finished product",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "productschemaId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllAssetsSupply": {
            "description": "Returns the state of all supplys, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllAssetsSupply"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of supply states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^SUP": {
                                "description": "A supply's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This supply's world state supply ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the supply's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The supply's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This supply has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "supply": {
                                                "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "makerproductId": {
                                                        "description": "A product's ID",
                                                        "type": "string"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "shippedDate": {
                                                        "description": "Date of shipment to maker",
                                                        "type": "string"
                                                    },
                                                    "supplierId": {
                                                        "description": "A supplier's ID",
                                                        "type": "string"
                                                    },
                                                    "supplyAggregationLevel": {
                                                        "description": "Supply type - RawMaterial / Component / Operational",
                                                        "type": "string"
                                                    },
                                                    "supplyAvailableDate": {
                                                        "description": "Date of manufacture",
                                                        "type": "string"
                                                    },
                                                    "supplyBatchId": {
                                                        "description": "Supplier's production batch under which this supply falls",
                                                        "type": "string"
                                                    },
                                                    "supplyCount": {
                                                        "description": "Number of assets",
                                                        "type": "number"
                                                    },
                                                    "supplyDescription": {
                                                        "description": "Description about the supply",
                                                        "type": "string"
                                                    },
                                                    "supplyId": {
                                                        "description": "A supply's ID",
                                                        "type": "string"
                                                    },
                                                    "supplyStatus": {
                                                        "description": "The current status of the supply",
                                                        "type": "string"
                                                    },
                                                    "supplyType": {
                                                        "description": "eg.LightingCircuits, EngineSubasembly etc ",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "supplyId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "supply": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The supply's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this supply",
                                        "properties": {
                                            "supply": {
                                                "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "makerproductId": {
                                                        "description": "A product's ID",
                                                        "type": "string"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "shippedDate": {
                                                        "description": "Date of shipment to maker",
                                                        "type": "string"
                                                    },
                                                    "supplierId": {
                                                        "description": "A supplier's ID",
                                                        "type": "string"
                                                    },
                                                    "supplyAggregationLevel": {
                                                        "description": "Supply type - RawMaterial / Component / Operational",
                                                        "type": "string"
                                                    },
                                                    "supplyAvailableDate": {
                                                        "description": "Date of manufacture",
                                                        "type": "string"
                                                    },
                                                    "supplyBatchId": {
                                                        "description": "Supplier's production batch under which this supply falls",
                                                        "type": "string"
                                                    },
                                                    "supplyCount": {
                                                        "description": "Number of assets",
                                                        "type": "number"
                                                    },
                                                    "supplyDescription": {
                                                        "description": "Description about the supply",
                                                        "type": "string"
                                                    },
                                                    "supplyId": {
                                                        "description": "A supply's ID",
                                                        "type": "string"
                                                    },
                                                    "supplyStatus": {
                                                        "description": "The current status of the supply",
                                                        "type": "string"
                                                    },
                                                    "supplyType": {
                                                        "description": "eg.LightingCircuits, EngineSubasembly etc ",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "supplyId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllParticipantsMaker": {
            "description": "Returns the state of all makers, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllParticipantsMaker"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of maker states, can mix participant classes",
                    "items": {
                        "patternProperties": {
                            "^SUPR": {
                                "description": "A maker's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This maker's world state maker ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the maker's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The maker's participant class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This maker has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateParticipantMaker",
                                        "properties": {
                                            "maker": {
                                                "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makerAddress": {
                                                        "description": "The makers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "sellers": {
                                                        "items": {
                                                            "properties": {
                                                                "seller": {
                                                                    "properties": {
                                                                        "sellerId": {
                                                                            "description": "A seller's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "suppliers": {
                                                        "items": {
                                                            "properties": {
                                                                "supplier": {
                                                                    "properties": {
                                                                        "supplierId": {
                                                                            "description": "A supplier's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    }
                                                },
                                                "required": [
                                                    "makerId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "maker": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The maker's participant class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this maker",
                                        "properties": {
                                            "maker": {
                                                "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makerAddress": {
                                                        "description": "The makers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "sellers": {
                                                        "items": {
                                                            "properties": {
                                                                "seller": {
                                                                    "properties": {
                                                                        "sellerId": {
                                                                            "description": "A seller's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "suppliers": {
                                                        "items": {
                                                            "properties": {
                                                                "supplier": {
                                                                    "properties": {
                                                                        "supplierId": {
                                                                            "description": "A supplier's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    }
                                                },
                                                "required": [
                                                    "makerId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllParticipantsSeller": {
            "description": "Returns the state of all sellers, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllParticipantsSeller"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of seller states, can mix participant classes",
                    "items": {
                        "patternProperties": {
                            "^SUPR": {
                                "description": "A seller's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This seller's world state seller ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the seller's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The seller's participant class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This seller has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateParticipantSeller",
                                        "properties": {
                                            "seller": {
                                                "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makers": {
                                                        "items": {
                                                            "properties": {
                                                                "maker": {
                                                                    "properties": {
                                                                        "makerId": {
                                                                            "description": "A maker's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "sellerAddress": {
                                                        "description": "The sellers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "sellerId": {
                                                        "description": "A seller's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "sellerId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "seller": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The seller's participant class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this seller",
                                        "properties": {
                                            "seller": {
                                                "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makers": {
                                                        "items": {
                                                            "properties": {
                                                                "maker": {
                                                                    "properties": {
                                                                        "makerId": {
                                                                            "description": "A maker's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "sellerAddress": {
                                                        "description": "The sellers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "sellerId": {
                                                        "description": "A seller's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "sellerId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllParticipantsSupplier": {
            "description": "Returns the state of all suppliers, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllParticipantsSupplier"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of supplier states, can mix participant classes",
                    "items": {
                        "patternProperties": {
                            "^SUPR": {
                                "description": "A supplier's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This supplier's world state supplier ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the supplier's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The supplier's participant class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This supplier has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateParticipantSupplier",
                                        "properties": {
                                            "supplier": {
                                                "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makers": {
                                                        "items": {
                                                            "properties": {
                                                                "maker": {
                                                                    "properties": {
                                                                        "makerId": {
                                                                            "description": "A maker's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "supplierAddress": {
                                                        "description": "The suppliers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "supplierId": {
                                                        "description": "A supplier's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "supplierId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "supplier": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The supplier's participant class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this supplier",
                                        "properties": {
                                            "supplier": {
                                                "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makers": {
                                                        "items": {
                                                            "properties": {
                                                                "maker": {
                                                                    "properties": {
                                                                        "makerId": {
                                                                            "description": "A maker's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "supplierAddress": {
                                                        "description": "The suppliers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "supplierId": {
                                                        "description": "A supplier's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "supplierId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllRoutes": {
            "description": "Returns an array of registered API calls by function (debugging)",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllRoutes"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "An array of routes",
                    "items": {
                        "description": "A route defines a contract API that can be called to perform a service",
                        "properties": {
                            "class": {
                                "description": "An asset's classifier definition",
                                "properties": {
                                    "assetidpath": "An asset's primary key, expressed as a qualified property path (see example contracts)",
                                    "name": "An asset's class name",
                                    "prefix": "An asset's world state prefix, used to allow iteration over all assets of a class"
                                },
                                "type": "object"
                            },
                            "functionname": {
                                "type": "string"
                            },
                            "method": {
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetDistributionDetail": {
            "description": "Returns the state a distributionDetail",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "distributionDetail": {
                                "properties": {
                                    "distributionDetailId": {
                                        "description": "Unique id of a dist. record",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A distribution record's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This distribution record's world state distribution record ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "An array of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the distribution record's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The distribution record's asset class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This distribution record has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAssetContainer",
                            "properties": {
                                "distributionDetail": {
                                    "description": "Distribution Checkpoints",
                                    "properties": {
                                        "checkPointEvents": {
                                            "description": "CheckPoint Events",
                                            "type": "string"
                                        },
                                        "checkPointShock": {
                                            "description": "Shock measured at the CheckPoint",
                                            "type": "string"
                                        },
                                        "checkPointTemperature": {
                                            "description": "Temperature measured at the checkpoint",
                                            "type": "string"
                                        },
                                        "checkpointDate": {
                                            "description": "Date-time at which the shipment is expected at the checkpoint",
                                            "type": "string"
                                        },
                                        "checkpointLocation": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "checkpointName": {
                                            "description": "Checkpoint lat longs of distributor",
                                            "type": "string"
                                        },
                                        "distributionDetailId": {
                                            "description": "Unique id of a dist. record",
                                            "type": "string"
                                        },
                                        "distributionId": {
                                            "description": "Participant Id of order distributor",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "distributionDetailId"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "distributionDetail": {
                                    "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                    "properties": {
                                        "name": {
                                            "default": "EVT.IOTCP.INVOKE.RESULT",
                                            "enum": [
                                                "EVT.IOTCP.INVOKE.RESULT"
                                            ],
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A map of contributed results",
                                            "properties": {
                                                "description": "the overall status of the invoke result, defined by err",
                                                "properties": {
                                                    "activeAlerts": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsCleared": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsRaised": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "invokeresult": {
                                                        "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                        "properties": {
                                                            "message": {
                                                                "type": "string"
                                                            },
                                                            "status": {
                                                                "enum": [
                                                                    "OK",
                                                                    "ERROR"
                                                                ],
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The distribution record's asset class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this distribution record",
                            "properties": {
                                "distribution record": {
                                    "description": "Distribution Checkpoints",
                                    "properties": {
                                        "checkPointEvents": {
                                            "description": "CheckPoint Events",
                                            "type": "string"
                                        },
                                        "checkPointShock": {
                                            "description": "Shock measured at the CheckPoint",
                                            "type": "string"
                                        },
                                        "checkPointTemperature": {
                                            "description": "Temperature measured at the checkpoint",
                                            "type": "string"
                                        },
                                        "checkpointDate": {
                                            "description": "Date-time at which the shipment is expected at the checkpoint",
                                            "type": "string"
                                        },
                                        "checkpointLocation": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "checkpointName": {
                                            "description": "Checkpoint lat longs of distributor",
                                            "type": "string"
                                        },
                                        "distributionDetailId": {
                                            "description": "Unique id of a dist. record",
                                            "type": "string"
                                        },
                                        "distributionId": {
                                            "description": "Participant Id of order distributor",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "distributionDetailId"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetOrder": {
            "description": "Returns the state a order",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "order": {
                                "properties": {
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetOrder"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A order's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This order's world state order ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "An array of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the order's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The order's asset class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This order has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAssetContainer",
                            "properties": {
                                "order": {
                                    "description": "The changeable properties for a order, also considered its 'event' as a partial state",
                                    "properties": {
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "goodsOrServicesType": {
                                            "description": "Goods / Services",
                                            "type": "string"
                                        },
                                        "industryType": {
                                            "description": "Industry type - to start with Auto / Pharma / Food",
                                            "type": "string"
                                        },
                                        "orderContentType": {
                                            "description": "Raw Material/ Component/Subassembly / Finished order",
                                            "type": "string"
                                        },
                                        "orderDeliveryLocation": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderDescription": {
                                            "description": "Description of the order",
                                            "type": "string"
                                        },
                                        "orderDistributor": {
                                            "description": "Order distributor details",
                                            "minItems": 1,
                                            "properties": {
                                                "bolId": {
                                                    "description": "Bill of Lading number",
                                                    "type": "string"
                                                },
                                                "distId": {
                                                    "description": "Participant Id of order distributor",
                                                    "type": "string"
                                                },
                                                "distName": {
                                                    "description": "Participant name of order distributor",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderFulfiller": {
                                            "description": "Order fulfiller details",
                                            "minItems": 1,
                                            "properties": {
                                                "Id": {
                                                    "description": "Participant Id of order fulfiller",
                                                    "type": "string"
                                                },
                                                "Type": {
                                                    "description": "Participant type of order fulfiller",
                                                    "type": "string"
                                                },
                                                "actualFulfilmentDate": {
                                                    "description": "Fulfilment date",
                                                    "type": "string"
                                                },
                                                "batchIds": {
                                                    "description": "batch ids used to fill order",
                                                    "type": "string"
                                                },
                                                "batchSizeDelivered": {
                                                    "description": "Number of assets",
                                                    "type": "number"
                                                },
                                                "committedDeliveryDate": {
                                                    "description": "Order due date",
                                                    "type": "string"
                                                },
                                                "orderReceivedDate": {
                                                    "description": "Date the order was received",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderFulfilmentLocation": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderId": {
                                            "description": "An order's ID",
                                            "type": "string"
                                        },
                                        "orderIssuer": {
                                            "description": "Order issuer details",
                                            "minItems": 1,
                                            "properties": {
                                                "batchSizeRequired": {
                                                    "description": "Number of assets",
                                                    "type": "number"
                                                },
                                                "oderIssuerId": {
                                                    "description": "Participant Id of order issuer",
                                                    "type": "string"
                                                },
                                                "oderIssuerType": {
                                                    "description": "Participant type of order issuer",
                                                    "type": "string"
                                                },
                                                "orderDueDate": {
                                                    "description": "Order due date",
                                                    "type": "string"
                                                },
                                                "orderIssueDate": {
                                                    "description": "Year of Manufacturing",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderStatus": {
                                            "description": "The current status of the order",
                                            "type": "string"
                                        },
                                        "orderType": {
                                            "description": "Manufacturing / Distribution/MRO",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "orderId",
                                        "industryType",
                                        "orderType",
                                        "orderContentType",
                                        "orderStatus",
                                        "orderDescription",
                                        "orderFulfilmentLocation",
                                        "orderDeliveryLocation"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "order": {
                                    "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                    "properties": {
                                        "name": {
                                            "default": "EVT.IOTCP.INVOKE.RESULT",
                                            "enum": [
                                                "EVT.IOTCP.INVOKE.RESULT"
                                            ],
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A map of contributed results",
                                            "properties": {
                                                "description": "the overall status of the invoke result, defined by err",
                                                "properties": {
                                                    "activeAlerts": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsCleared": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsRaised": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "invokeresult": {
                                                        "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                        "properties": {
                                                            "message": {
                                                                "type": "string"
                                                            },
                                                            "status": {
                                                                "enum": [
                                                                    "OK",
                                                                    "ERROR"
                                                                ],
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The order's asset class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this order",
                            "properties": {
                                "order": {
                                    "description": "The changeable properties for a order, also considered its 'event' as a partial state",
                                    "properties": {
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "goodsOrServicesType": {
                                            "description": "Goods / Services",
                                            "type": "string"
                                        },
                                        "industryType": {
                                            "description": "Industry type - to start with Auto / Pharma / Food",
                                            "type": "string"
                                        },
                                        "orderContentType": {
                                            "description": "Raw Material/ Component/Subassembly / Finished order",
                                            "type": "string"
                                        },
                                        "orderDeliveryLocation": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderDescription": {
                                            "description": "Description of the order",
                                            "type": "string"
                                        },
                                        "orderDistributor": {
                                            "description": "Order distributor details",
                                            "minItems": 1,
                                            "properties": {
                                                "bolId": {
                                                    "description": "Bill of Lading number",
                                                    "type": "string"
                                                },
                                                "distId": {
                                                    "description": "Participant Id of order distributor",
                                                    "type": "string"
                                                },
                                                "distName": {
                                                    "description": "Participant name of order distributor",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderFulfiller": {
                                            "description": "Order fulfiller details",
                                            "minItems": 1,
                                            "properties": {
                                                "Id": {
                                                    "description": "Participant Id of order fulfiller",
                                                    "type": "string"
                                                },
                                                "Type": {
                                                    "description": "Participant type of order fulfiller",
                                                    "type": "string"
                                                },
                                                "actualFulfilmentDate": {
                                                    "description": "Fulfilment date",
                                                    "type": "string"
                                                },
                                                "batchIds": {
                                                    "description": "batch ids used to fill order",
                                                    "type": "string"
                                                },
                                                "batchSizeDelivered": {
                                                    "description": "Number of assets",
                                                    "type": "number"
                                                },
                                                "committedDeliveryDate": {
                                                    "description": "Order due date",
                                                    "type": "string"
                                                },
                                                "orderReceivedDate": {
                                                    "description": "Date the order was received",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderFulfilmentLocation": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderId": {
                                            "description": "An order's ID",
                                            "type": "string"
                                        },
                                        "orderIssuer": {
                                            "description": "Order issuer details",
                                            "minItems": 1,
                                            "properties": {
                                                "batchSizeRequired": {
                                                    "description": "Number of assets",
                                                    "type": "number"
                                                },
                                                "oderIssuerId": {
                                                    "description": "Participant Id of order issuer",
                                                    "type": "string"
                                                },
                                                "oderIssuerType": {
                                                    "description": "Participant type of order issuer",
                                                    "type": "string"
                                                },
                                                "orderDueDate": {
                                                    "description": "Order due date",
                                                    "type": "string"
                                                },
                                                "orderIssueDate": {
                                                    "description": "Year of Manufacturing",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "orderStatus": {
                                            "description": "The current status of the order",
                                            "type": "string"
                                        },
                                        "orderType": {
                                            "description": "Manufacturing / Distribution/MRO",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "orderId",
                                        "industryType",
                                        "orderType",
                                        "orderContentType",
                                        "orderStatus",
                                        "orderDescription",
                                        "orderFulfilmentLocation",
                                        "orderDeliveryLocation"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetProduct": {
            "description": "Returns the state a product",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "product": {
                                "properties": {
                                    "productId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetProduct"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A product's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This product's world state product ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "An array of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the product's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The product's asset class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This product has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAssetContainer",
                            "properties": {
                                "product": {
                                    "description": "The changeable properties for a product, also considered its 'event' as a partial state",
                                    "properties": {
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "industryType": {
                                            "description": "Industry type - to start with Auto / Pharma / Food",
                                            "type": "string"
                                        },
                                        "madeOf": {
                                            "items": {
                                                "properties": {
                                                    "assetBatchId": {
                                                        "description": "Batch Id of supply / productModelId of product",
                                                        "type": "string"
                                                    },
                                                    "assetCount": {
                                                        "description": "Number of items in the product",
                                                        "type": "number"
                                                    },
                                                    "assetDescription": {
                                                        "description": "Description of the supply/supplies",
                                                        "type": "string"
                                                    },
                                                    "assetIds": {
                                                        "description": "Comma separated list of actual supplies",
                                                        "type": "string"
                                                    },
                                                    "assetType": {
                                                        "description": "Is this a product(subassembly) or a supply asset",
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "makerId": {
                                            "description": "A maker's ID",
                                            "type": "string"
                                        },
                                        "manufacturingDate": {
                                            "description": "Date of Manufacturing",
                                            "type": "string"
                                        },
                                        "orderId": {
                                            "description": "An order's ID",
                                            "type": "string"
                                        },
                                        "productBatchId": {
                                            "description": "Batch id to which the product belongs",
                                            "type": "string"
                                        },
                                        "productDescription": {
                                            "description": "Description of the finished product",
                                            "type": "string"
                                        },
                                        "productId": {
                                            "description": "A product's ID",
                                            "type": "string"
                                        },
                                        "productInternalId": {
                                            "description": "A product's Internal ID",
                                            "type": "string"
                                        },
                                        "productModel": {
                                            "description": "Product Model",
                                            "type": "string"
                                        },
                                        "productRecalls": {
                                            "items": {
                                                "properties": {
                                                    "assetType": {
                                                        "description": "Is this a product(subassembly) or a supply asset",
                                                        "type": "string"
                                                    },
                                                    "newassetIds": {
                                                        "description": "Asset Ids of supply / product in recall",
                                                        "type": "string"
                                                    },
                                                    "originalassetIds": {
                                                        "description": "Asset Ids of supply / product in original schema",
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "productSchemaId": {
                                            "description": "A product schema ID",
                                            "type": "string"
                                        },
                                        "productStatus": {
                                            "description": "The current status of the product",
                                            "type": "string"
                                        },
                                        "productType": {
                                            "description": "FP(FinishedProduct/Subassembly",
                                            "type": "string"
                                        },
                                        "productionLocation": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "sellerId": {
                                            "description": "A seller's ID",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "productId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "product": {
                                    "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                    "properties": {
                                        "name": {
                                            "default": "EVT.IOTCP.INVOKE.RESULT",
                                            "enum": [
                                                "EVT.IOTCP.INVOKE.RESULT"
                                            ],
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A map of contributed results",
                                            "properties": {
                                                "description": "the overall status of the invoke result, defined by err",
                                                "properties": {
                                                    "activeAlerts": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsCleared": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsRaised": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "invokeresult": {
                                                        "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                        "properties": {
                                                            "message": {
                                                                "type": "string"
                                                            },
                                                            "status": {
                                                                "enum": [
                                                                    "OK",
                                                                    "ERROR"
                                                                ],
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The product's asset class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this product",
                            "properties": {
                                "product": {
                                    "description": "The changeable properties for a product, also considered its 'event' as a partial state",
                                    "properties": {
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "industryType": {
                                            "description": "Industry type - to start with Auto / Pharma / Food",
                                            "type": "string"
                                        },
                                        "madeOf": {
                                            "items": {
                                                "properties": {
                                                    "assetBatchId": {
                                                        "description": "Batch Id of supply / productModelId of product",
                                                        "type": "string"
                                                    },
                                                    "assetCount": {
                                                        "description": "Number of items in the product",
                                                        "type": "number"
                                                    },
                                                    "assetDescription": {
                                                        "description": "Description of the supply/supplies",
                                                        "type": "string"
                                                    },
                                                    "assetIds": {
                                                        "description": "Comma separated list of actual supplies",
                                                        "type": "string"
                                                    },
                                                    "assetType": {
                                                        "description": "Is this a product(subassembly) or a supply asset",
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "makerId": {
                                            "description": "A maker's ID",
                                            "type": "string"
                                        },
                                        "manufacturingDate": {
                                            "description": "Date of Manufacturing",
                                            "type": "string"
                                        },
                                        "orderId": {
                                            "description": "An order's ID",
                                            "type": "string"
                                        },
                                        "productBatchId": {
                                            "description": "Batch id to which the product belongs",
                                            "type": "string"
                                        },
                                        "productDescription": {
                                            "description": "Description of the finished product",
                                            "type": "string"
                                        },
                                        "productId": {
                                            "description": "A product's ID",
                                            "type": "string"
                                        },
                                        "productInternalId": {
                                            "description": "A product's Internal ID",
                                            "type": "string"
                                        },
                                        "productModel": {
                                            "description": "Product Model",
                                            "type": "string"
                                        },
                                        "productRecalls": {
                                            "items": {
                                                "properties": {
                                                    "assetType": {
                                                        "description": "Is this a product(subassembly) or a supply asset",
                                                        "type": "string"
                                                    },
                                                    "newassetIds": {
                                                        "description": "Asset Ids of supply / product in recall",
                                                        "type": "string"
                                                    },
                                                    "originalassetIds": {
                                                        "description": "Asset Ids of supply / product in original schema",
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "productSchemaId": {
                                            "description": "A product schema ID",
                                            "type": "string"
                                        },
                                        "productStatus": {
                                            "description": "The current status of the product",
                                            "type": "string"
                                        },
                                        "productType": {
                                            "description": "FP(FinishedProduct/Subassembly",
                                            "type": "string"
                                        },
                                        "productionLocation": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "sellerId": {
                                            "description": "A seller's ID",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "productId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetProductSchema": {
            "description": "Returns the state a productschema",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "productschema": {
                                "properties": {
                                    "productschemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetProductSchema"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A productschema's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This productschema's world state productschema ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "An array of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the productschema's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The productschema's asset class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This productschema has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAssetContainer",
                            "properties": {
                                "productschema": {
                                    "description": "The changeable properties for a productschema, also considered its 'event' as a partial state",
                                    "properties": {
                                        "industryType": {
                                            "description": "Industry type - to start with Auto / Pharma / Food",
                                            "type": "string"
                                        },
                                        "makerId": {
                                            "description": "A maker's ID",
                                            "type": "string"
                                        },
                                        "productSchemaChange": {
                                            "items": {
                                                "properties": {
                                                    "assetType": {
                                                        "description": "Is this a product(subassembly) or a supply asset",
                                                        "type": "string"
                                                    },
                                                    "newassetBatchId": {
                                                        "description": "Batch Id of supply / productModelId of product in recall",
                                                        "type": "string"
                                                    },
                                                    "originalassetBatchId": {
                                                        "description": "Batch Id of supply / productModelId of product in original schema",
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "productSchemaContent": {
                                            "description": "Schema content",
                                            "type": "string"
                                        },
                                        "productcode": {
                                            "description": "Maker's code for this product type",
                                            "type": "string"
                                        },
                                        "productschemaCreationDate": {
                                            "description": "Date of schema definition",
                                            "type": "string"
                                        },
                                        "productschemaDescription": {
                                            "description": "Description about the productschema",
                                            "type": "string"
                                        },
                                        "productschemaId": {
                                            "description": "A product schema ID",
                                            "type": "string"
                                        },
                                        "productschemaStatus": {
                                            "description": "New / OnHold / InUse / Archived",
                                            "type": "string"
                                        },
                                        "productschemaType": {
                                            "description": "0 for subassembly, 1 for finished product",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "productschemaId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "productschema": {
                                    "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                    "properties": {
                                        "name": {
                                            "default": "EVT.IOTCP.INVOKE.RESULT",
                                            "enum": [
                                                "EVT.IOTCP.INVOKE.RESULT"
                                            ],
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A map of contributed results",
                                            "properties": {
                                                "description": "the overall status of the invoke result, defined by err",
                                                "properties": {
                                                    "activeAlerts": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsCleared": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsRaised": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "invokeresult": {
                                                        "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                        "properties": {
                                                            "message": {
                                                                "type": "string"
                                                            },
                                                            "status": {
                                                                "enum": [
                                                                    "OK",
                                                                    "ERROR"
                                                                ],
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The productschema's asset class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this productschema",
                            "properties": {
                                "productschema": {
                                    "description": "The changeable properties for a productschema, also considered its 'event' as a partial state",
                                    "properties": {
                                        "industryType": {
                                            "description": "Industry type - to start with Auto / Pharma / Food",
                                            "type": "string"
                                        },
                                        "makerId": {
                                            "description": "A maker's ID",
                                            "type": "string"
                                        },
                                        "productSchemaChange": {
                                            "items": {
                                                "properties": {
                                                    "assetType": {
                                                        "description": "Is this a product(subassembly) or a supply asset",
                                                        "type": "string"
                                                    },
                                                    "newassetBatchId": {
                                                        "description": "Batch Id of supply / productModelId of product in recall",
                                                        "type": "string"
                                                    },
                                                    "originalassetBatchId": {
                                                        "description": "Batch Id of supply / productModelId of product in original schema",
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "productSchemaContent": {
                                            "description": "Schema content",
                                            "type": "string"
                                        },
                                        "productcode": {
                                            "description": "Maker's code for this product type",
                                            "type": "string"
                                        },
                                        "productschemaCreationDate": {
                                            "description": "Date of schema definition",
                                            "type": "string"
                                        },
                                        "productschemaDescription": {
                                            "description": "Description about the productschema",
                                            "type": "string"
                                        },
                                        "productschemaId": {
                                            "description": "A product schema ID",
                                            "type": "string"
                                        },
                                        "productschemaStatus": {
                                            "description": "New / OnHold / InUse / Archived",
                                            "type": "string"
                                        },
                                        "productschemaType": {
                                            "description": "0 for subassembly, 1 for finished product",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "productschemaId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetStateHistoryDistributionDetail": {
            "description": "Returns history states for a distributionDetail",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "distributionDetail": {
                                "properties": {
                                    "distributionDetailId": {
                                        "description": "Unique id of a dist. record",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "distributionDetail"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetStateHistoryDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of distributionDetail states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^DISTDET": {
                                "description": "A distribution record's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This distribution record's world state distribution record ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the distribution record's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The distribution record's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This distribution record has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "distributionDetail": {
                                                "description": "Distribution Checkpoints",
                                                "properties": {
                                                    "checkPointEvents": {
                                                        "description": "CheckPoint Events",
                                                        "type": "string"
                                                    },
                                                    "checkPointShock": {
                                                        "description": "Shock measured at the CheckPoint",
                                                        "type": "string"
                                                    },
                                                    "checkPointTemperature": {
                                                        "description": "Temperature measured at the checkpoint",
                                                        "type": "string"
                                                    },
                                                    "checkpointDate": {
                                                        "description": "Date-time at which the shipment is expected at the checkpoint",
                                                        "type": "string"
                                                    },
                                                    "checkpointLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "checkpointName": {
                                                        "description": "Checkpoint lat longs of distributor",
                                                        "type": "string"
                                                    },
                                                    "distributionDetailId": {
                                                        "description": "Unique id of a dist. record",
                                                        "type": "string"
                                                    },
                                                    "distributionId": {
                                                        "description": "Participant Id of order distributor",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "distributionDetailId"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "distributionDetail": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The distribution record's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this distribution record",
                                        "properties": {
                                            "distribution record": {
                                                "description": "Distribution Checkpoints",
                                                "properties": {
                                                    "checkPointEvents": {
                                                        "description": "CheckPoint Events",
                                                        "type": "string"
                                                    },
                                                    "checkPointShock": {
                                                        "description": "Shock measured at the CheckPoint",
                                                        "type": "string"
                                                    },
                                                    "checkPointTemperature": {
                                                        "description": "Temperature measured at the checkpoint",
                                                        "type": "string"
                                                    },
                                                    "checkpointDate": {
                                                        "description": "Date-time at which the shipment is expected at the checkpoint",
                                                        "type": "string"
                                                    },
                                                    "checkpointLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "checkpointName": {
                                                        "description": "Checkpoint lat longs of distributor",
                                                        "type": "string"
                                                    },
                                                    "distributionDetailId": {
                                                        "description": "Unique id of a dist. record",
                                                        "type": "string"
                                                    },
                                                    "distributionId": {
                                                        "description": "Participant Id of order distributor",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "distributionDetailId"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetStateHistoryOrder": {
            "description": "Returns history states for a order",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "order": {
                                "properties": {
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "order"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetStateHistoryOrder"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of order states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^PROD": {
                                "description": "A order's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This order's world state order ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the order's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The order's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This order has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "order": {
                                                "description": "The changeable properties for a order, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "goodsOrServicesType": {
                                                        "description": "Goods / Services",
                                                        "type": "string"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "orderContentType": {
                                                        "description": "Raw Material/ Component/Subassembly / Finished order",
                                                        "type": "string"
                                                    },
                                                    "orderDeliveryLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderDescription": {
                                                        "description": "Description of the order",
                                                        "type": "string"
                                                    },
                                                    "orderDistributor": {
                                                        "description": "Order distributor details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "bolId": {
                                                                "description": "Bill of Lading number",
                                                                "type": "string"
                                                            },
                                                            "distId": {
                                                                "description": "Participant Id of order distributor",
                                                                "type": "string"
                                                            },
                                                            "distName": {
                                                                "description": "Participant name of order distributor",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderFulfiller": {
                                                        "description": "Order fulfiller details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "Id": {
                                                                "description": "Participant Id of order fulfiller",
                                                                "type": "string"
                                                            },
                                                            "Type": {
                                                                "description": "Participant type of order fulfiller",
                                                                "type": "string"
                                                            },
                                                            "actualFulfilmentDate": {
                                                                "description": "Fulfilment date",
                                                                "type": "string"
                                                            },
                                                            "batchIds": {
                                                                "description": "batch ids used to fill order",
                                                                "type": "string"
                                                            },
                                                            "batchSizeDelivered": {
                                                                "description": "Number of assets",
                                                                "type": "number"
                                                            },
                                                            "committedDeliveryDate": {
                                                                "description": "Order due date",
                                                                "type": "string"
                                                            },
                                                            "orderReceivedDate": {
                                                                "description": "Date the order was received",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderFulfilmentLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "orderIssuer": {
                                                        "description": "Order issuer details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "batchSizeRequired": {
                                                                "description": "Number of assets",
                                                                "type": "number"
                                                            },
                                                            "oderIssuerId": {
                                                                "description": "Participant Id of order issuer",
                                                                "type": "string"
                                                            },
                                                            "oderIssuerType": {
                                                                "description": "Participant type of order issuer",
                                                                "type": "string"
                                                            },
                                                            "orderDueDate": {
                                                                "description": "Order due date",
                                                                "type": "string"
                                                            },
                                                            "orderIssueDate": {
                                                                "description": "Year of Manufacturing",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderStatus": {
                                                        "description": "The current status of the order",
                                                        "type": "string"
                                                    },
                                                    "orderType": {
                                                        "description": "Manufacturing / Distribution/MRO",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "orderId",
                                                    "industryType",
                                                    "orderType",
                                                    "orderContentType",
                                                    "orderStatus",
                                                    "orderDescription",
                                                    "orderFulfilmentLocation",
                                                    "orderDeliveryLocation"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "order": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The order's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this order",
                                        "properties": {
                                            "order": {
                                                "description": "The changeable properties for a order, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "goodsOrServicesType": {
                                                        "description": "Goods / Services",
                                                        "type": "string"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "orderContentType": {
                                                        "description": "Raw Material/ Component/Subassembly / Finished order",
                                                        "type": "string"
                                                    },
                                                    "orderDeliveryLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderDescription": {
                                                        "description": "Description of the order",
                                                        "type": "string"
                                                    },
                                                    "orderDistributor": {
                                                        "description": "Order distributor details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "bolId": {
                                                                "description": "Bill of Lading number",
                                                                "type": "string"
                                                            },
                                                            "distId": {
                                                                "description": "Participant Id of order distributor",
                                                                "type": "string"
                                                            },
                                                            "distName": {
                                                                "description": "Participant name of order distributor",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderFulfiller": {
                                                        "description": "Order fulfiller details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "Id": {
                                                                "description": "Participant Id of order fulfiller",
                                                                "type": "string"
                                                            },
                                                            "Type": {
                                                                "description": "Participant type of order fulfiller",
                                                                "type": "string"
                                                            },
                                                            "actualFulfilmentDate": {
                                                                "description": "Fulfilment date",
                                                                "type": "string"
                                                            },
                                                            "batchIds": {
                                                                "description": "batch ids used to fill order",
                                                                "type": "string"
                                                            },
                                                            "batchSizeDelivered": {
                                                                "description": "Number of assets",
                                                                "type": "number"
                                                            },
                                                            "committedDeliveryDate": {
                                                                "description": "Order due date",
                                                                "type": "string"
                                                            },
                                                            "orderReceivedDate": {
                                                                "description": "Date the order was received",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderFulfilmentLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "orderIssuer": {
                                                        "description": "Order issuer details",
                                                        "minItems": 1,
                                                        "properties": {
                                                            "batchSizeRequired": {
                                                                "description": "Number of assets",
                                                                "type": "number"
                                                            },
                                                            "oderIssuerId": {
                                                                "description": "Participant Id of order issuer",
                                                                "type": "string"
                                                            },
                                                            "oderIssuerType": {
                                                                "description": "Participant type of order issuer",
                                                                "type": "string"
                                                            },
                                                            "orderDueDate": {
                                                                "description": "Order due date",
                                                                "type": "string"
                                                            },
                                                            "orderIssueDate": {
                                                                "description": "Year of Manufacturing",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "orderStatus": {
                                                        "description": "The current status of the order",
                                                        "type": "string"
                                                    },
                                                    "orderType": {
                                                        "description": "Manufacturing / Distribution/MRO",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "orderId",
                                                    "industryType",
                                                    "orderType",
                                                    "orderContentType",
                                                    "orderStatus",
                                                    "orderDescription",
                                                    "orderFulfilmentLocation",
                                                    "orderDeliveryLocation"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetStateHistoryProduct": {
            "description": "Returns history states for a product",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "product": {
                                "properties": {
                                    "productId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "product"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetStateHistoryProduct"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of product states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^PROD": {
                                "description": "A product's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This product's world state product ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the product's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The product's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This product has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "product": {
                                                "description": "The changeable properties for a product, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "madeOf": {
                                                        "items": {
                                                            "properties": {
                                                                "assetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product",
                                                                    "type": "string"
                                                                },
                                                                "assetCount": {
                                                                    "description": "Number of items in the product",
                                                                    "type": "number"
                                                                },
                                                                "assetDescription": {
                                                                    "description": "Description of the supply/supplies",
                                                                    "type": "string"
                                                                },
                                                                "assetIds": {
                                                                    "description": "Comma separated list of actual supplies",
                                                                    "type": "string"
                                                                },
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "manufacturingDate": {
                                                        "description": "Date of Manufacturing",
                                                        "type": "string"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "productBatchId": {
                                                        "description": "Batch id to which the product belongs",
                                                        "type": "string"
                                                    },
                                                    "productDescription": {
                                                        "description": "Description of the finished product",
                                                        "type": "string"
                                                    },
                                                    "productId": {
                                                        "description": "A product's ID",
                                                        "type": "string"
                                                    },
                                                    "productInternalId": {
                                                        "description": "A product's Internal ID",
                                                        "type": "string"
                                                    },
                                                    "productModel": {
                                                        "description": "Product Model",
                                                        "type": "string"
                                                    },
                                                    "productRecalls": {
                                                        "items": {
                                                            "properties": {
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                },
                                                                "newassetIds": {
                                                                    "description": "Asset Ids of supply / product in recall",
                                                                    "type": "string"
                                                                },
                                                                "originalassetIds": {
                                                                    "description": "Asset Ids of supply / product in original schema",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "productSchemaId": {
                                                        "description": "A product schema ID",
                                                        "type": "string"
                                                    },
                                                    "productStatus": {
                                                        "description": "The current status of the product",
                                                        "type": "string"
                                                    },
                                                    "productType": {
                                                        "description": "FP(FinishedProduct/Subassembly",
                                                        "type": "string"
                                                    },
                                                    "productionLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "sellerId": {
                                                        "description": "A seller's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "productId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "product": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The product's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this product",
                                        "properties": {
                                            "product": {
                                                "description": "The changeable properties for a product, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "madeOf": {
                                                        "items": {
                                                            "properties": {
                                                                "assetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product",
                                                                    "type": "string"
                                                                },
                                                                "assetCount": {
                                                                    "description": "Number of items in the product",
                                                                    "type": "number"
                                                                },
                                                                "assetDescription": {
                                                                    "description": "Description of the supply/supplies",
                                                                    "type": "string"
                                                                },
                                                                "assetIds": {
                                                                    "description": "Comma separated list of actual supplies",
                                                                    "type": "string"
                                                                },
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "manufacturingDate": {
                                                        "description": "Date of Manufacturing",
                                                        "type": "string"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "productBatchId": {
                                                        "description": "Batch id to which the product belongs",
                                                        "type": "string"
                                                    },
                                                    "productDescription": {
                                                        "description": "Description of the finished product",
                                                        "type": "string"
                                                    },
                                                    "productId": {
                                                        "description": "A product's ID",
                                                        "type": "string"
                                                    },
                                                    "productInternalId": {
                                                        "description": "A product's Internal ID",
                                                        "type": "string"
                                                    },
                                                    "productModel": {
                                                        "description": "Product Model",
                                                        "type": "string"
                                                    },
                                                    "productRecalls": {
                                                        "items": {
                                                            "properties": {
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                },
                                                                "newassetIds": {
                                                                    "description": "Asset Ids of supply / product in recall",
                                                                    "type": "string"
                                                                },
                                                                "originalassetIds": {
                                                                    "description": "Asset Ids of supply / product in original schema",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "productSchemaId": {
                                                        "description": "A product schema ID",
                                                        "type": "string"
                                                    },
                                                    "productStatus": {
                                                        "description": "The current status of the product",
                                                        "type": "string"
                                                    },
                                                    "productType": {
                                                        "description": "FP(FinishedProduct/Subassembly",
                                                        "type": "string"
                                                    },
                                                    "productionLocation": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "sellerId": {
                                                        "description": "A seller's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "productId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetStateHistoryProductSchema": {
            "description": "Returns history states for a productschema",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "productschema": {
                                "properties": {
                                    "productschemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "productschema"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetStateHistoryProductSchema"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of productschema states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^SUP": {
                                "description": "A productschema's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This productschema's world state productschema ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the productschema's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The productschema's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This productschema has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "productschema": {
                                                "description": "The changeable properties for a productschema, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "productSchemaChange": {
                                                        "items": {
                                                            "properties": {
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                },
                                                                "newassetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product in recall",
                                                                    "type": "string"
                                                                },
                                                                "originalassetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product in original schema",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "productSchemaContent": {
                                                        "description": "Schema content",
                                                        "type": "string"
                                                    },
                                                    "productcode": {
                                                        "description": "Maker's code for this product type",
                                                        "type": "string"
                                                    },
                                                    "productschemaCreationDate": {
                                                        "description": "Date of schema definition",
                                                        "type": "string"
                                                    },
                                                    "productschemaDescription": {
                                                        "description": "Description about the productschema",
                                                        "type": "string"
                                                    },
                                                    "productschemaId": {
                                                        "description": "A product schema ID",
                                                        "type": "string"
                                                    },
                                                    "productschemaStatus": {
                                                        "description": "New / OnHold / InUse / Archived",
                                                        "type": "string"
                                                    },
                                                    "productschemaType": {
                                                        "description": "0 for subassembly, 1 for finished product",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "productschemaId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "productschema": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The productschema's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this productschema",
                                        "properties": {
                                            "productschema": {
                                                "description": "The changeable properties for a productschema, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "productSchemaChange": {
                                                        "items": {
                                                            "properties": {
                                                                "assetType": {
                                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                                    "type": "string"
                                                                },
                                                                "newassetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product in recall",
                                                                    "type": "string"
                                                                },
                                                                "originalassetBatchId": {
                                                                    "description": "Batch Id of supply / productModelId of product in original schema",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "productSchemaContent": {
                                                        "description": "Schema content",
                                                        "type": "string"
                                                    },
                                                    "productcode": {
                                                        "description": "Maker's code for this product type",
                                                        "type": "string"
                                                    },
                                                    "productschemaCreationDate": {
                                                        "description": "Date of schema definition",
                                                        "type": "string"
                                                    },
                                                    "productschemaDescription": {
                                                        "description": "Description about the productschema",
                                                        "type": "string"
                                                    },
                                                    "productschemaId": {
                                                        "description": "A product schema ID",
                                                        "type": "string"
                                                    },
                                                    "productschemaStatus": {
                                                        "description": "New / OnHold / InUse / Archived",
                                                        "type": "string"
                                                    },
                                                    "productschemaType": {
                                                        "description": "0 for subassembly, 1 for finished product",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "productschemaId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetStateHistorySupply": {
            "description": "Returns history states for a supply",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "supply": {
                                "properties": {
                                    "supplyId": {
                                        "description": "A supply's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "supply"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetStateHistorySupply"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of supply states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^SUP": {
                                "description": "A supply's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This supply's world state supply ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the supply's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The supply's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This supply has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "supply": {
                                                "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "makerproductId": {
                                                        "description": "A product's ID",
                                                        "type": "string"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "shippedDate": {
                                                        "description": "Date of shipment to maker",
                                                        "type": "string"
                                                    },
                                                    "supplierId": {
                                                        "description": "A supplier's ID",
                                                        "type": "string"
                                                    },
                                                    "supplyAggregationLevel": {
                                                        "description": "Supply type - RawMaterial / Component / Operational",
                                                        "type": "string"
                                                    },
                                                    "supplyAvailableDate": {
                                                        "description": "Date of manufacture",
                                                        "type": "string"
                                                    },
                                                    "supplyBatchId": {
                                                        "description": "Supplier's production batch under which this supply falls",
                                                        "type": "string"
                                                    },
                                                    "supplyCount": {
                                                        "description": "Number of assets",
                                                        "type": "number"
                                                    },
                                                    "supplyDescription": {
                                                        "description": "Description about the supply",
                                                        "type": "string"
                                                    },
                                                    "supplyId": {
                                                        "description": "A supply's ID",
                                                        "type": "string"
                                                    },
                                                    "supplyStatus": {
                                                        "description": "The current status of the supply",
                                                        "type": "string"
                                                    },
                                                    "supplyType": {
                                                        "description": "eg.LightingCircuits, EngineSubasembly etc ",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "supplyId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "supply": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The supply's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this supply",
                                        "properties": {
                                            "supply": {
                                                "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "common": {
                                                        "description": "Common properties for all assets",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "deviceID": {
                                                                "description": "A unique identifier for the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "industryType": {
                                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                                        "type": "string"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "makerproductId": {
                                                        "description": "A product's ID",
                                                        "type": "string"
                                                    },
                                                    "orderId": {
                                                        "description": "An order's ID",
                                                        "type": "string"
                                                    },
                                                    "shippedDate": {
                                                        "description": "Date of shipment to maker",
                                                        "type": "string"
                                                    },
                                                    "supplierId": {
                                                        "description": "A supplier's ID",
                                                        "type": "string"
                                                    },
                                                    "supplyAggregationLevel": {
                                                        "description": "Supply type - RawMaterial / Component / Operational",
                                                        "type": "string"
                                                    },
                                                    "supplyAvailableDate": {
                                                        "description": "Date of manufacture",
                                                        "type": "string"
                                                    },
                                                    "supplyBatchId": {
                                                        "description": "Supplier's production batch under which this supply falls",
                                                        "type": "string"
                                                    },
                                                    "supplyCount": {
                                                        "description": "Number of assets",
                                                        "type": "number"
                                                    },
                                                    "supplyDescription": {
                                                        "description": "Description about the supply",
                                                        "type": "string"
                                                    },
                                                    "supplyId": {
                                                        "description": "A supply's ID",
                                                        "type": "string"
                                                    },
                                                    "supplyStatus": {
                                                        "description": "The current status of the supply",
                                                        "type": "string"
                                                    },
                                                    "supplyType": {
                                                        "description": "eg.LightingCircuits, EngineSubasembly etc ",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "supplyId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetSupply": {
            "description": "Returns the state a supply",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supply": {
                                "properties": {
                                    "supplyId": {
                                        "description": "A supply's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetSupply"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A supply's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This supply's world state supply ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "An array of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the supply's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The supply's asset class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This supply has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAssetContainer",
                            "properties": {
                                "supply": {
                                    "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
                                    "properties": {
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "industryType": {
                                            "description": "Industry type - to start with Auto / Pharma / Food",
                                            "type": "string"
                                        },
                                        "makerId": {
                                            "description": "A maker's ID",
                                            "type": "string"
                                        },
                                        "makerproductId": {
                                            "description": "A product's ID",
                                            "type": "string"
                                        },
                                        "orderId": {
                                            "description": "An order's ID",
                                            "type": "string"
                                        },
                                        "shippedDate": {
                                            "description": "Date of shipment to maker",
                                            "type": "string"
                                        },
                                        "supplierId": {
                                            "description": "A supplier's ID",
                                            "type": "string"
                                        },
                                        "supplyAggregationLevel": {
                                            "description": "Supply type - RawMaterial / Component / Operational",
                                            "type": "string"
                                        },
                                        "supplyAvailableDate": {
                                            "description": "Date of manufacture",
                                            "type": "string"
                                        },
                                        "supplyBatchId": {
                                            "description": "Supplier's production batch under which this supply falls",
                                            "type": "string"
                                        },
                                        "supplyCount": {
                                            "description": "Number of assets",
                                            "type": "number"
                                        },
                                        "supplyDescription": {
                                            "description": "Description about the supply",
                                            "type": "string"
                                        },
                                        "supplyId": {
                                            "description": "A supply's ID",
                                            "type": "string"
                                        },
                                        "supplyStatus": {
                                            "description": "The current status of the supply",
                                            "type": "string"
                                        },
                                        "supplyType": {
                                            "description": "eg.LightingCircuits, EngineSubasembly etc ",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "supplyId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "supply": {
                                    "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                    "properties": {
                                        "name": {
                                            "default": "EVT.IOTCP.INVOKE.RESULT",
                                            "enum": [
                                                "EVT.IOTCP.INVOKE.RESULT"
                                            ],
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A map of contributed results",
                                            "properties": {
                                                "description": "the overall status of the invoke result, defined by err",
                                                "properties": {
                                                    "activeAlerts": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsCleared": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsRaised": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "invokeresult": {
                                                        "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                        "properties": {
                                                            "message": {
                                                                "type": "string"
                                                            },
                                                            "status": {
                                                                "enum": [
                                                                    "OK",
                                                                    "ERROR"
                                                                ],
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The supply's asset class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this supply",
                            "properties": {
                                "supply": {
                                    "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
                                    "properties": {
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "industryType": {
                                            "description": "Industry type - to start with Auto / Pharma / Food",
                                            "type": "string"
                                        },
                                        "makerId": {
                                            "description": "A maker's ID",
                                            "type": "string"
                                        },
                                        "makerproductId": {
                                            "description": "A product's ID",
                                            "type": "string"
                                        },
                                        "orderId": {
                                            "description": "An order's ID",
                                            "type": "string"
                                        },
                                        "shippedDate": {
                                            "description": "Date of shipment to maker",
                                            "type": "string"
                                        },
                                        "supplierId": {
                                            "description": "A supplier's ID",
                                            "type": "string"
                                        },
                                        "supplyAggregationLevel": {
                                            "description": "Supply type - RawMaterial / Component / Operational",
                                            "type": "string"
                                        },
                                        "supplyAvailableDate": {
                                            "description": "Date of manufacture",
                                            "type": "string"
                                        },
                                        "supplyBatchId": {
                                            "description": "Supplier's production batch under which this supply falls",
                                            "type": "string"
                                        },
                                        "supplyCount": {
                                            "description": "Number of assets",
                                            "type": "number"
                                        },
                                        "supplyDescription": {
                                            "description": "Description about the supply",
                                            "type": "string"
                                        },
                                        "supplyId": {
                                            "description": "A supply's ID",
                                            "type": "string"
                                        },
                                        "supplyStatus": {
                                            "description": "The current status of the supply",
                                            "type": "string"
                                        },
                                        "supplyType": {
                                            "description": "eg.LightingCircuits, EngineSubasembly etc ",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "supplyId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readParticipantMaker": {
            "description": "Returns the state a maker",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "maker": {
                                "properties": {
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readParticipantMaker"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A maker's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This maker's world state maker ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "An array of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the maker's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The maker's participant class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This maker has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateParticipantMaker",
                            "properties": {
                                "maker": {
                                    "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
                                    "properties": {
                                        "industryType": {
                                            "description": "Industry, to start off with is auto, pharma or food",
                                            "type": "string"
                                        },
                                        "makerAddress": {
                                            "description": "The makers' physical address and other means of contacting",
                                            "properties": {
                                                "city": {
                                                    "type": "string"
                                                },
                                                "country": {
                                                    "type": "string"
                                                },
                                                "name": {
                                                    "type": "string"
                                                },
                                                "phone": {
                                                    "type": "string"
                                                },
                                                "street": {
                                                    "type": "string"
                                                },
                                                "website": {
                                                    "type": "string"
                                                },
                                                "zipcode": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "makerId": {
                                            "description": "A maker's ID",
                                            "type": "string"
                                        },
                                        "sellers": {
                                            "items": {
                                                "properties": {
                                                    "seller": {
                                                        "properties": {
                                                            "sellerId": {
                                                                "description": "A seller's ID",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "suppliers": {
                                            "items": {
                                                "properties": {
                                                    "supplier": {
                                                        "properties": {
                                                            "supplierId": {
                                                                "description": "A supplier's ID",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        }
                                    },
                                    "required": [
                                        "makerId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "maker": {
                                    "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                    "properties": {
                                        "name": {
                                            "default": "EVT.IOTCP.INVOKE.RESULT",
                                            "enum": [
                                                "EVT.IOTCP.INVOKE.RESULT"
                                            ],
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A map of contributed results",
                                            "properties": {
                                                "description": "the overall status of the invoke result, defined by err",
                                                "properties": {
                                                    "activeAlerts": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsCleared": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsRaised": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "invokeresult": {
                                                        "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                        "properties": {
                                                            "message": {
                                                                "type": "string"
                                                            },
                                                            "status": {
                                                                "enum": [
                                                                    "OK",
                                                                    "ERROR"
                                                                ],
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The maker's participant class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this maker",
                            "properties": {
                                "maker": {
                                    "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
                                    "properties": {
                                        "industryType": {
                                            "description": "Industry, to start off with is auto, pharma or food",
                                            "type": "string"
                                        },
                                        "makerAddress": {
                                            "description": "The makers' physical address and other means of contacting",
                                            "properties": {
                                                "city": {
                                                    "type": "string"
                                                },
                                                "country": {
                                                    "type": "string"
                                                },
                                                "name": {
                                                    "type": "string"
                                                },
                                                "phone": {
                                                    "type": "string"
                                                },
                                                "street": {
                                                    "type": "string"
                                                },
                                                "website": {
                                                    "type": "string"
                                                },
                                                "zipcode": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "makerId": {
                                            "description": "A maker's ID",
                                            "type": "string"
                                        },
                                        "sellers": {
                                            "items": {
                                                "properties": {
                                                    "seller": {
                                                        "properties": {
                                                            "sellerId": {
                                                                "description": "A seller's ID",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "suppliers": {
                                            "items": {
                                                "properties": {
                                                    "supplier": {
                                                        "properties": {
                                                            "supplierId": {
                                                                "description": "A supplier's ID",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        }
                                    },
                                    "required": [
                                        "makerId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readParticipantSeller": {
            "description": "Returns the state a seller",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "seller": {
                                "properties": {
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readParticipantSeller"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A seller's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This seller's world state seller ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "An array of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the seller's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The seller's participant class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This seller has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateParticipantSeller",
                            "properties": {
                                "seller": {
                                    "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
                                    "properties": {
                                        "industryType": {
                                            "description": "Industry, to start off with is auto, pharma or food",
                                            "type": "string"
                                        },
                                        "makers": {
                                            "items": {
                                                "properties": {
                                                    "maker": {
                                                        "properties": {
                                                            "makerId": {
                                                                "description": "A maker's ID",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "sellerAddress": {
                                            "description": "The sellers' physical address and other means of contacting",
                                            "properties": {
                                                "city": {
                                                    "type": "string"
                                                },
                                                "country": {
                                                    "type": "string"
                                                },
                                                "name": {
                                                    "type": "string"
                                                },
                                                "phone": {
                                                    "type": "string"
                                                },
                                                "street": {
                                                    "type": "string"
                                                },
                                                "website": {
                                                    "type": "string"
                                                },
                                                "zipcode": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "sellerId": {
                                            "description": "A seller's ID",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "sellerId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "seller": {
                                    "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                    "properties": {
                                        "name": {
                                            "default": "EVT.IOTCP.INVOKE.RESULT",
                                            "enum": [
                                                "EVT.IOTCP.INVOKE.RESULT"
                                            ],
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A map of contributed results",
                                            "properties": {
                                                "description": "the overall status of the invoke result, defined by err",
                                                "properties": {
                                                    "activeAlerts": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsCleared": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsRaised": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "invokeresult": {
                                                        "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                        "properties": {
                                                            "message": {
                                                                "type": "string"
                                                            },
                                                            "status": {
                                                                "enum": [
                                                                    "OK",
                                                                    "ERROR"
                                                                ],
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The seller's participant class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this seller",
                            "properties": {
                                "seller": {
                                    "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
                                    "properties": {
                                        "industryType": {
                                            "description": "Industry, to start off with is auto, pharma or food",
                                            "type": "string"
                                        },
                                        "makers": {
                                            "items": {
                                                "properties": {
                                                    "maker": {
                                                        "properties": {
                                                            "makerId": {
                                                                "description": "A maker's ID",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "sellerAddress": {
                                            "description": "The sellers' physical address and other means of contacting",
                                            "properties": {
                                                "city": {
                                                    "type": "string"
                                                },
                                                "country": {
                                                    "type": "string"
                                                },
                                                "name": {
                                                    "type": "string"
                                                },
                                                "phone": {
                                                    "type": "string"
                                                },
                                                "street": {
                                                    "type": "string"
                                                },
                                                "website": {
                                                    "type": "string"
                                                },
                                                "zipcode": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "sellerId": {
                                            "description": "A seller's ID",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "sellerId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readParticipantStateHistoryMaker": {
            "description": "Returns history states for a maker",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "maker": {
                                "properties": {
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "maker"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readParticipantStateHistoryMaker"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of maker states, can mix participant classes",
                    "items": {
                        "patternProperties": {
                            "^SUPR": {
                                "description": "A maker's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This maker's world state maker ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the maker's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The maker's participant class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This maker has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateParticipantMaker",
                                        "properties": {
                                            "maker": {
                                                "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makerAddress": {
                                                        "description": "The makers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "sellers": {
                                                        "items": {
                                                            "properties": {
                                                                "seller": {
                                                                    "properties": {
                                                                        "sellerId": {
                                                                            "description": "A seller's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "suppliers": {
                                                        "items": {
                                                            "properties": {
                                                                "supplier": {
                                                                    "properties": {
                                                                        "supplierId": {
                                                                            "description": "A supplier's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    }
                                                },
                                                "required": [
                                                    "makerId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "maker": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The maker's participant class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this maker",
                                        "properties": {
                                            "maker": {
                                                "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makerAddress": {
                                                        "description": "The makers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "makerId": {
                                                        "description": "A maker's ID",
                                                        "type": "string"
                                                    },
                                                    "sellers": {
                                                        "items": {
                                                            "properties": {
                                                                "seller": {
                                                                    "properties": {
                                                                        "sellerId": {
                                                                            "description": "A seller's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "suppliers": {
                                                        "items": {
                                                            "properties": {
                                                                "supplier": {
                                                                    "properties": {
                                                                        "supplierId": {
                                                                            "description": "A supplier's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    }
                                                },
                                                "required": [
                                                    "makerId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readParticipantStateHistorySeller": {
            "description": "Returns history states for a seller",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "seller": {
                                "properties": {
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "seller"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readParticipantStateHistorySeller"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of seller states, can mix participant classes",
                    "items": {
                        "patternProperties": {
                            "^SUPR": {
                                "description": "A seller's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This seller's world state seller ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the seller's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The seller's participant class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This seller has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateParticipantSeller",
                                        "properties": {
                                            "seller": {
                                                "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makers": {
                                                        "items": {
                                                            "properties": {
                                                                "maker": {
                                                                    "properties": {
                                                                        "makerId": {
                                                                            "description": "A maker's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "sellerAddress": {
                                                        "description": "The sellers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "sellerId": {
                                                        "description": "A seller's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "sellerId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "seller": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The seller's participant class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this seller",
                                        "properties": {
                                            "seller": {
                                                "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makers": {
                                                        "items": {
                                                            "properties": {
                                                                "maker": {
                                                                    "properties": {
                                                                        "makerId": {
                                                                            "description": "A maker's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "sellerAddress": {
                                                        "description": "The sellers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "sellerId": {
                                                        "description": "A seller's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "sellerId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readParticipantStateHistorySupplier": {
            "description": "Returns history states for a supplier",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "supplier": {
                                "properties": {
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "supplier"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readParticipantStateHistorySupplier"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of supplier states, can mix participant classes",
                    "items": {
                        "patternProperties": {
                            "^SUPR": {
                                "description": "A supplier's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This supplier's world state supplier ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "An array of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the supplier's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The supplier's participant class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This supplier has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateParticipantSupplier",
                                        "properties": {
                                            "supplier": {
                                                "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makers": {
                                                        "items": {
                                                            "properties": {
                                                                "maker": {
                                                                    "properties": {
                                                                        "makerId": {
                                                                            "description": "A maker's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "supplierAddress": {
                                                        "description": "The suppliers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "supplierId": {
                                                        "description": "A supplier's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "supplierId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "supplier": {
                                                "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                                "properties": {
                                                    "name": {
                                                        "default": "EVT.IOTCP.INVOKE.RESULT",
                                                        "enum": [
                                                            "EVT.IOTCP.INVOKE.RESULT"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A map of contributed results",
                                                        "properties": {
                                                            "description": "the overall status of the invoke result, defined by err",
                                                            "properties": {
                                                                "activeAlerts": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsCleared": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "alertsRaised": {
                                                                    "description": "An array of alert names",
                                                                    "items": {
                                                                        "description": "An alert name",
                                                                        "type": "string"
                                                                    },
                                                                    "type": "array"
                                                                },
                                                                "invokeresult": {
                                                                    "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                                    "properties": {
                                                                        "message": {
                                                                            "type": "string"
                                                                        },
                                                                        "status": {
                                                                            "enum": [
                                                                                "OK",
                                                                                "ERROR"
                                                                            ],
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The supplier's participant class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this supplier",
                                        "properties": {
                                            "supplier": {
                                                "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "industryType": {
                                                        "description": "Industry, to start off with is auto, pharma or food",
                                                        "type": "string"
                                                    },
                                                    "makers": {
                                                        "items": {
                                                            "properties": {
                                                                "maker": {
                                                                    "properties": {
                                                                        "makerId": {
                                                                            "description": "A maker's ID",
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 1,
                                                        "type": "array"
                                                    },
                                                    "supplierAddress": {
                                                        "description": "The suppliers' physical address and other means of contacting",
                                                        "properties": {
                                                            "city": {
                                                                "type": "string"
                                                            },
                                                            "country": {
                                                                "type": "string"
                                                            },
                                                            "name": {
                                                                "type": "string"
                                                            },
                                                            "phone": {
                                                                "type": "string"
                                                            },
                                                            "street": {
                                                                "type": "string"
                                                            },
                                                            "website": {
                                                                "type": "string"
                                                            },
                                                            "zipcode": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "supplierId": {
                                                        "description": "A supplier's ID",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "supplierId",
                                                    "industryType"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readParticipantSupplier": {
            "description": "Returns the state a supplier",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supplier": {
                                "properties": {
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readParticipantSupplier"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A supplier's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This supplier's world state supplier ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "An array of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the supplier's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The supplier's participant class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This supplier has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateParticipantSupplier",
                            "properties": {
                                "supplier": {
                                    "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
                                    "properties": {
                                        "industryType": {
                                            "description": "Industry, to start off with is auto, pharma or food",
                                            "type": "string"
                                        },
                                        "makers": {
                                            "items": {
                                                "properties": {
                                                    "maker": {
                                                        "properties": {
                                                            "makerId": {
                                                                "description": "A maker's ID",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "supplierAddress": {
                                            "description": "The suppliers' physical address and other means of contacting",
                                            "properties": {
                                                "city": {
                                                    "type": "string"
                                                },
                                                "country": {
                                                    "type": "string"
                                                },
                                                "name": {
                                                    "type": "string"
                                                },
                                                "phone": {
                                                    "type": "string"
                                                },
                                                "street": {
                                                    "type": "string"
                                                },
                                                "website": {
                                                    "type": "string"
                                                },
                                                "zipcode": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "supplierId": {
                                            "description": "A supplier's ID",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "supplierId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "supplier": {
                                    "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                    "properties": {
                                        "name": {
                                            "default": "EVT.IOTCP.INVOKE.RESULT",
                                            "enum": [
                                                "EVT.IOTCP.INVOKE.RESULT"
                                            ],
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A map of contributed results",
                                            "properties": {
                                                "description": "the overall status of the invoke result, defined by err",
                                                "properties": {
                                                    "activeAlerts": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsCleared": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsRaised": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "invokeresult": {
                                                        "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                        "properties": {
                                                            "message": {
                                                                "type": "string"
                                                            },
                                                            "status": {
                                                                "enum": [
                                                                    "OK",
                                                                    "ERROR"
                                                                ],
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The supplier's participant class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this supplier",
                            "properties": {
                                "supplier": {
                                    "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
                                    "properties": {
                                        "industryType": {
                                            "description": "Industry, to start off with is auto, pharma or food",
                                            "type": "string"
                                        },
                                        "makers": {
                                            "items": {
                                                "properties": {
                                                    "maker": {
                                                        "properties": {
                                                            "makerId": {
                                                                "description": "A maker's ID",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 1,
                                            "type": "array"
                                        },
                                        "supplierAddress": {
                                            "description": "The suppliers' physical address and other means of contacting",
                                            "properties": {
                                                "city": {
                                                    "type": "string"
                                                },
                                                "country": {
                                                    "type": "string"
                                                },
                                                "name": {
                                                    "type": "string"
                                                },
                                                "phone": {
                                                    "type": "string"
                                                },
                                                "street": {
                                                    "type": "string"
                                                },
                                                "website": {
                                                    "type": "string"
                                                },
                                                "zipcode": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "supplierId": {
                                            "description": "A supplier's ID",
                                            "type": "string"
                                        }
                                    },
                                    "required": [
                                        "supplierId",
                                        "industryType"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readRecentStates": {
            "description": "Returns the state of recently updated assets",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "begin": {
                                "description": "zero based beginning of range",
                                "type": "integer"
                            },
                            "end": {
                                "description": "zero based end of range, absence means to end",
                                "type": "integer"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readRecentStates"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of asset states, can mix asset classes",
                    "items": {
                        "description": "A asset's complete state",
                        "properties": {
                            "alerts": {
                                "description": "An array of alert names",
                                "items": {
                                    "description": "An alert name",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "assetID": {
                                "description": "This asset's world state asset ID",
                                "type": "string"
                            },
                            "class": {
                                "description": "An asset's classifier definition",
                                "properties": {
                                    "assetidpath": "An asset's primary key, expressed as a qualified property path (see example contracts)",
                                    "name": "An asset's class name",
                                    "prefix": "An asset's world state prefix, used to allow iteration over all assets of a class"
                                },
                                "type": "object"
                            },
                            "compliant": {
                                "description": "This asset has no active alerts",
                                "type": "boolean"
                            },
                            "eventin": {
                                "description": "The contract event that created this state, for example updateAsset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "eventout": {
                                "description": "The chaincode event emitted on invoke exit, if any",
                                "properties": {
                                    "asset": {
                                        "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                        "properties": {
                                            "name": {
                                                "default": "EVT.IOTCP.INVOKE.RESULT",
                                                "enum": [
                                                    "EVT.IOTCP.INVOKE.RESULT"
                                                ],
                                                "type": "string"
                                            },
                                            "payload": {
                                                "description": "A map of contributed results",
                                                "properties": {
                                                    "description": "the overall status of the invoke result, defined by err",
                                                    "properties": {
                                                        "activeAlerts": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "alertsCleared": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "alertsRaised": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "invokeresult": {
                                                            "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                            "properties": {
                                                                "message": {
                                                                    "type": "string"
                                                                },
                                                                "status": {
                                                                    "enum": [
                                                                        "OK",
                                                                        "ERROR"
                                                                    ],
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "state": {
                                "description": "Properties that have been received or calculated for this asset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "txnid": {
                                "description": "Transaction UUID matching the blockchain",
                                "type": "string"
                            },
                            "txnts": {
                                "description": "Transaction timestamp matching the blockchain",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readWorldState": {
            "description": "Returns the entire contents of world state",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readWorldState"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "properties": {},
                    "type": "object"
                }
            },
            "type": "object"
        },
        "replaceAssetDistributionDetail": {
            "description": "Replaces a DistributionDetail's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "distributionDetail": {
                                "description": "Distribution Checkpoints",
                                "properties": {
                                    "checkPointEvents": {
                                        "description": "CheckPoint Events",
                                        "type": "string"
                                    },
                                    "checkPointShock": {
                                        "description": "Shock measured at the CheckPoint",
                                        "type": "string"
                                    },
                                    "checkPointTemperature": {
                                        "description": "Temperature measured at the checkpoint",
                                        "type": "string"
                                    },
                                    "checkpointDate": {
                                        "description": "Date-time at which the shipment is expected at the checkpoint",
                                        "type": "string"
                                    },
                                    "checkpointLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "checkpointName": {
                                        "description": "Checkpoint lat longs of distributor",
                                        "type": "string"
                                    },
                                    "distributionDetailId": {
                                        "description": "Unique id of a dist. record",
                                        "type": "string"
                                    },
                                    "distributionId": {
                                        "description": "Participant Id of order distributor",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "distributionDetailId"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceAssetDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "replaceAssetOrder": {
            "description": "Replaces a Order's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "order": {
                                "description": "The changeable properties for a order, also considered its 'event' as a partial state",
                                "properties": {
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "goodsOrServicesType": {
                                        "description": "Goods / Services",
                                        "type": "string"
                                    },
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "orderContentType": {
                                        "description": "Raw Material/ Component/Subassembly / Finished order",
                                        "type": "string"
                                    },
                                    "orderDeliveryLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderDescription": {
                                        "description": "Description of the order",
                                        "type": "string"
                                    },
                                    "orderDistributor": {
                                        "description": "Order distributor details",
                                        "minItems": 1,
                                        "properties": {
                                            "bolId": {
                                                "description": "Bill of Lading number",
                                                "type": "string"
                                            },
                                            "distId": {
                                                "description": "Participant Id of order distributor",
                                                "type": "string"
                                            },
                                            "distName": {
                                                "description": "Participant name of order distributor",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderFulfiller": {
                                        "description": "Order fulfiller details",
                                        "minItems": 1,
                                        "properties": {
                                            "Id": {
                                                "description": "Participant Id of order fulfiller",
                                                "type": "string"
                                            },
                                            "Type": {
                                                "description": "Participant type of order fulfiller",
                                                "type": "string"
                                            },
                                            "actualFulfilmentDate": {
                                                "description": "Fulfilment date",
                                                "type": "string"
                                            },
                                            "batchIds": {
                                                "description": "batch ids used to fill order",
                                                "type": "string"
                                            },
                                            "batchSizeDelivered": {
                                                "description": "Number of assets",
                                                "type": "number"
                                            },
                                            "committedDeliveryDate": {
                                                "description": "Order due date",
                                                "type": "string"
                                            },
                                            "orderReceivedDate": {
                                                "description": "Date the order was received",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderFulfilmentLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    },
                                    "orderIssuer": {
                                        "description": "Order issuer details",
                                        "minItems": 1,
                                        "properties": {
                                            "batchSizeRequired": {
                                                "description": "Number of assets",
                                                "type": "number"
                                            },
                                            "oderIssuerId": {
                                                "description": "Participant Id of order issuer",
                                                "type": "string"
                                            },
                                            "oderIssuerType": {
                                                "description": "Participant type of order issuer",
                                                "type": "string"
                                            },
                                            "orderDueDate": {
                                                "description": "Order due date",
                                                "type": "string"
                                            },
                                            "orderIssueDate": {
                                                "description": "Year of Manufacturing",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderStatus": {
                                        "description": "The current status of the order",
                                        "type": "string"
                                    },
                                    "orderType": {
                                        "description": "Manufacturing / Distribution/MRO",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "orderId",
                                    "industryType",
                                    "orderType",
                                    "orderContentType",
                                    "orderStatus",
                                    "orderDescription",
                                    "orderFulfilmentLocation",
                                    "orderDeliveryLocation"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceAssetOrder"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "replaceAssetProduct": {
            "description": "Replaces a Product's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "product": {
                                "description": "The changeable properties for a product, also considered its 'event' as a partial state",
                                "properties": {
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "madeOf": {
                                        "items": {
                                            "properties": {
                                                "assetBatchId": {
                                                    "description": "Batch Id of supply / productModelId of product",
                                                    "type": "string"
                                                },
                                                "assetCount": {
                                                    "description": "Number of items in the product",
                                                    "type": "number"
                                                },
                                                "assetDescription": {
                                                    "description": "Description of the supply/supplies",
                                                    "type": "string"
                                                },
                                                "assetIds": {
                                                    "description": "Comma separated list of actual supplies",
                                                    "type": "string"
                                                },
                                                "assetType": {
                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "manufacturingDate": {
                                        "description": "Date of Manufacturing",
                                        "type": "string"
                                    },
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    },
                                    "productBatchId": {
                                        "description": "Batch id to which the product belongs",
                                        "type": "string"
                                    },
                                    "productDescription": {
                                        "description": "Description of the finished product",
                                        "type": "string"
                                    },
                                    "productId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    },
                                    "productInternalId": {
                                        "description": "A product's Internal ID",
                                        "type": "string"
                                    },
                                    "productModel": {
                                        "description": "Product Model",
                                        "type": "string"
                                    },
                                    "productRecalls": {
                                        "items": {
                                            "properties": {
                                                "assetType": {
                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                    "type": "string"
                                                },
                                                "newassetIds": {
                                                    "description": "Asset Ids of supply / product in recall",
                                                    "type": "string"
                                                },
                                                "originalassetIds": {
                                                    "description": "Asset Ids of supply / product in original schema",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "productSchemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    },
                                    "productStatus": {
                                        "description": "The current status of the product",
                                        "type": "string"
                                    },
                                    "productType": {
                                        "description": "FP(FinishedProduct/Subassembly",
                                        "type": "string"
                                    },
                                    "productionLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "productId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceAssetProduct"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "replaceAssetProductSchema": {
            "description": "Replaces a ProductSchema's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "productschema": {
                                "description": "The changeable properties for a productschema, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "productSchemaChange": {
                                        "items": {
                                            "properties": {
                                                "assetType": {
                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                    "type": "string"
                                                },
                                                "newassetBatchId": {
                                                    "description": "Batch Id of supply / productModelId of product in recall",
                                                    "type": "string"
                                                },
                                                "originalassetBatchId": {
                                                    "description": "Batch Id of supply / productModelId of product in original schema",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "productSchemaContent": {
                                        "description": "Schema content",
                                        "type": "string"
                                    },
                                    "productcode": {
                                        "description": "Maker's code for this product type",
                                        "type": "string"
                                    },
                                    "productschemaCreationDate": {
                                        "description": "Date of schema definition",
                                        "type": "string"
                                    },
                                    "productschemaDescription": {
                                        "description": "Description about the productschema",
                                        "type": "string"
                                    },
                                    "productschemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    },
                                    "productschemaStatus": {
                                        "description": "New / OnHold / InUse / Archived",
                                        "type": "string"
                                    },
                                    "productschemaType": {
                                        "description": "0 for subassembly, 1 for finished product",
                                        "type": "number"
                                    }
                                },
                                "required": [
                                    "productschemaId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceAssetProductSchema"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "replaceAssetSupply": {
            "description": "Replaces a Supply's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supply": {
                                "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
                                "properties": {
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "makerproductId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    },
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    },
                                    "shippedDate": {
                                        "description": "Date of shipment to maker",
                                        "type": "string"
                                    },
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    },
                                    "supplyAggregationLevel": {
                                        "description": "Supply type - RawMaterial / Component / Operational",
                                        "type": "string"
                                    },
                                    "supplyAvailableDate": {
                                        "description": "Date of manufacture",
                                        "type": "string"
                                    },
                                    "supplyBatchId": {
                                        "description": "Supplier's production batch under which this supply falls",
                                        "type": "string"
                                    },
                                    "supplyCount": {
                                        "description": "Number of assets",
                                        "type": "number"
                                    },
                                    "supplyDescription": {
                                        "description": "Description about the supply",
                                        "type": "string"
                                    },
                                    "supplyId": {
                                        "description": "A supply's ID",
                                        "type": "string"
                                    },
                                    "supplyStatus": {
                                        "description": "The current status of the supply",
                                        "type": "string"
                                    },
                                    "supplyType": {
                                        "description": "eg.LightingCircuits, EngineSubasembly etc ",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "supplyId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceAssetSupply"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "replaceParticipantMaker": {
            "description": "Replaces a maker's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "maker": {
                                "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry, to start off with is auto, pharma or food",
                                        "type": "string"
                                    },
                                    "makerAddress": {
                                        "description": "The makers' physical address and other means of contacting",
                                        "properties": {
                                            "city": {
                                                "type": "string"
                                            },
                                            "country": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            },
                                            "phone": {
                                                "type": "string"
                                            },
                                            "street": {
                                                "type": "string"
                                            },
                                            "website": {
                                                "type": "string"
                                            },
                                            "zipcode": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "sellers": {
                                        "items": {
                                            "properties": {
                                                "seller": {
                                                    "properties": {
                                                        "sellerId": {
                                                            "description": "A seller's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "suppliers": {
                                        "items": {
                                            "properties": {
                                                "supplier": {
                                                    "properties": {
                                                        "supplierId": {
                                                            "description": "A supplier's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    }
                                },
                                "required": [
                                    "makerId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceParticipantMaker"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "replaceParticipantSeller": {
            "description": "Replaces a seller's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "seller": {
                                "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry, to start off with is auto, pharma or food",
                                        "type": "string"
                                    },
                                    "makers": {
                                        "items": {
                                            "properties": {
                                                "maker": {
                                                    "properties": {
                                                        "makerId": {
                                                            "description": "A maker's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "sellerAddress": {
                                        "description": "The sellers' physical address and other means of contacting",
                                        "properties": {
                                            "city": {
                                                "type": "string"
                                            },
                                            "country": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            },
                                            "phone": {
                                                "type": "string"
                                            },
                                            "street": {
                                                "type": "string"
                                            },
                                            "website": {
                                                "type": "string"
                                            },
                                            "zipcode": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "sellerId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceParticipantSeller"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "replaceParticipantSupplier": {
            "description": "Replaces a supplier's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supplier": {
                                "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry, to start off with is auto, pharma or food",
                                        "type": "string"
                                    },
                                    "makers": {
                                        "items": {
                                            "properties": {
                                                "maker": {
                                                    "properties": {
                                                        "makerId": {
                                                            "description": "A maker's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "supplierAddress": {
                                        "description": "The suppliers' physical address and other means of contacting",
                                        "properties": {
                                            "city": {
                                                "type": "string"
                                            },
                                            "country": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            },
                                            "phone": {
                                                "type": "string"
                                            },
                                            "street": {
                                                "type": "string"
                                            },
                                            "website": {
                                                "type": "string"
                                            },
                                            "zipcode": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "supplierId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceParticipantSupplier"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setCreateOnFirstUpdate": {
            "description": "Allow updateAsset to create an asset upon receipt of its first event",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "setCreateOnFirstUpdate": {
                                "description": "Allows updates to create missing assets on first event",
                                "type": "boolean"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "setCreateOnFirstUpdate"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setLoggingLevel": {
            "description": "Sets the logging level for the contract",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "logLevel": {
                                "enum": [
                                    "CRITICAL",
                                    "ERROR",
                                    "WARNING",
                                    "NOTICE",
                                    "INFO",
                                    "DEBUG"
                                ],
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "setLoggingLevel"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAssetDistributionDetail": {
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "distributionDetail": {
                                "description": "Distribution Checkpoints",
                                "properties": {
                                    "checkPointEvents": {
                                        "description": "CheckPoint Events",
                                        "type": "string"
                                    },
                                    "checkPointShock": {
                                        "description": "Shock measured at the CheckPoint",
                                        "type": "string"
                                    },
                                    "checkPointTemperature": {
                                        "description": "Temperature measured at the checkpoint",
                                        "type": "string"
                                    },
                                    "checkpointDate": {
                                        "description": "Date-time at which the shipment is expected at the checkpoint",
                                        "type": "string"
                                    },
                                    "checkpointLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "checkpointName": {
                                        "description": "Checkpoint lat longs of distributor",
                                        "type": "string"
                                    },
                                    "distributionDetailId": {
                                        "description": "Unique id of a dist. record",
                                        "type": "string"
                                    },
                                    "distributionId": {
                                        "description": "Participant Id of order distributor",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "distributionDetailId"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateAssetDistributionDetail"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAssetOrder": {
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "order": {
                                "description": "The changeable properties for a order, also considered its 'event' as a partial state",
                                "properties": {
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "goodsOrServicesType": {
                                        "description": "Goods / Services",
                                        "type": "string"
                                    },
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "orderContentType": {
                                        "description": "Raw Material/ Component/Subassembly / Finished order",
                                        "type": "string"
                                    },
                                    "orderDeliveryLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderDescription": {
                                        "description": "Description of the order",
                                        "type": "string"
                                    },
                                    "orderDistributor": {
                                        "description": "Order distributor details",
                                        "minItems": 1,
                                        "properties": {
                                            "bolId": {
                                                "description": "Bill of Lading number",
                                                "type": "string"
                                            },
                                            "distId": {
                                                "description": "Participant Id of order distributor",
                                                "type": "string"
                                            },
                                            "distName": {
                                                "description": "Participant name of order distributor",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderFulfiller": {
                                        "description": "Order fulfiller details",
                                        "minItems": 1,
                                        "properties": {
                                            "Id": {
                                                "description": "Participant Id of order fulfiller",
                                                "type": "string"
                                            },
                                            "Type": {
                                                "description": "Participant type of order fulfiller",
                                                "type": "string"
                                            },
                                            "actualFulfilmentDate": {
                                                "description": "Fulfilment date",
                                                "type": "string"
                                            },
                                            "batchIds": {
                                                "description": "batch ids used to fill order",
                                                "type": "string"
                                            },
                                            "batchSizeDelivered": {
                                                "description": "Number of assets",
                                                "type": "number"
                                            },
                                            "committedDeliveryDate": {
                                                "description": "Order due date",
                                                "type": "string"
                                            },
                                            "orderReceivedDate": {
                                                "description": "Date the order was received",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderFulfilmentLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    },
                                    "orderIssuer": {
                                        "description": "Order issuer details",
                                        "minItems": 1,
                                        "properties": {
                                            "batchSizeRequired": {
                                                "description": "Number of assets",
                                                "type": "number"
                                            },
                                            "oderIssuerId": {
                                                "description": "Participant Id of order issuer",
                                                "type": "string"
                                            },
                                            "oderIssuerType": {
                                                "description": "Participant type of order issuer",
                                                "type": "string"
                                            },
                                            "orderDueDate": {
                                                "description": "Order due date",
                                                "type": "string"
                                            },
                                            "orderIssueDate": {
                                                "description": "Year of Manufacturing",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "orderStatus": {
                                        "description": "The current status of the order",
                                        "type": "string"
                                    },
                                    "orderType": {
                                        "description": "Manufacturing / Distribution/MRO",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "orderId",
                                    "industryType",
                                    "orderType",
                                    "orderContentType",
                                    "orderStatus",
                                    "orderDescription",
                                    "orderFulfilmentLocation",
                                    "orderDeliveryLocation"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateAssetOrder"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAssetProduct": {
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "product": {
                                "description": "The changeable properties for a product, also considered its 'event' as a partial state",
                                "properties": {
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "madeOf": {
                                        "items": {
                                            "properties": {
                                                "assetBatchId": {
                                                    "description": "Batch Id of supply / productModelId of product",
                                                    "type": "string"
                                                },
                                                "assetCount": {
                                                    "description": "Number of items in the product",
                                                    "type": "number"
                                                },
                                                "assetDescription": {
                                                    "description": "Description of the supply/supplies",
                                                    "type": "string"
                                                },
                                                "assetIds": {
                                                    "description": "Comma separated list of actual supplies",
                                                    "type": "string"
                                                },
                                                "assetType": {
                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "manufacturingDate": {
                                        "description": "Date of Manufacturing",
                                        "type": "string"
                                    },
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    },
                                    "productBatchId": {
                                        "description": "Batch id to which the product belongs",
                                        "type": "string"
                                    },
                                    "productDescription": {
                                        "description": "Description of the finished product",
                                        "type": "string"
                                    },
                                    "productId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    },
                                    "productInternalId": {
                                        "description": "A product's Internal ID",
                                        "type": "string"
                                    },
                                    "productModel": {
                                        "description": "Product Model",
                                        "type": "string"
                                    },
                                    "productRecalls": {
                                        "items": {
                                            "properties": {
                                                "assetType": {
                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                    "type": "string"
                                                },
                                                "newassetIds": {
                                                    "description": "Asset Ids of supply / product in recall",
                                                    "type": "string"
                                                },
                                                "originalassetIds": {
                                                    "description": "Asset Ids of supply / product in original schema",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "productSchemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    },
                                    "productStatus": {
                                        "description": "The current status of the product",
                                        "type": "string"
                                    },
                                    "productType": {
                                        "description": "FP(FinishedProduct/Subassembly",
                                        "type": "string"
                                    },
                                    "productionLocation": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "productId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateAssetProduct"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAssetProductSchema": {
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "productschema": {
                                "description": "The changeable properties for a productschema, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "productSchemaChange": {
                                        "items": {
                                            "properties": {
                                                "assetType": {
                                                    "description": "Is this a product(subassembly) or a supply asset",
                                                    "type": "string"
                                                },
                                                "newassetBatchId": {
                                                    "description": "Batch Id of supply / productModelId of product in recall",
                                                    "type": "string"
                                                },
                                                "originalassetBatchId": {
                                                    "description": "Batch Id of supply / productModelId of product in original schema",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "productSchemaContent": {
                                        "description": "Schema content",
                                        "type": "string"
                                    },
                                    "productcode": {
                                        "description": "Maker's code for this product type",
                                        "type": "string"
                                    },
                                    "productschemaCreationDate": {
                                        "description": "Date of schema definition",
                                        "type": "string"
                                    },
                                    "productschemaDescription": {
                                        "description": "Description about the productschema",
                                        "type": "string"
                                    },
                                    "productschemaId": {
                                        "description": "A product schema ID",
                                        "type": "string"
                                    },
                                    "productschemaStatus": {
                                        "description": "New / OnHold / InUse / Archived",
                                        "type": "string"
                                    },
                                    "productschemaType": {
                                        "description": "0 for subassembly, 1 for finished product",
                                        "type": "number"
                                    }
                                },
                                "required": [
                                    "productschemaId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateAssetProductSchema"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAssetSupply": {
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supply": {
                                "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
                                "properties": {
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "industryType": {
                                        "description": "Industry type - to start with Auto / Pharma / Food",
                                        "type": "string"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "makerproductId": {
                                        "description": "A product's ID",
                                        "type": "string"
                                    },
                                    "orderId": {
                                        "description": "An order's ID",
                                        "type": "string"
                                    },
                                    "shippedDate": {
                                        "description": "Date of shipment to maker",
                                        "type": "string"
                                    },
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    },
                                    "supplyAggregationLevel": {
                                        "description": "Supply type - RawMaterial / Component / Operational",
                                        "type": "string"
                                    },
                                    "supplyAvailableDate": {
                                        "description": "Date of manufacture",
                                        "type": "string"
                                    },
                                    "supplyBatchId": {
                                        "description": "Supplier's production batch under which this supply falls",
                                        "type": "string"
                                    },
                                    "supplyCount": {
                                        "description": "Number of assets",
                                        "type": "number"
                                    },
                                    "supplyDescription": {
                                        "description": "Description about the supply",
                                        "type": "string"
                                    },
                                    "supplyId": {
                                        "description": "A supply's ID",
                                        "type": "string"
                                    },
                                    "supplyStatus": {
                                        "description": "The current status of the supply",
                                        "type": "string"
                                    },
                                    "supplyType": {
                                        "description": "eg.LightingCircuits, EngineSubasembly etc ",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "supplyId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateAssetSupply"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateParticipantMaker": {
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "maker": {
                                "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry, to start off with is auto, pharma or food",
                                        "type": "string"
                                    },
                                    "makerAddress": {
                                        "description": "The makers' physical address and other means of contacting",
                                        "properties": {
                                            "city": {
                                                "type": "string"
                                            },
                                            "country": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            },
                                            "phone": {
                                                "type": "string"
                                            },
                                            "street": {
                                                "type": "string"
                                            },
                                            "website": {
                                                "type": "string"
                                            },
                                            "zipcode": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    },
                                    "sellers": {
                                        "items": {
                                            "properties": {
                                                "seller": {
                                                    "properties": {
                                                        "sellerId": {
                                                            "description": "A seller's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "suppliers": {
                                        "items": {
                                            "properties": {
                                                "supplier": {
                                                    "properties": {
                                                        "supplierId": {
                                                            "description": "A supplier's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    }
                                },
                                "required": [
                                    "makerId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateParticipantMaker"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateParticipantSeller": {
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "seller": {
                                "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry, to start off with is auto, pharma or food",
                                        "type": "string"
                                    },
                                    "makers": {
                                        "items": {
                                            "properties": {
                                                "maker": {
                                                    "properties": {
                                                        "makerId": {
                                                            "description": "A maker's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "sellerAddress": {
                                        "description": "The sellers' physical address and other means of contacting",
                                        "properties": {
                                            "city": {
                                                "type": "string"
                                            },
                                            "country": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            },
                                            "phone": {
                                                "type": "string"
                                            },
                                            "street": {
                                                "type": "string"
                                            },
                                            "website": {
                                                "type": "string"
                                            },
                                            "zipcode": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "sellerId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateParticipantSeller"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateParticipantSupplier": {
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "supplier": {
                                "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
                                "properties": {
                                    "industryType": {
                                        "description": "Industry, to start off with is auto, pharma or food",
                                        "type": "string"
                                    },
                                    "makers": {
                                        "items": {
                                            "properties": {
                                                "maker": {
                                                    "properties": {
                                                        "makerId": {
                                                            "description": "A maker's ID",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "minItems": 1,
                                        "type": "array"
                                    },
                                    "supplierAddress": {
                                        "description": "The suppliers' physical address and other means of contacting",
                                        "properties": {
                                            "city": {
                                                "type": "string"
                                            },
                                            "country": {
                                                "type": "string"
                                            },
                                            "name": {
                                                "type": "string"
                                            },
                                            "phone": {
                                                "type": "string"
                                            },
                                            "street": {
                                                "type": "string"
                                            },
                                            "website": {
                                                "type": "string"
                                            },
                                            "zipcode": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "supplierId",
                                    "industryType"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateParticipantSupplier"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "Model": {
        "maker": {
            "description": "The changeable properties for a maker, also considered its 'event' as a partial state",
            "properties": {
                "industryType": {
                    "description": "Industry, to start off with is auto, pharma or food",
                    "type": "string"
                },
                "makerAddress": {
                    "description": "The makers' physical address and other means of contacting",
                    "properties": {
                        "city": {
                            "type": "string"
                        },
                        "country": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        },
                        "phone": {
                            "type": "string"
                        },
                        "street": {
                            "type": "string"
                        },
                        "website": {
                            "type": "string"
                        },
                        "zipcode": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "makerId": {
                    "description": "A maker's ID",
                    "type": "string"
                },
                "sellers": {
                    "items": {
                        "properties": {
                            "seller": {
                                "properties": {
                                    "sellerId": {
                                        "description": "A seller's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 1,
                    "type": "array"
                },
                "suppliers": {
                    "items": {
                        "properties": {
                            "supplier": {
                                "properties": {
                                    "supplierId": {
                                        "description": "A supplier's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 1,
                    "type": "array"
                }
            },
            "required": [
                "makerId",
                "industryType"
            ],
            "type": "object"
        },
        "order": {
            "description": "The changeable properties for a order, also considered its 'event' as a partial state",
            "properties": {
                "common": {
                    "description": "Common properties for all assets",
                    "properties": {
                        "appdata": {
                            "description": "Application managed information as an array of key:value pairs",
                            "items": {
                                "properties": {
                                    "K": {
                                        "type": "string"
                                    },
                                    "V": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "deviceID": {
                            "description": "A unique identifier for the device that sent the current event",
                            "type": "string"
                        },
                        "devicetimestamp": {
                            "description": "A timestamp recoded by the device that sent the current event",
                            "type": "string"
                        },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "goodsOrServicesType": {
                    "description": "Goods / Services",
                    "type": "string"
                },
                "industryType": {
                    "description": "Industry type - to start with Auto / Pharma / Food",
                    "type": "string"
                },
                "orderContentType": {
                    "description": "Raw Material/ Component/Subassembly / Finished order",
                    "type": "string"
                },
                "orderDeliveryLocation": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "orderDescription": {
                    "description": "Description of the order",
                    "type": "string"
                },
                "orderDistributor": {
                    "description": "Order distributor details",
                    "minItems": 1,
                    "properties": {
                        "bolId": {
                            "description": "Bill of Lading number",
                            "type": "string"
                        },
                        "distId": {
                            "description": "Participant Id of order distributor",
                            "type": "string"
                        },
                        "distName": {
                            "description": "Participant name of order distributor",
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "orderFulfiller": {
                    "description": "Order fulfiller details",
                    "minItems": 1,
                    "properties": {
                        "Id": {
                            "description": "Participant Id of order fulfiller",
                            "type": "string"
                        },
                        "Type": {
                            "description": "Participant type of order fulfiller",
                            "type": "string"
                        },
                        "actualFulfilmentDate": {
                            "description": "Fulfilment date",
                            "type": "string"
                        },
                        "batchIds": {
                            "description": "batch ids used to fill order",
                            "type": "string"
                        },
                        "batchSizeDelivered": {
                            "description": "Number of assets",
                            "type": "number"
                        },
                        "committedDeliveryDate": {
                            "description": "Order due date",
                            "type": "string"
                        },
                        "orderReceivedDate": {
                            "description": "Date the order was received",
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "orderFulfilmentLocation": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "orderId": {
                    "description": "An order's ID",
                    "type": "string"
                },
                "orderIssuer": {
                    "description": "Order issuer details",
                    "minItems": 1,
                    "properties": {
                        "batchSizeRequired": {
                            "description": "Number of assets",
                            "type": "number"
                        },
                        "oderIssuerId": {
                            "description": "Participant Id of order issuer",
                            "type": "string"
                        },
                        "oderIssuerType": {
                            "description": "Participant type of order issuer",
                            "type": "string"
                        },
                        "orderDueDate": {
                            "description": "Order due date",
                            "type": "string"
                        },
                        "orderIssueDate": {
                            "description": "Year of Manufacturing",
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "orderStatus": {
                    "description": "The current status of the order",
                    "type": "string"
                },
                "orderType": {
                    "description": "Manufacturing / Distribution/MRO",
                    "type": "string"
                }
            },
            "required": [
                "orderId",
                "industryType",
                "orderType",
                "orderContentType",
                "orderStatus",
                "orderDescription",
                "orderFulfilmentLocation",
                "orderDeliveryLocation"
            ],
            "type": "object"
        },
        "product": {
            "description": "The changeable properties for a product, also considered its 'event' as a partial state",
            "properties": {
                "common": {
                    "description": "Common properties for all assets",
                    "properties": {
                        "appdata": {
                            "description": "Application managed information as an array of key:value pairs",
                            "items": {
                                "properties": {
                                    "K": {
                                        "type": "string"
                                    },
                                    "V": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "deviceID": {
                            "description": "A unique identifier for the device that sent the current event",
                            "type": "string"
                        },
                        "devicetimestamp": {
                            "description": "A timestamp recoded by the device that sent the current event",
                            "type": "string"
                        },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "industryType": {
                    "description": "Industry type - to start with Auto / Pharma / Food",
                    "type": "string"
                },
                "madeOf": {
                    "items": {
                        "properties": {
                            "assetBatchId": {
                                "description": "Batch Id of supply / productModelId of product",
                                "type": "string"
                            },
                            "assetCount": {
                                "description": "Number of items in the product",
                                "type": "number"
                            },
                            "assetDescription": {
                                "description": "Description of the supply/supplies",
                                "type": "string"
                            },
                            "assetIds": {
                                "description": "Comma separated list of actual supplies",
                                "type": "string"
                            },
                            "assetType": {
                                "description": "Is this a product(subassembly) or a supply asset",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 1,
                    "type": "array"
                },
                "makerId": {
                    "description": "A maker's ID",
                    "type": "string"
                },
                "manufacturingDate": {
                    "description": "Date of Manufacturing",
                    "type": "string"
                },
                "orderId": {
                    "description": "An order's ID",
                    "type": "string"
                },
                "productBatchId": {
                    "description": "Batch id to which the product belongs",
                    "type": "string"
                },
                "productDescription": {
                    "description": "Description of the finished product",
                    "type": "string"
                },
                "productId": {
                    "description": "A product's ID",
                    "type": "string"
                },
                "productInternalId": {
                    "description": "A product's Internal ID",
                    "type": "string"
                },
                "productModel": {
                    "description": "Product Model",
                    "type": "string"
                },
                "productRecalls": {
                    "items": {
                        "properties": {
                            "assetType": {
                                "description": "Is this a product(subassembly) or a supply asset",
                                "type": "string"
                            },
                            "newassetIds": {
                                "description": "Asset Ids of supply / product in recall",
                                "type": "string"
                            },
                            "originalassetIds": {
                                "description": "Asset Ids of supply / product in original schema",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 1,
                    "type": "array"
                },
                "productSchemaId": {
                    "description": "A product schema ID",
                    "type": "string"
                },
                "productStatus": {
                    "description": "The current status of the product",
                    "type": "string"
                },
                "productType": {
                    "description": "FP(FinishedProduct/Subassembly",
                    "type": "string"
                },
                "productionLocation": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "sellerId": {
                    "description": "A seller's ID",
                    "type": "string"
                }
            },
            "required": [
                "productId",
                "industryType"
            ],
            "type": "object"
        },
        "seller": {
            "description": "The changeable properties for a seller, also considered its 'event' as a partial state",
            "properties": {
                "industryType": {
                    "description": "Industry, to start off with is auto, pharma or food",
                    "type": "string"
                },
                "makers": {
                    "items": {
                        "properties": {
                            "maker": {
                                "properties": {
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 1,
                    "type": "array"
                },
                "sellerAddress": {
                    "description": "The sellers' physical address and other means of contacting",
                    "properties": {
                        "city": {
                            "type": "string"
                        },
                        "country": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        },
                        "phone": {
                            "type": "string"
                        },
                        "street": {
                            "type": "string"
                        },
                        "website": {
                            "type": "string"
                        },
                        "zipcode": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "sellerId": {
                    "description": "A seller's ID",
                    "type": "string"
                }
            },
            "required": [
                "sellerId",
                "industryType"
            ],
            "type": "object"
        },
        "supplier": {
            "description": "The changeable properties for a supplier, also considered its 'event' as a partial state",
            "properties": {
                "industryType": {
                    "description": "Industry, to start off with is auto, pharma or food",
                    "type": "string"
                },
                "makers": {
                    "items": {
                        "properties": {
                            "maker": {
                                "properties": {
                                    "makerId": {
                                        "description": "A maker's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 1,
                    "type": "array"
                },
                "supplierAddress": {
                    "description": "The suppliers' physical address and other means of contacting",
                    "properties": {
                        "city": {
                            "type": "string"
                        },
                        "country": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        },
                        "phone": {
                            "type": "string"
                        },
                        "street": {
                            "type": "string"
                        },
                        "website": {
                            "type": "string"
                        },
                        "zipcode": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "supplierId": {
                    "description": "A supplier's ID",
                    "type": "string"
                }
            },
            "required": [
                "supplierId",
                "industryType"
            ],
            "type": "object"
        },
        "supply": {
            "description": "The changeable properties for a supply, also considered its 'event' as a partial state",
            "properties": {
                "common": {
                    "description": "Common properties for all assets",
                    "properties": {
                        "appdata": {
                            "description": "Application managed information as an array of key:value pairs",
                            "items": {
                                "properties": {
                                    "K": {
                                        "type": "string"
                                    },
                                    "V": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "deviceID": {
                            "description": "A unique identifier for the device that sent the current event",
                            "type": "string"
                        },
                        "devicetimestamp": {
                            "description": "A timestamp recoded by the device that sent the current event",
                            "type": "string"
                        },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "industryType": {
                    "description": "Industry type - to start with Auto / Pharma / Food",
                    "type": "string"
                },
                "makerId": {
                    "description": "A maker's ID",
                    "type": "string"
                },
                "makerproductId": {
                    "description": "A product's ID",
                    "type": "string"
                },
                "orderId": {
                    "description": "An order's ID",
                    "type": "string"
                },
                "shippedDate": {
                    "description": "Date of shipment to maker",
                    "type": "string"
                },
                "supplierId": {
                    "description": "A supplier's ID",
                    "type": "string"
                },
                "supplyAggregationLevel": {
                    "description": "Supply type - RawMaterial / Component / Operational",
                    "type": "string"
                },
                "supplyAvailableDate": {
                    "description": "Date of manufacture",
                    "type": "string"
                },
                "supplyBatchId": {
                    "description": "Supplier's production batch under which this supply falls",
                    "type": "string"
                },
                "supplyCount": {
                    "description": "Number of assets",
                    "type": "number"
                },
                "supplyDescription": {
                    "description": "Description about the supply",
                    "type": "string"
                },
                "supplyId": {
                    "description": "A supply's ID",
                    "type": "string"
                },
                "supplyStatus": {
                    "description": "The current status of the supply",
                    "type": "string"
                },
                "supplyType": {
                    "description": "eg.LightingCircuits, EngineSubasembly etc ",
                    "type": "string"
                }
            },
            "required": [
                "supplyId",
                "industryType"
            ],
            "type": "object"
        }
    }
}`


	var readAssetSchemas iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(schemas), nil
	}
	func init() {
		iot.AddRoute("readAssetSchemas", "query", iot.SystemClass, readAssetSchemas)
	}
	