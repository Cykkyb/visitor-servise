package service

import (
	"log/slog"
	"visitor/internal/entity"
	"visitor/internal/repository/postgres"
)

type VisitorService struct {
	repository repository.Visitor
	log        *slog.Logger
}

func NewVisitorService(repo repository.Visitor, log *slog.Logger) *VisitorService {
	return &VisitorService{
		repository: repo,
		log:        log,
	}
}

func (s *VisitorService) CreateUser(user *entity.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	id, err := s.repository.CreateUser(*user)
	if err != nil {
		return err
	}

	user.Id = id
	return nil
}

func (s *VisitorService) GetUser(id int) (entity.User, error) {
	return s.repository.GetUser(id)
}

func (s *VisitorService) UpdateUser(user entity.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := s.repository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *VisitorService) DeleteUser(id int) error {
	_, err := s.GetUser(id)
	if err != nil {
		return err
	}

	return s.repository.DeleteUser(id)
}
