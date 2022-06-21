package auto

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func GetConfigPath() (string, error) {
	fullexecpath, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir, _ := filepath.Split(fullexecpath)
	return filepath.Join(dir, "service.json"), nil
}

type Config struct {
	CmdList []Cmd
}

type Cmd struct {
	Name string `json:"name"`
	Start string `json:"start"`
	Stop string `json:"stop"`
	Status bool `json:"status"`
}

func GetConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	conf := &Config{}

	r := json.NewDecoder(f)
	err = r.Decode(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
