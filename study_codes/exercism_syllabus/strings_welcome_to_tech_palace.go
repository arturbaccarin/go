package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	message := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	stringOutput := stars + "\n" + welcomeMsg + "\n" + stars
	return stringOutput
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newString := strings.ReplaceAll(oldMsg, "*", "")
	newString = strings.TrimSpace(newString)
	return newString
}
