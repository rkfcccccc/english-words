package dsync

type stubMutex struct{}

func (m *stubMutex) Lock() error {
	return nil
}

func (m *stubMutex) Unlock() (bool, error) {
	return true, nil
}

type stubClient struct{}

func (c *stubClient) NewMutex(name string) Mutex {
	return &stubMutex{}
}

func NewStubClient() Client {
	return &stubClient{}
}
