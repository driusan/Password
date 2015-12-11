package Password

import "golang.org/x/crypto/bcrypt"

type PlainTextPassword string
type PasswordHash string

func Hash(password PlainTextPassword, algo int) (PasswordHash, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	// The security of the algorithm "$2a" is ambiguous.
	// It should be $2y since the Go implementation of bcrypt
	// wasn't the implementation which had a bug causing them
	// to change the prefix to make it easier to identify
	// vulnerable passwords.
	// This should be fixed upstream, but for now just change the
	// prefix here so that NeedsRehash doesn't constantly return
	// true
	hashedPassword[2] = 'y'

	return PasswordHash(hashedPassword), err

}
