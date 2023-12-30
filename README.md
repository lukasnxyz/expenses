### expenses
Just a poor cs student trying to keep track of his money.

#### Quick Start
Move expenses.csv to your desired directory. Make sure your `$GOPATH` and `$GOBIN` are set
```bash
$ go install
```
Adding an expense is as easy as:
```bash
$ expenses.sh -f /path/to/file.csv -a
```
Seeing expenses and total spent this month:
```bash
$ expenses.sh -f /path/to/file.csv -l
```
<!--Seeing expenses and total spent for specific month and year:
```bash
# Example: August 2023
$ expenses.sh -ld 8 2023
```-->

#### Commands
```text
Expenses
-f --- expenses file path
-h --- this message
-a --- add an expense
-l --- list all expenses this month with total
-ld --- <month> <year> list all expenses with total
```
