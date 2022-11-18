package password

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) ([]byte, error) {

	bytePass := []byte(password)
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil

}

func Compare(hashedPassword string, password string) error {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
