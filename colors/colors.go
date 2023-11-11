package colors

import "fmt"

var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

func Red(str string) string {
	return fmt.Sprintf(red + str)
}

func Green(str string) string {
	return fmt.Sprintf(green + str)
}

func Yellow(str string) string {
	return fmt.Sprintf(yellow + str)
}

func Blue(str string) string {
	return fmt.Sprintf(blue + str)
}

func Purple(str string) string {
	return fmt.Sprintf(purple + str)
}

func Cyan(str string) string {
	return fmt.Sprintf(cyan + str)
}

func Gray(str string) string {
	return fmt.Sprintf(gray + str)
}

func White(str string) string {
	return fmt.Sprintf(white + str)
}
