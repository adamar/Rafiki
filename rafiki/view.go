package rafiki

import (
	"fmt"
	"os"
	"os/exec"
)

// ClearScreen clears the screen
//
func ClearScreen() {

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

}

// PrintOrange prints a message to the screen
// in orange
//
func PrintOrange(msg string) {

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", msg)

}
