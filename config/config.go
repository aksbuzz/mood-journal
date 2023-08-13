package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Filename string `yaml:"filename"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	App struct {
		Version     string `yaml:"version"`
		Environment string `yaml:"environment"`
	} `yaml:"app"`
}

func New(cfgFile string) (*Config, error) {
	file, err := os.Open(cfgFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	if err := yd.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
