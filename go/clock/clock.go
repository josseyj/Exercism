//Package clock provides a clock
package clock

import (
	"fmt"
	"math"
)

//Clock type
type Clock struct {
	hour int
	min  int
}

//New creates a new clock
func New(hour int, min int) Clock {
	return Clock{hour: 0, min: 0}.Add(60*hour + min)
}

//Add adds mins to the clock
func (clock Clock) Add(min int) Clock {
	min = min % (24 * 60)
	totalMins := clock.hour*60 + clock.min + min
	clock.hour = (24 + int(math.Floor(float64(totalMins)/60))) % 24
	clock.min = (60 + totalMins%60) % 60
	return clock
}

//Subtract subtracts mins from the clock
func (clock Clock) Subtract(min int) Clock {
	return clock.Add(-min)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.min)
}
