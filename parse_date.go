package main

import (
	"errors"
	"time"

	"github.com/bcampbell/fuzzytime"
)

func parseDate(s string) (*time.Time, error) {
	t, _, err := fuzzytime.Extract(s)

	if err != nil || t.Empty() {
		return nil, errors.New("Couldn't parse the date")
	}

	now := time.Now()

	if !t.Date.HasDay() {
		t.SetDay(now.Day())
	}

	if !t.HasMonth() {
		t.SetMonth(int(now.Month()))
	}

	if !t.HasYear() {
		t.SetYear(now.Year())
	}

	if !t.HasHour() {
		t.SetHour(10)
	}

	if !t.HasMinute() {
		t.SetMinute(0)
	}

	if !t.HasSecond() {
		t.SetSecond(0)
	}

	d := time.Date(t.Year(), time.Month(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, now.Location())
	return &d, nil
}
