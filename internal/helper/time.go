package helper

import "time"

func IsoToUnix(iso string) (int64, error) {
	t, err := time.Parse(time.RFC3339, iso)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func UnixToIso(unixTime int64) string {
	t := time.Unix(unixTime, 0).UTC()
	return t.Format(time.RFC3339)
}
