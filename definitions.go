package main

// SupplierList - List of Suppliers in the system
type SupplierList struct {
	Suppliers []string `json:"suppliers"`
}
//IndustryTypes that the supplychain demo is addressing
type IndustryTypes string
const (
    Auto IndustryTypes = "AUTO"
    Pharma  ="PHARMA"
    Food = "FOOD"
) 
//EntityTypes :  Types of Entities in the system.
type EntityTypes string
const (
	Asset EntityTypes = "ASSET"
    Participant = "PARTICIPANT"
)
//AssetTypes :  Types of assets in the system.
type AssetTypes string
const (
    Supply AssetTypes = "SUPPLY"
    Product  ="PRODUCT"
)

//RoleTypes : Roles played by the Participants
type RoleTypes string
const (
    Supplier RoleTypes = "SUPPLIER"
    Maker  ="MAKER"
    Seller = "SELLER"
	Buyer = "BUYER" // Buyer not implemented in this iteration
)

