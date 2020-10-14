package requestor

import (
	"../colorformat"
	"fmt"
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
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/CETUS")) + "\n")
			var textobjc string = colorformat.ForeCynBold(jsonobjc.Get("shortString").String()) + "\n" +
				colorformat.ForeGrnRglr("ID") + " .............. " + jsonobjc.Get("id").String() + "\n" +
				colorformat.ForeGrnRglr("Expiry") + " .......... " + jsonobjc.Get("expiry").String() + "\n" +
				colorformat.ForeGrnRglr("Activation") + " ...... " + jsonobjc.Get("activation").String() + "\n" +
				colorformat.ForeGrnRglr("Is it day now?") + " .. " + strings.Title(jsonobjc.Get("isDay").String()) + "\n" +
				colorformat.ForeGrnRglr("Cetus state") + " ..... " + strings.Title(jsonobjc.Get("state").String()) + "\n" +
				colorformat.ForeGrnRglr("Time left") + " ....... " + jsonobjc.Get("timeLeft").String() + "\n" + "\n"
			fmt.Print(textobjc)
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
			"The information could not be fetched")
	}
}

func VallisCycle(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/vallisCycle")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/FORTUNA")) + "\n")
			var textobjc string = colorformat.ForeCynBold(jsonobjc.Get("shortString").String()) + "\n" +
				colorformat.ForeGrnRglr("ID") + " .............. " + jsonobjc.Get("id").String() + "\n" +
				colorformat.ForeGrnRglr("Expiry") + " .......... " + jsonobjc.Get("expiry").String() + "\n" +
				colorformat.ForeGrnRglr("Activation") + " ...... " + jsonobjc.Get("activation").String() + "\n" +
				colorformat.ForeGrnRglr("Is it warm now?") + " . " + strings.Title(jsonobjc.Get("isWarm").String()) + "\n" +
				colorformat.ForeGrnRglr("Cetus state") + " ..... " + strings.Title(jsonobjc.Get("state").String()) + "\n" +
				colorformat.ForeGrnRglr("Time left") + " ....... " + jsonobjc.Get("timeLeft").String() + "\n" + "\n"
			fmt.Print(textobjc)
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
			"The information could not be fetched")
	}
}

func ConclaveChallenges(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/conclaveChallenges")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/CONCLAVE")) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = colorformat.ForeCynBold(singchal.Get("title").String()) + "\n" +
					singchal.Get("asString").String() + "\n" +
					colorformat.ForeGrnRglr("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					colorformat.ForeGrnRglr("Expiry") + " .......... " + singchal.Get("expiry").String() + "\n" +
					colorformat.ForeGrnRglr("Activation") + " ...... " + singchal.Get("activation").String() + "\n" +
					colorformat.ForeGrnRglr("Mode") + " ............ " + singchal.Get("mode").String() + "\n" +
					colorformat.ForeGrnRglr("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
			"The information could not be fetched")
	}
}

func ConstructionProgress(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/constructionProgress")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/CONSTRUCTION")) + "\n")
			var textobjc string = colorformat.ForeCynBold("Progress so far") + "\n" +
				colorformat.ForeGrnRglr("Fomorian") + " ........ " + jsonobjc.Get("fomorianProgress").String() + "\n" +
				colorformat.ForeGrnRglr("Razorback") + " ....... " + jsonobjc.Get("razorbackProgress").String() + "\n" + "\n"
			fmt.Print(textobjc)
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
			"The information could not be fetched")
	}
}

func DailyDeals(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/dailyDeals")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/DEALS")) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = colorformat.ForeCynBold(singchal.Get("item").String()) + "\n" +
					"Price " + singchal.Get("salePrice").String() + "P (Was " + singchal.Get("originalPrice").String() + "P) @ " + singchal.Get("discount").String() + "% discount" + "\n" +
					colorformat.ForeGrnRglr("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					colorformat.ForeGrnRglr("Expiry") + " .......... " + singchal.Get("expiry").String() + "\n" +
					colorformat.ForeGrnRglr("Activation") + " ...... " + singchal.Get("activation").String() + "\n" +
					colorformat.ForeGrnRglr("Sold") + " ............ " + singchal.Get("sold").String() + " of " + singchal.Get("total").String() + " items \n" +
					colorformat.ForeGrnRglr("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
			"The information could not be fetched")
	}
}

func VoidFissures(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/fissures")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/FISSURES")) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = colorformat.ForeCynBold(singchal.Get("node").String()) + "\n" +
					singchal.Get("missionType").String() + " @ " + singchal.Get("tier").String() + " " + singchal.Get("tierNum").String() + "\n" +
					colorformat.ForeGrnRglr("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					colorformat.ForeGrnRglr("Expiry") + " .......... " + singchal.Get("expiry").String() + "\n" +
					colorformat.ForeGrnRglr("Activation") + " ...... " + singchal.Get("activation").String() + "\n" +
					colorformat.ForeGrnRglr("Enemy") + " ........... " + singchal.Get("enemy").String() + "\n" +
					colorformat.ForeGrnRglr("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
			"The information could not be fetched")
	}
}

func InvasionData(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/invasions")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/INVASIONS")) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = colorformat.ForeCynBold(singchal.Get("node").String()) + "\n" +
					singchal.Get("desc").String() + " - " + colorformat.ForeMgtRglr(singchal.Get("attackingFaction").String()) + " vs " + colorformat.ForeBluRglr(singchal.Get("defendingFaction").String()) + "\n" +
					colorformat.ForeGrnRglr("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					colorformat.ForeGrnRglr("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" +
					colorformat.ForeGrnRglr("Activation") + " ...... " + singchal.Get("activation").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
			"The information could not be fetched")
	}
}

func NewsEvents(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/news")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/NEWS")) + "\n")
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = colorformat.ForeCynBold(singchal.Get("message").String()) + "\n" +
					singchal.Get("link").String() + "\n" +
					colorformat.ForeGrnRglr("ID") + " .............. " + singchal.Get("id").String() + "\n" +
					colorformat.ForeGrnRglr("Date") + " ............ " + singchal.Get("date").String() + "\n" +
					colorformat.ForeGrnRglr("Estimated time") + " .. " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Print(textobjc)
				return true
			})
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
			"The information could not be fetched")
	}
}

func NightWaveChallenges(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/nightwave")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/NIGHTWAVE")) + "\n")
			fmt.Print(colorformat.ForeCynBold("Following are the nightwave challenges") + "\n" +
				colorformat.ForeGrnRglr("Activation") + " ...... " + jsonobjc.Get("activation").String() + "\n" +
				colorformat.ForeGrnRglr("Expiry") + "........... " + jsonobjc.Get("expiry").String() + "\n" +
				colorformat.ForeGrnRglr("Season") + "........... " + jsonobjc.Get("season").String() + "\n")
			jsonobjc.Get("activeChallenges").ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = colorformat.ForeCynBold(singchal.Get("activeChallenges").String()) + "\n" +
					colorformat.ForeCynBold(singchal.Get("desc").String()) + "\n" +
					colorformat.ForeGrnRglr("Activation") + " ...... " + singchal.Get("activation").String() + "\n" +
					colorformat.ForeGrnRglr("Expiry") + " .......... " + singchal.Get("expiry").String() + "\n" +
					colorformat.ForeGrnRglr("Started at ") + " ..... " + singchal.Get("startString").String() + "\n" +
					colorformat.ForeGrnRglr("Reputation") + " ...... " + singchal.Get("reputation").String() + " standing" + "\n"
				fmt.Print(textobjc)
				return true
			})
			fmt.Println()
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
				"The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")) + "\n" +
			"The information could not be fetched")
	}
}
