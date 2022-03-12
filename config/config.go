package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Camera struct {
	Name             string `mapstructure:"name"`
	Address          string `mapstructure:"address"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	LowStreamQuality bool   `mapstructure:"low_stream_quality"`
}

type Config struct {
	Port               string        `mapstructure:"port"`
	DataDir            string        `mapstructure:"data_dir"`
	CleanFilesInterval time.Duration `mapstructure:"clean_files_interval"`
	SyncInterval       time.Duration `mapstructure:"sync_interval"`
	Cameras            []Camera      `mapstructure:"cameras"`
}

func NewConfig(configPath string) (*Config, error) {
	if configPath != "" {
		viper.SetConfigName(filepath.Base(configPath))
		viper.AddConfigPath(filepath.Dir(configPath))
	} else {
		viper.SetConfigName("reoui")
		viper.AddConfigPath("/etc/reoui/")
		viper.AddConfigPath("$HOME/etc/reoui")
		viper.AddConfigPath("./etc")
	}
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("REOUI")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	var cameras map[string]*Camera = make(map[string]*Camera)
	for _, env := range os.Environ() {
		idx := strings.Index(env, "=")
		key := env[0:idx]
		val := env[idx+1:]

		if strings.HasPrefix(key, "REOUI_CAMERA_") {
			// Extract the index from the env key.
			// Expects to be a string like REOUI_CAMERA_0_NAME
			l := len("REOUI_CAMERA_")
			camIdx := key[l : strings.Index(key[l:], "_")+l]
			if _, ok := cameras[camIdx]; !ok {
				cameras[camIdx] = &Camera{}
			}

			switch key[l+len(camIdx)+1:] {
			case "NAME":
				cameras[camIdx].Name = val
				break
			case "ADDRESS":
				cameras[camIdx].Address = val
				break
			case "USERNAME":
				cameras[camIdx].Username = val
				break
			case "PASSWORD":
				cameras[camIdx].Password = val
				break
			case "LOW_STREAM_QUALITY":
				if val == "true" {
					cameras[camIdx].LowStreamQuality = true
				}
				break
			default:
				return nil, fmt.Errorf("error: unknown variable name %s=%s\n", key, val)
			}
		}
	}

	if len(cameras) > 0 {
		conf.Cameras = []Camera{}
		for _, cam := range cameras {
			conf.Cameras = append(conf.Cameras, *cam)
		}
	}

	return &conf, nil
}
