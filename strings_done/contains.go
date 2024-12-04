package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	if s == "" {
		return false
	}
	if !StartsWith(s, seq) {
		return Contains(s[1:], seq)
	} else {
		return true
	}
}
