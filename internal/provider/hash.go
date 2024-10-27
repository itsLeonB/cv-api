package provider

import (
	"github.com/itsLeonB/cv-api/internal/apperror"
	"golang.org/x/crypto/bcrypt"
)

type HashProvider interface {
	HashPassword(pwd string) (string, error)
	CheckPassword(pwd string, hash string) error
}

type hashBcrypt struct {
	structName string
	cost       int
}

func NewBcryptHashProvider(cost int) *hashBcrypt {
	return &hashBcrypt{"hashBcrypt", cost}
}

func (h *hashBcrypt) HashPassword(pwd string) (string, error) {
	methodName := "HashPassword()"
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), h.cost)
	if err != nil {
		return "", apperror.NewAppError(
			err, h.structName, methodName,
			"bcrypt.GenerateFromPassword()",
		)
	}

	return string(hash), nil
}

func (h *hashBcrypt) CheckPassword(pwd string, hash string) error {
	methodName := "CheckPassword()"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return apperror.NewAppError(
			err, h.structName, methodName,
			"bcrypt.CompareHashAndPassword()",
		)
	}

	return nil
}
