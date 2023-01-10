package main

import (
	"os"

	"github.com/skykosiner/control-lights/pkg/lights"
)

func main() {
	url := "http://10.0.0.2/api/vGeourmApBqx37QJaJUQ4AxboqUjli1Fj3LtTQdY/"
	args := os.Args[1:]

	switch args[0] {
	case "ceiling":
		lights.ToggleLightsCeiling(url)
	case "others":
		lights.ToggleOthers(url)
	case "all":
		lights.ToggleLightsCeiling(url)
		lights.ToggleOthers(url)
	}
}
