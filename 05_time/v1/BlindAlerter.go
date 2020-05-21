package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

/*
type function implements the interface, thus concrete func
is implementor of the interface

That way users of your interface have the option to implement
your interface with just a function; rather than having to create
an empty struct type.

*/
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

/*
	This function bears the signature of function BlindAlerterFunc
	thus implementer the BlindAlerter interface.
*/
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
