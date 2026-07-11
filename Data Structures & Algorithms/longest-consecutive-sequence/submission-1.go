func longestConsecutive(nums []int) int {
	consequitives := make(map[int]int, len(nums))
	longest := 0
	for _, num := range nums {
		if _, ok := consequitives[num]; !ok {
			left := consequitives[num-1]
			right := consequitives[num+1]
			
			sum := left + right + 1

			consequitives[num] = sum
			consequitives[num+right] = sum
			consequitives[num-left] = sum

			if sum > longest {
				longest = sum
			}
		}
	}

	return longest
}
