package configtomlreader

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/configuration"
)

func ConfigTomlReader() configuration.Configuration {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	projEnv := os.Getenv("PROJENV")
	var tomlpath string
	//check if in local environment
	if projEnv == "local" {
		tomlpath = filepath.Join(workingDir, "config.toml")
	} else {
		tomlpath = filepath.Join(workingDir, "deploymentConfig.toml")
	}
	var config configuration.Configuration
	toml.DecodeFile(tomlpath, &config)
	return config
}
