package utilites

import "golang.org/x/crypto/bcrypt"

func EncryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
