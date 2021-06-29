package driver

import "go.uber.org/zap"

func NewMidiInput(opts Opts) (Input, error) {
	opts.Logger.Debug("NewMidiInput", zap.Any("args", opts.Args))
	return nil, nil
}

func NewMidiOutput(opts Opts) (Output, error) {
	return nil, nil
}
