package requestor

import (
	"fmt"
	"github.com/t0xic0der/spectrum"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
)

func CetusCycle(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/cetusCycle")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(spectrum.BG_GREEN((spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/CETUS"))) + "\n"))
			var textobjc string = spectrum.TX_BOLD(spectrum.FR_CYAN(jsonobjc.Get("shortString").String())) + "\n" +
				spectrum.FR_GREEN("ID") + " .............. " + jsonobjc.Get("id").String() + "\n" +
				spectrum.FR_GREEN("Expiry") + " .......... " + jsonobjc.Get("expiry").String() + "\n" +
				spectrum.FR_GREEN("Activation") + " ...... " + jsonobjc.Get("activation").String() + "\n" +
				spectrum.FR_GREEN("Is it day now?") + " .. " + strings.Title(jsonobjc.Get("isDay").String()) + "\n" +
				spectrum.FR_GREEN("Cetus state") + " ..... " + strings.Title(jsonobjc.Get("state").String()) + "\n" +
				spectrum.FR_GREEN("Time left") + " ....... " + jsonobjc.Get("timeLeft").String() + "\n" + "\n"
			fmt.Print(textobjc)
			defer resp.Body.Close()
		} else {
			fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
			"The information could not be fetched")
	}
}

func VallisCycle(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/vallisCycle")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(spectrum.BG_GREEN(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/FORTUNA"))) + "\n")
			var textobjc string = spectrum.TX_BOLD(spectrum.FR_CYAN(jsonobjc.Get("shortString").String())) + "\n" +
				spectrum.FR_GREEN("ID") + " .............. " + jsonobjc.Get("id").String() + "\n" +
				spectrum.FR_GREEN("Expiry") + " .......... " + jsonobjc.Get("expiry").String() + "\n" +
				spectrum.FR_GREEN("Activation") + " ...... " + jsonobjc.Get("activation").String() + "\n" +
				spectrum.FR_GREEN("Is it warm now?") + " . " + strings.Title(jsonobjc.Get("isWarm").String()) + "\n" +
				spectrum.FR_GREEN("Cetus state") + " ..... " + strings.Title(jsonobjc.Get("state").String()) + "\n" +
				spectrum.FR_GREEN("Time left") + " ....... " + jsonobjc.Get("timeLeft").String() + "\n" + "\n"
			fmt.Print(textobjc)
			defer resp.Body.Close()
		} else {
			fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
			"The information could not be fetched")
	}
}

func ConclaveChallenges(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/conclaveChallenges")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(spectrum.BG_GREEN(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/CONCLAVE"))) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = spectrum.TX_BOLD(spectrum.FR_CYAN(singchal.Get("title").String())) + "\n" +
					singchal.Get("asString").String() + "\n" +
					spectrum.FR_GREEN("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					spectrum.FR_GREEN("Expiry") + " .......... " + singchal.Get("expiry").String() + "\n" +
					spectrum.FR_GREEN("Activation") + " ...... " + singchal.Get("activation").String() + "\n" +
					spectrum.FR_GREEN("Mode") + " ............ " + singchal.Get("mode").String() + "\n" +
					spectrum.FR_GREEN("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
			"The information could not be fetched")
	}
}

func ConstructionProgress(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/constructionProgress")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(spectrum.BG_GREEN(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/CONSTRUCTION"))) + "\n")
			var textobjc string = spectrum.TX_BOLD(spectrum.FR_CYAN("Progress so far")) + "\n" +
				spectrum.FR_GREEN("Fomorian") + " ........ " + jsonobjc.Get("fomorianProgress").String() + "\n" +
				spectrum.FR_GREEN("Razorback") + " ....... " + jsonobjc.Get("razorbackProgress").String() + "\n" + "\n"
			fmt.Print(textobjc)
			defer resp.Body.Close()
		} else {
			fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
			"The information could not be fetched")
	}
}

