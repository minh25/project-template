package constant

import "regexp"

var (
	FileNameRegex    = regexp.MustCompile(`(?:.+/)?(?P<fileName>(?P<Name>[\w\p{P}\p{S} ]+)\.(?P<Extension>\w+))$`)
	PhoneNumberRegex = regexp.MustCompile(`^(03|05|07|08|09)\d{8}$`)
)
