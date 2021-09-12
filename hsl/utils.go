package hsl

import (
	"time"
)

func NextMonday() string {
	//today := time.Now().Format("2006-01-02")
	today := time.Now()
	var nextMonday time.Time
	currentWeekday := today.Weekday()
	switch currentWeekday {
	case time.Monday:
		nextMonday = today.Add(24 * time.Hour * 7)
	case time.Tuesday:
		nextMonday = today.Add(24 * time.Hour * 6)
	case time.Wednesday:
		nextMonday = today.Add(24 * time.Hour * 5)
	case time.Thursday:
		nextMonday = today.Add(24 * time.Hour * 4)
	case time.Friday:
		nextMonday = today.Add(24 * time.Hour * 3)
	case time.Saturday:
		nextMonday = today.Add(24 * time.Hour * 2)
	case time.Sunday:
		nextMonday = today.Add(24 * time.Hour * 1)
	}
	return nextMonday.Format("2006-01-02")
}
