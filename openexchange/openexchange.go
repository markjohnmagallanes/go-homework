package openexchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// APICall for openexchangerates
const APICall string = "https://openexchangerates.org/api/latest.json?app_id=%v"

// APIKey for openexchangerates
const APIKey string = "5f8245af19724a29bddfc5cd23c6596b"

var (
	instance *OERates
	mutex    sync.Mutex
)

//OpenExchangeAPI implementation
type OpenExchangeAPI struct{}

//APICalls interfaces for getting currency rates
type APICalls interface {
	GetCurrencyRates() OERates
}

// GetCurrencyRates return rates inside the OERates object
func (api OpenExchangeAPI) GetCurrencyRates() OERates {
	return getInstance().clone()
}

func getInstance() *OERates {
	if instance == nil || instance.Timestamp.Add(time.Hour).After(time.Now()) {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil || instance.Timestamp.Add(time.Hour).After(time.Now()) {
			requestLatestRates()
		}
	}
	return instance
}

// GetConversion provided from code and to code returns convertion rate
func GetConversion(from string, to string) (float64, bool) {
	oeRates := getInstance()
	rates := oeRates.Rates
	rate := 0.0
	okay := true
	if !strings.EqualFold(oeRates.Base, from) {
		returnValue, hasCurr := rates[strings.ToUpper(from)]
		if hasCurr {
			rate = 1 / returnValue
		} else {
			okay = false
		}
	}

	returnValue, hasCurr := rates[strings.ToUpper(to)]
	if hasCurr {
		rate = rate * returnValue
	} else {
		okay = false
	}

	return rate, okay
}

func requestLatestRates() {
	res, err := http.Get(fmt.Sprintf(APICall, APIKey))

	if err != nil {
		return
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(jsonData), &instance)

	if err != nil {
		fmt.Println(err)
	}

}

// OERates mapping object for openexchangerates JSON
type OERates struct {
	Base      string             `json:"base"`
	Timestamp OETimestamp        `json:"timestamp"`
	Rates     map[string]float64 `json:"rates"`
}

// OETimestamp created to handle Unix time sent by openexchangerates
type OETimestamp struct {
	time.Time
}

// UnmarshalJSON handle timestamp sent by openexchangerates
func (t *OETimestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	t.Time = time.Unix(int64(ts), 0)
	return nil
}

func (o *OERates) clone() OERates {
	if o != nil {
		return OERates{o.Base, o.Timestamp, o.Rates}
	}
	return OERates{}
}
