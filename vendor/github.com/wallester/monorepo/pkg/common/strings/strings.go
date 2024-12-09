package strings

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/wallester/monorepo/pkg/common/slices"
)

func Unique(all []string) []string {
	return NewUniqueSlice(all...).Items()
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

func IsLowercase(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			return false
		}
	}

	return true
}

func IsHexadecimal(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) && (uint32(r) < uint32('a') || uint32(r) > uint32('f')) && (uint32(r) < uint32('A') || uint32(r) > uint32('F')) {
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

func FindElement(obj map[string]any, key string) (any, bool) {
	for k, v := range obj {
		if strings.EqualFold(k, key) {
			return v, true
		}

		if m, ok := v.(map[string]any); ok {
			if res, ok := FindElement(m, key); ok {
				return res, true
			}
		}
	}

	return nil, false
}

func Filter(list []string, condition func(string) bool) []string {
	return slices.Select(list, condition)
}

func FirstNotEmpty(strings ...string) string {
	for _, s := range strings {
		if s != "" {
			return s
		}
	}

	return ""
}

// Abbreviate truncates the string to the specified length and adds postfix if the string is longer than the specified length.
// Special cases:
//   - If the specified length is less than or equal to 0, the string is returned unchanged.
//   - If the specified length is greater than the length of the string, the string is returned unchanged.
//   - If the specified length is less than the length of the postfix, the string is returned unchanged.
//
// Example:
//
//	Abbreviate("Hello, World!", 5) => "He [...]"
//	Abbreviate("Hello, World!", 2) => "He"
//	Abbreviate("Hello, World!", 0) => "Hello, World!"
func Abbreviate(s string, maxLength int) string {
	if maxLength <= len(AbbreviationPostfix) || len(s) <= maxLength {
		return s
	}

	return s[:maxLength-len(AbbreviationPostfix)] + AbbreviationPostfix
}

const AbbreviationPostfix = " [...]"

func Shorten(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}

	return s[:maxLength] + AbbreviationPostfix
}

func TrimSliceFromString(s, separator, cutSet string) []string {
	slice := strings.Split(s, separator)
	for i := range slice {
		slice[i] = strings.Trim(slice[i], cutSet)
	}

	return slice
}

// AppendIfMissing appends an element to the slice if it is not already present.
// Example:
//
//	AppendIfMissing([]string{"foo", "bar"}, "baz") => []string{"foo", "bar", "baz"}
//	AppendIfMissing([]string{"foo", "bar"}, "foo") => []string{"foo", "bar"}
//	AppendIfMissing([]string{"foo", "bar"}, "baz", "foo") => []string{"foo", "bar", "baz"}
func AppendIfMissing(list []string, items ...string) []string {
	for _, item := range items {
		if Contains(list, item) {
			continue
		}

		list = append(list, item)
	}

	return list
}

func ToModels[T ~string](list []string) []T {
	result := make([]T, len(list))
	for i, v := range list {
		result[i] = T(v)
	}

	return result
}

func ToStrings[T ~string](list []T) []string {
	result := make([]string, len(list))
	for i, v := range list {
		result[i] = string(v)
	}

	return result
}

func TrimLength(s string, maxLength int) string {
	if len(s) > maxLength {
		return s[:maxLength]
	}

	return s
}

func FromCamelToSnakeCase(s string) string {
	snake := matchFirstWordFromCamelToSnakeCase.ReplaceAllString(s, "${1}_${2}")
	snake = matchRestWordsFromCamelToSnakeCase.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// private

var (
	matchFirstWordFromCamelToSnakeCase = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchRestWordsFromCamelToSnakeCase = regexp.MustCompile("([a-z0-9])([A-Z])")
)
