package main

import (
	"anictl/models"
	"anictl/queries"
	"anictl/styles"
	"context"
	"database/sql"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type synchronizeDoneMsg struct {
	err error
}

func synchronizeCmd(db *sql.DB, q *queries.Queries) tea.Cmd {
	return func() tea.Msg {
		if err := synchronize(context.Background(), db, q); err != nil {
			return synchronizeDoneMsg{err: err}
		}
		return synchronizeDoneMsg{}
	}
}

type Model struct {
	db            *sql.DB
	entriesTable  models.EntriesTable
	err           error
	overviewPanel models.OverviewPanel
	queries       *queries.Queries
	synchronizing bool
	tabBar        models.TabBar
}

func (model Model) Init() tea.Cmd { return nil }

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	model.overviewPanel, cmd = model.overviewPanel.Update(msg)

	switch msg := msg.(type) {
	case synchronizeDoneMsg:
		model.synchronizing = false
		if msg.err != nil {
			model.err = msg.err
			return model, tea.Quit
		}
		return model, model.entriesTable.ReloadCmd(model.queries)

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return model, tea.Quit
		case "d":
			if model.synchronizing {
				return model, cmd
			}
			model.synchronizing = true
			return model, tea.Batch(cmd, synchronizeCmd(model.db, model.queries))
		}
	}

	model.tabBar, cmd = model.tabBar.Update(msg)
	cmd = tea.Batch(cmd, cmd)

	model.entriesTable, cmd = model.entriesTable.Update(
		msg,
		model.tabBar.ActiveTab,
		model.overviewPanel.Width-model.overviewPanel.Width/4,
	)
	cmd = tea.Batch(cmd, cmd)

	return model, cmd
}

func (model Model) View() tea.View {
	entriesTableView := styles.Base.
		Width(model.overviewPanel.Width - model.overviewPanel.Width/4).
		Render(model.entriesTable.View(model.tabBar.ActiveTab))
	overviewPanelView := model.overviewPanel.View(
		model.overviewPanel.Width - lipgloss.Width(entriesTableView),
	)
	tabBarView := model.tabBar.View()

	layout := lipgloss.JoinVertical(
		lipgloss.Left,
		tabBarView,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			entriesTableView,
			overviewPanelView,
		),
	)

	view := tea.NewView(layout)
	view.AltScreen = true
	return view
}

func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	q, db, err := openDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	entriesTable, err := models.EntriesTableNew(q)
	if err != nil {
		return err
	}

	model := Model{
		db:            db,
		entriesTable:  entriesTable,
		overviewPanel: models.OverviewPanelNew(),
		queries:       q,
		tabBar:        models.TabBarNew(),
	}
	result, err := tea.NewProgram(model).Run()
	if err != nil {
		return err
	}
	if finalModel, ok := result.(Model); ok {
		return finalModel.err
	}
	return nil
}
