package types

import (
	"fmt"
	"time"
)

type Session struct {
	Count int       `json:"count"`
	Date  time.Time `json:"data"`
}

func (s *Session) String() string {
	return fmt.Sprintf("Date: %s Count: %d", s.Date.Format("2006-01-02"), s.Count)
}

func NewSession(count int) *Session {
	return &Session{
		Count: count,
		Date:  time.Now(),
	}
}
