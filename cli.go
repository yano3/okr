package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		command string
		format  string

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&command, "command", "", "Convert command")
	flags.StringVar(&command, "c", "", "Convert command (Short)")

	flags.StringVar(&format, "format", "jpg", "Image file format")
	flags.StringVar(&format, "f", "jpg", "Image file format (Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	parsedArgs := flags.Args()
	if len(parsedArgs) != 1 {
		fmt.Println("Invalid argument: Specify original image URL.")
		return ExitCodeError
	}
	originalURL := parsedArgs[0]

	okaraHost := os.Getenv("OKARA_HOST")
	okaraService := os.Getenv("OKARA_SERVICE")
	okaraType := os.Getenv("OKARA_TYPE")
	okaraToken := os.Getenv("OKARA_SECRET_TOKEN")

	isHTTPS := strings.HasPrefix(originalURL, "https://")

	originalURL = strings.TrimPrefix(originalURL, "https://")
	originalURL = strings.TrimPrefix(originalURL, "http://")

	if isHTTPS {
		originalURL = "s/" + originalURL
	}

	text := okaraToken + okaraService + okaraType + command + format + originalURL

	h := sha1.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)

	signedParam := fmt.Sprintf("%x", bs)

	okaraURL := fmt.Sprintf("https://%s/%s/%s/%s/%s.%s/%s", okaraHost, okaraService, okaraType, command, signedParam, format, originalURL)

	fmt.Println(okaraURL)

	_ = command

	_ = format

	_ = originalURL

	return ExitCodeOK
}
