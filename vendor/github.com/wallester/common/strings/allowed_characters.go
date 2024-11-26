package strings

import (
	"strings"

	"github.com/juju/errors"
	commonrune "github.com/wallester/common/rune"
)

const PersonalizationCenterAllowedCharactersPattern = `(?i)^[a-z0-9'\-./,&æäåáéíøöóõšüúž ]+$`

func ValidatePersonalizationCenterAllowedCharacters(str string) error {
	str = strings.ToLower(str)

	for _, character := range str {
		if !commonrune.Contains(personalizationCenterAllowedCharacters, character) {
			return errors.Errorf("unallowed character found: string_runes=%v allowed_characters_runes=%v", []rune(str), personalizationCenterAllowedCharacters)
		}
	}

	return nil
}

// private

var personalizationCenterAllowedCharacters = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'æ', 'ä', 'å', 'á', 'é', 'í', 'ø', 'ö', 'ó', 'õ', 'š', 'ü', 'ú', 'ž',
	'\'', '-', '/', '.', ',', '&', ' ',
}
