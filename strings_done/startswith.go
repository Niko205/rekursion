package strings

// StartsWith liefert true, falls der string s mit der Sequenz seq beginnt.
func StartsWith(s, seq string) bool {
	if seq == "" {
		return true
	}
	if s == "" {
		return false
	}
	if s[0] == seq[0] {
		return true && StartsWith(s[1:], seq[1:])
	}
	return false
}
