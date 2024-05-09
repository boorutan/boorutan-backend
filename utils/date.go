package utils

import "time"

func ParseDanbooruTime(st string) (time.Time, error) {
	date, err := time.Parse("2006-01-02T15:04:05.000-07:00", st)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
