package lights

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type hueAPI struct {
	On bool `json:"on"`
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

func opisateState(state bool) hueAPI {
	if state {
		return hueAPI{On: false}
	}

	return hueAPI{On: true}
}

func toggleLight(light int, url string, state bool) {
	jsonReq, err := json.Marshal(opisateState(state))

	if err != nil {
		log.Fatal("There was an error toggling the light", err)
	}

	url = fmt.Sprintf("%slights/%d/state", url, light)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))

	if err != nil {
		log.Fatal("There was an issue turnning off the lights", err)
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
	lights := map[string]int{
		"ceilingLightOne":   20,
		"ceilingLightTwo":   21,
		"ceilingLightThree": 22,
		"ceilingLightFour":  23,
		"ceilingLightFive":  24,
		"ceilingLightSix":   25,
	}

	for k, v := range lights {
		fmt.Println("changeing light", k)
		toggleLight(v, url, GetCurrentState(url, v).On)
	}
}

func ToggleOthers(url string) {
	lights := map[string]int{
		"deskLamp": 16,
		// "deskPlug": 6,
		"headBoard": 1,
		"underSide": 7,
		"desk":      14,
		"lamp":      5,
	}

	for k, v := range lights {
		fmt.Println("changeing light", k)
		toggleLight(v, url, GetCurrentState(url, v).On)
	}
}
