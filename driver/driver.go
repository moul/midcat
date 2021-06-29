package driver

import (
	"context"

	"go.uber.org/zap"
)

type (
	// Message contains data passed from In to Out.
	Message struct {
		Bytes []byte
	}

	// Input defines the interface for an input driver.
	Input interface {
		Open(context.Context, chan Message) error
		Close() error
		String() string
	}

	// Output defines the interface for an output driver.
	Output interface {
		Open() error
		Close() error
		String() string
		Send(Message) error
	}

	// Opts defines the options passed to an input or output constructor.
	Opts struct {
		Args   string
		Logger *zap.Logger
	}

	// InputConstructor is a generic callback constructing an Input.
	InputConstructor func(Opts) (Input, error)

	// OutputConstructor is a generic callback constructing an Input.
	OutputConstructor func(Opts) (Output, error)
)

var (
	// InputDrivers is the list of available In drivers.
	InputDrivers = map[string]InputConstructor{
		"midi": NewMidiInput,
		"nop":  NewNopInput,
		"-":    NewStdioInput,
	}

	// OutputDrivers is the list of available Out drivers.
	OutputDrivers = map[string]OutputConstructor{
		"midi": NewMidiOutput,
		"nop":  NewNopOutput,
		"-":    NewStdioOutput,
	}
)
