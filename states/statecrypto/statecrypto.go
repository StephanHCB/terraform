package statecrypto

import (
	"github.com/hashicorp/terraform/states/statecrypto/cryptoconfig"
	"github.com/hashicorp/terraform/states/statecrypto/implementations/aes256state"
	"github.com/hashicorp/terraform/states/statecrypto/implementations/passthrough"
	"log"
)

func StateCryptoWrapper() StateCryptoProvider {
	var implementation StateCryptoProvider
	var err error = nil

	config := cryptoconfig.Configuration()
	switch name := cryptoconfig.ConfiguredImplementation(config); name {
	case "AES256":
		implementation, err = aes256state.New(config)
	default:
		implementation, err = passthrough.New(config)
	}

	if err != nil {
		log.Fatalf("error configuring state file crypto: %v", err)
	}

	return implementation
}
