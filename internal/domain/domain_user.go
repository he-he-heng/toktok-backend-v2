package domain

import (
	"context"
	"time"
)

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

type UserRepository interface {
	CreateUser(ctx context.Context, loginID string, hashedPassword string) (*User, error)
	GetUser(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, loginID *string, hashedPassword *string, email *string) (*User, error)
	DeleteUser(ctx context.Context, id int) (*User, error)
}

type UserUsecase interface {
	CreateUser(ctx context.Context, loginID string, hashedPassword string) (*User, error)
	GetUser(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, loginID *string, hashedPassword *string, email *string) (*User, error)
	DeleteUser(ctx context.Context, id int) (*User, error)
}
