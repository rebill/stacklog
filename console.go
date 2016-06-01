package stacklog

type Brush func(string) string

func NewBrush(color string) Brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

var colors = []Brush{
	NewBrush("1;37"), // Panic	 white
	NewBrush("1;35"), // Fatal   magenta
	NewBrush("1;31"), // Error   red
	NewBrush("1;33"), // Warning yellow
	NewBrush("1;34"), // Info	 blue
	NewBrush("1;32"), // Debug   green
}
