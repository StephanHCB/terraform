package cryptoconfig

import (
	"os"
	"strings"
)

// TODO clean this up and use json configuration as sketched out
// TODO document use
// TODO mark as experimental feature for now

// Set this environment variable to "ALGORITHM:KEY". Supported algorithms:
//   AES256, then KEY must be exactly 64 hexadecimal lower case characters for the 32 byte AES256 key
var KeyEnvName = "TF_REMOTE_STATE_ENCRYPTION"

// Set this environment variable to "ALGORITHM:KEY". Supported algorithms:
//   AES256, then KEY must be exactly 64 hexadecimal lower case characters for the 32 byte AES256 key
var FallbackKeyEnvName = "TF_REMOTE_STATE_DECRYPTION_FALLBACK"

func Configuration() []string {
	return strings.Split(os.Getenv(KeyEnvName), ":")
}

func ConfiguredImplementation(config []string) string {
	if len(config) > 0 {
		return config[0]
	}
	return ""
}
