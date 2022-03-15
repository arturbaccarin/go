package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var tableString string = strconv.Itoa(table)
	var lenTable int = len(tableString)

	if lenTable == 1 {
		return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table 00%d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
	} else if lenTable == 2 {
		return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table 0%d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
	} else {
		return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
	}
}
