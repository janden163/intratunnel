package client

import (
	"encoding/json"
	"io/ioutil"
)

type ClientConfig struct {
	ServerAddr string `json:"server_addr"`
	LocalAddr  string `json:"local_addr"`
}

func LoadConfig(path string) (*ClientConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config ClientConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
