package user

type User struct {
	UserID      int64
	Name        string
	Role        int
	Email       string
	PhoneNumber string
}

type CreateUserRequest struct {
	Name        string
	Role        int
	Email       string
	PhoneNumber string
}

type TopUpUserBalanceRequest struct {
	UserID      int64
	UserRole    int
	TopUpAmount int64
}

type TopUpUserBalanceResult struct {
	TotalBalanceAmount int64
}
