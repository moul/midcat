package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
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
	debug  bool
	logger *zap.Logger
}

func run(args []string) error {
	// setup CLI
	rootFs := flag.NewFlagSet("midcat", flag.ExitOnError)
	rootFs.BoolVar(&opts.debug, "debug", opts.debug, "debug mode")
	root := &ffcli.Command{
		Name:       "midcat",
		FlagSet:    rootFs,
		ShortUsage: "midcat [options] <address> <address>",
		Exec:       doRoot,
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

func doRoot(ctx context.Context, args []string) error {
	opts.logger.Debug("args", zap.Strings("args", args))
	fmt.Print(motd.Default())
	return nil
}
