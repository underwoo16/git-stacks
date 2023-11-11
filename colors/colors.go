package colors

import "fmt"

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

func Red(str string) string {
	return fmt.Sprintf(red + str + reset)
}

func Green(str string) string {
	return fmt.Sprintf(green + str + reset)
}

func Yellow(str string) string {
	return fmt.Sprintf(yellow + str + reset)
}

func Blue(str string) string {
	return fmt.Sprintf(blue + str + reset)
}

func Purple(str string) string {
	return fmt.Sprintf(purple + str + reset)
}

func Cyan(str string) string {
	return fmt.Sprintf(cyan + str + reset)
}

func Gray(str string) string {
	return fmt.Sprintf(gray + str + reset)
}

func White(str string) string {
	return fmt.Sprintf(white + str + reset)
}
