package usecases

import "context"

// UseCase interface represents command usecase.
type UseCase[Command any] interface {
	Handle(ctx context.Context, cmd Command) error
}
