package domain

import "time"

type UserRoleType string

const (
	Guest   UserRoleType = "GUEST"
	Manager UserRoleType = "MANAGER"
	Admin   UserRoleType = "ADMIN"
)

type User struct {
	ID        uint         `gorm:"primaryKey"`
	LoginID   string       `gorm:"type:varchar(255);unique;not null"`
	Password  string       `gorm:"type:varchar(255);not null"`
	Email     *string      `gorm:"type:varchar(255)"`
	Role      UserRoleType `gorm:"type:varchar(255);default:GUEST"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
