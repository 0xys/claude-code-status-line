package encoder

import "fmt"

func ClickableLink(text, url string) string {
	return fmt.Sprintf("\033]8;;%s\033\x07%s\033]8;;\033\x07", url, text)
}
