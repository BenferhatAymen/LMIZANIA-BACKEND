package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type PasswordHelper struct{}

func (p *PasswordHelper) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	  if err != nil {
        log.Println(err)
    }
	return string(bytes), err
}

func (p *PasswordHelper) CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
