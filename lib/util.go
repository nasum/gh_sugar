package lib

import "time"

func parseDateFromString(dateSt string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z"
	date, err := time.Parse(layout, dateSt)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
