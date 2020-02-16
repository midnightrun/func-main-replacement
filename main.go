package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	exitFail = 1
	argCount = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	var (
		verbose = flags.Bool("v", false, "verbose logging")
		format  = flags.String("f", "Hi %s\n", "greeting format")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if *verbose {
		fmt.Println("verbose mode activated")
	}

	if len(flags.Args()) < argCount {
		return errors.New("no names provided")
	}

	for _, name := range flags.Args() {
		fmt.Fprintf(stdout, *format, name)
	}

	return nil
}
