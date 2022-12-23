package main

import (
	"fmt"
	"time"
)

type familyMember struct {
	Name              string
	Relation          string
	Date, Month, Year int
}

func main() {
	famMember1 := familyMember{
		Name:     "Omkar",
		Relation: "Brother",
		Date:     02,
		Month:    04,
		Year:     2002,
	}
	famMember2 := familyMember{
		Name:     "Priyanka",
		Relation: "self",
		Date:     18,
		Month:    02,
		Year:     2000,
	}
	famMember3 := familyMember{
		Name:     "Nirmala",
		Relation: "Mother",
		Date:     24,
		Month:    06,
		Year:     1972,
	}
	famMember4 := familyMember{
		Name:     "BigBoss",
		Relation: "Father",
		Date:     01,
		Month:    06,
		Year:     1967,
	}
	membersList := []familyMember{
		famMember1,
		famMember2,
		famMember3,
		famMember4,
	}
	for _, eachMember := range membersList {
		findAge(eachMember.Name, eachMember.Year)
	}
}

func findAge(name string, birthYear int) {
	currentYear := time.Now().Year()
	ageOfMember := currentYear - birthYear
	fmt.Printf("Name: %s | Age: %dyr.\n", name, ageOfMember)
}
