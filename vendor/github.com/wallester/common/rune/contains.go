package rune

func Contains(haystack []rune, needle rune) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}
