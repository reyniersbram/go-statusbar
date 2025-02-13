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
	// GetFrequency returns how many seconds should be between each refresh of the
	// component
	GetFrequency() time.Duration
}

type Ticker struct {
	Frequency time.Duration
}

func (t *Ticker) GetFrequency() time.Duration {
	return t.Frequency
}
