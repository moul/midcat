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
	"gitlab.com/gomidi/portmididrv"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"moul.io/midcat/driver"
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
	Debug      bool
	rootLogger *zap.Logger
	mainLogger *zap.Logger
	midiDriver midi.Driver
}

func run(args []string) error {
	// setup CLI
	rootFs := flag.NewFlagSet("midcat", flag.ExitOnError)
	rootFs.BoolVar(&opts.Debug, "debug", opts.Debug, "debug mode")
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
		if opts.Debug {
			config = config.SetLevel(zapcore.DebugLevel)
		} else {
			config = config.SetLevel(zapcore.InfoLevel)
		}
		var err error
		opts.rootLogger, err = config.Build()
		if err != nil {
			return fmt.Errorf("logger init: %w", err)
		}
		opts.mainLogger = opts.rootLogger.Named("main")
	}

	// run
	{
		if err := root.Run(context.Background()); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func usage(c *ffcli.Command) string {
	var b strings.Builder

	fmt.Fprint(&b, motd.Default()+"\n\n")

	fmt.Fprint(&b, ffcli.DefaultUsageFunc(c)+"\n\n")

	fmt.Fprintf(&b, "ADDRESS\n")
	// FIXME: dynamic
	fmt.Fprintf(&b, "  midi        midi port id=FIRST\n")
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
	drv, err := portmididrv.New()
	if err != nil {
		fmt.Fprintf(&b, "  error: %v\n", err)
	} else {
		var errs error
		ins, err := drv.Ins()
		if err != nil {
			errs = multierr.Append(errs, err)
		} else {
			for _, in := range ins {
				fmt.Fprintf(&b, "  IN:  id=%d is-open=%v name=%q\n", in.Number(), in.IsOpen(), in.String())
			}
		}

		outs, err := drv.Outs()
		if err != nil {
			errs = multierr.Append(errs, err)
		} else {
			for _, out := range outs {
				fmt.Fprintf(&b, "  OUT: id=%d is-open=%v name=%q\n", out.Number(), out.IsOpen(), out.String())
			}
		}

		if errs != nil {
			fmt.Fprintf(&b, "error: %v\n", errs)
		}
	}

	return strings.TrimSpace(b.String())
}

func doRoot(ctx context.Context, args []string) error {
	if len(args) != 2 {
		return flag.ErrHelp
	}

	opts.mainLogger.Debug(
		"init",
		zap.Strings("args", args),
		zap.Any("opts", opts),
	)

	// init IO drivers
	var (
		input  driver.Input
		output driver.Output
	)
	{
		var errs error

		var err error
		input, err = initInputInstance(args[0])
		if err != nil {
			errs = multierr.Append(errs, fmt.Errorf("input: %w", err))
		}
		output, err = initOutputInstance(args[1])
		if err != nil {
			errs = multierr.Append(errs, fmt.Errorf("output: %w", err))
		}

		if errs != nil {
			return fmt.Errorf("init drivers: %w", errs)
		}
	}

	// pipe
	{
		fmt.Println(input, output)
	}

	_ = writer.New
	return nil
}

func initInputInstance(arg string) (driver.Input, error) {
	parts := strings.SplitN(arg, ",", 2)
	var driverName, driverOpts string
	if len(parts) == 2 {
		driverName, driverOpts = parts[0], parts[1]
	} else {
		driverName = parts[0]
	}
	opts.mainLogger.Debug(
		"input",
		zap.String("driver", driverName),
		zap.String("opts", driverOpts),
	)
	constructor, found := driver.InputDrivers[driverName]
	if !found {
		return nil, fmt.Errorf("unknown driver name: %q", driverName)
	}
	instance, err := constructor(driver.Opts{
		Logger: opts.rootLogger.Named("out"),
		Args:   driverOpts,
	})
	if err != nil {
		return nil, fmt.Errorf("constructor error: %w", err)
	}
	return instance, nil
}

func initOutputInstance(arg string) (driver.Output, error) {
	parts := strings.SplitN(arg, ",", 2)
	var driverName, driverOpts string
	if len(parts) == 2 {
		driverName, driverOpts = parts[0], parts[1]
	} else {
		driverName = parts[0]
	}
	opts.mainLogger.Debug(
		"output",
		zap.String("driver", driverName),
		zap.String("opts", driverOpts),
	)
	constructor, found := driver.OutputDrivers[driverName]
	if !found {
		return nil, fmt.Errorf("unknown driver name: %q", driverName)
	}
	instance, err := constructor(driver.Opts{
		Logger: opts.rootLogger.Named("out"),
		Args:   driverOpts,
	})
	if err != nil {
		return nil, fmt.Errorf("constructor error: %w", err)
	}
	return instance, nil
}
