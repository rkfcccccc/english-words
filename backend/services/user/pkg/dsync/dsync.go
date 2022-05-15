package dsync

type Mutex interface {
	Lock() error
	Unlock() (bool, error)
}

type Client interface {
	NewMutex(name string) Mutex
}
