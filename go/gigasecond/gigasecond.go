// functions for working with gigaseconds
package gigasecond

import "time"

// adds 1 gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
