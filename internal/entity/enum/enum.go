package enum

import (
	"database/sql/driver"
	"errors"
)

type UserStatus string

const (
	Student UserStatus = "student"
	Teacher UserStatus = "teacher"
)

func (status *UserStatus) Scan(value interface{}) error {
	if value != nil {
		if str, err := driver.String.ConvertValue(value); err == nil {
			switch str.(string) {
			case "student":
				*status = UserStatus(Student)
				return nil
			case "teacher":
				*status = UserStatus(Teacher)
				return nil
			}
		}
	}
	return errors.New("unable to map user status enum")

}

func (status UserStatus) Value() (driver.Value, error) {
	return string(status), nil
}

type LoginProvider string

const (
	Google LoginProvider = "google"
)

func (status *LoginProvider) Scan(value interface{}) error {
	if value != nil {
		if str, err := driver.String.ConvertValue(value); err == nil {
			switch str.(string) {
			case "google":
				*status = LoginProvider(Google)
				return nil
			}
		}
	}
	return errors.New("unable to map user status enum")

}

func (status LoginProvider) Value() (driver.Value, error) {
	return string(status), nil
}
