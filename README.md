### Expenses
Just a poor computer science student trying to keep track of his money.

#### Quick Start
Move expenses.csv to your desired directory.
```bash
$ sudo make install
```
I would suggest adding expeses.sh to your `$PATH` so that you don't have to specify the entire path
to your .csv file.

Adding an expense is as easy as:
```bash
$ expenses.sh -a
```
Seeing total spent this month:
```bash
$ expenses.sh -t
```

#### Commands
```text
Expenses
-f --- expenses file path
-h --- this message
-a --- add an expense
-l --- list all expenses this month with total
-ld --- <month> <yead> list all expenses with total
-r --- remove an expense (unimplemented)
```
