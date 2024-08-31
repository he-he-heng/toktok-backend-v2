package domain

type UserRoleType string

const (
	Guest   UserRoleType = "GUEST"
	Manager UserRoleType = "MANAGER"
	ADMIN   UserRoleType = "ADMIN"
)

type User struct {
	ID        int
	LoginID   int
	Password  string
	Email     *string // nullable
	Role      UserRoleType
	CreatedAt string
	UpdatedAt string
}
