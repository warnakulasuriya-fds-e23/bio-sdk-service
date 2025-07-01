package initializer

import (
	"log"

	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/configtomlreader"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

func Initializer() (sdk *core.SDKCore) {
	config := configtomlreader.ConfigTomlReader()
	s, err := core.NewSDKCore(config.ImagesDir, config.CborDir)
	if err != nil {
		log.Fatal(err.Error())
	}
	sdk = s
	return

}
