package lists

// Liefert zwei Listen:
// - Eine, die alle Elemente aus list enthält, die kleiner oder gleich key sind.
// - Eine, die alle übrigen Elemente aus list enthält.
func Partition(list []int, key int) ([]int, []int) {
	if Empty(list) {
		return []int{}, []int{}
	}
	smaller, larger := Partition(list[1:], key)
	if list[0] <= key {
		return append([]int{list[0]}, smaller...), larger
	}
	return smaller, append([]int{list[0]}, larger...)
}
