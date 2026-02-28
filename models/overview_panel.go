package models

import (
	"anictl/styles"

	tea "charm.land/bubbletea/v2"
)

type OverviewPanel struct {
	Height int
	Width  int
}

func OverviewPanelNew() OverviewPanel {
	return OverviewPanel{}
}

func (o OverviewPanel) Update(msg tea.Msg) (OverviewPanel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		o.Height = msg.Height - 3
		o.Width = msg.Width
	}
	return o, nil
}

func (o OverviewPanel) View(width int) string {
	return styles.Base.
		Height(o.Height).
		Width(width).
		Render("")
}
