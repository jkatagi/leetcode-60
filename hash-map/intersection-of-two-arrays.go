/**
* Runtime: 0 ms, faster than 100.00% of Go online submissions for Intersection of Two Arrays.
* Memory Usage: 3 MB, less than 60.53% of Go online submissions for Intersection of Two Arrays.
 */
func intersection(nums1 []int, nums2 []int) []int {
	var uniq []int
	uniq1 := make(map[int]bool, len(nums1))
	uniq2 := make(map[int]bool, len(nums2))

	for _, num1 := range nums1 {
		if !uniq1[num1] {
			uniq1[num1] = true
		}
	}

	for _, num2 := range nums2 {
		if !uniq2[num2] {
			uniq2[num2] = true
		}
	}

	for num1, _ := range uniq1 {
		if uniq2[num1] {
			uniq = append(uniq, num1)
		}
	}
	return uniq
}
