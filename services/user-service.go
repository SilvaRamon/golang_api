package service

import (
	"hello/core"
	"hello/repository"
)

type IService interface {
	GetUser(mail string) core.User
	GetUsers() []core.User
}

type service struct {
	repo repository.IUserRepository
}

func Initialize(repo repository.IUserRepository) *service {
	return &service{repo}
}

func (s *service) GetUsers() []core.User {
	return s.repo.GetUsers()
}

func (s *service) GetUser(mail string) core.User {
	return s.repo.GetUser(mail)
}
