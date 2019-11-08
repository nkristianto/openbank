// Code generated by "scopegen"; DO NOT EDIT.
package counterpartymetadata

type AuthScope string

const (
	Scope_Read AuthScope = "https://auth.bnk.to/counterpartymetadata.read"
	Scope_Write AuthScope = "https://auth.bnk.to/counterpartymetadata.write"
)

var AuthScopes = map[string][]AuthScope{
	"/counterpartymetadata.CounterpartyMetadataService/CreateCorporateLocation": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/CreateImageURL": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/CreateMoreInfo": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/CreateOpenCorporatesURL": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/CreatePhysicalLocation": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/CreatePrivateAlias": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/CreatePublicAlias": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/CreateURL": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/DeleteCorporateLocation": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/DeleteImageURL": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/DeleteMoreInfo": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/DeleteOpenCorporatesURL": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/DeletePhysicalLocation": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/DeletePrivateAlias": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/DeletePublicAlias": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/DeleteURL": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/GetCorporateLocation": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetCorporateLocations": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetImageURL": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetImageURLs": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetMoreInfo": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetMoreInfos": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetOpenCorporatesURL": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetOpenCorporatesURLs": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetOtherAccountMetadata": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetPhysicalLocation": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetPhysicalLocations": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetPrivateAlias": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetPrivateAliases": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetPublicAlias": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetPublicAliases": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetURL": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/GetURLs": []AuthScope{Scope_Read},
	"/counterpartymetadata.CounterpartyMetadataService/UpdateCorporateLocation": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/UpdateImageURL": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/UpdateMoreInfo": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/UpdateOpenCorporatesURL": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/UpdatePhysicalLocation": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/UpdatePrivateAlias": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/UpdatePublicAlias": []AuthScope{Scope_Write},
	"/counterpartymetadata.CounterpartyMetadataService/UpdateURL": []AuthScope{Scope_Write},
}
