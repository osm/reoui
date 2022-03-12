package reolink

import (
	"time"
)

func toTime(t *Time) time.Time {
	local, _ := time.LoadLocation("Local")

	return time.Date(
		t.Year,
		time.Month(t.Month),
		t.Day,
		t.Hour,
		t.Minute,
		t.Second,
		0,
		local,
	)
}
