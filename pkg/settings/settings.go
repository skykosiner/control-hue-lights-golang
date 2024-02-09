package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Url    string `json:"url"`
	Lights struct {
		Bedroom struct {
			CeilingLights []int `json:"ceilingLights"`
			Others        []int `json:"others"`
			Studio []int `json:"studio"`
		} `json:"bedroom"`
	} `json:"lights"`
}

func SetupConfig() {
	// Check if the config existis
	path := fmt.Sprintf("%s/.config/lights/config.json", os.Getenv("HOME"))

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the file with the defualts from the config.json in the base dir
		src := "./config.json"

		bytesRead, err := ioutil.ReadFile(src)

		if err != nil {
			log.Fatal("Error reading config.json file for config setup", err)
		}

		err = os.MkdirAll(fmt.Sprintf("%s/.config/lights", os.Getenv("HOME")), 0700)

		if err != nil {
			log.Fatal("Error creating new dir", err)
		}

		file, err := os.Create(path)

		if err != nil {
			log.Fatal("Error creating config.json file", err, file)
		}

		err = ioutil.WriteFile(path, bytesRead, 0644)

		if err != nil {
			log.Fatal("Error setting up help file ", err)
		}

	}
}

func ReadConfig() Config {
	var config Config

	bytesRead, err := ioutil.ReadFile(fmt.Sprintf("%s/.config/lights/config.json", os.Getenv("HOME")))

	if err != nil {
		log.Fatal("Error reading config.json file", err)
	}

	err = json.Unmarshal(bytesRead, &config)

	if err != nil {
		log.Fatal("Error getting json config", err)
	}

	return config
}
