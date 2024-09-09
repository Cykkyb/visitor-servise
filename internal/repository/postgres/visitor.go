package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"visitor/internal/entity"
)

type VisitorPostgres struct {
	db *sqlx.DB
}

func NewVisitorPostgres(db *sqlx.DB) *VisitorPostgres {
	return &VisitorPostgres{
		db: db,
	}
}

func (r *VisitorPostgres) CreateUser(user entity.User) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name, surname, email, phone, country_code) VALUES ($1, $2, $3, $4, $5) returning id`, usersTable)

	var id int
	err := r.db.QueryRow(query, user.Name, user.Surname, user.Email, user.Phone, user.Country.Code).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *VisitorPostgres) GetUser(id int) (entity.User, error) {
	query := fmt.Sprintf(`SELECT id, name, surname, email, phone, country_code FROM %s WHERE id = $1`, usersTable)

	var user entity.User

	err := r.db.Get(&user, query, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *VisitorPostgres) UpdateUser(user entity.User) error {
	query := fmt.Sprintf(`UPDATE %s SET name = $1, surname = $2, email = $3, phone = $4, country_code = $5 WHERE id = $6`, usersTable)

	_, err := r.db.Exec(query, user.Name, user.Surname, user.Email, user.Phone, user.Country.Code, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *VisitorPostgres) DeleteUser(id int) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, usersTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
