func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1

	maxVolume := 0
	for left != right {
		volume := 0
		if heights[right] < heights[left] {
			volume = (right - left) * heights[right]
		} else {
			volume = (right - left) * heights[left]
		}
		
		if volume > maxVolume {
			maxVolume = volume
		}

		if heights[right] < heights[left] {
			right--
		} else {
			left++
		}
	}

	return maxVolume
}
