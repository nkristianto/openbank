// Code generated by "scopegen"; DO NOT EDIT.
package entitlement

type AuthScope string

const (
	Scope_Read AuthScope = "https://auth.bnk.to/entitlement.read"
	Scope_Write AuthScope = "https://auth.bnk.to/entitlement.write"
)

var AuthScopes = map[string][]AuthScope{
	"/entitlement.EntitlementService/AddEntitlementRequestForCurrentUser": []AuthScope{Scope_Write},
	"/entitlement.EntitlementService/AddEntitlementRequestForUser": []AuthScope{Scope_Write},
	"/entitlement.EntitlementService/DeleteEntitlement": []AuthScope{Scope_Write},
	"/entitlement.EntitlementService/DeleteEntitlementRequest": []AuthScope{Scope_Write},
	"/entitlement.EntitlementService/GetAllEntitlementRequests": []AuthScope{Scope_Read},
	"/entitlement.EntitlementService/GetAllEntitlements": []AuthScope{Scope_Read},
	"/entitlement.EntitlementService/GetEntitlementForCurrentUser": []AuthScope{Scope_Read},
	"/entitlement.EntitlementService/GetEntitlementForUser": []AuthScope{Scope_Read},
	"/entitlement.EntitlementService/GetEntitlementForUserAtBank": []AuthScope{Scope_Read},
	"/entitlement.EntitlementService/GetEntitlementRequestForCurrentUser": []AuthScope{Scope_Read},
	"/entitlement.EntitlementService/GetEntitlementRequestForUser": []AuthScope{Scope_Read},
	"/entitlement.EntitlementService/GetRoles": []AuthScope{Scope_Read},
}
