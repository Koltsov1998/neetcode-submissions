type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	delimeter := rune(256)

	result := make([]rune, 0, len(strs[0])/2*len(strs))

	result = append(result, rune(len(strs)))

	for _, str := range strs {
		result = append(result, delimeter)
		result = append(result, rune(len(str)))
		for _, symbol := range str {
			result = append(result, symbol)
		}
	}

	return string(result)
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	delimeter := rune(256)

	totalLen := 0
	result := []string{}
	word := []rune{}
	readingLen := false
	firstWord := true

	for _, symbol := range encoded {
		if totalLen == 0 {
			totalLen = int(symbol)
			result = make([]string, 0, totalLen)
			continue
		}
		if readingLen {
			word = make([]rune, 0, int(symbol))
			readingLen = false
			continue
		}
		if symbol == delimeter {
			if !firstWord {
				result = append(result, string(word))
			} else {
				firstWord = false
			}
			readingLen = true
		} else {
			word = append(word, symbol)
		}
	}

	result = append(result, string(word))
	return result
}
