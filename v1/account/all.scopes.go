// Code generated by "scopegen"; DO NOT EDIT.
package accounts

type AuthScope string

const (
	Scope_Read AuthScope = "https://auth.bnk.to/account.read"
	Scope_Write AuthScope = "https://auth.bnk.to/account.write"
)

var AuthScopes = map[string][]AuthScope{
	"/accounts.AccountService/CheckAccount": []AuthScope{Scope_Read},
	"/accounts.AccountService/CreateAccount": []AuthScope{Scope_Write},
	"/accounts.AccountService/DeleteAccount": []AuthScope{Scope_Write},
	"/accounts.AccountService/GetAccount": []AuthScope{Scope_Read},
	"/accounts.AccountService/GetAccountStatus": []AuthScope{Scope_Read},
	"/accounts.AccountService/GetAccounts": []AuthScope{Scope_Read},
	"/accounts.AccountService/UpdateAccount": []AuthScope{Scope_Write},
	"/accounts.AccountService/UpdateAccountStatus": []AuthScope{Scope_Write},
}
