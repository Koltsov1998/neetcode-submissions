func twoSum(nums []int, target int) []int {
    checkedNumbers := make(map[int]int)
    
    for i, n := range nums {
        if j, ok := checkedNumbers[target - n]; ok {
            return []int{j, i}
        } else {
            checkedNumbers[n] = i
        }
    }

    return []int{}
}
