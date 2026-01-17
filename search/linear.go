package search


func LinearSearch(arr []int, target int) int {
	// Super slow linear search implementation
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
