package testhelper

import "time"

func RequireTimePtr(s string) *time.Time {
	parse, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		panic(err)
	}
	return &parse
}
