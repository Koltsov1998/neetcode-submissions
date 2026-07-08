func isAnagram(s string, t string) bool {
    characters := make(map[rune]int, len(s))

    for _, sRune := range s {
        if c, ok := characters[sRune]; ok {
            characters[sRune] = c + 1
        } else {
            characters[sRune] = 1
        }
    }

    for _, tRune := range t {
        if c, ok := characters[tRune]; !ok {
            return false
        } else {
            if c == 1 {
                delete(characters, tRune)
            } else {
                characters[tRune] = c - 1
            }
        }
    }

    if len(characters) > 0 {
        return false
    }

    return true
}
