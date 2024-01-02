package expense

import (
    "fmt"
    "os"
    "strconv"

    "github.com/fxtlabs/date"
)

type Expense struct {
    name string
    cost float64
    year int
    month int
    day int
}

// [...] is an array as well [4] and [n]
// [] is a slice
var (
    Months = [...]string {
    "Jan", "Feb", "Mar",
    "Apr", "May", "Jun",
    "Jul", "Aug", "Sep",
    "Oct", "Nov", "Dec",
    }
)

func (exp Expense) Add(fName string) (error) {
    f, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }

    defer f.Close()

    // user Expense.Parse here
    expCost := strconv.FormatFloat(exp.cost, 'f', 2, 64)
    expenseText := exp.name + ";" +  expCost + ";" + strconv.Itoa(exp.day) + ";" +
        strconv.Itoa(exp.month) + ";" + strconv.Itoa(exp.year) + "\n"

    _, err = f.WriteString(expenseText)
    if err != nil {
        return err
    }

    return err
}

func TotalExpenses(expenses []Expense) (total float64) {
    for _, exp := range expenses {
        total += exp.cost
    }

    return
}

func (exp *Expense) Parse(data []string) {
    exp.name = data[0]
    exp.cost, _ = strconv.ParseFloat(data[1], 64) // interfaces here for any type?
    exp.year, _ = strconv.Atoi(data[2]) // check for errors next 3 lines
    exp.month, _= strconv.Atoi(data[3])
    exp.day, _ = strconv.Atoi(data[4])
}

// Cli dependent functions below here
func (exp *Expense) UserInit() {
    today := date.Today()
    var expName string
    var expCost float64

    fmt.Print("Expense name: ")
    fmt.Scan(&expName)

    fmt.Print("Expense cost: ")
    // check that it's a posotive number (float or int / +)
    fmt.Scan(&expCost)

    exp.name = expName
    exp.cost = expCost
    exp.year = int(today.Year())
    exp.month = int(today.Month())
    exp.day = int(today.Day())
}

func (exp Expense) Print() {
    fmt.Printf("%s: €%.2f on %d %s %d\n", exp.name, exp.cost, exp.day,
        Months[exp.month - 1], exp.year)
}
