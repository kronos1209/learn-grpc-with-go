package services

import (
	"github.com/kronos1209/learn-grpc-with-go/internal/entities"
	"github.com/kronos1209/learn-grpc-with-go/internal/repository/user"
)

type UserServiceBuilder struct {
	repository user.UserRepository
}

func NewUserServiceBuilder() *UserServiceBuilder {
	return &UserServiceBuilder{}
}

func (b *UserServiceBuilder) Repository(repository user.UserRepository) *UserServiceBuilder {
	b.repository = repository
	return b
}

func (b *UserServiceBuilder) Build() *UserService {
	return &UserService{
		repository: b.repository,
	}
}

type UserService struct {
	repository user.UserRepository
}

func (s *UserService) CreateUser(userId string, name string, passwd string) (*entities.User, error) {
	u, err := entities.NewUser(userId, name, passwd)
	if err != nil {
		return nil, err
	}
	if err := s.repository.Add(u); err != nil {
		return nil, err
	}
	return u, nil
}
