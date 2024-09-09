package entity

import (
	"errors"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var PhoneCountriesMap = map[string]Country{
	"+1":  {"US"},
	"+44": {"GB"},
	"+7":  {"RU"},
}

type User struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name" validate:"required"`
	Surname string `json:"surname" db:"surname" validate:"required"`
	Email   string `json:"email" db:"email" validate:"email"`
	Phone   string `json:"phone" db:"phone" validate:"required,phone"`
	Country
}

func (u *User) Validate() error {
	validate := validator.New()

	if err := validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		phoneNumber := fl.Field().String()

		var phoneRegex = regexp.MustCompile(`^\+[0-9]+$`)
		if !phoneRegex.MatchString(phoneNumber) {
			return false
		}

		return true
	}); err != nil {
		return err
	}

	if err := validate.Struct(u); err != nil {
		return err
	}

	if u.Code == "" {
		if err := u.Country.SetCountryByPhoneNumber(u.Phone); err != nil {
			return err
		}
	}

	return nil
}

type Country struct {
	Code string `json:"country_code" db:"country_code"`
}

func (c *Country) SetCountryByPhoneNumber(phoneNumber string) error {
	if !strings.HasPrefix(phoneNumber, "+") {
		return errors.New("invalid phone number")
	}

	for prefix, country := range PhoneCountriesMap {
		if strings.HasPrefix(phoneNumber, prefix) {
			c.Code = country.Code
			return nil
		}
	}

	return errors.New("country not found")
}
