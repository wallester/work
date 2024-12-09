package strings

import (
	"strings"

	"github.com/juju/errors"
	"github.com/wallester/monorepo/pkg/common/runes"
)

const PersonalizationCenterAllowedCharactersPattern = `^[a-zA-Z0-9'\-.\/,&æäåáéíøöóõšüúžàèìòùâêîôûãñëïÿėçćńśźāēīōūąęįųčģķļņŗỳǹẁýǵḱĺḿṕŕẃŷĉĝĥĵŝŵẑẽĩũỹṽḧẅẍẗẘẙȳḡăĕĭŏŭğǫǎěǐǒǔďǧȟǰǩľňřťȩḑḩşţőűůǖǘǚǜÆÄÅÁÉÍØÖÓÕŠÜÚŽÀÈÌÒÙÂÊÎÔÛÃÑËÏŸĖÇĆŃŚŹĀĒĪŌŪĄĘĮŲČĢĶĻŅŖỲǸẀÝǴḰĹḾṔŔẂŶĈĜĤĴŜŴẐẼĨŨỸṼḦẄẌT̈W̊Y̊ȲḠĂĔĬŎŬĞǪǍĚǏǑǓĎǦȞJ̌ǨĽŇŘŤȨḐḨŞŢŐŰŮǕǗǙǛ ]+$`

func ValidatePersonalizationCenterAllowedCharacters(str string) error {
	str = strings.ToLower(str)

	for _, character := range str {
		if !runes.Contains(PersonalizationCenterAllowedCharacters, character) {
			return errors.Errorf("unallowed character found: string_runes=%v allowed_characters_runes=%v", []rune(str), PersonalizationCenterAllowedCharacters)
		}
	}

	return nil
}

var PersonalizationCenterAllowedCharacters = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'æ', 'ä', 'å', 'á', 'é', 'í', 'ø', 'ö', 'ó', 'õ', 'š', 'ü', 'ú', 'ž',
	'à', 'è', 'ì', 'ò', 'ù', 'â', 'ê', 'î', 'ô', 'û', 'ã', 'ñ', 'ë', 'ï', 'ÿ', 'ė', 'ç',
	'ć', 'ń', 'ś', 'ź', 'ā', 'ē', 'ī', 'ō', 'ū', 'ą', 'ę', 'į', 'ų', 'č', 'ģ', 'ķ', 'ļ', 'ņ', 'ŗ',
	'ỳ', 'ǹ', 'ẁ', 'ý', 'ǵ', 'ḱ', 'ĺ', 'ḿ', 'ṕ', 'ŕ', 'ẃ', 'ŷ', 'ĉ', 'ĝ', 'ĥ', 'ĵ', 'ŝ', 'ŵ', 'ẑ', 'ẽ', 'ĩ', 'ũ', 'ỹ',
	'ṽ', 'ḧ', 'ẅ', 'ẍ', 'ẗ', 'ẘ', 'ẙ', 'ȳ', 'ḡ', 'ă', 'ĕ', 'ĭ', 'ŏ', 'ŭ', 'ğ', 'ǫ', 'ǎ', 'ě', 'ǐ', 'ǒ', 'ǔ', 'ď', 'ǧ', 'ȟ',
	'ǰ', 'ǩ', 'ľ', 'ň', 'ř', 'ť', 'ȩ', 'ḑ', 'ḩ', 'ş', 'ţ', 'ő', 'ű', 'ů', 'ǖ', 'ǘ', 'ǚ', 'ǜ',
	'\'', '-', '/', '.', ',', '&', ' ',
}
