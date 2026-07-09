func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1

	for i := 0; i < len(result); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1

	for i := len(result) - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
