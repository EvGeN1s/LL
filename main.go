package main

import (
	"LL/table"
	"LL/validation"

	"fmt"
)

const tableFilePath = "table.csv"
const line = "-8*-(-3+a*b*-8*(a+-3))"

func main() {
	t := table.IntiTable(tableFilePath)

	fmt.Println(validation.IsLineValid(line, t))
}
