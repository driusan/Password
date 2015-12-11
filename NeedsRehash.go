package Password

const PASSWORD_BCRYPT int = 1
const PASSWORD_DEFAULT int = PASSWORD_BCRYPT

func NeedsRehash(hash PasswordHash, algo int, options ...map[string]interface{}) bool {
	return hash[:3] != "$2y"
}
