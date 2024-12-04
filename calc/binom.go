package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	if k == 0 {
		return 1
	} else if k < n/2 {
		return n / BinomialCoefficient(n, k-1)
	}
	return n * BinomialCoefficient(n, k-1)
}
