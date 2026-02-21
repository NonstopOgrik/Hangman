package readFromConsoleService

import (
	"fmt"
	"os"
	"strings"
)

func ConsoleReadWord() string {
	var word string
	fmt.Fscan(os.Stdin, &word)
	word = strings.ToLower(word)
	return word
}
