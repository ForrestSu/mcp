package battery

import (
	"fmt"
	"strings"
)

// GetBatteryInfo 格式化电池信息
func GetBatteryInfo() (string, error) {
	// GetAll() copy from github.com/distatus/battery v0.11.0
	batteries, err := GetAll()
	if err != nil {
		return "", fmt.Errorf("failed to get battery info: %v", err)
	}
	buff := strings.Builder{}
	for i, b := range batteries {
		health := fmt.Sprintf("%.2f%%", 100*float64(b.MaxCapacity)/float64(b.DesignCapacity))
		one := fmt.Sprintf("<%d>: MaxCapcity %d (desgin: %d), "+
			"State: %s, "+
			"CycleCount: %d, "+
			"health: %s\n",
			i, b.MaxCapacity, b.DesignCapacity,
			b.State,
			b.CycleCount,
			health,
		)
		buff.WriteString(one)
	}
	return buff.String(), nil
}
