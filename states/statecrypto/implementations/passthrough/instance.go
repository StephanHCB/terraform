package passthrough

// create a new passthrough state encryption wrapper (that does nothing)
func New(configuration []string) (*PassthroughStateWrapper, error) {
	return &PassthroughStateWrapper{}, nil
}
