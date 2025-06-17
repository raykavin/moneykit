// Package moneykit provides precise monetary calculations and currency handling.
//
// MoneyKit is designed to handle money safely and accurately by using integer arithmetic
// instead of floating-point numbers, which can introduce precision errors. All monetary
// values are stored as integers in the currency's smallest unit (e.g., cents for USD).
//
// Basic Usage:
//
//	// Create money instances
//	price := moneykit.New(2550, "USD")  // $25.50
//	tax := moneykit.New(255, "USD")     // $2.55
//
//	// Perform safe arithmetic
//	total, err := price.Add(tax)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(total.Display()) // $28.05
//
//	// Split amounts fairly
//	shares, err := total.Split(3)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// shares[0]: $9.35, shares[1]: $9.35, shares[2]: $9.35
//
// Currency Support:
//
// MoneyKit includes built-in support for all active ISO 4217 currencies with proper
// formatting, decimal places, and symbols. You can also add custom currencies:
//
//	moneykit.AddCurrency("BTC", "₿", "₿1", ".", ",", 8)
//	bitcoin := moneykit.New(100000000, "BTC") // 1.00000000 BTC
//
// Database Integration:
//
// Money values can be stored and retrieved from databases using the built-in
// sql.Scanner and driver.Valuer interfaces:
//
//	var money moneykit.Money
//	err := db.QueryRow("SELECT amount FROM orders WHERE id = ?", 1).Scan(&money)
//
// JSON Serialization:
//
// Money values can be marshaled to and from JSON:
//
//	money := moneykit.New(1000, "USD")
//	data, err := json.Marshal(money)
//	// {"amount":1000,"currency":"USD"}
package moneykit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

// Injection points for backward compatibility.
// If you need to keep your JSON marshal/unmarshal way, overwrite them like below.
//
//	currency.UnmarshalJSON = func (m *Money, b []byte) error { ... }
//	currency.MarshalJSON = func (m Money) ([]byte, error) { ... }
var (

	// UnmarshalJSON is an injection point for customizing JSON unmarshaling behavior.
	// Override this function to implement custom JSON formats.
	//
	// Example:
	//	moneykit.UnmarshalJSON = func(m *moneykit.Money, b []byte) error {
	//		// Custom unmarshaling logic
	//		return nil
	//	}
	UnmarshalJSON = defaultUnmarshalJSON

	// MarshalJSON is an injection point for customizing JSON marshaling behavior.
	// Override this function to implement custom JSON formats.
	//
	// Example:
	//	moneykit.MarshalJSON = func(m moneykit.Money) ([]byte, error) {
	//		return json.Marshal(map[string]interface{}{
	//			"value":    m.AsMajorUnits(),
	//			"currency": m.Currency().Code,
	//		})
	//	}
	MarshalJSON = defaultMarshalJSON

	// ErrCurrencyMismatch is returned when attempting operations between
	// Money instances with different currencies.
	ErrCurrencyMismatch = errors.New("currencies don't match")

	// ErrInvalidJSONUnmarshal is returned when JSON unmarshaling fails
	// due to invalid or malformed data.
	ErrInvalidJSONUnmarshal = errors.New("invalid json unmarshal")
)

func defaultUnmarshalJSON(m *Money, b []byte) error {
	data := make(map[string]interface{})
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	var amount float64
	if amountRaw, ok := data["amount"]; ok {
		amount, ok = amountRaw.(float64)
		if !ok {
			return ErrInvalidJSONUnmarshal
		}
	}

	var currency string
	if currencyRaw, ok := data["currency"]; ok {
		currency, ok = currencyRaw.(string)
		if !ok {
			return ErrInvalidJSONUnmarshal
		}
	}

	var ref *Money
	if amount == 0 && currency == "" {
		ref = &Money{}
	} else {
		ref = New(int64(amount), currency)
	}

	*m = *ref
	return nil
}

func defaultMarshalJSON(m Money) ([]byte, error) {
	if m == (Money{}) {
		m = *New(0, "")
	}

	buff := bytes.NewBufferString(fmt.Sprintf(`{"amount": %d, "currency": "%s"}`, m.Amount(), m.Currency().Code))
	return buff.Bytes(), nil
}

// Amount represents a monetary amount as an integer in the currency's smallest unit.
// For example, for USD this would be cents, for EUR this would be euro cents.
// This type alias is used throughout the package to maintain clarity about
// what integer values represent.
type Amount = int64

// Money represents a monetary value with its associated currency.
// It stores the amount as an integer to avoid floating-point precision issues.
// All arithmetic operations maintain currency safety by ensuring operations
// are only performed between Money instances of the same currency.
//
// Example:
//
//	money := moneykit.New(2550, "USD") // $25.50
//	fmt.Println(money.Display())       // $25.50
//	fmt.Println(money.Amount())        // 2550
//	fmt.Println(money.Currency().Code) // USD
type Money struct {
	amount   Amount    `db:"amount"`
	currency *Currency `db:"currency"`
}

