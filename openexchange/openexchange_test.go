package openexchange

import (
	"encoding/json"
	"testing"
	"time"
)

func TestUnmarshalJSON(t *testing.T) {
	expectedTimestamp := time.Unix(1494255600, 0)
	expectedBase := "USD"
	var oeRates OERates

	json.Unmarshal([]byte(`{
    "base":"USD",
    "timestamp": 1494255600,
    "rates":{"USD":1.5,"AUD":2.5}
    }`),
		&oeRates)

	if !oeRates.Timestamp.Equal(expectedTimestamp) {
		t.Errorf("json convertion incorrect timestamp expected: %v but actual: %v", expectedTimestamp, oeRates.Timestamp)
	}
	if oeRates.Base != expectedBase {
		t.Errorf("json convertion incorrect base expected: %v but actual: %v", expectedBase, oeRates.Base)
	}

	if len(oeRates.Rates) != 2 {
		t.Errorf("json convertion incorrect expected: %v but actual: %v", 2, len(oeRates.Rates))
	}
}

func TestUnmarshalFail(t *testing.T) {
	var oeRates OERates
	expected := OETimestamp{time.Time{}}
	json.Unmarshal([]byte(`{
    "timestamp": aaaaaaa,
    }`),
		&oeRates)
	if oeRates.Timestamp != expected {
		t.Errorf("json convertion incorrect timestamp expected: %v but actual: %v", expected, oeRates.Timestamp)
	}

}
