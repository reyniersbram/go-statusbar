package components

import (
	"fmt"
	"time"
)

// Component represents a component in the status bar.
// Each module is responsible for holding and updating its own data.
type Component interface {
	fmt.Stringer
	// Refresh updates the component's internal data.
	// It reports whether the update was successful.
	Refresh() bool
	// GetDuration returns how many seconds should be between each refresh of the
	// component
	GetDuration() time.Duration
}

type Ticker struct {
	Duration time.Duration
}

func (t *Ticker) GetDuration() time.Duration {
	return t.Duration
}
