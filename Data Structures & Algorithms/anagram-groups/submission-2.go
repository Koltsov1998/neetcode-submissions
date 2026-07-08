func groupAnagrams(strs []string) [][]string {
	groups := make(map[int64][]string)

	for _, str := range strs {
		code := calcStrCode(str)

		if stsGroup, ok := groups[code]; ok {
			groups[code] = append(stsGroup, str)
		} else {
			groups[code] = []string {str}
		}
	}

	result := make([][]string, 0)

	for _, g := range groups {
		result = append(result, g)
	}

	return result
}

func calcStrCode(str string) int64 {
	simpleNumbers := map[rune]int64{
		'a': 2,
		'b': 3,
		'c': 5,
		'd': 7,
		'e': 11,
		'f': 13,
		'g': 17,
		'h': 19,
		'i': 29,
		'j': 31,
		'k': 37,
		'l': 41,
		'm': 43,
		'n': 47,
		'o': 53,
		'p': 59,
		'q': 61,
		'r': 67,
		's': 71,
		't': 73,
		'u': 79,
		'v': 83,
		'w': 89,
		'x': 97,
		'y': 101,
		'z': 103,
	}

	result := int64(1)
	for _, s :=range str {
		result *= simpleNumbers[s]
	}

	return result
}
