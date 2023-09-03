#!/bin/sh

# Change this to your expenses file path. It needs to be the full path.
expenses_path=/home/ln/doc/expenses.csv

expenses -f $expenses_path $1 $2 $3
