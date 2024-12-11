package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	if len(list) == 0 {
		return -1
	}
	mid := len(list) / 2
	if list[mid] == x {
		return mid
	}
	if x < list[mid] {
		return FindSorted(list[:mid], x)
	}
	result := FindSorted(list[mid+1:], x)

	if result == -1 {
		return -1
	}
	return result + mid + 1
}
