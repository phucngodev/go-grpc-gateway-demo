package service

import (
	"context"

	pb "kratos-blog/api/blog/v1"
	"kratos-blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	log  *log.Helper
	user *biz.UserUsercase
}

func NewUserService(user *biz.UserUsercase, logger log.Logger) *UserService {
	return &UserService{
		log:  log.NewHelper(logger),
		user: user,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	s.log.Infof("call login %v", req)
	return &pb.AuthResponse{}, nil
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{}, nil
}
