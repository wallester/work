package strings

import (
	"strings"
	"unicode/utf8"
)

type IFullNameFormatter interface {
	FormatFirstNameInitialLastName(firstNameInitial byte, lastName string) string
	FormatFirstNameInitialMiddleNameInitialLastName(firstNameInitial byte, middleNameInitial byte, lastName string) string
	FormatFirstNameLastName(firstName, lastName string) string
	FormatFirstNameMiddleNameInitialLastName(firstName string, middleNameInitial byte, lastName string) string
	FormatFirstNameMiddleNameLastName(firstName, middleName, lastName string) string
	FormatLastName(lastName string) string

	FormatLastNameEqualFullNameAllowedNumberOfChars(lastName string) string
}

// FormatFullName formats full name according to formats in formatter and returns uppercase string.
func FormatFullName(formatter IFullNameFormatter, firstName, middleName, lastName string) string {
	var fullName string

	if middleName == "" {
		fullName = formatFirstNameLastName(formatter, firstName, lastName)
	} else {
		fullName = formatFullName(formatter, firstName, middleName, lastName)
	}

	return strings.TrimSpace(strings.ToUpper(fullName))
}

// private

const (
	FullNameAllowedNumberOfChars              = 26
	FullNameAllowedNumberOfCharsOneReserved   = 25
	FullNameAllowedNumberOfCharsTwoReserved   = 24
	FullNameAllowedNumberOfCharsThreeReserved = 23
)

func formatFirstNameLastName(formatter IFullNameFormatter, firstName, lastName string) string {
	if utf8.RuneCountInString(firstName+lastName) < FullNameAllowedNumberOfChars {
		return formatter.FormatFirstNameLastName(firstName, lastName)
	}

	switch lastNameLen := utf8.RuneCountInString(lastName); {
	case lastNameLen < FullNameAllowedNumberOfCharsOneReserved:
		return formatter.FormatFirstNameInitialLastName(firstName[0], lastName)

	case lastNameLen > FullNameAllowedNumberOfChars:
		return formatLastNameLongerFullNameAllowedNumberOfChars(formatter, lastName)

	case lastNameLen == FullNameAllowedNumberOfChars:
		return formatter.FormatLastNameEqualFullNameAllowedNumberOfChars(lastName)

	default:
		return formatter.FormatLastName(lastName)
	}
}

func formatFullName(formatter IFullNameFormatter, firstName, middleName, lastName string) string {
	if utf8.RuneCountInString(firstName+middleName+lastName) < FullNameAllowedNumberOfCharsOneReserved {
		return formatter.FormatFirstNameMiddleNameLastName(firstName, middleName, lastName)
	}

	if utf8.RuneCountInString(firstName+lastName) < FullNameAllowedNumberOfCharsTwoReserved {
		return formatter.FormatFirstNameMiddleNameInitialLastName(firstName, middleName[0], lastName)
	}

	switch lastNameLen := utf8.RuneCountInString(lastName); {
	case lastNameLen < FullNameAllowedNumberOfCharsThreeReserved:
		return formatter.FormatFirstNameInitialMiddleNameInitialLastName(firstName[0], middleName[0], lastName)

	case lastNameLen == FullNameAllowedNumberOfCharsThreeReserved || lastNameLen == FullNameAllowedNumberOfCharsTwoReserved:
		return formatter.FormatFirstNameInitialLastName(firstName[0], lastName)

	case lastNameLen > FullNameAllowedNumberOfChars:
		return formatLastNameLongerFullNameAllowedNumberOfChars(formatter, lastName)

	case lastNameLen == FullNameAllowedNumberOfChars:
		return formatter.FormatLastNameEqualFullNameAllowedNumberOfChars(lastName)

	default:
		return formatter.FormatLastName(lastName)
	}
}

func formatLastNameLongerFullNameAllowedNumberOfChars(formatter IFullNameFormatter, lastName string) string {
	char := lastName[FullNameAllowedNumberOfCharsTwoReserved]
	if char == ' ' || char == '-' {
		return formatter.FormatLastName(lastName[:FullNameAllowedNumberOfCharsTwoReserved])
	}

	return formatter.FormatLastName(lastName[:FullNameAllowedNumberOfCharsOneReserved])
}
