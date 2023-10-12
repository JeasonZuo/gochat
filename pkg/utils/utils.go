package utils

import "github.com/spf13/viper"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(viper.GetString("JwtSecret"))
}
