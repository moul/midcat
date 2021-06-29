package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/portmididrv"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"moul.io/motd"
	"moul.io/srand"
	"moul.io/zapconfig"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		}
		os.Exit(1)
	}
}

var opts struct {
	debug      bool
	logger     *zap.Logger
	midiDriver midi.Driver
}

func run(args []string) error {
	// setup CLI
	rootFs := flag.NewFlagSet("midcat", flag.ExitOnError)
	rootFs.BoolVar(&opts.debug, "debug", opts.debug, "debug mode")
	root := &ffcli.Command{
		Name:       "midcat",
		FlagSet:    rootFs,
		ShortUsage: "midcat [FLAGS] <ADDRESS>[,OPTS] <ADDRESS>[,OPTS]",
		Exec:       doRoot,
		UsageFunc:  usage,
	}
	if err := root.Parse(args); err != nil {
		return fmt.Errorf("parse error: %w", err)
	}

	// init logger
	{
		rand.Seed(srand.Fast())
		config := zapconfig.New().SetPreset("light-console")
		if opts.debug {
			config = config.SetLevel(zapcore.DebugLevel)
		} else {
			config = config.SetLevel(zapcore.InfoLevel)
		}
		var err error
		opts.logger, err = config.Build()
		if err != nil {
			return fmt.Errorf("logger init: %w", err)
		}
	}

	// run
	{
		if err := root.Run(context.Background()); err != nil {
			return fmt.Errorf("run error: %w", err)
		}
	}

	return nil
}

func usage(c *ffcli.Command) string {
	var b strings.Builder

	fmt.Fprint(&b, ffcli.DefaultUsageFunc(c)+"\n\n")

	fmt.Fprintf(&b, "ADDRESS\n")
	fmt.Fprintf(&b, "  -           stdio\n")
	fmt.Fprintf(&b, "  pipe        echo/fifo\n")
	fmt.Fprintf(&b, "  tcp         ...\n")
	fmt.Fprintf(&b, "  tick        ...\n")
	fmt.Fprintf(&b, "  rand        ...\n")
	fmt.Fprintf(&b, "  udp         ...\n")
	fmt.Fprintf(&b, "  websocket   ...\n")
	fmt.Fprintf(&b, "\n")

	fmt.Fprintf(&b, "OPTS\n")
	fmt.Fprintf(&b, "  debug       ...\n")
	fmt.Fprintf(&b, "  reconnect   ...\n")
	fmt.Fprintf(&b, "  bpm         ...\n")
	fmt.Fprintf(&b, "  quantify    ...\n")
	fmt.Fprintf(&b, "  filter      ...\n")
	fmt.Fprintf(&b, "\n")

	fmt.Fprintf(&b, "HARDWARE\n")
	drv, err := driver.New()
	if err != nil {
		fmt.Fprintf(&b, "  error: %v\n", err)
	} else {
		var errs error
		ins, err := drv.Ins()
		if err != nil {
			errs = multierr.Append(errs, err)
		} else {
			for _, in := range ins {
				fmt.Fprintf(&b, "  IN:  number=%d is-open=%v name=%q\n", in.Number(), in.IsOpen(), in.String())
			}
		}

		outs, err := drv.Outs()
		if err != nil {
			errs = multierr.Append(errs, err)
		} else {
			for _, out := range outs {
				fmt.Fprintf(&b, "  OUT: number=%d is-open=%v name=%q\n", out.Number(), out.IsOpen(), out.String())
			}
		}

		if errs != nil {
			fmt.Fprintf(&b, "error: %v\n", errs)
		}
	}

	return strings.TrimSpace(b.String())
}

func doRoot(ctx context.Context, args []string) error {
	opts.logger.Debug("args", zap.Strings("args", args))
	fmt.Print(motd.Default())
	_ = writer.New
	return nil
}
