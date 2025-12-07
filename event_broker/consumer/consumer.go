package consumer

type Consumer interface {
	Consume(event interface{}) error
}

type consumer struct {
}

func NewEventConsumer() Consumer {
	return &consumer{}
}

func (c *consumer) Consume(event interface{}) error {
	return nil
}
