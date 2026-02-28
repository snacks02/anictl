package models

import (
	"anictl/colors"
	"anictl/queries"
	"context"
	"database/sql"
	"fmt"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type EntriesTable struct {
	Height int
	Width  int
	tables []table.Model
}

type EntriesReloadedMsg struct {
	err  error
	rows [][]table.Row
}

func EntriesTableNew(q *queries.Queries) (EntriesTable, error) {
	anime, err := entriesTableAnimeEntries(q)
	if err != nil {
		return EntriesTable{}, err
	}
	lightNovels, err := entriesTableLightNovelsEntries(q)
	if err != nil {
		return EntriesTable{}, err
	}
	manga, err := entriesTableMangaEntries(q)
	if err != nil {
		return EntriesTable{}, err
	}
	return EntriesTable{
		tables: []table.Model{anime, lightNovels, manga},
	}, nil
}

func (e EntriesTable) Update(msg tea.Msg, activeTab, width int) (EntriesTable, tea.Cmd) {
	switch msg := msg.(type) {
	case EntriesReloadedMsg:
		if msg.err == nil {
			for index := range e.tables {
				e.tables[index].SetRows(msg.rows[index])
			}
		}
	case tea.WindowSizeMsg:
		for index := range e.tables {
			e.tables[index].SetWidth(msg.Width)
			e.tables[index].SetHeight(msg.Height - 5)
		}
	}

	for index := range e.tables {
		count := len(e.tables[index].Rows())
		switch index {
		case 0:
			const favorite, progress, overhead = 8, 8, 10
			titleWidth := width - overhead - favorite - progress
			e.tables[index].SetColumns([]table.Column{
				{Title: fmt.Sprintf("Title (%d)", count), Width: titleWidth},
				{Title: "Favorite", Width: favorite},
				{Title: "Progress", Width: progress},
			})

		case 1:
			const favorite, chapters, volumes, overhead = 8, 10, 10, 12
			titleWidth := width - overhead - favorite - chapters - volumes
			e.tables[index].SetColumns([]table.Column{
				{Title: fmt.Sprintf("Title (%d)", count), Width: titleWidth},
				{Title: "Favorite", Width: favorite},
				{Title: "Chapters", Width: chapters},
				{Title: "Volumes", Width: volumes},
			})

		case 2:
			const favorite, chapters, volumes, overhead = 8, 10, 10, 12
			titleWidth := width - overhead - favorite - chapters - volumes
			e.tables[index].SetColumns([]table.Column{
				{Title: fmt.Sprintf("Title (%d)", count), Width: titleWidth},
				{Title: "Favorite", Width: favorite},
				{Title: "Chapters", Width: chapters},
				{Title: "Volumes", Width: volumes},
			})
		}
	}

	var cmd tea.Cmd
	e.tables[activeTab], cmd = e.tables[activeTab].Update(msg)
	return e, cmd
}

func (e EntriesTable) View(activeTab int) string {
	return e.tables[activeTab].View()
}

func (e EntriesTable) ReloadCmd(q *queries.Queries) tea.Cmd {
	return func() tea.Msg {
		newTable, err := EntriesTableNew(q)
		if err != nil {
			return EntriesReloadedMsg{err: err}
		}
		rows := make([][]table.Row, len(newTable.tables))
		for index := range newTable.tables {
			rows[index] = newTable.tables[index].Rows()
		}
		return EntriesReloadedMsg{rows: rows}
	}
}

func entriesTableAnimeEntries(q *queries.Queries) (table.Model, error) {
	list, err := q.ListAnime(context.Background())
	if err != nil {
		return table.Model{}, err
	}
	rows := make([]table.Row, len(list))
	for index, row := range list {
		rows[index] = table.Row{
			row.Title,
			formatFavorite(row.Favorite),
			formatProgress(row.ProgressCurrent, row.ProgressTotal),
		}
	}
	return styledTableNew(rows, 3), nil
}

func entriesTableLightNovelsEntries(q *queries.Queries) (table.Model, error) {
	list, err := q.ListLightNovels(context.Background())
	if err != nil {
		return table.Model{}, err
	}
	rows := make([]table.Row, len(list))
	for index, row := range list {
		rows[index] = table.Row{
			row.Title,
			formatFavorite(row.Favorite),
			formatProgress(row.ChaptersCurrent, row.ChaptersTotal),
			formatProgress(row.VolumesCurrent, row.VolumesTotal),
		}
	}
	return styledTableNew(rows, 4), nil
}

func entriesTableMangaEntries(q *queries.Queries) (table.Model, error) {
	list, err := q.ListManga(context.Background())
	if err != nil {
		return table.Model{}, err
	}
	rows := make([]table.Row, len(list))
	for index, row := range list {
		rows[index] = table.Row{
			row.Title,
			formatFavorite(row.Favorite),
			formatProgress(row.ChaptersCurrent, row.ChaptersTotal),
			formatProgress(row.VolumesCurrent, row.VolumesTotal),
		}
	}
	return styledTableNew(rows, 4), nil
}

func formatFavorite(favorite int64) string {
	if favorite == 0 {
		return ""
	}
	return "★"
}

func formatProgress(current int64, total sql.NullInt64) string {
	if !total.Valid {
		return fmt.Sprintf("%d", current)
	}
	return fmt.Sprintf("%d/%d", current, total.Int64)
}

func styledTableNew(rows []table.Row, numColumns int) table.Model {
	columns := make([]table.Column, numColumns)
	for index := range columns {
		columns[index] = table.Column{Title: "", Width: 10}
	}

	entriesTable := table.New(
		table.WithRows(rows),
		table.WithColumns(columns),
		table.WithFocused(true),
	)

	entriesTableStyles := table.DefaultStyles()
	entriesTableStyles.Header = entriesTableStyles.Header.
		Bold(false).
		BorderBottom(true).
		BorderForeground(colors.Border).
		BorderStyle(lipgloss.RoundedBorder())
	entriesTableStyles.Selected = entriesTableStyles.Selected.
		Background(colors.Accent).
		Bold(false).
		Foreground(colors.SelectedForeground)
	entriesTable.SetStyles(entriesTableStyles)
	return entriesTable
}
