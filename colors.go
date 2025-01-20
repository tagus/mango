package mango

type ColorCode string

const (
	ColorCodeReset  ColorCode = "\033[0m"
	ColorCodeRed    ColorCode = "\033[38;2;214;48;49m"
	ColorCodeGreen  ColorCode = "\033[38;2;46;204;113m"
	ColorCodeYellow ColorCode = "\033[38;2;241;196;15m"
	ColorCodeBlue   ColorCode = "\033[38;2;9;132;227m"
	ColorCodePurple ColorCode = "\033[38;2;155;89;182m"
	ColorCodeWhite  ColorCode = "\033[37m"
)

func Colorize(color ColorCode, s string) string {
	return string(color) + s + string(ColorCodeReset)
}

func ColorizeRed(s string) string {
	return Colorize(ColorCodeRed, s)
}

func ColorizeGreen(s string) string {
	return Colorize(ColorCodeGreen, s)
}

func ColorizeYellow(s string) string {
	return Colorize(ColorCodeYellow, s)
}

func ColorizeBlue(s string) string {
	return Colorize(ColorCodeBlue, s)
}

func ColorizePurple(s string) string {
	return Colorize(ColorCodePurple, s)
}

func ColorizeWhite(s string) string {
	return Colorize(ColorCodeWhite, s)
}
