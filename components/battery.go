package components

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Battery struct {
	duration time.Duration
	Capacity int
	Status   string
	bat      string
	tmpl     *template.Template
}

func NewBattery(bat string, duration time.Duration) *Battery {
	tmpl := template.Must(
		template.New("battery").
			Parse("{{.Icon}} {{.Capacity}}%"))
	battery := &Battery{
		duration: duration,
		bat:      bat,
		tmpl:     tmpl,
	}
	return battery
}

func (b Battery) GetDuration() time.Duration {
	return b.duration
}

func (b Battery) String() string {
	var buf bytes.Buffer
	err := b.tmpl.Execute(&buf, b)
	if err != nil {
		log.Println("template execution error: ", err)
	}
	return buf.String()
}

func readValue(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(data)), nil
}

func (b *Battery) Refresh() bool {
	base_path := "/sys/class/power_supply"
	capacityPath := fmt.Sprintf("%s/%s/%s", base_path, b.bat, "capacity")
	statusPath := fmt.Sprintf("%s/%s/%s", base_path, b.bat, "status")
	capRaw, err := readValue(capacityPath)
	if err != nil {
		log.Printf("Could not read battery capacity: %s\n", err)
	}
	cap, err := strconv.Atoi(capRaw)
	if err != nil {
		log.Printf("Battery capacity is not a number: %s\n", err)
	}
	b.Capacity = cap
	status, err := readValue(statusPath)
	if err != nil {
		log.Printf("Could not read battery status: %s\n", err)
	}
	b.Status = status
	return true
}

// Icon returns the icon to display, depends on capacity and status.
// Possible values for Status are [Charging, Not charging, Discharging, Full,
// Unknown]
// See https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-power
func (b Battery) Icon() string {
	switch {
	case b.Status == "Charging":
		return "\U000f0084"
	case b.Status == "Not charging":
		return "\U000f1211"
	case b.Status == "Unknown":
		return "\U000f0091"
	case b.Capacity < 10:
		return "\U000f007a"
	case b.Capacity < 20:
		return "\U000f007b"
	case b.Capacity < 30:
		return "\U000f007c"
	case b.Capacity < 40:
		return "\U000f007d"
	case b.Capacity < 50:
		return "\U000f007e"
	case b.Capacity < 60:
		return "\U000f007f"
	case b.Capacity < 70:
		return "\U000f0080"
	case b.Capacity < 80:
		return "\U000f0081"
	case b.Capacity < 90:
		return "\U000f0082"
	default:
		return "\U000f0079"
	}
}
