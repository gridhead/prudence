package main

import (
	"./requestor"
	"flag"
)

var cetscycl bool
var orbvcycl bool
var conclave bool

func main() {
	flag.BoolVar(&cetscycl, "cetscycl", false, "Know the state of Plains of Eidolon")
	flag.BoolVar(&orbvcycl, "orbvcycl", false, "Know the state of Orb Vallis")
	flag.BoolVar(&conclave, "conclave", false, "Know the details about Conclave")
	flag.Parse()
	if cetscycl {
		requestor.CetusCycle("pc")
	} else if orbvcycl {
		requestor.VallisCycle("pc")
	} else if conclave {
		requestor.ConclaveChallenges("pc")
	}
}
