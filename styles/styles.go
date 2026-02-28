package styles

import (
	"anictl/colors"

	"charm.land/lipgloss/v2"
)

var Base = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(colors.Border)

var TabActive = lipgloss.NewStyle().
	Foreground(colors.SelectedForeground).
	Background(colors.Accent).
	Padding(0, 1)

var TabInactive = lipgloss.NewStyle().
	Padding(0, 1)
