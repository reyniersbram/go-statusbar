package components

import (
	"strconv"
	"sync"
	"time"
)

type Counter struct {
	Ticker
	sync.Mutex
	count int
}

func (c *Counter) String() string {
	c.Lock()
	defer c.Unlock()
	return strconv.Itoa(c.count)
}

func (c *Counter) Refresh() bool {
	c.Lock()
	defer c.Unlock()
	c.count += 1
	return true
}

type DelayedCounter struct {
	Counter
	delay time.Duration
}

func (c *DelayedCounter) Refresh() bool {
	time.Sleep(c.delay)
	c.Lock()
	defer c.Unlock()
	c.count += 1
	return true
}

func NewCounter(duration time.Duration) *Counter {
	return &Counter{Ticker: Ticker{Frequency: duration}}
}

func NewDelayedCounter(delay, duration time.Duration) *DelayedCounter {
	return &DelayedCounter{
		Counter: Counter{Ticker: Ticker{Frequency: duration}},
		delay: delay,
	}
}
