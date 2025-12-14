package domain

import "context"

type Consumer interface {
	GetEvent(ctx context.Context, eventType string) (uint, interface{}, error)
	MarkEventAsProcessing(ctx context.Context, id uint) error
	MarkEventAsProcessed(ctx context.Context, id uint) error
	MarkEventAsFailed(ctx context.Context, id uint) error
	MarkEventAsPending(ctx context.Context, id uint) error
}
