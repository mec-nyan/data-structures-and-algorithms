package printer

import "fmt"

const (
	black = iota
	red
	green
	yellow
	blue
	pink
	cyan
	grey
)

func printInColour(text string, colour int) string {
	return fmt.Sprintf("\x1b[38:5:%dm%s[0m", colour, text)
}

func Black(text string) string {
	return printInColour(text, black)
}

func Red(text string) string {
	return printInColour(text, red)
}

func Green(text string) string {
	return printInColour(text, green)
}

func Yellow(text string) string {
	return printInColour(text, yellow)
}

func Blue(text string) string {
	return printInColour(text, blue)
}

func Pink(text string) string {
	return printInColour(text, pink)
}

func Cyan(text string) string {
	return printInColour(text, cyan)
}

func Grey(text string) string {
	return printInColour(text, grey)
}
