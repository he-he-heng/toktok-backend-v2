package dto

type CreateUserRequest struct {
	LoginID  string `json:"login_id" validate:"gte=4,lte=18"`
	Password string `json:"password" validate:"gte=6,lte=32"`
}
