/**
* Runtime: 20 ms, faster than 69.03% of Go online submissions for First Unique Character in a String.
* Memory Usage: 8.6 MB, less than 5.20% of Go online submissions for First Unique Character in a String.
 */
func firstUniqChar(s string) int {
	counts := make(map[string]int, len(s))
	for _, r := range s {
		counts[string(r)] += 1
	}
	for i, r := range s {
		if counts[string(r)] == 1 {
			return i
		}
	}
	return -1
}
