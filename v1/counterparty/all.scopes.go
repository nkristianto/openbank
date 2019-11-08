// Code generated by "scopegen"; DO NOT EDIT.
package counterparty

type AuthScope string

const (
	Scope_Read AuthScope = "https://auth.bnk.to/counterparty.read"
	Scope_Write AuthScope = "https://auth.bnk.to/counterparty.write"
)

var AuthScopes = map[string][]AuthScope{
	"/counterparty.CounterPartyService/CreateCounterParty": []AuthScope{Scope_Write},
	"/counterparty.CounterPartyService/GetCounterParties": []AuthScope{Scope_Read},
	"/counterparty.CounterPartyService/GetCounterParty": []AuthScope{Scope_Read},
	"/counterparty.CounterPartyService/GetOtherAccountByID": []AuthScope{Scope_Read},
	"/counterparty.CounterPartyService/GetOtherAccounts": []AuthScope{Scope_Read},
}