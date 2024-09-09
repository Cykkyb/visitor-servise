package service

import (
	"log/slog"
	"visitor/internal/entity"
	"visitor/internal/repository/postgres"
)

type Visitor interface {
	CreateUser(*entity.User) error
	UpdateUser(entity.User) error
	GetUser(int) (entity.User, error)
	DeleteUser(int) error
}

type Service struct {
	Visitor
}

func NewService(repo *repository.Repository, log *slog.Logger) *Service {
	return &Service{
		Visitor: NewVisitorService(repo.Visitor, log),
	}
}
