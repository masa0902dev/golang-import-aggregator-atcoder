package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserConfirmation(prompt string) bool {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()) == "y"
	}
	return false
}
