package constant

const (
	UserRoleAdmin    = 1
	UserRoleBorrower = 2
	UserRoleInvestor = 3
	UserRoleSystem   = 4
)

var (
	UserRoleText = map[int]string{
		UserRoleAdmin:    "Admin",
		UserRoleBorrower: "Borrower",
		UserRoleInvestor: "Investor",
		UserRoleSystem:   "System",
	}
)
