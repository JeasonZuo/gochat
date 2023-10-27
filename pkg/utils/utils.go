package utils

import (
	"os"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}
