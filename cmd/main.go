package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jamiekieranmartin/bluetooth"
)

const cliVersion = "0.0.2"

const helpMessage = `
bluetooth is a minimal bluetooth scanner.
	bluetooth v%s

Usage.
	bluetooth [flags]
`

func main() {
	flag.Usage = func() {
		fmt.Printf(helpMessage, cliVersion)
		flag.PrintDefaults()
	}

	version := flag.Bool("version", false, "Print version string and exit")
	help := flag.Bool("help", false, "Print help message and exit")

	duration := flag.Int("duration", int(10*time.Second), "Duration to scan for (ms)")

	flag.Parse()

	// if asked for version, disregard everything else
	if *version {
		fmt.Printf("bluetooth v%s\n", cliVersion)
		return
	} else if *help {
		flag.Usage()
		return
	}

	scanner := bluetooth.NewScanner(&bluetooth.Config{
		Duration: time.Duration(*duration),
		Debug:    true,
	})

	events := scanner.Start()

	for event := range events {
		fmt.Printf("%+v\n", event)
	}
}
