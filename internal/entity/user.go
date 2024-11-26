package entity

type User struct {
	UserID      int64  `gorm:"primaryKey" json:"user_id"`
	Name        string `json:"name"`
	Role        int    `json:"role"` // assumption: one user only has one role (e.g. admin, borrower or investor)
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
