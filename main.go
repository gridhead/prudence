package main

import (
	"flag"
	"fmt"
	"github.com/t0xic0der/spectrum"
	"github.com/t0xic0der/prudence/requestor"
)

var cetscycl, orbvcycl, conclave, cnstruct, drvodeal, fissures, invasion, newsevnt, nitewave bool

func main() {
	flag.BoolVar(&cetscycl, "cetscycl", false, "Know the state of Plains of Eidolon")
	flag.BoolVar(&orbvcycl, "orbvcycl", false, "Know the state of Orb Vallis")
	flag.BoolVar(&conclave, "conclave", false, "Know the details about Conclave")
	flag.BoolVar(&cnstruct, "cnstruct", false, "Know the progress of construction")
	flag.BoolVar(&drvodeal, "drvodeal", false, "Know the deals available from Darvo")
	flag.BoolVar(&fissures, "fissures", false, "Know the details of current void fissure missions")
	flag.BoolVar(&invasion, "invasion", false, "Know the details about current invasions")
	flag.BoolVar(&newsevnt, "newsevnt", false, "Know the details about current news")
	flag.BoolVar(&nitewave, "nitewave", false, "Know the details about nightwave missions")
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
	} else if fissures {
		requestor.VoidFissures("pc")
	} else if invasion {
		requestor.InvasionData("pc")
	} else if newsevnt {
		requestor.NewsEvents("pc")
	} else if nitewave {
		requestor.NightWaveChallenges("pc")
	} else {
		fmt.Print(spectrum.BG_GREEN(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/USAGE"))) + " v22102020 by t0xic0der" + "\n" +
			"Visit https://github.com/t0xic0der/prudence for updates" + "\n" + "\n" +
			spectrum.TX_BOLD(spectrum.FR_CYAN("The following usages are available")) + "\n" +
			spectrum.FR_CYAN("--cetscycl") + " ...... " + "Know the state of Plains of Eidolon" + "\n" +
			spectrum.FR_CYAN("--orbvcycl") + " ...... " + "Know the state of Orb Vallis" + "\n" +
			spectrum.FR_CYAN("--conclave") + " ...... " + "Know the details of Conclave" + "\n" +
			spectrum.FR_CYAN("--cnstruct") + " ...... " + "Know the progress of construction" + "\n" +
			spectrum.FR_CYAN("--drvodeal") + " ...... " + "Know the deals available from Darvo" + "\n" +
			spectrum.FR_CYAN("--fissures") + " ...... " + "Know the details of current void fissure missions" + "\n" +
			spectrum.FR_CYAN("--invasion") + " ...... " + "Know the details about current invasions" + "\n" +
			spectrum.FR_CYAN("--newsevnt") + " ...... " + "Know the details about current news" + "\n" +
			spectrum.FR_CYAN("--nitewave") + " ...... " + "Know the details about nightwave missions" + "\n" + "\n")
	}
}
