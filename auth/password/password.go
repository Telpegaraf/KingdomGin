package password

import "golang.org/x/crypto/bcrypt"

func CreatePassword(password string, strength int) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), strength)
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func ComparePassword(hashedPassword, password []byte) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, password) == nil
}
