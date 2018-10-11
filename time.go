package fake

import (
	"fmt"
	"time"
)

// These example values must use the reference time "Mon Jan 2 15:04:05 MST 2006"
// as described at https://gobyexample.com/time-formatting-parsing
const (
	BaseDate       = "2006-01-02"
	BaseTime       = "15:04:05"
	BaseMonth      = "January"
	BaseYear       = "2006"
	BaseDay        = "Monday"
	BaseDayOfMonth = "_2"
	BaseTimePeriod = "PM"
)

// RandomUnixTime is a helper function returning random Unix time
// within last one year
func UnixTime() int64 {
	return time.Now().Unix() - 31536000 + r.Int63n(31536000)
}

// Timestamp formats DateTime using example Timestamp const
func Timestamp() string {
	return time.Unix(UnixTime(), 0).Format(fmt.Sprintf("%s %s", BaseDate, BaseTime))
}

// Time generates time within last one year
func Time() time.Time {
	return time.Unix(UnixTime(), 0)

}
