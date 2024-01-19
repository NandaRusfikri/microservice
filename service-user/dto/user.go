package dto

type SchemaUser struct {
	ID       uint64 `json:"id" `
	Fullname string `json:"fullname,omitempty" validate:"required,lowercase"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	IsActive *bool  `json:"is_active,omitempty" validate:"required"`
	Balance  int64  `json:"balance,omitempty" validate:"required"`
}
type CutBalanceRequest struct {
	UserId  uint64 `json:"user_id" `
	Balance int64  `json:"balance" validate:"required"`
}
