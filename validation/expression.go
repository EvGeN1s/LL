package validation

import (
	"LL/table"
	"container/list"
	"fmt"
)

func IsLineValid(line string, table []table.State) bool {
	stack := list.New()
	index := 0
	ch := 0

	for true {
		symbol := getSymbol(line, ch)

		if symbol == "." {
			return true
		}

		state := table[index]
		fmt.Println(index, state.Name, symbol, state.NextStateID)
		if !hasSymbol(symbol, state) {
			if state.Error {
				if stack.Len() > 0 {
					back := stack.Back()
					stack.Remove(back)
					index = back.Value.(int)
					continue
				}

				return false
			} else {
				index++
				continue
			}
		}
		if state.PushToStack {
			stack.PushBack(index + 1)
		}
		if state.Shift {
			ch++
			symbol = getSymbol(line, ch)
		}
		if state.End {
			break
		}
		if state.NextStateID != -1 {
			index = state.NextStateID
		} else {
			back := stack.Back()
			stack.Remove(back)
			index = back.Value.(int)
		}
	}

	return true
}

func getSymbol(line string, ch int) string {
	if ch >= len(line) {
		return "."
	} else {
		return fmt.Sprintf("%c", line[ch])
	}

}

func hasSymbol(symbol string, state table.State) bool {
	for _, ch := range state.GuideCharacters {
		if ch == symbol {
			return true
		}
	}

	return false
}
