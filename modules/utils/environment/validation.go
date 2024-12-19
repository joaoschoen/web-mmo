package environment

import (
	"errors"
	"os"
)

func Validate() error {
	var variable = os.Getenv("DATABASE_URL")
	if len(variable) == 0 {
		return errors.New("Empty environment variable DATABASE_URL")
	}
	variable = os.Getenv("TOKEN_SECRET")
	if len(variable) == 0 {
		return errors.New("Empty environment variable TOKEN_SECRET")
	}

	return nil
}
