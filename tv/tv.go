package tv

import "context"

type TV interface {
	TurnOn(ctx context.Context) error
	TurnOff(ctx context.Context) error
	VolumeUp(ctx context.Context) error
	VolumeDown(ctx context.Context) error
}
