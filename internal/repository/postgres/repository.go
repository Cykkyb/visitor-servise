package repository

import (
	"github.com/jmoiron/sqlx"
	"visitor/internal/entity"
)

type Visitor interface {
	CreateUser(entity.User) (int, error)
	UpdateUser(entity.User) error
	GetUser(int) (entity.User, error)
	DeleteUser(int) error
}

type Repository struct {
	Visitor
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Visitor: NewVisitorPostgres(db),
	}
}
