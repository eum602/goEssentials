package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	definitions()
	correctPassword := `passwd123`
	hashedPassword := example1Encrypt(correctPassword)
	fmt.Println("\nTesting a correct password")
	example2Compare(hashedPassword, correctPassword)
	fmt.Println("\nTesting a wrong password")
	example2Compare(hashedPassword, `someFakePassword`)
}

func definitions() {
	fmt.Println(`
	https://godoc.org/?q=bcrypt
	`)
}

func example1Encrypt(password string)  []byte {
	s1 := []byte(password)
	hashedPassword,err := bcrypt.GenerateFromPassword(s1, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("err",err)
	}
	fmt.Println("initial password ==>",password)
	fmt.Println("encrypted password ==>",hashedPassword)
	return hashedPassword
}

func example2Compare(hashedPassword []byte, password string)  {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("PASSWORD IS WRONG")
	}else{
		fmt.Println("Password and hashed Password were successfully verified")
	}	
}
