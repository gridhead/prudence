package main

import (
	"flag"
	"prudence/requestor"
)

var cetscycl bool
var orbvcycl bool
var conclave bool
var cnstruct bool
var drvodeal bool

func main() {
	EnableVirtualTerminalProcessing()
	flag.BoolVar(&cetscycl, "cetscycl", false, "Know the state of Plains of Eidolon")
	flag.BoolVar(&orbvcycl, "orbvcycl", false, "Know the state of Orb Vallis")
	flag.BoolVar(&conclave, "conclave", false, "Know the details about Conclave")
	flag.BoolVar(&cnstruct, "cnstruct", false, "Know the progress of construction")
	flag.BoolVar(&drvodeal, "drvodeal", false, "Know the deals available from Darvo")
	flag.Parse()
	if cetscycl {
		requestor.CetusCycle("pc")
	} else if orbvcycl {
		requestor.VallisCycle("pc")
	} else if conclave {
		requestor.ConclaveChallenges("pc")
	} else if cnstruct {
		requestor.ConstructionProgress("pc")
	} else if drvodeal {
		requestor.DailyDeals("pc")
	}
}
