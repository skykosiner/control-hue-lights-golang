package main

import (
	"fmt"
	"os"

	"github.com/skykosiner/control-lights/pkg/lights"
	"github.com/skykosiner/control-lights/pkg/settings"
	"github.com/skykosiner/control-lights/pkg/status"
)

func main() {
	config := settings.ReadConfig()
	args := os.Args[1:]

	switch args[0] {
	case "ceiling":
		lights.ToggleLightsCeiling(config.Url)
	case "others":
		lights.ToggleOthers(config.Url)
	case "all":
		lights.ToggleLightsCeiling(config.Url)
		lights.ToggleOthers(config.Url)
	case "setupConfig":
		settings.SetupConfig()
	case "status":
		fmt.Println(status.GetStatus())
	}
}
