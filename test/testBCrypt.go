package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("012345678901234567890123456789012345678901234567890123456789012345678912")
	//newPassword := []byte("user_password_new")
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	} else {
		hashString := string(hash)
		fmt.Println(len(hashString))
		err := bcrypt.CompareHashAndPassword([]byte(hashString), password)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Password is correct")
		}
	}
}
