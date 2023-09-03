### Expenses
Just a poor computer science student trying to keep track of his money.

#### Quick Start
Move expenses.csv to your desired directory.
```bash
$ make
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
```bash
-f -> filename
-h -> this message
-a -> add an expense
-r -> remove an expense (unimplemented)
-l -> list all expenses this month (unimplemented)
-t -> show current total spent this month
```
