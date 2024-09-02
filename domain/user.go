package domain

import "gorm.io/gorm"

type UserRoleType string

const (
	Guest   UserRoleType = "GUEST"
	Manager UserRoleType = "MANAGER"
	Admin   UserRoleType = "ADMIN"
)

type User struct {
	gorm.Model

	LoginID  string       `gorm:"type:varchar(255);unique;not null"`
	Password string       `gorm:"type:varchar(255);not null"`
	Email    *string      `gorm:"type:varchar(255)"`
	Role     UserRoleType `gorm:"type:varchar(255);default:GUEST;not null"`
}
