package status

import (
	"github.com/skykosiner/control-lights/pkg/lights"
	"github.com/skykosiner/control-lights/pkg/settings"
)

func GetStatus() (string, string) {
	CeilingString := "C: On"
	OtherString := "O: On"

	CeilightState := make(map[int]bool)
	OtherState := make(map[int]bool)

	ceilingLights := settings.ReadConfig().Lights.Bedroom.CeilingLights
	otherLights := settings.ReadConfig().Lights.Bedroom.Others

	for _, v := range ceilingLights {
		power := lights.GetCurrentState(settings.ReadConfig().Url, v).On
		CeilightState[v] = power
	}

	for _, v := range otherLights {
		power := lights.GetCurrentState(settings.ReadConfig().Url, v).On
		OtherState[v] = power
	}

	CeilingLightTrue := true

	for _, v := range CeilightState {
		if !v {
			CeilingLightTrue = false
			break
		}
	}

	OtherLightTrue := true

	for _, v := range OtherState {
		if !v {
			OtherLightTrue = false
			break
		}
	}

	if !CeilingLightTrue {
		CeilingString = "C: Off"
	}

	if !OtherLightTrue {
		OtherString = "O: Off"
	}

	return CeilingString, OtherString
}
