package types

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

type Log struct {
	Exercises []*Exercise `json:"exercises"`
}

func (l *Log) String() string {
	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Exercise", Width: 14},
		{Title: "Goal", Width: 6},
		{Title: "Record", Width: 6},
		{Title: "Percentage", Width: 14},
	}

	rows := make([]table.Row, 0)
	for i, entry := range l.Exercises {
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", i+1),
			entry.Name,
			fmt.Sprintf("%d", entry.Goal),
			fmt.Sprintf("%d", entry.Record()),
			fmt.Sprintf("%.0f%%", float64(entry.Record())/float64(entry.Goal)*100)})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(len(l.Exercises)+1),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#008080")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.Foreground(lipgloss.NoColor{}).Bold(false)

	t.SetStyles(s)

	return t.View()
}

func NewLog() *Log {
	return &Log{
		Exercises: make([]*Exercise, 0),
	}
}

func (l *Log) Add(exercise *Exercise) {
	if l.Exercises == nil {
		l.Exercises = make([]*Exercise, 0)
	}

	l.Exercises = append(l.Exercises, exercise)
}

func (l *Log) Remove(id int) {
	id -= 1
	if l.indexInRange(id) {
		l.Exercises = append(l.Exercises[:id], l.Exercises[id+1:]...)
	}
}

func (l *Log) indexInRange(id int) bool {
	return id > 0 || id <= len(l.Exercises)-1
}
