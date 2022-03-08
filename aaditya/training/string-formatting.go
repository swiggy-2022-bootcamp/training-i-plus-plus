package partyrobot

//Concepts - String formatting  
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	var s string = fmt.Sprintf("Welcome to my party, %s!",name)
    return s
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	var s string = fmt.Sprintf("Happy birthday %s! You are now %d years old!",name,age)
    return s
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var s string = fmt.Sprintf("Welcome to my party, %s!\n", name)
    s+=fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n",table,direction,distance)
    s+=fmt.Sprintf("You will be sitting next to %s.",neighbor)
    return s
}