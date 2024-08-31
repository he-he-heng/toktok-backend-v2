package frame

type GetUser struct {
	ID int
}

type CreateUser struct {
	LoginID  string
	Password string
}

type UpdateUser struct {
	LoginID         *string
	Password        *string
	Email           *string
	ConfirmPassword *string
}

type DeleteUser struct {
	ID int
}
