func isPalindrome(s string) bool {
	prepared := []rune{}
	caseDiff := 'A' - 'a'
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			prepared = append(prepared, r - caseDiff)
		}

		if r>= 'a' && r <= 'z' || r>= '0' && r <= '9' {
			prepared = append(prepared, r)
		}
	}

	for i := 0; i < len(prepared) / 2; i++ {
		if prepared[i] != prepared[len(prepared) - 1 - i] {
			return false
		}
	}

	return true
}
