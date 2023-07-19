package clio

import (
	"github.com/sulaiman-coder/goeventbus"
)

type BusConstructor func(Config) *eventbus.Bus

var _ BusConstructor = newBus

func newBus(_ Config) *eventbus.Bus {
	return eventbus.NewBus()
}
