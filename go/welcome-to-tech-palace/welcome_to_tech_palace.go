package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLine := fmt.Sprintf("%s", strings.Repeat("*", numStarsPerLine))
	return fmt.Sprintf("%s\n%s\n%s", starLine, welcomeMsg, starLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	result := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(result)
}
