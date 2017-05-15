package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markjohnmagallanes/go-homework/openexchange"
)

type MockOpenExchangeAPI struct {
	oeRates openexchange.OERates
}

func (api MockOpenExchangeAPI) GetCurrencyRates() openexchange.OERates {
	return api.oeRates
}

func TestRatesHandlerGetLatestCurrency(t *testing.T) {

	mockAPICalls := MockOpenExchangeAPI{}
	mockAPICalls.oeRates = openexchange.OERates{
		Base: "USD",
		Rates: map[string]float64{
			"AUD": 1.347864,
			"SGD": 1.397609,
		},
	}

	req, err := http.NewRequest("GET", "/current_rates", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	handler := ratesHandler(mockAPICalls)
	handler.ServeHTTP(res, req)
	resData := res.Result()
	body, _ := ioutil.ReadAll(resData.Body)

	if resData.StatusCode != http.StatusOK {
		t.Error("expected http response code 200")
	}

	var resDataMap = make(map[string]interface{})
	json.Unmarshal([]byte(body), &resDataMap)

	if val, ok := resDataMap["error"]; !ok || val.(bool) {
		t.Error("JSON respond expected to have the key [error] with value[false]")
	}

	if val, ok := resDataMap["msg"]; !ok || val != "" {
		t.Error("JSON respond expected to have the key [msg] with Value []")
	}

	if val, ok := resDataMap["rates"]; !ok || len(val.(map[string]interface{})) == 0 {
		t.Error("JSON respond expected to have the key [rates] with value [unempty map]")
	}

	if _, ok := resDataMap["timestamp"]; !ok {
		t.Error("JSON respond expected to have the key [timestamp]")
	}
	if val, ok := resDataMap["base"]; !ok || val != "USD" {
		t.Error("JSON respond expected to have the key [base] with value[USD]")
	}

}

func TestRatesHandlerUnableToRetrievedRates(t *testing.T) {

	req, err := http.NewRequest("GET", "/current_rates", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	handler := ratesHandler(MockOpenExchangeAPI{})
	handler.ServeHTTP(res, req)
	resData := res.Result()
	body, _ := ioutil.ReadAll(resData.Body)

	if resData.StatusCode != http.StatusOK {
		t.Error("expected http response code 200")
	}

	var resDataMap = make(map[string]interface{})
	json.Unmarshal([]byte(body), &resDataMap)

	if val, ok := resDataMap["error"]; !ok || !val.(bool) {
		t.Error("JSON respond expected to have the key [error] with value[true]")
	}

	if val, ok := resDataMap["msg"]; !ok || val != "Unable to retrieved rates" {
		t.Error("JSON respond expected to have the key [msg] with Value [Unable to retrieved rates]")
	}

	if val, ok := resDataMap["rates"]; !ok || val != nil {
		t.Error("JSON respond expected to have the key [rates] with value [nil]")
	}

	if _, ok := resDataMap["timestamp"]; !ok {
		t.Error("JSON respond expected to have the key [timestamp]")
	}
	if val, ok := resDataMap["base"]; !ok || val != "" {
		t.Error("JSON respond expected to have the key [base] with value[empty]")
	}

}

func TestRatesHandlerWithParam(t *testing.T) {

	mockAPICalls := MockOpenExchangeAPI{}
	mockAPICalls.oeRates = openexchange.OERates{
		Base: "USD",
		Rates: map[string]float64{
			"AUD": 1.347864,
			"SGD": 1.397609,
		},
	}

	req, err := http.NewRequest("GET", "/current_rates?currency=SGD", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	handler := ratesHandler(mockAPICalls)
	handler.ServeHTTP(res, req)
	resData := res.Result()
	body, _ := ioutil.ReadAll(resData.Body)

	if resData.StatusCode != http.StatusOK {
		t.Error("expected http response code 200")
	}

	var resDataMap = make(map[string]interface{})

	json.Unmarshal([]byte(body), &resDataMap)

	if val, ok := resDataMap["error"]; !ok || val.(bool) {
		t.Error("JSON respond expected to have the key [error] with value[false]")
	}

	if val, ok := resDataMap["msg"]; !ok || val != "" {
		t.Error("JSON respond expected to have the key [msg] with Value []")
	}

	if _, ok := resDataMap["timestamp"]; !ok {
		t.Error("JSON respond expected to have the key [timestamp]")
	}

	if val, ok := resDataMap["base"]; !ok || val != "USD" {
		t.Error("JSON respond expected to have the key [base] with value[USD]")
	}

	if val, ok := resDataMap["code"]; !ok || val != "SGD" {
		t.Error("JSON respond expected to have the key [base] with value[SGD]")
	}

	if val, ok := resDataMap["rate"]; !ok || val.(float64) != 1.397609 {
		t.Error("JSON respond expected to have the key [base] with value[1.397609]")
	}
}

func TestRatesHandlerWithParamUnsupportedCurrency(t *testing.T) {
	req, err := http.NewRequest("GET", "/current_rates?currency=SGD", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	handler := ratesHandler(MockOpenExchangeAPI{})
	handler.ServeHTTP(res, req)
	resData := res.Result()
	body, _ := ioutil.ReadAll(resData.Body)

	if resData.StatusCode != http.StatusOK {
		t.Error("expected http response code 200")
	}

	var resDataMap = make(map[string]interface{})

	json.Unmarshal([]byte(body), &resDataMap)

	if val, ok := resDataMap["error"]; !ok || !val.(bool) {
		t.Error("JSON respond expected to have the key [error] with value[true]")
	}

	if val, ok := resDataMap["msg"]; !ok || val != "Unsupported Currency" {
		t.Error("JSON respond expected to have the key [msg] with Value [Unsupported Currency]")
	}

	if _, ok := resDataMap["timestamp"]; !ok {
		t.Error("JSON respond expected to have the key [timestamp]")
	}

	if val, ok := resDataMap["base"]; !ok || val != "" {
		t.Error("JSON respond expected to have the key [base] with value[]")
	}

	if val, ok := resDataMap["code"]; !ok || val != "SGD" {
		t.Error("JSON respond expected to have the key [base] with value[SGD]")
	}

	if val, ok := resDataMap["rate"]; !ok || val.(float64) != 0.0 {
		t.Error("JSON respond expected to have the key [base] with value[0.0]")
	}
}

func TestDefaultHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/unsupport_path", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	defaultHandler(res, req)

	if res.Code != http.StatusNotFound {
		t.Error("http response expected to be 404")
	}
}
