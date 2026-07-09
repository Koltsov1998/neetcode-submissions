type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	delimeter := rune(256)

	result := []rune{}

	for _, str := range strs {
		for _, symbol := range str {
			result = append(result, symbol)
		}
		result = append(result, delimeter)
	}

	return string(result)
}

func (s *Solution) Decode(encoded string) []string {
	delimeter := rune(256)

	result := []string{}
	word := []rune{}
	for _, symbol := range encoded {
		if symbol == delimeter {
			result = append(result, string(word))
			word = []rune{}
		} else {
			word = append(word, symbol)
		}
	}

	return result
}
