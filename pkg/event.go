package pkg

type Event interface {
	ID() string

	Topic() string

	Content() []byte

	Unmarshal([]byte) error
}
