package moneykit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrency_Get(t *testing.T) {
	tcs := []struct {
		code     string
		expected string
	}{
		{EUR, "EUR"},
		{"EUR", "EUR"},
		{"Eur", "EUR"},
	}

	for _, tc := range tcs {
		c := newCurrency(tc.code).get()
		assert.Equal(t, tc.expected, c.Code, "Currency code should match expected value")
	}
}

func TestCurrency_Get1(t *testing.T) {
	code := "RANDOM"
	c := newCurrency(code).get()

	assert.Equal(t, code, c.Grapheme, "Currency grapheme should match the provided code")
}

func TestCurrency_Equals(t *testing.T) {
	tcs := []struct {
		code  string
		other string
	}{
		{EUR, "EUR"},
		{"EUR", "EUR"},
		{"Eur", "EUR"},
		{"usd", "USD"},
	}

	for _, tc := range tcs {
		c := newCurrency(tc.code).get()
		oc := newCurrency(tc.other).get()

		assert.True(t, c.equals(oc), "Currencies %v and %v should be equal", c, oc)
	}
}

func TestCurrency_AddCurrency(t *testing.T) {
	tcs := []struct {
		code     string
		template string
	}{
		{"GOLD", "1$"},
	}

	for _, tc := range tcs {
		AddCurrency(tc.code, "", tc.template, "", "", 0)
		c := newCurrency(tc.code).get()

		assert.Equal(t, tc.template, c.Template, "Currency template should match expected value")
	}
}

func TestCurrency_GetCurrency(t *testing.T) {
	code := "KLINGONDOLLAR"
	desired := Currency{Decimal: ".", Thousand: ",", Code: code, Fraction: 2, Grapheme: "$", Template: "$1"}
	AddCurrency(desired.Code, desired.Grapheme, desired.Template, desired.Decimal, desired.Thousand, desired.Fraction)
	currency := GetCurrency(code)

	assert.Equal(t, &desired, currency, "Retrieved currency should match the added currency")
}

func TestCurrency_GetNonExistingCurrency(t *testing.T) {
	currency := GetCurrency("I*am*Not*a*Currency")
	assert.Nil(t, currency, "Non-existing currency should return nil")
}

func TestCurrencies(t *testing.T) {
	const currencyFooCode = "FOO"
	const currencyFooNumericCode = "1234"
	curFoo := &Currency{
		Code:        currencyFooCode,
		NumericCode: currencyFooNumericCode,
		Fraction:    10,
		Grapheme:    "1",
		Template:    "2",
		Decimal:     "3",
		Thousand:    "4",
	}
	var cs = Currencies{
		currencyFooCode: curFoo,
	}
	const currencyBarCode = "BAR"
	const currencyBarNumericCode = "4321"
	curBar := &Currency{
		Code:        currencyBarCode,
		NumericCode: currencyBarNumericCode,
		Fraction:    1,
		Grapheme:    "2",
		Template:    "3",
		Decimal:     "4",
		Thousand:    "5",
	}
	cs = cs.Add(curBar)

	ac := cs.CurrencyByCode(currencyFooCode)
	assert.True(t, curFoo.equals(ac), "Currency retrieved by code should equal expected currency. Expected: %v, Got: %v", curFoo, ac)

	ac = cs.CurrencyByNumericCode(currencyFooNumericCode)
	assert.True(t, curFoo.equals(ac), "Currency retrieved by numeric code should equal expected currency. Expected: %v, Got: %v", curFoo, ac)

	ac = cs.CurrencyByCode(currencyBarCode)
	assert.True(t, curBar.equals(ac), "Currency retrieved by code should equal expected currency. Expected: %v, Got: %v", curBar, ac)

	ac = cs.CurrencyByNumericCode(currencyBarNumericCode)
	assert.True(t, curBar.equals(ac), "Currency retrieved by numeric code should equal expected currency. Expected: %v, Got: %v", curBar, ac)
}

func TestCurrency_GetCurrencyByNumericCode(t *testing.T) {
	code := "986"
	expected := GetCurrency(BRL)
	got := GetCurrencyByNumericCode(code)

	assert.True(t, expected.equals(got), "Currency retrieved by numeric code should equal expected currency. Expected: %v, Got: %v", expected, got)
}

func TestCurrency_GetCurrencyByNumericCodeNonExistingCurrency(t *testing.T) {
	currency := GetCurrencyByNumericCode("I*am*Not*a*Valid*Numeric*Code")
	assert.Nil(t, currency, "Non-existing numeric code should return nil")
}
