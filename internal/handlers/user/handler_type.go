package user

type CreateUserRequest struct {
	Name        string `json:"name"`
	Role        int    `json:"role"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type CreateUserResponse struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"user_role"`
}

type TopUpUserBalanceRequest struct {
	TopUpAmount int64 `json:"top_up_amount"`
}

type TopUpUserBalanceResponse struct {
	BalanceAmount int64 `json:"balance_amount"`
}
