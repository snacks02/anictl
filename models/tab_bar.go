package models

import (
	"anictl/styles"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var tabBarNames = []string{"Anime", "Light Novels", "Manga"}

type TabBar struct {
	ActiveTab int
}

func TabBarNew() TabBar {
	return TabBar{}
}

func (t TabBar) Update(msg tea.Msg) (TabBar, tea.Cmd) {
	if msg, ok := msg.(tea.KeyPressMsg); ok {
		switch msg.String() {
		case "1", "2", "3":
			t.ActiveTab = int(msg.String()[0] - '1')
		}
	}
	return t, nil
}

func (t TabBar) View() string {
	renderedTabs := make([]string, len(tabBarNames))
	for index, name := range tabBarNames {
		style := styles.TabInactive
		if index == t.ActiveTab {
			style = styles.TabActive
		}
		renderedTabs[index] = style.Render(name)
	}
	return styles.Base.Render(
		lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...),
	)
}
