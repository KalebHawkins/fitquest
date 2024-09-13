package types

import (
	"testing"
	"time"
)

func TestExerciseString(t *testing.T) {
	testCases := []struct {
		name         string
		want         string
		exerciseName string
		goal         int
		record       int
		sessions     []*Session
	}{
		{
			name:         "TestNoSessionsLogged",
			exerciseName: "Push-ups",
			goal:         100,
			record:       0,
			sessions:     []*Session{},
			want:         "Push-ups:\n  Goal: 100\n  Record: 0 (No sessions)\n  Percentage Complete: 0%\n\n  Started: N/A\n\n  Sessions:\n    No sessions logged\n",
		},
		{
			name:         "TestNoSessionsNoGoal",
			exerciseName: "Push-ups",
			goal:         0,
			record:       0,
			sessions:     []*Session{},
			want:         "Push-ups:\n  Goal: 0\n  Record: 0 (No sessions)\n  Percentage Complete: N/A (Goal is zero)\n\n  Started: N/A\n\n  Sessions:\n    No sessions logged\n",
		},
		{
			name:         "TestSingleSessionRecord",
			exerciseName: "Push-ups",
			goal:         100,
			record:       20,
			sessions: []*Session{
				{Count: 20, Date: time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC)},
			},
			want: "Push-ups:\n  Goal: 100\n  Record: 20 (2006-01-02)\n  Percentage Complete: 20%\n\n  Started: 2006-01-02\n\n  Sessions:\n    Date: 2006-01-02 Count: 20\n",
		},
		{
			name:         "TestMultiSessionRecord",
			exerciseName: "Push-ups",
			goal:         100,
			record:       75,
			sessions: []*Session{
				{Count: 20, Date: time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC)},
				{Count: 75, Date: time.Date(2006, 01, 06, 0, 0, 0, 0, time.UTC)},
				{Count: 45, Date: time.Date(2006, 01, 04, 0, 0, 0, 0, time.UTC)},
			},
			want: "Push-ups:\n  Goal: 100\n  Record: 75 (2006-01-06)\n  Percentage Complete: 75%\n\n  Started: 2006-01-02\n\n  Sessions:\n    Date: 2006-01-06 Count: 75\n    Date: 2006-01-04 Count: 45\n    Date: 2006-01-02 Count: 20\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := (&Exercise{
				Name:    tc.exerciseName,
				Goal:    tc.goal,
				Session: tc.sessions,
			}).String()

			if tc.want != got {
				t.Errorf("expected `%s`, got `%s`", tc.want, got)
			}
		})
	}
}
