package table

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadCSV(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func IntiTable(filePath string) []State {
	table := ReadCSV(filePath)

	states := make([]State, 0, len(table))

	for _, row := range table {
		state := State{
			Name:            row[0],
			GuideCharacters: getSymbols(row[1]),
			Shift:           stringToBool(row[2]),
			Error:           stringToBool(row[3]),
			NextStateID:     stringToInt(row[4]),
			PushToStack:     stringToBool(row[5]),
			End:             stringToBool(row[6]),
		}

		states = append(states, state)
	}
	return states
}

func stringToInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf(`can not conver %v to int`, str)
	}
	return res
}

func stringToBool(str string) bool {
	return str == "true"
}

func getSymbols(str string) []string {
	argsString := str[1 : len(str)-1]

	return strings.Split(argsString, ", ")
}
