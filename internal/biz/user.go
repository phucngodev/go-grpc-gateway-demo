package biz

import (
	"context"
	pb "kratos-blog/api/blog/v1"
	"kratos-blog/internal/conf"
	"kratos-blog/pkg/hash"
	JWT "kratos-blog/pkg/jwt"
	"net/mail"
	"time"
)

const (
	DefaultImage = "https://seccdn.libravatar.org/gravatarproxy/515e2b667cc65fac595640adbf6bfd76?s=80"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Image    string
	Password string
}

type AuthResponse struct {
	Name  string
	Email string
	Token string
}

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
}

type UserUsercase struct {
	userRepo UserRepo
	cfg      *conf.Bootstrap
}

func NewUserUsercase(repo UserRepo, cfg *conf.Bootstrap) *UserUsercase {
	return &UserUsercase{
		userRepo: repo,
		cfg:      cfg,
	}
}

func (uc *UserUsercase) Login(ctx context.Context, email string, password string) (*AuthResponse, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, pb.ErrorEmailNotVerify("email is not valid")
	}

	u, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, pb.ErrorEmailNotAvailable("email is not available for register")
	}

	if err := hash.CompareHashAndPassword(u.Password, password); err != nil {
		return nil, pb.ErrorPasswordNotVerify("password is not correct")
	}

	jwt := JWT.New([]byte(uc.cfg.Jwt.Secret), "kratos-blog.com")

	token, err := jwt.CreateToken(u.ID, time.Minute*5)
	if err != nil {
		return nil, pb.ErrorTokenNotCreated("could not create token")
	}

	return &AuthResponse{
		Name:  u.Name,
		Email: u.Email,
		Token: token,
	}, nil
}

func (uc *UserUsercase) Register(ctx context.Context, email string, password string) (*AuthResponse, error) {
	return nil, nil
}
