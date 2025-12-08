package domain

type Consumer interface {
	GetEvent(eventType string) (uint, interface{}, error)
	MarkEventAsProcessing(id uint) error
	MarkEventAsProcessed(id uint) error
	MarkEventAsFailed(id uint) error
	MarkEventAsPending(id uint) error
}
