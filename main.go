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
	m1 := make(db.M1)
	m2 := make(db.M2)
	var trans db.Stack // stack of transactions

	// m1["a"] = "foo"
	// m1["b"] = "foo"
	// m2["foo"] = []string{"a", "b"}
	// trans.Push(m1, m2)
	// element, ok := trans.Pop()
	// if ok {
	// 	fmt.Printf("%+v %+v\n", element.MainMap, element.ReverseMap)
	// }
	// fmt.Println(len(trans))

	// fmt.Printf("%+v\n", trans[0].Mm1)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(prompt)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := strings.TrimSpace(scanner.Text())
		fmt.Println("trans outside PRE: ", trans)
		output, err := db.Eval(line, m1, m2, &trans)
		fmt.Println("trans outside: ", trans)
		if output == "END" {
			os.Exit(0)
		}

		if err != nil {
			fmt.Println(err.Error())
		}

		if output == "HELP" {
			fmt.Println(commands)
		}

		fmt.Println(output)
	}
}
