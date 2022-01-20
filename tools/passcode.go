package tools

import "golang.org/x/crypto/bcrypt"

func HashGenerate(passcode string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(passcode), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func HashCompare(passcode string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passcode))
	if err != nil {
		return false
	}
	return true
}
