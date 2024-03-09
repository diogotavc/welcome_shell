package main

import (
	"fmt"
	"os/user"
	"welcome_shell/pkg/power"
)

func main() {
	fmt.Printf("%s\n\n%s\n", userString(), consumptionString())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func userString() string {
	user, err := user.Current()
	checkError(err)
	return fmt.Sprintf("Welcome back, \033[1m%s\033[0m!", user.Name)
}

func consumptionString() string {
	var powerMetering [3]string = power.Consumption()
	if powerMetering[0] == "Charging" || powerMetering[0] == "Discharging" {
		return fmt.Sprintf("🗲  %s at about %s.%s Watts", powerMetering[0], powerMetering[1], powerMetering[2])
	}
	return fmt.Sprintf("🗲  %s", powerMetering[0])
}
