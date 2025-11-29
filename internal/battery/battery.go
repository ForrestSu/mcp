package battery

import (
	"fmt"
)

// ErrNotFound Only ever returned wrapped in ErrFatal.
var ErrNotFound = fmt.Errorf("not found")

// AgnosticState type enumerates possible battery states, using platform agnostic naming.
type AgnosticState int8

const (
	// Undefined specifies a state that was returned by the controller, but there is no
	// platform agnostic mapping for it.
	// This generally shouldn't happen, if it does consider opening a report for the library
	// (ideally with the contents of `State.Explain()` call as well).
	Undefined AgnosticState = -1
	// Unknown specifies a literal unknown state returned by the controller.
	// This state is also considered the "default", therefore it will be set when an Error
	// was returned from the Get/GetAll call.
	Unknown AgnosticState = iota - 1
	Empty
	Full
	Charging
	Discharging
	// Idle specifies a state where battery is in "capacity saving" mode.
	// It usually means that it sits idle at around 80% charge while power source is plugged in.
	Idle
)

var states = map[AgnosticState]string{
	Undefined:   "Undefined",
	Unknown:     "Unknown",
	Empty:       "Empty",
	Full:        "Full",
	Charging:    "Charging",
	Discharging: "Discharging",
	Idle:        "Idle",
}

func (s AgnosticState) String() string {
	return states[s]
}

type State struct {
	Raw      AgnosticState
	specific string
}

func (s State) Explain() string {
	return s.specific
}

func (s State) String() string {
	return s.Raw.String()
}

func (s State) GoString() string {
	return fmt.Sprintf("%s (%s)", s.Raw, s.specific)
}

type Battery struct {
	CurrentCapacity int
	MaxCapacity     int
	DesignCapacity  int
	State           State
	ChargeRate      float64
	Voltage         float64
	CycleCount      int // 循环次数
}

// GetAll returns information about all batteries in the system.
//
// If error != nil, it will be either ErrFatal or Errors.
// If error is of type Errors, it is guaranteed that length of both returned slices is the same and that i-th error coresponds with i-th battery structure.
func GetAll() ([]*Battery, error) {
	return systemGetAll()
}
