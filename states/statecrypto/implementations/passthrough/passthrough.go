package passthrough

type PassthroughStateWrapper struct {
}

func (a *PassthroughStateWrapper) Encrypt(data []byte) ([]byte, error) {
	return data, nil
}

func (a *PassthroughStateWrapper) Decrypt(data []byte) ([]byte, error) {
	return data, nil
}
