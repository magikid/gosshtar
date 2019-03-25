package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	initLogging()
	options := initFlags()

	if options.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debugging mode on")
	}

	var messages []string

	if options.Path != "" {
		messages = parseFile(options.Path)
	}

	StartServer(options.Server, messages)
}

func initLogging() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}

func initFlags() AppOptions {
	hostPtr := flag.String("host", "0.0.0.0", "Restrict the server to a specific IP")
	portPtr := flag.Int("port", 2222, "The port to run on")
	filePtr := flag.String("file", "", "A file of text to send to the bad actor")
	helpPtr := flag.Bool("help", false, "Print this help and exit")
	debugPtr := flag.Bool("debug", false, "Print lots of log messages")
	flag.Parse()

	if *helpPtr {
		printHelp()
	}

	log.Debugf("Parsed host: %v, parsed port: %v, parsed file: %v, debug mode: %v", *hostPtr, *portPtr, *filePtr, *debugPtr)

	server := SSHServer{*hostPtr, *portPtr}
	return AppOptions{server, *filePtr, *debugPtr}
}

type AppOptions struct {
	Server SSHServer
	Path   string
	Debug  bool
}

func printHelp() {
	fmt.Printf("%v [args]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(0)
}
