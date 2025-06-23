package initializer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/configuration"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

func Initializer() (sdk *core.SDKCore) {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	tomlpath := filepath.Join(workingDir, "config.toml")
	var config configuration.Configuration
	toml.DecodeFile(tomlpath, &config)
	s, err := core.NewSDKCore(config.ImagesDir, config.CborDir)
	if err != nil {
		log.Fatal(err.Error())
	}
	sdk = s
	return

}
