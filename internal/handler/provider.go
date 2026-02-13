package handler

import "github.com/google/wire"

type Handlers struct {
	PingHandler *PingHandler
}

func NewHandlers(
	ping *PingHandler,
) *Handlers {
	return &Handlers{
		PingHandler: ping,
	}
}

var ProviderSet = wire.NewSet(
	NewHandlers,
	NewPingHandler,
)
