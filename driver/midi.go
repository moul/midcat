package driver

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/rtmididrv"
	"go.uber.org/zap"
)

func NewMidiInput(opts Opts) (Input, error) {
	opts.Logger.Debug("NewMidiInput", zap.Any("args", opts.Args))
	d := &midiInputDriver{}

	var err error
	// FIXME: configure portmididrv.SleepingTime
	d.driver, err = rtmididrv.New()
	if err != nil {
		return nil, fmt.Errorf("portmididrv.New: %w", err)
	}

	d.portNum = 0   // FIXME: parse args
	d.portName = "" // FIXME: parse args

	return d, nil
}

type midiInputDriver struct {
	driver   midi.Driver
	portNum  int
	portName string
	port     midi.In
	lock     sync.Mutex
}

func (d *midiInputDriver) Open(ctx context.Context, ch chan Message) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	var err error
	d.port, err = midi.OpenIn(d.driver, d.portNum, d.portName)
	if err != nil {
		return fmt.Errorf("midi.OpenIn: %w", err)
	}

	err = d.port.SetListener(func(data []byte, deltaMicroseconds int64) {
		ch <- Message{
			Bytes: data,
		}
	})
	if err != nil {
		return fmt.Errorf("recv error: %w", err)
	}

	// FIXME: wait for other causes of issues
	<-ctx.Done()

	return nil
}

func (d *midiInputDriver) Close() error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.port == nil {
		return nil
	}
	if err := d.port.StopListening(); err != nil {
		return fmt.Errorf("port.StopListening: %w", err)
	}
	if err := d.port.Close(); err != nil {
		return fmt.Errorf("port.Close: %w", err)
	}
	d.port = nil
	d.driver.Close()
	return nil
}

func (d *midiInputDriver) String() string {
	return fmt.Sprintf("midi,id=%d", d.port.Number())
}

func NewMidiOutput(opts Opts) (Output, error) {
	opts.Logger.Debug("NewMidiOutput", zap.Any("args", opts.Args))
	d := &midiOutputDriver{logger: opts.Logger}

	var err error
	// FIXME: configure portmididrv.SleepingTime
	d.driver, err = rtmididrv.New()
	if err != nil {
		return nil, fmt.Errorf("portmididrv.New: %w", err)
	}

	d.portNum = 0   // FIXME: parse args
	d.portName = "" // FIXME: parse args

	return d, nil
}

type midiOutputDriver struct {
	driver   midi.Driver
	portNum  int
	portName string
	port     midi.Out
	lock     sync.Mutex
	logger   *zap.Logger
}

func (d *midiOutputDriver) Close() error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.port == nil {
		return nil
	}
	if err := d.port.Close(); err != nil {
		return fmt.Errorf("port.Close: %w", err)
	}
	d.port = nil
	d.driver.Close()
	return nil
}

func (d *midiOutputDriver) Open() error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.port != nil {
		return fmt.Errorf("already open")
	}

	var err error
	d.port, err = midi.OpenOut(d.driver, d.portNum, d.portName)
	if err != nil {
		return fmt.Errorf("midi.OpenOut: %w", err)
	}
	return nil
}

func (d *midiOutputDriver) Send(msg Message) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.port == nil {
		return nil
	}

	d.logger.Debug("write", zap.Any("bytes", msg.Bytes))
	_, err := d.port.Write(msg.Bytes)
	if err != nil {
		return fmt.Errorf("write error: %w", err)
	}
	runtime.Gosched()
	return nil
}

func (d *midiOutputDriver) String() string {
	return fmt.Sprintf("midi,id=%d", d.port.Number())
}