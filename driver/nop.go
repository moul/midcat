package driver

func NewNopInput(opts Opts) (Input, error) {
	return &nopInputDriver{}, nil
}

type nopInputDriver struct {
	isOpen bool
}

func (d *nopInputDriver) Open() error    { d.isOpen = true; return nil }
func (d *nopInputDriver) Close() error   { return nil }
func (d *nopInputDriver) IsOpen() bool   { return d.isOpen }
func (d *nopInputDriver) String() string { return "nop" }

func NewNopOutput(opts Opts) (Output, error) {
	return &nopOutputDriver{}, nil
}

type nopOutputDriver struct {
	isOpen bool
}

func (d *nopOutputDriver) Open() error                 { d.isOpen = true; return nil }
func (d *nopOutputDriver) Close() error                { return nil }
func (d *nopOutputDriver) IsOpen() bool                { return d.isOpen }
func (d *nopOutputDriver) String() string              { return "nop" }
func (d *nopOutputDriver) Write(b []byte) (int, error) { return len(b), nil }
