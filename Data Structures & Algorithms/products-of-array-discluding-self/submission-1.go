func productExceptSelf(nums []int) []int {
	pref := make([]int, len(nums))
	suff := make([]int, len(nums))

	n := len(nums)

	pref[0] = 1
	suff[n-1] = 1

	for i := 1; i < len(pref); i++ {
		pref[i] = pref[i-1]*nums[i-1]
	}
	for i := len(pref) - 2; i >= 0; i-- {
		suff[i] = suff[i+1]*nums[i+1]
	}

	result := make([]int, len(nums))
	for i:=0; i< len(nums);i++{
		result[i] = pref[i] * suff[i]
	}

	return result
}
