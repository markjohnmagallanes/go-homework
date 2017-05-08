package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRatesHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/current_rates", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	ratesHandler(res, req)
	resData := res.Result()
	body, _ := ioutil.ReadAll(resData.Body)

	if resData.StatusCode != http.StatusOK {
		t.Error("expected http response code 200")
	}

	var resDataMap = make(map[string]interface{})
	json.Unmarshal([]byte(body), &resDataMap)

	if _, ok := resDataMap["error"]; !ok {
		t.Error("JSON respond expected to have the key [error]")
	}
	if _, ok := resDataMap["msg"]; !ok {
		t.Error("JSON respond expected to have the key [msg]")
	}

	if _, ok := resDataMap["rates"]; !ok {
		t.Error("JSON respond expected to have the key [rates]")
	}

	if _, ok := resDataMap["timestamp"]; !ok {
		t.Error("JSON respond expected to have the key [timestamp]")
	}
	if _, ok := resDataMap["base"]; !ok {
		t.Error("JSON respond expected to have the key [base]")
	}

}

func TestRatesHandlerWithParam(t *testing.T) {

	req, err := http.NewRequest("GET", "/current_rates?currency=SGD", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	ratesHandler(res, req)
	resData := res.Result()
	body, _ := ioutil.ReadAll(resData.Body)

	if resData.StatusCode != http.StatusOK {
		t.Error("expected http response code 200")
	}

	var resDataMap = make(map[string]interface{})

	json.Unmarshal([]byte(body), &resDataMap)

	if _, ok := resDataMap["error"]; !ok {
		t.Error("JSON respond expected to have the key [error]")
	}
	if _, ok := resDataMap["msg"]; !ok {
		t.Error("JSON respond expected to have the key [msg]")
	}

	if _, ok := resDataMap["rate"]; !ok {
		t.Error("JSON respond expected to have the key [rate]")
	}

	if _, ok := resDataMap["timestamp"]; !ok {
		t.Error("JSON respond expected to have the key [timestamp]")
	}
	if _, ok := resDataMap["base"]; !ok {
		t.Error("JSON respond expected to have the key [base]")
	}

}
