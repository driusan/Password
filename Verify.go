package Password

import "golang.org/x/crypto/bcrypt"

func Verify(password PlainTextPassword, hash PasswordHash) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err == nil {
		return true, nil
	}
	return false, err

}
