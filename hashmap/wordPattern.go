package hashmap

import "strings"

func wordPattern(pattern string, s string) bool {
	mpWords := make(map[int32]string)
	mpSymbols := make(map[string]int32)

	mas := strings.Split(s, " ")

	if len(mas) != len(pattern) {
		return false
	}

	for index, symbol := range pattern {
		valString, okSymbol := mpWords[symbol]
		if !okSymbol {
			mpWords[symbol] = mas[index]
		} else {
			if valString != mas[index] {
				return false
			}
		}
		valSymbol, okWord := mpSymbols[mas[index]]
		if !okWord {
			mpSymbols[mas[index]] = symbol
		} else {
			if valSymbol != symbol {
				return false
			}
		}

	}
	return true
}