// New creates a new Money instance with the specified amount and currency code.
// The amount should be provided in the currency's smallest unit (e.g., cents for USD).
//
// Parameters:
//   - amount: The monetary amount in the currency's smallest unit
//   - code: The ISO 4217 currency code (e.g., "USD", "EUR", "JPY")
//
// Example:
//
//	usd := moneykit.New(2550, "USD")  // $25.50
//	eur := moneykit.New(1000, "EUR")  // €10.00
//	jpy := moneykit.New(1000, "JPY")  // ¥1000 (no decimals)
func New(amount int64, code string) *Money {
	return &Money{
		amount:   amount,
		currency: newCurrency(code).get(),
	}
}

// NewFromFloat creates a new Money instance from a floating-point number.
// The float is automatically converted to the currency's smallest unit.
// This method should be used sparingly as it can introduce precision issues
// for very large numbers or numbers with many decimal places.
//
// Parameters:
//   - amount: The monetary amount as a floating-point number
//   - code: The ISO 4217 currency code
//
// Example:
//
//	money := moneykit.NewFromFloat(25.50, "USD") // $25.50
//	fmt.Println(money.Amount()) // 2550
func NewFromFloat(amount float64, code string) *Money {
	currencyDecimals := math.Pow10(newCurrency(code).get().Fraction)
	return New(int64(amount*currencyDecimals), code)
}

// Currency returns the Currency information associated with this Money instance.
// This includes details like the currency code, symbol, decimal places, and formatting rules.
//
// Example:
//
//	money := moneykit.New(1000, "USD")
//	currency := money.Currency()
//	fmt.Println(currency.Code)     // USD
//	fmt.Println(currency.Grapheme) // $
//	fmt.Println(currency.Fraction) // 2
func (m *Money) Currency() *Currency {
	return m.currency
}

// Amount returns the monetary amount as an integer in the currency's smallest unit.
// This is a copy of the internal value, so modifying it won't affect the Money instance.
//
// Example:
//
//	money := moneykit.New(2550, "USD")
//	fmt.Println(money.Amount()) // 2550 (cents)
func (m *Money) Amount() int64 {
	return m.amount
}

// SameCurrency checks if this Money instance has the same currency as another Money instance.
// This is used internally to ensure currency safety in arithmetic operations.
//
// Example:
//
//	usd1 := moneykit.New(100, "USD")
//	usd2 := moneykit.New(200, "USD")
//	eur := moneykit.New(100, "EUR")
//
//	fmt.Println(usd1.SameCurrency(usd2)) // true
//	fmt.Println(usd1.SameCurrency(eur))  // false
func (m *Money) SameCurrency(om *Money) bool {
	return m.currency.equals(om.currency)
}

func (m *Money) assertSameCurrency(om *Money) error {
	if !m.SameCurrency(om) {
		return ErrCurrencyMismatch
	}

	return nil
}

func (m *Money) compare(om *Money) int {
	switch {
	case m.amount > om.amount:
		return 1
	case m.amount < om.amount:
		return -1
	}

	return 0
}

// Equals checks if this Money instance is equal to another Money instance.
// Both the amount and currency must match.
//
// Returns:
//   - bool: true if amounts and currencies are equal
//   - error: ErrCurrencyMismatch if currencies don't match
//
// Example:
//
//	money1 := moneykit.New(1000, "USD")
//	money2 := moneykit.New(1000, "USD")
//	money3 := moneykit.New(1000, "EUR")
//
//	equal, err := money1.Equals(money2)
//	fmt.Println(equal, err) // true, nil
//
//	equal, err = money1.Equals(money3)
//	fmt.Println(equal, err) // false, currencies don't match
func (m *Money) Equals(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) == 0, nil
}

// GreaterThan checks if this Money instance is greater than another Money instance.
//
// Returns:
//   - bool: true if this amount is greater
//   - error: ErrCurrencyMismatch if currencies don't match
func (m *Money) GreaterThan(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) == 1, nil
}

// GreaterThanOrEqual checks if this Money instance is greater than or equal to another Money instance.
func (m *Money) GreaterThanOrEqual(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) >= 0, nil
}

// LessThan checks if this Money instance is less than another Money instance.
func (m *Money) LessThan(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) == -1, nil
}

// LessThanOrEqual checks if this Money instance is less than or equal to another Money instance.
func (m *Money) LessThanOrEqual(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) <= 0, nil
}

// IsZero returns true if the monetary amount is zero.
func (m *Money) IsZero() bool {
	return m.amount == 0
}

// IsPositive returns true if the monetary amount is greater than zero.
func (m *Money) IsPositive() bool {
	return m.amount > 0
}

