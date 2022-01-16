func sortedSquares(nums []int) []int {
	left, right := 0, len(nums) - 1
    ans := make([]int, len(nums))
	for i := len(nums)-1; i >= 0; i-- {
		if abs(nums[left]) > abs(nums[right]) {
			ans[i] = nums[left] * nums[left]
			left += 1
		} else {
			ans[i] = nums[right] * nums[right]
			right -= 1
		}
	}

	return ans
}

func abs(num int) int {

	if num > 0 {
		return num
	}

	return -num
}