/**
* Runtime: 44 ms, faster than 21.83% of Go online submissions for Two Sum.
* Memory Usage: 2.9 MB, less than 91.48% of Go online submissions for Two Sum.
 */
func twoSum(nums []int, target int) []int {
	var num1, num2 int
	for i, num := range nums {
		diff := target - num
		for j := i + 1; j < len(nums); j++ {
			if diff == nums[j] {
				num1 = i
				num2 = j
			}
		}
	}
	return []int{num1, num2}
}