// IsNegative returns true if the monetary amount is less than zero.
func (m *Money) IsNegative() bool {
	return m.amount < 0
}

// Absolute returns a new Money instance with the absolute value of this Money.
//
// Example:
//
//	debt := moneykit.New(-500, "USD")
//	amount := debt.Absolute()
//	fmt.Println(amount.Display()) // $5.00
func (m *Money) Absolute() *Money {
	return &Money{amount: mutate.calc.absolute(m.amount), currency: m.currency}
}

// Negative returns a new Money instance with the negative value of this Money.
// If the money is already negative, it remains negative (idempotent).
//
// Example:
//
//	positive := moneykit.New(500, "USD")
//	negative := positive.Negative()
//	fmt.Println(negative.Display()) // -$5.00
func (m *Money) Negative() *Money {
	return &Money{amount: mutate.calc.negative(m.amount), currency: m.currency}
}

// Add returns a new Money instance representing the sum of this Money and one or more other Money instances.
// All Money instances must have the same currency, otherwise an ErrCurrencyMismatch error is returned.
//
// Parameters:
//   - ms: One or more Money instances to add
//
// Returns:
//   - *Money: A new Money instance with the sum
//   - error: ErrCurrencyMismatch if currencies don't match
//
// Example:
//
//	base := moneykit.New(1000, "USD")
//	tip := moneykit.New(150, "USD")
//	tax := moneykit.New(80, "USD")
//
//	total, err := base.Add(tip, tax)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(total.Display()) // $12.30
func (m *Money) Add(ms ...*Money) (*Money, error) {
	if len(ms) == 0 {
		return m, nil
	}

	k := New(0, m.currency.Code)

	for _, m2 := range ms {
		if err := m.assertSameCurrency(m2); err != nil {
			return nil, err
		}

		k.amount = mutate.calc.add(k.amount, m2.amount)
	}

	return &Money{amount: mutate.calc.add(m.amount, k.amount), currency: m.currency}, nil
}

// Subtract returns a new Money instance representing the difference between this Money
// and one or more other Money instances. All Money instances must have the same currency.
//
// Parameters:
//   - ms: One or more Money instances to subtract
//
// Returns:
//   - *Money: A new Money instance with the difference
//   - error: ErrCurrencyMismatch if currencies don't match
//
// Example:
//
//	total := moneykit.New(2550, "USD")
//	discount := moneykit.New(255, "USD")
//	tax := moneykit.New(51, "USD")
//
//	final, err := total.Subtract(discount, tax)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(final.Display()) // $22.44
func (m *Money) Subtract(ms ...*Money) (*Money, error) {
	if len(ms) == 0 {
		return m, nil
	}

	k := New(0, m.currency.Code)

	for _, m2 := range ms {
		if err := m.assertSameCurrency(m2); err != nil {
			return nil, err
		}

		k.amount = mutate.calc.add(k.amount, m2.amount)
	}

	return &Money{amount: mutate.calc.subtract(m.amount, k.amount), currency: m.currency}, nil
}

// Multiply returns a new Money instance representing this Money multiplied by one or more integers.
// This method panics if no multipliers are provided.
//
// Parameters:
//   - muls: One or more integers to multiply by
//
// Example:
//
//	price := moneykit.New(1000, "USD")
//	doubled := price.Multiply(2)
//	fmt.Println(doubled.Display()) // $20.00
//
//	// Chain multiplications: 10 * 2 * 3 = 60
//	result := price.Multiply(2, 3)
//	fmt.Println(result.Display()) // $60.00
func (m *Money) Multiply(muls ...int64) *Money {
	if len(muls) == 0 {
		panic("At least one multiplier is required to multiply")
	}

	k := New(1, m.currency.Code)

	for _, m2 := range muls {
		k.amount = mutate.calc.multiply(k.amount, m2)
	}

	return &Money{amount: mutate.calc.multiply(m.amount, k.amount), currency: m.currency}
}

// Round returns a new Money instance with the amount rounded to the currency's
// standard precision (number of decimal places).
//
// Example:
//
//	money := moneykit.New(1567, "USD") // $15.67
//	rounded := money.Round()           // Rounds to nearest dollar
func (m *Money) Round() *Money {
	return &Money{amount: mutate.calc.round(m.amount, m.currency.Fraction), currency: m.currency}
}

