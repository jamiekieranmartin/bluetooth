package main

import (
	"flag"
	"fmt"

	"github.com/jamiekieranmartin/bluetooth"
)

const cliVersion = "0.0.1"

const helpMessage = `
bluetooth v%s

`

func main() {
	flag.Usage = func() {
		fmt.Printf(helpMessage, cliVersion)
		flag.PrintDefaults()
	}

	// cli arguments
	version := flag.Bool("version", false, "Print version string and exit")
	help := flag.Bool("help", false, "Print help message and exit")

	flag.Parse()

	// if asked for version, disregard everything else
	if *version {
		fmt.Printf("bluetooth v%s\n", cliVersion)
		return
	} else if *help {
		flag.Usage()
		return
	}

	bluetooth.StartScanning()
}
