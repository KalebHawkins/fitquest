package types

import (
	"fmt"
	"sort"
	"strings"
)

type Exercise struct {
	Name    string     `json:"name"`
	Goal    int        `json:"goal"`
	Session []*Session `json:"sessions"`
}

func (e *Exercise) String() string {
	var sb strings.Builder

	sb.WriteString(
		fmt.Sprintf("%s:\n  Goal: %d\n",
			e.Name, e.Goal),
	)

	if len(e.Session) > 0 {
		recordSession := e.Session[0]
		for _, session := range e.Session {
			if session.Count > recordSession.Count {
				recordSession = session
			}
		}
		sb.WriteString(fmt.Sprintf("  Record: %d (%s)\n", recordSession.Count, recordSession.Date.Format("2006-01-02")))
	} else {
		sb.WriteString("  Record: 0 (No sessions)\n")
	}

	if e.Goal > 0 {
		sb.WriteString(fmt.Sprintf("  Percentage Complete: %.0f%%\n", float64(e.Record())/float64(e.Goal)*100))
	} else {
		sb.WriteString("  Percentage Complete: N/A (Goal is zero)\n")
	}

	sb.WriteRune('\n')

	if len(e.Session) > 0 {

		e.SortSessions()

		firstSessionDate := e.Session[len(e.Session)-1].Date
		sb.WriteString(fmt.Sprintf("  Started: %s\n", firstSessionDate.Format("2006-01-02")))

	} else {
		sb.WriteString("  Started: N/A\n")
	}

	sb.WriteRune('\n')
	sb.WriteString("  Sessions:\n")

	if len(e.Session) == 0 {
		sb.WriteString("    No sessions logged\n")
	} else {
		for _, session := range e.Session {
			sb.WriteString(fmt.Sprintf("    Date: %s Count: %d\n", session.Date.Format("2006-01-02"), session.Count))
		}
	}

	return sb.String()
}

func (e *Exercise) Record() int {
	if len(e.Session) > 0 {
		recordSession := e.Session[0]
		for _, session := range e.Session {
			if session.Count > recordSession.Count {
				recordSession = session
			}
		}
		return recordSession.Count
	}

	return 0
}

func (e *Exercise) SortSessions() {
	sort.Slice(e.Session, func(i, j int) bool {
		if e.Session[i].Date != e.Session[j].Date {
			return e.Session[i].Date.After(e.Session[j].Date)
		}
		return e.Session[i].Count > e.Session[j].Count
	})
}
