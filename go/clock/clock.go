// type Clock                      // define the clock type
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock

package clock

import "fmt"

// Clock is a date-less time store. We'll store time as seconds, under the hood.
type Clock uint

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * 60
	minutesPerHour   = 60
	secondsPerDay    = 24 * secondsPerHour
)

// New creates a new clock!
func New(hour, minute int) Clock {
	var c Clock
	return c.Add((hour * minutesPerHour) + minute)
}

// Add minutes to the clock
func (c Clock) Add(minute int) Clock {
	seconds := int(c) + (minute * secondsPerMinute)

	// Make sure we're within 24-hour bounds
	seconds %= secondsPerDay

	// Handle negatives
	if seconds < 0 {
		seconds += secondsPerDay
	}

	return Clock(seconds)
}

// String representation of the clock
func (c Clock) String() string {
	hours := c / secondsPerHour
	minutes := (c % secondsPerHour) / secondsPerMinute

	return fmt.Sprintf(
		"%02d:%02d",
		hours,
		minutes,
	)
}
