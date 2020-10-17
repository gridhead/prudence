package requestor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../colorformat"
	"github.com/tidwall/gjson"
)

func CetusCycle(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/cetusCycle")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/CETUS")))
			var textobjc string = colorformat.ForeGrnBold(jsonobjc.Get("shortString").String()) + "\n" +
				colorformat.ForeGrnRglr("ID") + " ............. " + jsonobjc.Get("id").String() + "\n" +
				colorformat.ForeGrnRglr("Expiry") + " ......... " + jsonobjc.Get("expiry").String() + "\n" +
				colorformat.ForeGrnRglr("Activation") + " ..... " + jsonobjc.Get("activation").String() + "\n" +
				colorformat.ForeGrnRglr("Is it day now?") + " . " + jsonobjc.Get("isDay").String() + "\n" +
				colorformat.ForeGrnRglr("Cetus state") + " .... " + jsonobjc.Get("state").String() + "\n" +
				colorformat.ForeGrnRglr("Time left") + " ...... " + jsonobjc.Get("timeLeft").String() + "\n"
			fmt.Printf(textobjc)
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
			fmt.Println("The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
		fmt.Println("The information could not be fetched")
	}
}

func VallisCycle(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/vallisCycle")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/FORTUNA")))
			var textobjc string = colorformat.ForeGrnBold(jsonobjc.Get("shortString").String()) + "\n" +
				colorformat.ForeGrnRglr("ID") + " ............. " + jsonobjc.Get("id").String() + "\n" +
				colorformat.ForeGrnRglr("Expiry") + " ......... " + jsonobjc.Get("expiry").String() + "\n" +
				colorformat.ForeGrnRglr("Activation") + " ..... " + jsonobjc.Get("activation").String() + "\n" +
				colorformat.ForeGrnRglr("Is it day now?") + " . " + jsonobjc.Get("isWarm").String() + "\n" +
				colorformat.ForeGrnRglr("Cetus state") + " .... " + jsonobjc.Get("state").String() + "\n" +
				colorformat.ForeGrnRglr("Time left") + " ...... " + jsonobjc.Get("timeLeft").String() + "\n"
			fmt.Printf(textobjc)
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
			fmt.Println("The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
		fmt.Println("The information could not be fetched")
	}
}

func ConclaveChallenges(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/conclaveChallenges")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/CONCLAVE")))
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = colorformat.ForeGrnBold(singchal.Get("title").String()) + "\n" +
					singchal.Get("asString").String() + "\n" +
					colorformat.ForeGrnRglr("ID") + " ............. " + singchal.Get("id").String() + "\n" +
					colorformat.ForeGrnRglr("Expiry") + " ......... " + singchal.Get("expiry").String() + "\n" +
					colorformat.ForeGrnRglr("Activation") + " ..... " + singchal.Get("activation").String() + "\n" +
					colorformat.ForeGrnRglr("Mode") + " ........... " + singchal.Get("mode").String() + "\n" +
					colorformat.ForeGrnRglr("Estimated time") + " . " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Printf(textobjc)
				return true
			})
			fmt.Println("\n" + colorformat.ForeGrnBold("Total number of records fetched: "+
				strconv.Itoa(len(jsonobjc.Array()))) + "\n")
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
			fmt.Println("The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
		fmt.Println("The information could not be fetched")
	}
}

func ConstructionProgress(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/constructionProgress")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/CONSTRUCTION")))
			var textobjc string = colorformat.ForeGrnBold("Progress so far") + "\n" +
				colorformat.ForeGrnRglr("Fomorian") + " ....... " + jsonobjc.Get("fomorianProgress").String() + "\n" +
				colorformat.ForeGrnRglr("Razorback") + " ...... " + jsonobjc.Get("razorbackProgress").String() + "\n"
			fmt.Printf(textobjc)
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
			fmt.Println("The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
		fmt.Println("The information could not be fetched")
	}
}

func DailyDeals(platform string) {
	resp, eror := http.Get("https://api.warframestat.us/" + platform + "/dailyDeals")
	if eror == nil {
		body, eror := ioutil.ReadAll(resp.Body)
		if eror == nil {
			var jsonobjc gjson.Result = gjson.Parse(string(body))
			fmt.Println(colorformat.BackGrnBold(colorformat.ForeBlkBold("PRUDENCE/DEALS")))
			jsonobjc.ForEach(func(keyd, valu gjson.Result) bool {
				singchal := gjson.Parse(valu.String())
				var textobjc string = colorformat.ForeGrnBold(singchal.Get("item").String()) + "\n" +
					singchal.Get("salePrice").String() + "P (Was " + singchal.Get("originalPrice").String() + "P) @ " + singchal.Get("discount").String() + " percent discount" + "\n" +
					colorformat.ForeGrnRglr("ID") + " ............. " + singchal.Get("id").String() + "\n" +
					colorformat.ForeGrnRglr("Expiry") + " ......... " + singchal.Get("expiry").String() + "\n" +
					colorformat.ForeGrnRglr("Activation") + " ..... " + singchal.Get("activation").String() + "\n" +
					colorformat.ForeGrnRglr("Sold") + " ........... " + singchal.Get("sold").String() + " of " + singchal.Get("total").String() + " items \n" +
					colorformat.ForeGrnRglr("Estimated time") + " . " + singchal.Get("eta").String() + "\n" + "\n"
				fmt.Printf(textobjc)
				return true
			})
			fmt.Println("\n" + colorformat.ForeGrnBold("Total number of records fetched: "+
				strconv.Itoa(len(jsonobjc.Array()))) + "\n")
			defer resp.Body.Close()
		} else {
			fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
			fmt.Println("The information could not be fetched")
		}
	} else {
		fmt.Println(colorformat.BackRedBold(colorformat.ForeBlkBold("PRUDENCE/WARNING")))
		fmt.Println("The information could not be fetched")
	}
}
