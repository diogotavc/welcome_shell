package power

import (
	"errors"
	"os"
	"path"
	"strconv"
	"strings"
)

func Consumption() [3]string {
	var formattedConsumption [3]string

	status, err := os.ReadFile(path.Join(detectBattery(), "status"))
	checkError(err)
	consumption := getPowerConsumption()

	formattedConsumption[0], formattedConsumption[1], formattedConsumption[2] =
		strings.TrimSuffix(string(status), "\n"), strconv.Itoa(consumption/10), strconv.Itoa(consumption%10)

	return formattedConsumption
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func detectBattery() string {
	subpath := "/sys/class/power_supply/"
	for _, bat := range [3]string{"BAT0", "BAT1", "BAT2"} {
		if _, err := os.Stat(path.Join(subpath, bat)); err == nil {
			return path.Join(subpath, bat)
		}
	}
	return ""
}

func readInt(filename string) int {
	file, err := os.ReadFile(filename)
	checkError(err)
	read, err := strconv.Atoi(strings.TrimSpace(string(file)))
	checkError(err)
	return read
}

func getPowerConsumption() int {
	batteryPath := detectBattery()
	powerPath := path.Join(batteryPath, "power_now")
	currentPath := path.Join(batteryPath, "current_now")
	energyPath := path.Join(batteryPath, "energy_now")

	if _, err := os.Stat(powerPath); err == nil {
		return readInt(powerPath) / 100000
	} else if errors.Is(err, os.ErrNotExist) {
		return (readInt(currentPath) * readInt(energyPath)) / 100000
	} else {
		return 0
	}
}
