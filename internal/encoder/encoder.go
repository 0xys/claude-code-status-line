package encoder

import "strings"

func Encode(element ...string) string {
	return strings.Join(element, " ")
}
