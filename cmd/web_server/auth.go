package main

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"
	userv1 "github.com/Vaziria/pkl/backend/gen/catur/user/v1"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"size:254;uniqueIndex;not null"`
	Username     string `gorm:"size:32;uniqueIndex;not null"`
	PasswordHash string `gorm:"size:255;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type AuthService struct {
	db *gorm.DB
}

// Logout implements [userv1connect.AuthServiceHandler].
func (a *AuthService) Logout(context.Context, *connect.Request[userv1.LogoutRequest]) (*connect.Response[userv1.LogoutResponse], error) {
	panic("unimplemented")
}

// Login implements [userv1connect.AuthServiceHandler].
func (a *AuthService) Login(context.Context, *connect.Request[userv1.LoginRequest]) (*connect.Response[userv1.LoginResponse], error) {
	return nil, errors.New("Programmer males, belum ditulis")
}

// Register implements [userv1connect.AuthServiceHandler].
func (a *AuthService) Register(ctx context.Context, req *connect.Request[userv1.RegisterRequest]) (*connect.Response[userv1.RegisterResponse], error) {
	var err error
	resp := &userv1.RegisterResponse{
		Name: req.Msg.Username,
	}

	payload := req.Msg

	user := User{
		Email:        payload.Email,
		Username:     payload.Username,
		PasswordHash: payload.Password,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = a.db.Save(&user).Error

	return connect.NewResponse(resp), err

}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db,
	}

}
