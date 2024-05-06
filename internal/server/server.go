package server

import (
	"context"

	"github.com/kronos1209/learn-grpc-with-go/internal/repository/user"
	"github.com/kronos1209/learn-grpc-with-go/internal/services"
	"github.com/kronos1209/learn-grpc-with-go/pkg/proto"
)

type GrpcServer struct {
	proto.UnimplementedUserServiceServer
	userService *services.UserService
}

var _ (proto.UserServiceServer) = (*GrpcServer)(nil)

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{
		userService: services.NewUserServiceBuilder().Repository(user.NewMemoryRepository()).Build(),
	}
}

func (s *GrpcServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	u, err := s.userService.CreateUser(req.UserId, req.Name, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.CreateUserResponse{
		AccountId: u.ID,
	}, nil
}
