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

const batteryTemplateName = "battery"

// Battery represents a system battery and provides method to query and format
// its status. It reads information from the Linux power supply sysfs interface
// located at `/sys/class/power_supply`.
//
// The following template variables are available:
//
//   - Capacity: The current battery capacity (integer 0-100)
//   - ChargingStatus: The charging status of the battery, one of:
//     ["Charging", "Not charging", "Discharging", "Full", "Unknown"]
//   - BatteryName: The system name of the battery (e.g. "BAT0")
//   - Icon: A Unicode icon representing the current battery state
//
// Reference: https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-power
type Battery struct {
	duration       time.Duration
	tmpl           *template.Template
	Capacity       int
	ChargingStatus string
	BatteryName    string
}

// NewBattery initializes a new Battery component that queries the system
// battery with the given name `bat` (e.g. "BAT0")
// The duration specifies the polling interval that should be used.
//
// The tmplString is a Go text/template string that can include Battery fields
// to format its display output.
//
// Example:
//
//	tmpl :=	"{{.Icon}} {{.Capacity}}%"
//	bat := components.NewBattery(30 * time.Second, "BAT0", tmpl)
//
// Returns a pointer to a Battery instance.
func NewBattery(
	duration time.Duration,
	bat string,
	tmplString string,
) *Battery {
	battery := &Battery{
		duration:    duration,
		BatteryName: bat,
		tmpl: template.Must(
			template.New(batteryTemplateName).Parse(tmplString),
		),
	}
	return battery
}

func (b Battery) GetDuration() time.Duration {
	return b.duration
}

// String executes the configured template using the latest battery data and
// returns the formatted output.
//
// If the template fails to execute, the error is logged an an empty string is
// returned.
func (b Battery) String() string {
	return ExecuteTemplate(*b.tmpl, b)
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
	capacityPath := fmt.Sprintf("%s/%s/%s", base_path, b.BatteryName, "capacity")
	statusPath := fmt.Sprintf("%s/%s/%s", base_path, b.BatteryName, "status")
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
	b.ChargingStatus = status
	return true
}

// Icon returns the icon to display, depends on capacity and status.
func (b Battery) Icon() string {
	switch {
	case b.ChargingStatus == "Charging":
		return "\U000f0084"
	case b.ChargingStatus == "Not charging":
		return "\U000f1211"
	case b.ChargingStatus == "Unknown":
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
