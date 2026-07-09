func topKFrequent(nums []int, k int) []int {
	indices := make(map[int]int)
	sorted := make([]int, 0)
	frequencies := make(map[int]int)

	for _, num := range nums {
		if _, ok := frequencies[num]; ok {
			frequencies[num] += 1
			i := indices[num]
			for i > 0 && frequencies[sorted[i-1]] <= frequencies[sorted[i]] {
					temp := sorted[i-1]
					sorted[i-1] = sorted[i]
					sorted[i] = temp
					indices[sorted[i-1]] = i-1
					indices[sorted[i]] = i
					i = indices[num]
				}
		} else {
			sorted = append(sorted, num)
			indices[num] = len(sorted) - 1
			frequencies[num] = 1
		}
	}

	result := make([]int, 0, k)

	for i := 0; i < k; i++ {
		result = append(result, sorted[i])
	}

	return result
}
