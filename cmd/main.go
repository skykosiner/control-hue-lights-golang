package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/skykosiner/control-lights/pkg/lights"
	"github.com/skykosiner/control-lights/pkg/settings"
	"github.com/skykosiner/control-lights/pkg/status"
)

func main() {
	if len(os.Args[1:]) <= 0 {
		fmt.Println("No command line arguments passed in")
		return
	}

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
		// Make sure connocted to the same wifi as the aircon
		networkName := exec.Command("iwgetid", "-r")
		stdOout, err := networkName.Output()

		if err != nil {
			log.Fatal("Error getting network name")
		}

		if strings.TrimSuffix(string(stdOout), "\n") == "The Kosiner's wifi" {
			fmt.Println(status.GetStatus())
		} else {
			fmt.Println("Not connected to correct wifi")
		}
	case "reputationEra":
        lights.ReputationEra(config.Url)
	default:

		number, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatal("Error getting number percent", err)
		}

		lights.SetBright(config.Url, number)
	}
}
