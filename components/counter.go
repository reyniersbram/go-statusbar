package components

import (
	"strconv"
	"time"
)

type Counter struct {
	duration time.Duration
	count    int
}

func (c *Counter) String() string {
	return strconv.Itoa(c.count)
}

func (c *Counter) Refresh() bool {
	c.count += 1
	return true
}

func NewCounter(duration time.Duration) *Counter {
	return &Counter{duration: duration, count: 0}
}

func (c *Counter) GetDuration() time.Duration {
	return c.duration
}
