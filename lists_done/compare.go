package lists

// Liefert true, falls die beiden Listen gleich sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func ListsEqual(list1, list2 []int) bool {
	if Empty(list1) && Empty(list2) {
		return true
	} else if Empty(list1) != Empty(list2) {
		return false
	}
	if list1[0] != list2[0] {
		return false
	}
	return true && ListsEqual(list1[1:], list2[1:])
}