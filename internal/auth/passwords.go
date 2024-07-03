package auth

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
)

/*
Password complexity requirements

- Min: 8 Characters
- Max: 32 Characters

- ASCII characters only a-zA-Z0-9

- 72 Bytes is max limit
- Pepper will be 12-32 characters
- If the password is over 40 bytes we should prevent it going any further

- over 72 Bytes disable it.
*/

var pepper = os.Getenv("sitePepper")

var iterations = 14

// Hashing

func hashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+pepper), iterations)
	return hashedPassword, err
}

func verifyPassword(storedPassword []byte, password string) bool {
	passwordCompare := bcrypt.CompareHashAndPassword(storedPassword, []byte(password+pepper))
	return passwordCompare == nil
}

// To be used after the bcrypt algorithm has hashed a password.
func isBytesTooLarge(hashedPassword []byte) bool {
	return 60 < len(hashedPassword)
}

// Validation
func validPassword(password string) bool {
	if isPasswordTooLarge(password) {
		return false
	}
	if !isValidCharactersAndLength(password) {
		return false
	}
	return true
}

func validUsername(username string) bool {
	// Define the regex pattern
	pattern := `^[a-zA-Z0-9]{1,32}$`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	return re.MatchString(username)
}

func isPasswordTooLarge(password string) bool {
	// Convert the password to bytes and get its length
	length := len([]byte(password))

	// Check if the length exceeds 40 bytes
	return length > 40
}

func isValidCharactersAndLength(password string) bool {
	// Define the regex pattern
	pattern := `^[a-zA-Z0-9]{8,32}$`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	return re.MatchString(password)
}
