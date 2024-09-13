package types

import (
	"testing"
	"time"
)

func TestSessionString(t *testing.T) {
	testcases := []struct {
		name  string
		want  string
		count int
		date  time.Time
	}{
		{
			name:  "TestPrintSession",
			want:  "Date: 2006-01-02 Count: 20",
			count: 20,
			date:  time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := (&Session{Count: tc.count, Date: tc.date}).String()

			if tc.want != got {
				t.Errorf("expected `%s`, got `%s`", tc.want, got)
			}
		})
	}
}
