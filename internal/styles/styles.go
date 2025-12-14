package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var HearderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("16")).
	Background(lipgloss.Color("123")).
	PaddingLeft(2).
	PaddingRight(2)

var DocStyle = lipgloss.NewStyle().
	Margin(1, 2)

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("123"))
