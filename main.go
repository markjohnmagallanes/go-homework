package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/markjohnmagallanes/go-homework/openexchange"
)

// CurrParam attribute name for request parameter
const CurrParam string = "currency"

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/current_rates", ratesHandler)
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func ratesHandler(res http.ResponseWriter, req *http.Request) {
	_, hasParam := req.URL.Query()[CurrParam]
	var resData interface{}
	currency := req.URL.Query().Get(CurrParam)
	oeRates := openexchange.GetCurrencyRates()

	if hasParam {
		rates := oeRates.Rates
		rate, hasCurr := rates[strings.ToUpper(currency)]
		msg := ""
		if !hasCurr {
			msg = "Unsupported Currency"
		}

		resData = map[string]interface{}{
			"error":     !hasCurr,
			"msg":       msg,
			"timestamp": oeRates.Timestamp,
			"base":      oeRates.Base,
			"code":      currency,
			"rate":      rate,
		}

	} else {

		hasErr := len(oeRates.Rates) == 0
		msg := ""
		if hasErr {
			msg = "Unabled to retrieved rates"
		}
		resData = map[string]interface{}{
			"error":     hasErr,
			"msg":       msg,
			"timestamp": oeRates.Timestamp,
			"base":      oeRates.Base,
			"rates":     oeRates.Rates,
		}

	}

	js, err := json.Marshal(resData)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	http.NotFound(res, req)
}
