package data

import (
	"context"
	"kratos-blog/internal/biz"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	Emage        string
	Username     string
	PasswordHash string
}

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	return nil, nil
}