// Split divides this Money into n equal parts, distributing any remainder
// using a round-robin approach. The first parties in the slice will receive
// any extra pennies.
//
// Parameters:
//   - n: Number of parts to split into (must be > 0)
//
// Returns:
//   - []*Money: Slice of Money instances representing the split amounts
//   - error: Error if n <= 0
//
// Example:
//
//	bill := moneykit.New(1000, "USD") // $10.00
//	shares, err := bill.Split(3)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// shares[0]: $3.34 (gets extra penny)
//	// shares[1]: $3.33
//	// shares[2]: $3.33
func (m *Money) Split(n int) ([]*Money, error) {
	if n <= 0 {
		return nil, errors.New("split must be higher than zero")
	}

	a := mutate.calc.divide(m.amount, int64(n))
	ms := make([]*Money, n)

	for i := 0; i < n; i++ {
		ms[i] = &Money{amount: a, currency: m.currency}
	}

	r := mutate.calc.modulus(m.amount, int64(n))
	l := mutate.calc.absolute(r)
	// Add leftovers to the first parties.

	v := int64(1)
	if m.amount < 0 {
		v = -1
	}
	for p := 0; l != 0; p++ {
		ms[p].amount = mutate.calc.add(ms[p].amount, v)
		l--
	}

	return ms, nil
}

// Allocate divides this Money according to the provided ratios, distributing
// any remainder using a round-robin approach. This is useful for proportional
// distribution based on percentages or weights.
//
// Parameters:
//   - rs: Variable number of integers representing allocation ratios
//
// Returns:
//   - []*Money: Slice of Money instances allocated according to ratios
//   - error: Error if no ratios provided, negative ratios, or ratio sum overflow
//
// Example:
//
//	revenue := moneykit.New(10000, "USD") // $100.00
//	portions, err := revenue.Allocate(50, 30, 20) // 50%, 30%, 20%
//	if err != nil {
//		log.Fatal(err)
//	}
//	// portions[0]: $50.00
//	// portions[1]: $30.00
//	// portions[2]: $20.00
//
//	// Handle remainders
//	amount := moneykit.New(100, "USD") // $1.00
//	parts, err := amount.Allocate(33, 33, 33)
//	// parts[0]: $0.34 (gets extra penny)
//	// parts[1]: $0.33
//	// parts[2]: $0.33
func (m *Money) Allocate(rs ...int) ([]*Money, error) {
	if len(rs) == 0 {
		return nil, errors.New("no ratios specified")
	}

	// Calculate sum of ratios.
	var sum int64
	for _, r := range rs {
		if r < 0 {
			return nil, errors.New("negative ratios not allowed")
		}
		if int64(r) > (math.MaxInt64 - sum) {
			return nil, errors.New("sum of given ratios exceeds max int")
		}
		sum += int64(r)
	}

	var total int64
	ms := make([]*Money, 0, len(rs))
	for _, r := range rs {
		party := &Money{
			amount:   mutate.calc.allocate(m.amount, int64(r), sum),
			currency: m.currency,
		}

		ms = append(ms, party)
		total += party.amount
	}

	// if the sum of all ratios is zero, then we just returns zeros and don't do anything
	// with the leftover
	if sum == 0 {
		return ms, nil
	}

	// Calculate leftover value and divide to first parties.
	lo := m.amount - total
	sub := int64(1)
	if lo < 0 {
		sub = -sub
	}

	for p := 0; lo != 0; p++ {
		ms[p].amount = mutate.calc.add(ms[p].amount, sub)
		lo -= sub
	}

	return ms, nil
}

// Display returns a formatted string representation of the Money using the currency's
// formatting rules. This includes the proper currency symbol, decimal places,
// and thousands separators according to the currency's conventions.
//
// Example:
//
//	usd := moneykit.New(123456, "USD")
//	fmt.Println(usd.Display()) // $1,234.56
//
//	eur := moneykit.New(123456, "EUR")
//	fmt.Println(eur.Display()) // €1,234.56
//
//	jpy := moneykit.New(12345, "JPY")
//	fmt.Println(jpy.Display()) // ¥12,345
func (m *Money) Display() string {
	c := m.currency.get()
	return c.Formatter().Format(m.amount)
}

// AsMajorUnits returns the monetary value as a floating-point number in the currency's
// major units (e.g., dollars instead of cents). This is useful for display purposes
// or when interfacing with systems that expect decimal values.
//
// Example:
//
//	money := moneykit.New(2550, "USD")
//	fmt.Printf("%.2f", money.AsMajorUnits()) // 25.50
func (m *Money) AsMajorUnits() float64 {
	c := m.currency.get()
	return c.Formatter().ToMajorUnits(m.amount)
}

// Compare compares this Money instance with another and returns:
//   - 1 if this Money is greater than the other
//   - 0 if they are equal
//   - -1 if this Money is less than the other
//   - error if currencies don't match
//
// Example:
//
//	money1 := moneykit.New(1000, "USD")
//	money2 := moneykit.New(1500, "USD")
//
//	result, err := money1.Compare(money2)
//	fmt.Println(result) // -1 (money1 < money2)
func (m *Money) Compare(om *Money) (int, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return int(m.amount), err
	}

	return m.compare(om), nil
}
