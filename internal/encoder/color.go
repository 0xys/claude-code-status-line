package encoder

import "fmt"

func White(text string) string {
	return fmt.Sprintf("\033[97m%s\033[0m", text)
}

func Gray(text string) string {
	return fmt.Sprintf("\033[38;2;128;128;128m%s\033[0m", text)
}

func LightGray(text string) string {
	return fmt.Sprintf("\033[38;2;192;192;192m%s\033[0m", text)
}

func DarkGray(text string) string {
	return fmt.Sprintf("\033[38;2;64;64;64m%s\033[0m", text)
}

func Black(text string) string {
	return fmt.Sprintf("\033[38;2;0;0;0m%s\033[0m", text)
}

func Red(text string) string {
	return fmt.Sprintf("\033[38;2;255;0;0m%s\033[0m", text)
}

func LightRed(text string) string {
	return fmt.Sprintf("\033[38;2;255;128;128m%s\033[0m", text)
}

func Green(text string) string {
	return fmt.Sprintf("\033[38;2;0;255;0m%s\033[0m", text)
}

func LightGreen(text string) string {
	return fmt.Sprintf("\\033[38;2;128;255;128m%s\033[0m", text)
}

func Yellow(text string) string {
	return fmt.Sprintf("\033[38;2;255;255;0m%s\033[0m", text)
}

func Orange(text string) string {
	return fmt.Sprintf("\033[38;2;255;165;0m%s\033[0m", text)
}

func DarkOrange(text string) string {
	return fmt.Sprintf("\033[38;2;200;130;0m%s\033[0m", text)
}

func Blue(text string) string {
	return fmt.Sprintf("\033[38;2;0;0;255m%s\033[0m", text)
}

func LightBlue(text string) string {
	return fmt.Sprintf("\033[38;2;128;128;255m%s\033[0m", text)
}

func Magenta(text string) string {
	return fmt.Sprintf("\033[38;2;255;0;255m%s\033[0m", text)
}

func Cyan(text string) string {
	return fmt.Sprintf("\033[38;2;0;255;255m%s\033[0m", text)
}
