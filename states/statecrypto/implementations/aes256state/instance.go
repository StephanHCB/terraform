package aes256state

// New creates a new AES256 state encryption wrapper.
func New(configuration []string) (*AES256StateWrapper, error) {
	instance := &AES256StateWrapper{}
	err := instance.parseKeysFromConfiguration(configuration)
	return instance, err
}
