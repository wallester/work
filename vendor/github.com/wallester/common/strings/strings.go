package strings

import (
	"unicode"
)

func Unique(all []string) []string {
	known := make(map[string]bool)
	var result []string
	for _, s := range all {
		if !known[s] {
			result = append(result, s)
			known[s] = true
		}
	}

	return result
}

func Contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}

func IsUppercase(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}

	return true
}

// Remove removes an element(string) from the slice (case sensitive).
func Remove(slice []string, item string) []string {
	for i, other := range slice {
		if other == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

// RemoveBulk removes multiple elements(string) from the slice (case sensitive).
func RemoveBulk(items []string, toRemove []string) []string {
	itemsToRemove := make(map[string]struct{})
	var result []string

	for _, toRemoveValue := range toRemove {
		itemsToRemove[toRemoveValue] = struct{}{}
	}

	for _, itemValue := range items {
		if _, exists := itemsToRemove[itemValue]; !exists {
			result = append(result, itemValue)
		}
	}

	return result
}

func DuplicateInArray(array []string) bool {
	visited := make(map[string]bool)
	for i := 0; i < len(array); i++ {
		if visited[array[i]] {
			return true
		}

		visited[array[i]] = true
	}

	return false
}
