package driver

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func NewStdioInput(opts Opts) (Input, error) {
	return &stdioInputDriver{}, nil
}

type stdioInputDriver struct{}

func (d *stdioInputDriver) Open(ctx context.Context, ch chan Message) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		out := make([]byte, len(text)/2)
		_, err := fmt.Sscanf(text, "%X", &out)
		if err != nil {
			return fmt.Errorf("scan error: %w", err)
		}
		ch <- Message{
			Bytes: out,
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	return nil
}
func (d *stdioInputDriver) Close() error   { return nil }
func (d *stdioInputDriver) String() string { return "stdio" }

func NewStdioOutput(opts Opts) (Output, error) {
	return &stdioOutputDriver{}, nil
}

type stdioOutputDriver struct{}

func (d *stdioOutputDriver) Open() error    { return nil }
func (d *stdioOutputDriver) Close() error   { return nil }
func (d *stdioOutputDriver) String() string { return "stdio" }
func (d *stdioOutputDriver) Send(msg Message) error {
	_, err := fmt.Fprintf(os.Stdout, "%X\n", msg.Bytes)
	return err
}
