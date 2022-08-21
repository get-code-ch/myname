package config

import (
	"encoding/json"
	"errors"
	"os"
	"runtime"
)

type Config struct {
	ZonesFilesPath string `json:"zones_files_path"`
}

func NewConfig(configFile string) (*Config, error) {
	// Creating new config object
	config := new(Config)

	// Checking config file parameters
	cnf := configFile
	if cnf == "" {
		switch runtime.GOOS {
		case "linux", "darwin":
			cnf = "/etc/myname.json"
			break
		case "windows":
			appdata := os.Getenv("LOCALAPPADATA")
			cnf = appdata + "/myname/myname.json"
			break
		default:
			return nil, errors.New("error missing config file (" + configFile + ")")
		}
	}
	if _, err := os.Stat(cnf); os.IsNotExist(err) {
		return nil, errors.New("error loading config file (" + cnf + ")")
	}

	if configData, err := os.ReadFile(cnf); err == nil {
		if err := json.Unmarshal(configData, config); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return config, nil
}
