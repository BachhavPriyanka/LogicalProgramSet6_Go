package main
import (
	"fmt"
	"regexp"
) 

func main(){
	var email string
	fmt.Println("Enter Email id :")
	fmt.Scanln(&email)
	checkEmail := isEmailValid(email)
	fmt.Println(checkEmail)
}

func isEmailValid(userEmail string) bool{
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(userEmail)
}