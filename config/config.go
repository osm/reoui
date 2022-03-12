package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Camera struct {
	Name             string `yaml:"name"`
	Address          string `yaml:"address"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	LowStreamQuality bool   `yaml:"low_stream_quality"`
}

type Config struct {
	Port               string        `yaml:"port"`
	DataDir            string        `yaml:"data_dir"`
	CleanFilesInterval time.Duration `yaml:"clean_files_interval"`
	SyncInterval       time.Duration `yaml:"sync_interval"`
	Cameras            []Camera      `yaml:"cameras"`
}

func NewConfig(configPath string) (*Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d := yaml.NewDecoder(f)
	config := &Config{}
	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
