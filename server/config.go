package server

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConfig struct {
	ListenAddr string `json:"listen_addr"`
	TLSCert    string `json:"tls_cert"`
	TLSKey     string `json:"tls_key"`
}

func LoadConfig(path string) (*ServerConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config ServerConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
