package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Hotkey       string
	SavedScripts map[string]string
}

const configFilename = "go_macros.yaml"

func loadConfig() *Config {
	conf := &Config{
		Hotkey:       "f2",
		SavedScripts: make(map[string]string),
	}
	if _, err := os.Stat(configFilename); os.IsNotExist(err) {
		return conf
	}
	file, _ := os.ReadFile(configFilename)
	yaml.Unmarshal(file, conf)
	return conf
}

func saveConfig(conf *Config) {
	bytes, _ := yaml.Marshal(conf)
	os.WriteFile(configFilename, bytes, 0644)
}
