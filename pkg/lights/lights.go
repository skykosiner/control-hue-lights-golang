package lights

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/skykosiner/control-lights/pkg/settings"
)

type hueAPI struct {
	On bool `json:"on"`
}

type hueAPiColor struct {
	On  bool `json:"on"`
	Hue int  `json:"hue"`
}

type hueAPIBrightness struct {
	On bool `json:"on"`
	Brightness int `json:"bri"`
}

type hueAPICt struct {
	On bool `json:"on"`
	Ct int `json:"ct"`
}

type State struct {
	On         bool   `json:"on"`
	Brightness int    `json:"bri"`
	Alert      string `json:"alert"`
	Mode       string `json:"mode"`
	Reachable  bool   `json:"reachable"`
}

type Swupdate struct {
	State       string `json:"state"`
	Lastinstall string `json:"lastinstall"`
}

type Capabilities struct {
	Certiffied bool `json:"certiffied"`
	Control    struct {
		MindimLevel int `json:"mindimlevel"`
		MaxLuemen   int `json:"maxluemen"`
	} `json:"control"`
	Streaming struct {
		Renderer bool `json:"renderer"`
		Proxy    bool `json:"porxy"`
	} `json:"streaming"`
}

type Config struct {
	ArcheType string `json:"archetype"`
	Function  string `json:"function"`
	Direction string `json:"direction"`
	Startup   struct {
		Mode       string `json:"mode"`
		Configured bool   `json:"configured"`
	}
}

type ApiResponse struct {
	State           State        `json:"state"`
	Swupdate        Swupdate     `json:"swupdate"`
	Type            string       `json:"string"`
	Name            string       `json:"name"`
	Modelid         string       `json:"modelid"`
	ManufactureName string       `json:"manufacturename"`
	ProductName     string       `json:"productname"`
	Capabilities    Capabilities `json:"capabilities"`
	Config          Config       `json:"config"`
	UniqueId        string       `json:"uniqueid"`
	SwversIon       string       `json:"swversion"`
	SwconfigId      string       `json:"swconfigid"`
	ProductId       string       `json:"productid"`
}

func makeRequest(light int, url string, jsonReq []byte) {
	url = fmt.Sprintf("%slights/%d/state", url, light)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatal("Error making request to lights", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		log.Fatal(string(body))
	}

	defer resp.Body.Close()
}

func setCT(light int, url string, ct int) {
	jsonReq, err := json.Marshal(hueAPICt{On: true, Ct: ct})

	if err != nil {
		log.Fatal("Error setting brightness of lights", err)
	}

	makeRequest(light, url, jsonReq)
}

func setBrightness(light int, url string, brightness int) {
	jsonReq, err := json.Marshal(hueAPIBrightness{On: true, Brightness: brightness})

	if err != nil {
		log.Fatal("Error setting brightness of lights", err)
	}

	makeRequest(light, url, jsonReq)
}

func toggleLight(light int, url string, state bool) {
	jsonReq, err := json.Marshal(hueAPI{On: !state})

	if err != nil {
		log.Fatal("There was an error toggling the light", err)
	}

	makeRequest(light, url, jsonReq)
}

func setColor(light int, url string) {
	jsonReq, err := json.Marshal(hueAPiColor{On: true, Hue: 0})

	if err != nil {
		log.Fatal("There was an error toggling the light", err)
	}

	makeRequest(light, url, jsonReq)
}

func GetCurrentState(url string, light int) State {
	var Response ApiResponse
	url = fmt.Sprintf("%slights/%d", url, light)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error getting current light state", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &Response)

	if err != nil {
		log.Fatal("Error turning string into json", err)
	}

	return Response.State
}

func ToggleLightsCeiling(url string) {
	lights := settings.ReadConfig().Lights.Bedroom.CeilingLights

	for _, v := range lights {
		toggleLight(v, url, GetCurrentState(url, v).On)
	}
}

func ToggleOthers(url string) {
	lights := settings.ReadConfig().Lights.Bedroom.Others

	for _, v := range lights {
		toggleLight(v, url, GetCurrentState(url, v).On)
	}
}

func ReputationEra(url string) {
	lights := settings.ReadConfig().Lights.Bedroom.Others

	for _, v := range lights {
		// Set the lights to red
		setColor(v, url)
	}
}

func SetBright(url string, brightness int) {
	ceilingLights := settings.ReadConfig().Lights.Bedroom.CeilingLights

	for _, light := range ceilingLights {
		setBrightness(light, url, brightness)
	}
}

func SetCt(url string, ct int) {
	ceilingLights := settings.ReadConfig().Lights.Bedroom.CeilingLights

	for _, light := range ceilingLights {
		setCT(light, url, ct)
	}
}
