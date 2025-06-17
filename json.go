package moneykit

// JSON Serialization
//
// Money implements json.Marshaler and json.Unmarshaler interfaces.
// The default format is {"amount": 1000, "currency": "USD"}.

// MarshalJSON implements json.Marshaler interface.
// Uses the global MarshalJSON function which can be customized.
//
// Default format: {"amount": 1000, "currency": "USD"}
//
// Example:
//
//	money := moneykit.New(1000, "USD")
//	data, err := json.Marshal(money)
//	// {"amount":1000,"currency":"USD"}
func (m *Money) UnmarshalJSON(b []byte) error {
	return UnmarshalJSON(m, b)
}

// Uses the global UnmarshalJSON function which can be customized.
//
// Example:
//
//	var money moneykit.Money
//	err := json.Unmarshal([]byte(`{"amount":1000,"currency":"USD"}`), &money)
func (m Money) MarshalJSON() ([]byte, error) {
	return MarshalJSON(m)
}
