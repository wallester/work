package strings

import "regexp"

const (
	clientNameRestrictedCharacters  string = `[\\?\/*\[\]]+`
	productCodeRestrictedCharacters string = `[\\?\/*\[\]]+`
)

var (
	RxClientNameRestrictedCharacters  = regexp.MustCompile(clientNameRestrictedCharacters)
	RxProductCodeRestrictedCharacters = regexp.MustCompile(productCodeRestrictedCharacters)
)
