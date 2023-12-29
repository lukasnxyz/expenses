package main

import (
    "fmt"
    //"errors"
    "os"
    "strconv"
    //"encoding/csv"
)

type Expense struct {
    name string
    cost float64
    year int
    month int
    day int
}

var months = [12]string {
    "Jan", "Feb", "Mar",
    "Apr", "May", "Jun",
    "Jul", "Aug", "Sep",
    "Oct", "Nov", "Dec",
}

func main() {
    argv := os.Args[1:]

    if len(argv) <= 0 {
        fmt.Println("You must pass an argument!")
        help()

        return
    }

    err := argsHandle(argv)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
}

func argsHandle(argv []string) (error) {
    var fName string

    for i, arg := range argv {
        if arg == "-f" {
            fName = argv[i + 1]

        } else if arg == "-h" {
            help()
            return nil
        } else if arg == "-a" {
            newExpense := Expense {
                "test",
                33.23,
                2023,
                12,
                29,
            }

            err := newExpense.Add(fName)
            if err != nil {
                return err
            }
        } else if arg == "-l" {
            err := listExpenses(fName, 0, 0)
            if err != nil {
                return err
            }
        }
    }

    return nil
}

func listExpenses(fName string, month, year uint) (error) {
    // read
//    csvReader := csv.NewReader(file)
 //   data, err := csvReader.ReadAll()
    //if err != nil {
        //return err
    //}

    //expCost := strconv.FormatFloat(exp.cost, 'f', 2, 64)
    //expenseText := "- " + exp.name + ": $" +  expCost + " on " + strconv.Itoa(exp.day) + " " + strconv.Itoa(exp.month) + " " + strconv.Itoa(exp.year)

    //fmt.Println(data)

    return nil
}

func (exp *Expense) Add(fName string) (error) {
    // prompt user for input

    f, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }

    defer f.Close()

    expCost := strconv.FormatFloat(exp.cost, 'f', 2, 64)
    expenseText := exp.name + ";" +  expCost + ";" + strconv.Itoa(exp.day) + ";" + strconv.Itoa(exp.month) + ";" + strconv.Itoa(exp.year) + "\n"

    _, err = f.WriteString(expenseText)
    if err != nil {
        return err
    }

    return err
}

func (exp *Expense) Print() {
    fmt.Printf("%s: €%.2f on %d %s %d\n", exp.name, exp.cost, exp.day, months[exp.month - 1], exp.year);
}

func help() {
    fmt.Println("Expenses")
    fmt.Println("-f --- expenses file path")
    fmt.Println("-h --- this message")
    fmt.Println("-a --- add an expense")
    fmt.Println("-l --- list all expenses this month with total")
    fmt.Println("-ld --- <month> <yead> list all expenses with total")
    fmt.Println("-r --- remove an expense (unimplemented)")
}
