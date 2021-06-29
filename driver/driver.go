package driver

import (
	"go.uber.org/zap"
)

type Input interface {
	Open() error
	Close() error
	IsOpen() bool
	String() string
}

type Output interface {
	Open() error
	Close() error
	IsOpen() bool
	String() string

	Write(b []byte) (int, error)
}

type Opts struct {
	Args   string
	Logger *zap.Logger
}

type (
	InputConstructor  func(Opts) (Input, error)
	OutputConstructor func(Opts) (Output, error)
)

var InputDrivers = map[string]InputConstructor{
	"midi": NewMidiInput,
	"nop":  NewNopInput,
}

var OutputDrivers = map[string]OutputConstructor{
	"midi": NewMidiOutput,
	"nop":  NewNopOutput,
}
