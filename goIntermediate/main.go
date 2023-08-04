package main

import (
	"fmt"
	"goIntermediate/go_buffers"
	"goIntermediate/go_channels"
	"goIntermediate/go_context"
	"goIntermediate/go_interfaces"
	"goIntermediate/go_mutex"
	"goIntermediate/go_pointers"
	"goIntermediate/go_routines"
	"goIntermediate/go_select"
	"os"
)

func main() {
	runLoop()
}

func runLoop() {
	for i := 1; i < len(os.Args); i++ {
		systemArgument := os.Args[i]

		fmt.Println(i, ".CURRENT METHOD:", systemArgument)

		checkSysArg(systemArgument)

		fmt.Println()
		fmt.Println()
	}
}

func checkSysArg(systemArgument string) {
	if systemArgument == "BFF" {
		go_buffers.INIT()
	} else if systemArgument == "CNTXT" {
		go_context.INIT()
	} else if systemArgument == "INTF" {
		go_interfaces.INIT()
	} else if systemArgument == "MTX" || systemArgument == "RWMTX" || systemArgument == "MTXI" {
		go_mutex.INIT(&systemArgument)
	} else if systemArgument == "PNTRS" {
		go_pointers.INIT()
	} else if systemArgument == "GRTS" {
		go_routines.INIT()
	} else if systemArgument == "SLCT" {
		go_select.INIT()
	} else if systemArgument == "CHNL" {
		go_channels.INIT()
	}
}
