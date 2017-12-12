// Package gigasecond fiddles with time by gigaseconds!
package gigasecond

import "time"

// AddGigasecond adds a gigasecond to the given time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Duration(1e9) * time.Second)
	return t
}
