package driver

import "context"

type nopInputDriver struct{}

func NewNopInput(opts Opts) (Input, error)                         { return &nopInputDriver{}, nil }
func (d *nopInputDriver) Open(context.Context, chan Message) error { return nil }
func (d *nopInputDriver) Close() error                             { return nil }
func (d *nopInputDriver) String() string                           { return "nop" }

type nopOutputDriver struct{}

func NewNopOutput(opts Opts) (Output, error)  { return &nopOutputDriver{}, nil }
func (d *nopOutputDriver) Open() error        { return nil }
func (d *nopOutputDriver) Close() error       { return nil }
func (d *nopOutputDriver) String() string     { return "nop" }
func (d *nopOutputDriver) Send(Message) error { return nil }
