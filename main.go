package main

import (
	"bufio"
	"db"
	"fmt"
	"os"
	"strings"
)

const prompt = ">> "
const commands = `SET [name] [value]: Sets name in database to given value.
Name and value are case-sensitive strings with no spaces.

GET [name]: Prints value for given name. If value not in database, prints N​ULL
DELETE [name]: Deletes name/value pair from database
COUNT [value]: Returns number of names with given value assigned to them. If value not assigned anywhere, prints 0​
END: Exits database

The database supports transactions:

BEGIN: Begins new transaction
ROLLBACK: Rolls back most recent transaction. If no transaction to rollback, prints T​RANSACTION NOT FOUND.
COMMIT: Commits ​all​ open transactions`

// REPL: Read from input source till we see \n
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(prompt)

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := strings.TrimSpace(scanner.Text())

		output, err := db.Eval(line)

		if output == "END" {
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		if output == "HELP" {
			fmt.Println(commands)
		}
	}
}
