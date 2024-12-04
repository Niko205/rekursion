package strings

// ContainsChain liefert true, falls s eine Kette von count aufeinanderfolgenden
// Vorkommen von symbol enthält.
func ContainsChain(s, symbol string, count int) bool {
	if symbol == "" || count == 0 {
		return true
	}
	chain := Chain(symbol, count)
	return Contains(s, chain)
}

/*
func ContainsChain(s, symbol string, count int) bool {
	if s == "" && symbol == "" || count == 0 {
		return true
	} else if s == "" {
		return false
	}
	if LengthChain(s, symbol) == count {
		return true
	} else {
		return ContainsChain(s[1:], symbol, count)
	}
}

func LengthChain(s, symbol string) int {
	if s == "" {
		return 0
	}
	if s[0] == symbol[0] {
		return LengthChain(s[1:], symbol) + 1
	} else {
		return 0
	}
}
*/
