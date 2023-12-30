package cli

import (
    "fmt"
    "os"
    "strconv"
    "encoding/csv"

    "github.com/systemb4/expenses/expense"

    "github.com/fxtlabs/date"
)

func ArgsHandle(argv []string) (error) {
    var fName string

    for i, arg := range argv {
        if arg == "-f" {
            fName = argv[i + 1]

        } else if arg == "-h" {
            Help()
            return nil
        } else if arg == "-a" {
            var newExpense expense.Expense
            newExpense.UserInit()

            err := newExpense.Add(fName)
            if err != nil {
                return err
            }
            fmt.Println("Expense added!")
            newExpense.Print()
        } else if arg == "-l" {
            err := ListExpenses(fName, 0, 0)
            if err != nil {
                return err
            }
        }
    }

    return nil
}

func ListExpenses(fName string, month, year uint) (error) {
    f, err := os.Open(fName)
    if err != nil {
        return err
    }

    defer f.Close()

    _, err = f.Seek(1, 0)
    if err != nil {
        return nil
    }

    // removes first line of file
    b := make([]byte, 1)
    for string(b) != "\n" {
        f.Read(b)
    }

    csvReader := csv.NewReader(f)
    csvReader.Comma = ';'
    data, err := csvReader.ReadAll()
    if err != nil {
        return err
    }

    today := date.Today()

    var expenses []expense.Expense
    for _, row := range data {
        expMonth, _ := strconv.Atoi(row[3])
        expYear, _ := strconv.Atoi(row[4])
        if expMonth == int(today.Month()) && expYear == int(today.Year()) {
            var exp expense.Expense
            exp.Parse(row)
            expenses = append(expenses, exp)
        }
    }

    fmt.Printf("Expenses for %s %d: €%.2f\n", expense.Months[today.Month() - 1], today.Year(),
        expense.TotalExpenses(expenses))

    for _, exp := range expenses {
        fmt.Print(" - ")
        exp.Print()
    }

    return nil
}

func Help() {
    fmt.Println("Expenses")
    fmt.Println("-f --- expenses file path")
    fmt.Println("-h --- this message")
    fmt.Println("-a --- add an expense")
    fmt.Println("-l --- list all expenses this month with total")
    fmt.Println("-ld --- <month> <year> list all expenses with total")
}
