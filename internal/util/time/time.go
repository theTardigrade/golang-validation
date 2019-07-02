package time

import (
	"time"
)

func Age(date time.Time) int {
	now := time.Now().UTC()
	age := now.Year() - date.Year()

	dateYearDay := date.YearDay()
	nowYearDay := now.YearDay()

	const leapYearDay = 60

	if isLeapYear(date) && !isLeapYear(now) && dateYearDay >= leapYearDay {
		dateYearDay--
	} else if isLeapYear(now) && !isLeapYear(date) && nowYearDay >= leapYearDay {
		dateYearDay++
	}

	if nowYearDay < dateYearDay {
		age--
	}

	return age
}

func isLeapYear(date time.Time) bool {
	year := date.Year()

	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}

	return false
}