func DailyDeals(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/dailyDeals")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(spectrum.BG_GREEN(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/DEALS"))) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = spectrum.TX_BOLD(spectrum.FR_CYAN(singchal.Get("item").String())) + "\n" +
					"Price " + singchal.Get("salePrice").String() + "P (Was " + singchal.Get("originalPrice").String() + "P) @ " + singchal.Get("discount").String() + "% discount" + "\n" +
					spectrum.FR_GREEN("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					spectrum.FR_GREEN("Expiry") + " .......... " + singchal.Get("expiry").String() + "\n" +
					spectrum.FR_GREEN("Activation") + " ...... " + singchal.Get("activation").String() + "\n" +
					spectrum.FR_GREEN("Sold") + " ............ " + singchal.Get("sold").String() + " of " + singchal.Get("total").String() + " items \n" +
					spectrum.FR_GREEN("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
			"The information could not be fetched")
	}
}

func VoidFissures(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/fissures")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(spectrum.BG_GREEN(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/FISSURES"))) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = spectrum.TX_BOLD(spectrum.FR_CYAN(singchal.Get("node").String())) + "\n" +
					singchal.Get("missionType").String() + " @ " + singchal.Get("tier").String() + " " + singchal.Get("tierNum").String() + "\n" +
					spectrum.FR_GREEN("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					spectrum.FR_GREEN("Expiry") + " .......... " + singchal.Get("expiry").String() + "\n" +
					spectrum.FR_GREEN("Activation") + " ...... " + singchal.Get("activation").String() + "\n" +
					spectrum.FR_GREEN("Enemy") + " ........... " + singchal.Get("enemy").String() + "\n" +
					spectrum.FR_GREEN("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
			"The information could not be fetched")
	}
}

func InvasionData(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/invasions")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(spectrum.BG_GREEN(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/INVASIONS"))) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = spectrum.TX_BOLD(spectrum.FR_CYAN(singchal.Get("node").String())) + "\n" +
					singchal.Get("desc").String() + " - " + spectrum.FR_MAGENTA(singchal.Get("attackingFaction").String()) + " vs " + spectrum.FR_BLUE(singchal.Get("defendingFaction").String()) + "\n" +
					spectrum.FR_GREEN("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					spectrum.FR_GREEN("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" +
					spectrum.FR_GREEN("Activation") + " ...... " + singchal.Get("activation").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
			"The information could not be fetched")
	}
}

func NewsEvents(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/news")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(spectrum.BG_GREEN(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/NEWS"))) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = spectrum.TX_BOLD(spectrum.FR_CYAN(singchal.Get("message").String())) + "\n" +
					singchal.Get("link").String() + "\n" +
					spectrum.FR_GREEN("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					spectrum.FR_GREEN("Date") + " ............ " + singchal.Get("date").String() + "\n" +
					spectrum.FR_GREEN("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
			"The information could not be fetched")
	}
}

func NightWaveChallenges(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/nightwave")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(spectrum.BG_GREEN(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/NIGHTWAVE"))) + "\n")
			fmt.Print(spectrum.TX_BOLD(spectrum.FR_CYAN("Following are the nightwave challenges")) + "\n" +
				spectrum.FR_GREEN("Activation") + " ...... " + jsonobjc.Get("activation").String() + "\n" +
				spectrum.FR_GREEN("Expiry") + "........... " + jsonobjc.Get("expiry").String() + "\n" +
				spectrum.FR_GREEN("Season") + "........... " + jsonobjc.Get("season").String() + "\n")
			jsonobjc.Get("activeChallenges").ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = spectrum.TX_BOLD(spectrum.FR_CYAN(singchal.Get("activeChallenges").String())) + "\n" +
					spectrum.TX_BOLD(spectrum.FR_CYAN(singchal.Get("desc").String())) + "\n" +
					spectrum.FR_GREEN("Activation") + " ...... " + singchal.Get("activation").String() + "\n" +
					spectrum.FR_GREEN("Expiry") + " .......... " + singchal.Get("expiry").String() + "\n" +
					spectrum.FR_GREEN("Started at ") + " ..... " + singchal.Get("startString").String() + "\n" +
					spectrum.FR_GREEN("Reputation") + " ...... " + singchal.Get("reputation").String() + " standing" + "\n"
				fmt.Print(textobjc)
				return true
			})
			fmt.Println()
			defer resp.Body.Close()
		} else {
			fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(spectrum.BG_RED(spectrum.TX_BOLD(spectrum.FR_BLACK("PRUDENCE/WARNING"))) + "\n" +
			"The information could not be fetched")
	}
}
