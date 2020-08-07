/**
* Runtime: 224 ms, faster than 16.00% of Go online submissions for Kth Largest Element in a Stream.
* Memory Usage: 7.3 MB, less than 100.00% of Go online submissions for Kth Largest Element in a Stream.
 */

// This problem was solved with insertion sort, but "コーディング面接対策のために解きたいLeetCode 60問" classified this in "Heap, PriorityQueue" section, so I should rewrite by that.

type KthLargest struct {
	K     int
	Array []int
}

func Constructor(k int, nums []int) KthLargest {
	heap := KthLargest{
		K:     k,
		Array: insertionSort(nums),
	}
	return heap
}

func insertionSort(nums []int) []int {
	for j := 1; j < len(nums); j++ {
		key := nums[j]
		i := j - 1
		for i >= 0 && nums[i] > key {
			nums[i+1] = nums[i]
			i = i - 1
		}
		nums[i+1] = key
	}
	return nums
}
func (this *KthLargest) Add(val int) int {
	this.Array = append(this.Array, val)
	this.Array = insertionSort(this.Array)
	return this.Array[len(this.Array)-this.K]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
