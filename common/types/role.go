package types

import "slices"

const RoleSuperAdmin = "superAdmin"
const RoleAdmin = "admin"
const RoleManager = "manager"
const RoleManagerAssist = "managerAssist"
const RoleSeller = "seller"
const RoleSellerAssist = "sellerAssist"
const RoleDriver = "driver"
const RoleDriverAssist = "driverAssist"
const RoleCustomer = "customer"
const RoleCustomerService = "customerService"

var AllRoles = []string{
	RoleSuperAdmin,
	RoleAdmin,
	RoleManager, RoleManagerAssist,
	RoleSeller, RoleSellerAssist,
	RoleDriver, RoleDriverAssist,
	RoleCustomer, RoleCustomerService,
}

func IsValidRole(role string) bool {
	return slices.Contains(AllRoles, role)
}
