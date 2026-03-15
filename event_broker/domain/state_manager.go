package domain

import "context"

type StateManager interface {
	MarkEventAsProcessing(ctx context.Context, id uint) error
	MarkEventAsProcessed(ctx context.Context, id uint) error
	MarkEventAsFailed(ctx context.Context, id uint) error
	MarkEventAsPending(ctx context.Context, id uint) error
}
