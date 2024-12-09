package strings

import (
	"sort"
	"strings"
)

// JoinMapStringValues joins all string values of a specified map using ', ' delimiter.
func JoinMapStringValues(m map[string]string) string {
	result := make([]string, 0, len(m))
	for _, k := range m {
		result = append(result, k)
	}

	// Make output field order predictable
	sort.Strings(result)

	return strings.Join(result, ", ")
}
