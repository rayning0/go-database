# Golang In-Memory Database

**[Read PDF of tech challenge](https://github.com/rayning0/go-database/blob/master/TechAssessment.pdf)**. To run, type `go run main.go`

Type each command in command line. Hit ENTER after each. Commands are case-insensitive.

See **[sample screen output](https://github.com/rayning0/go-database/blob/master/output.txt)**. See **[software tests](https://github.com/rayning0/go-database/blob/master/src/db/db_test.go)**.

## Commands:

`SET [name] [value]`
- Sets name in the database to given value. Name and value are case-sensitive strings with no spaces.

`GET [name]`
- Prints value for given name. If value not in database, prints N​ULL

`DELETE [name]`
- Deletes name/value pair from database

`COUNT [value]`
- Returns number of names with given value assigned to them. If value not assigned anywhere, prints 0​

`END`
- Exits database

`?`
- Prints list of commands

### The database also supports transactions:

`BEGIN`
- Begins new transaction

`ROLLBACK`
- Rolls back most recent transaction. If no transaction to rollback, prints T​RANSACTION NOT FOUND.

`COMMIT`
- Commits ​all​ open transactions. May not ROLLBACK after COMMIT.
