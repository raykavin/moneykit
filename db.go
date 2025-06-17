package moneykit

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

var (
	// DBMoneyValueSeparator is used to join Amount and Currency when storing Money
	// as strings in databases. Can be customized to use different separators.
	// Default: "|"
	//
	// Example:
	//	moneykit.DBMoneyValueSeparator = ":"
	//	// Now Money values are stored as "1000:USD" instead of "1000|USD"
	DBMoneyValueSeparator = DefaultDBMoneyValueSeparator
)

// Database Integration
//
// Money implements both sql.Scanner and driver.Valuer interfaces for seamless
// database integration. Values are stored as strings in the format "amount|currency".

// Value implements driver.Valuer interface to serialize Money for database storage.
// The Money instance is converted to a string in the format "amount|currency_code".
//
// Example database value: "2550|USD" represents $25.50
//
// Example:
//
//	money := moneykit.New(2550, "USD")
//	value, err := money.Value() // "2550|USD"
func (m *Money) Value() (driver.Value, error) {
	return fmt.Sprintf("%d%s%s", m.amount, DBMoneyValueSeparator, m.Currency().Code), nil
}

// Scan implements sql.Scanner interface to deserialize Money from database storage.
// Expects a string in the format "amount|currency_code".
//
// Parameters:
//   - src: Source value from database (should be string)
//
// Example:
//
//	var money moneykit.Money
//	err := money.Scan("2550|USD") // Creates $25.50
func (m *Money) Scan(src any) error {
	var amount Amount
	currency := &Currency{}

	// let's support string and int64
	switch s := src.(type) {
	case string:
		parts := strings.Split(s, DBMoneyValueSeparator)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return fmt.Errorf("%#v is not valid to scan into Money; update your query to return a currency.DBMoneyValueSeparator-separated pair of \"amount%scurrency_code\"", src.(string), DBMoneyValueSeparator)
		}

		if a, err := strconv.ParseInt(parts[0], 10, 64); err == nil {
			amount = a
		} else {
			return fmt.Errorf("scanning %#v into an Amount: %v", parts[0], err)
		}

		if err := currency.Scan(parts[1]); err != nil {
			return fmt.Errorf("scanning %#v into a Currency: %v", parts[1], err)
		}
	default:
		return fmt.Errorf("don't know how to scan %T into Money; update your query to return a currency.DBMoneyValueSeparator-separated pair of \"amount%scurrency_code\"", src, DBMoneyValueSeparator)
	}

	// allocate new Money with the scanned amount and currency
	*m = Money{
		amount:   amount,
		currency: currency,
	}

	return nil
}

// Value implements driver.Valuer to serialize a Currency code into a string for saving to a database
func (c Currency) Value() (driver.Value, error) {
	return c.Code, nil
}

// Scan implements sql.Scanner to deserialize a Currency from a string value read from a database
func (c *Currency) Scan(src any) error {
	var val *Currency
	// let's support string only
	switch s := src.(type) {
	case string:
		val = GetCurrency(s)
	default:
		return fmt.Errorf("%T is not a supported type for a Currency (store the Currency.Code value as a string only)", src)
	}

	if val == nil {
		return fmt.Errorf("GetCurrency(%#v) returned nil", src)
	}

	// copy the value
	*c = *val

	return nil
}
