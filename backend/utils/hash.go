// Package utils holds small helpers shared across the app.
package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword turns a plain-text password into a secure, salted hash.
// bcrypt automatically handles the "salt" (random data mixed in so two
// users with the same password don't get the same hash) — you never
// manage that yourself.
func HashPassword(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPassword compares a plain-text password against a stored hash.
// Returns true if they match.
func CheckPassword(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}