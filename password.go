package main

import (
	"fmt"
	"regexp"
)

func main() {
	var pass string
	fmt.Println("Enter the Password")
	fmt.Scanln(&pass)
	passwordCheck := regexp.MustCompile("([!@#$%^&*.?-])+").MatchString(pass)

	if passwordCheck == true {
		fmt.Println("Correct Password")
	} else {
		fmt.Println("Incorect Password")
	}

}
