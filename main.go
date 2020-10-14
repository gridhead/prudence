package main

import (
	"./colorformat"
	"./requestor"
	"flag"
	"fmt"
)

var cetscycl, orbvcycl, conclave, cnstruct, drvodeal, fissures, invasion, newsevnt bool

func main() {
	flag.BoolVar(&cetscycl, "cetscycl", false, "Know the state of Plains of Eidolon")
	flag.BoolVar(&orbvcycl, "orbvcycl", false, "Know the state of Orb Vallis")
	flag.BoolVar(&conclave, "conclave", false, "Know the details about Conclave")
	flag.BoolVar(&cnstruct, "cnstruct", false, "Know the progress of construction")
	flag.BoolVar(&drvodeal, "drvodeal", false, "Know the deals available from Darvo")
	flag.BoolVar(&fissures, "fissures", false, "Know the details of current void fissure missions")
	flag.BoolVar(&invasion, "invasion", false, "Know the details about current invasions")
	flag.BoolVar(&newsevnt, "newsevnt", false, "Know the details about current news")
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
	} else {
		fmt.Print(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/USAGE")) + " v14102020 by t0xic0der" + "\n" +
			"Visit https://github.com/t0xic0der/prudence for updates" + "\n" + "\n" +
			colorformat.ForeCynBold("The following usages are available") + "\n" +
			colorformat.ForeGrnRglr("--cetscycl") + " ...... " + "Know the state of Plains of Eidolon" + "\n" +
			colorformat.ForeGrnRglr("--orbvcycl") + " ...... " + "Know the state of Orb Vallis" + "\n" +
			colorformat.ForeGrnRglr("--conclave") + " ...... " + "Know the details of Conclave" + "\n" +
			colorformat.ForeGrnRglr("--cnstruct") + " ...... " + "Know the progress of construction" + "\n" +
			colorformat.ForeGrnRglr("--drvodeal") + " ...... " + "Know the deals available from Darvo" + "\n" +
			colorformat.ForeGrnRglr("--fissures") + " ...... " + "Know the details of current void fissure missions" + "\n" +
			colorformat.ForeGrnRglr("--invasion") + " ...... " + "Know the details about current invasions" + "\n" +
			colorformat.ForeGrnRglr("--newsevnt") + " ...... " + "Know the details about current news" + "\n" + "\n")
	}
}
