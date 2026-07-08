func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    letters := [26]rune{}

    for i := 0; i < len(s); i++ {
        letters[s[i] - 'a']++
        letters[t[i] - 'a']--
    }

    for _, c := range letters {
        if c != 0 {
            return false
        }
    }

    return true
}
