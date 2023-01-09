package lights

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type hueAPI struct {
	On bool `json:"on"`
}

func toggleLight(light int, url string, state bool) {
	jsonReq, err := json.Marshal(&hueAPI{On: state})

	if err != nil {
		log.Fatal("There was an error toggling the light", err)
	}

	url = fmt.Sprintf("%slights/%d/state", url, light)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))

	if err != nil {
		log.Fatal("There was an issue turnning off the lights `hue.go`", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
}

func getCurrentState() bool {
}

func ToggleLights() {
	url := os.Getenv("HUE_URL")
	lights := map[string]int{
		"ceilingLightOne":   10,
		"ceilingLightTwo":   8,
		"ceilingLightThree": 18,
		"ceilingLightFour":  17,
		"ceilingLightFive":  9,
		"ceilingLightSix":   11,
	}

	for _, v := range lights {
		toggleLight(v, url, getCurrentState())
	}
}
