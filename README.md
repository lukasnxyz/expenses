### Expenses
Just a poor cs student trying to keep track of his money.

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
Seeing expenses and total spent this month:
```bash
$ expenses.sh -l
```
Seeing expenses and total spent for specific month and year:
```bash
# Example: August 2023
$ expenses.sh -ld 8 2023
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

#### Todo
- [ ] Implement config.h
	- currency ($, €)
	- .csv location
	- default option to run with no options given
- [ ] See average daily/monthly spending
- [ ] Ncurses interface?
- [ ] Add support for spaces in name of expense
