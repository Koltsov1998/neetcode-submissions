func hasDuplicate(nums []int) bool {
    numbers := make(map[int]int, len(nums))

    for _, n := range nums {
        if _, ok := numbers[n]; ok {
            return true
        } else {
            numbers[n] = 1
        }
    }

    return false
}
