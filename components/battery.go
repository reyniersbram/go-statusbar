package components

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

)

type Battery struct {
	Ticker
	capacity int
	bat      string
}

func NewBattery(bat string, duration time.Duration) *Battery {
	battery := &Battery{
		Ticker: Ticker{Frequency: duration},
		bat: bat,
	}
	return battery
}

func (b Battery) String() string {
	return strconv.Itoa(b.capacity)
}

func (b *Battery) Refresh() bool {
	base_path := "/sys/class/power_supply"
	path := fmt.Sprintf("%s/%s/%s", base_path, b.bat, "capacity")
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return false
	}
	defer f.Close()
	n, err := fmt.Fscanf(f, "%d", &b.capacity)
	if err != nil || n < 1 {
		log.Println(err)
		return false
	}
	return true
}

