package config

import (
	"encoding/json"
	"errors"
	"os"
	"ticket-auction-manager/openstub/internal/env"
)

func GenDefault() Config {
	return Config{
		DeviceName: env.GetHostname(),
		DeviceID:   env.GenStrBytes(32),
	}
}

func GetConfig() (*Config, error) {
	dataPath := env.GetDataPath()
	os.MkdirAll(dataPath, 0755)

	configPath := env.GetSysConfigPath()
	contents, err := os.ReadFile(configPath)
	if errors.Is(err, os.ErrNotExist) {
		defaultConfig := GenDefault()
		defaultFile, _ := json.MarshalIndent(defaultConfig, "", "")
		err = os.WriteFile(configPath, defaultFile, 0640)
		if err != nil {
			return nil, err
		}
		return &defaultConfig, nil
	} else if err != nil {
		return nil, err
	}

	rtnConfig := new(Config)
	json.Unmarshal(contents, &rtnConfig)

	return rtnConfig, nil
}
