package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CommanderSettings type for settings.json
type CommanderSettings struct {
	APIKey        string `json:"apiKey"`
	MaxMemory     string `json:"maxMemory"`
	MinMemory     string `json:"minMemory"`
	CommanderPort string `json:"commanderPort"`
	StartCommand  string `json:"startCommand"`
	NoAuth        bool   `json:"noAuth"`
}

// GetSettingsFromFile converts json file to CommanderSettings type
func GetSettingsFromFile(filepath string) CommanderSettings {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("error:", err)
	}

	var settings CommanderSettings

	err = json.Unmarshal(data, &settings)
	if err != nil {
		fmt.Println("error:", err)
	}

	return settings
}
