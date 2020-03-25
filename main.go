package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const PROMPT = ">> "

// Read from input source till we see \n
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		if strings.TrimSpace(strings.ToLower(line)) == "end" {
			break
		}
		fmt.Println(line)
	}
}
