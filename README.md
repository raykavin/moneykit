# MoneyKit

[![Go Reference](https://pkg.go.dev/badge/github.com/raykavin/moneykit.svg)](https://pkg.go.dev/github.com/raykavin/moneykit)
[![Go Report Card](https://goreportcard.com/badge/github.com/raykavin/moneykit)](https://goreportcard.com/report/github.com/raykavin/moneykit)
[![License](https://img.shields.io/badge/license-GPL-blue.svg)](LICENSE)

MoneyKit is a comprehensive Go library for handling monetary values with precision. It provides safe arithmetic operations, currency support, formatting, and database integration while avoiding floating-point precision issues.

## Features

- 🔢 **Precise calculations** - Uses integer arithmetic to avoid floating-point precision issues
- 💰 **165+ currencies** - Built-in support for all active ISO 4217 currency codes
- 🧮 **Safe operations** - Currency-aware arithmetic with mismatch detection
- 📊 **Smart allocation** - Split amounts with proper remainder distribution
- 🎨 **Flexible formatting** - Localized display with customizable templates
- 🗃️ **Database ready** - Built-in SQL driver support for easy persistence
- 🔄 **JSON support** - Seamless marshaling/unmarshaling with customization hooks
- ⚡ **Zero dependencies** - Pure Go implementation

## Installation

```bash
go get github.com/raykavin/moneykit
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/raykavin/moneykit"
)

func main() {
    // Create money instances
    price := moneykit.New(2500, "USD")  // $25.00
    tax := moneykit.New(250, "USD")     // $2.50
    
    // Perform calculations
    total, err := price.Add(tax)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(total.Display()) // $27.50
    
    // Split bills
    parties, err := total.Split(3)
    if err != nil {
        log.Fatal(err)
    }
    
    for i, party := range parties {
        fmt.Printf("Person %d pays: %s\n", i+1, party.Display())
    }
    // Person 1 pays: $9.17
    // Person 2 pays: $9.17
    // Person 3 pays: $9.16
}
```

## Core Concepts

### Money Representation

MoneyKit represents money as integers in the currency's smallest unit (e.g., cents for USD, pence for GBP):

```go
// $10.50 represented as 1050 cents
money := moneykit.New(1050, "USD")

// £5.99 represented as 599 pence  
price := moneykit.New(599, "GBP")

// ¥1000 (whole yen, no decimals)
yen := moneykit.New(1000, "JPY")
```

### Creating Money from Floats

```go
// Create from float (automatically converts to proper units)
price := moneykit.NewFromFloat(19.99, "USD") // $19.99
fmt.Println(price.Amount()) // 1999 (cents)
```

## Basic Operations

### Arithmetic

```go
base := moneykit.New(1000, "USD")
tip := moneykit.New(150, "USD")

// Addition
total, err := base.Add(tip)
// Result: $11.50

// Multiple additions
meal := moneykit.New(2500, "USD")
tax := moneykit.New(200, "USD")
tip := moneykit.New(375, "USD")
total, err := meal.Add(tax, tip)
// Result: $30.75

// Subtraction
discount := moneykit.New(500, "USD")
final, err := total.Subtract(discount)
// Result: $25.75

// Multiplication
doubled := base.Multiply(2)
// Result: $20.00

// Chain operations
result := base.Multiply(3, 2) // 3 * 2 = 6
// Result: $60.00
```

### Comparisons

```go
price1 := moneykit.New(1000, "USD")
price2 := moneykit.New(1500, "USD")

// Equality
equal, err := price1.Equals(price2)
// false, nil

// Comparisons
greater, err := price2.GreaterThan(price1)
// true, nil

less, err := price1.LessThan(price2)
// true, nil

// Generic comparison (-1, 0, 1)
result, err := price1.Compare(price2)
// -1, nil
```

### Status Checks

```go
zero := moneykit.New(0, "USD")
positive := moneykit.New(100, "USD")
negative := moneykit.New(-50, "USD")

fmt.Println(zero.IsZero())       // true
fmt.Println(positive.IsPositive()) // true
fmt.Println(negative.IsNegative()) // true
```

## Advanced Features

### Smart Splitting

Split amounts evenly with automatic remainder distribution:

```go
bill := moneykit.New(1000, "USD") // $10.00

// Split among 3 people
shares, err := bill.Split(3)
// shares[0]: $3.34
// shares[1]: $3.33  
// shares[2]: $3.33
```

### Proportional Allocation

Allocate amounts based on ratios:

```go
revenue := moneykit.New(10000, "USD") // $100.00

// Allocate 50%, 30%, 20%
portions, err := revenue.Allocate(50, 30, 20)
// portions[0]: $50.00
// portions[1]: $30.00
// portions[2]: $20.00

// Handle remainders automatically
amount := moneykit.New(100, "USD") // $1.00
parts, err := amount.Allocate(33, 33, 33)
// parts[0]: $0.34 (gets the extra penny)
// parts[1]: $0.33
// parts[2]: $0.33
```

### Rounding

```go
money := moneykit.New(1567, "USD") // $15.67
rounded := money.Round()           // $16.00 (rounds to currency precision)
```

### Absolute and Negative Values

```go
debt := moneykit.New(-500, "USD")
absolute := debt.Absolute() // $5.00
negative := debt.Negative() // -$5.00 (idempotent for negative values)
```

## Currency Support

### Built-in Currencies

MoneyKit includes all active ISO 4217 currencies:

```go
usd := moneykit.New(100, "USD")     // US Dollar
eur := moneykit.New(100, "EUR")     // Euro  
jpy := moneykit.New(100, "JPY")     // Japanese Yen (no decimals)
bhd := moneykit.New(100, "BHD")     // Bahraini Dinar (3 decimals)
```

### Custom Currencies

```go
// Add custom currency
moneykit.AddCurrency("BTC", "₿", "₿1", ".", ",", 8)

bitcoin := moneykit.New(100000000, "BTC") // 1.00000000 BTC
fmt.Println(bitcoin.Display()) // ₿1.00000000
```

### Currency Information

```go
currency := moneykit.GetCurrency("USD")
fmt.Println(currency.Code)        // USD
fmt.Println(currency.Grapheme)    // $
fmt.Println(currency.Fraction)    // 2 (decimal places)

// Get by numeric code
eur := moneykit.GetCurrencyByNumericCode("978") // EUR
```

## Formatting and Display

### Default Formatting

```go
money := moneykit.New(123456, "USD")
fmt.Println(money.Display())      // $1,234.56
fmt.Println(money.AsMajorUnits()) // 1234.56
```

### Different Currency Formats

```go
eur := moneykit.New(123456, "EUR")
fmt.Println(eur.Display()) // €1,234.56

jpy := moneykit.New(1234, "JPY")  
fmt.Println(jpy.Display()) // ¥1,234

dkk := moneykit.New(123456, "DKK")
fmt.Println(dkk.Display()) // kr 1.234,56 (Danish format)
```

### Custom Formatting

```go
formatter := moneykit.NewFormatter(
    2,      // decimal places
    ".",    // decimal separator  
    ",",    // thousands separator
    "$",    // currency symbol
    "$1",   // template
)

formatted := formatter.Format(123456) // $1,234.56
majorUnits := formatter.ToMajorUnits(123456) // 1234.56
```

## Database Integration

MoneyKit provides seamless database integration with the `sql/driver` interface:

### Saving to Database

```go
money := moneykit.New(2550, "USD")

// Money implements driver.Valuer
_, err := db.Exec("INSERT INTO orders (total) VALUES (?)", money)
// Stores as: "2550|USD"
```

### Loading from Database

```go
var money moneykit.Money

// Money implements sql.Scanner  
err := db.QueryRow("SELECT total FROM orders WHERE id = ?", 1).Scan(&money)

fmt.Println(money.Display()) // $25.50
```

### Custom Separator

```go
// Change the default separator
moneykit.DBMoneyValueSeparator = ":"

money := moneykit.New(1000, "EUR")
value, _ := money.Value() // "1000:EUR"
```

## JSON Serialization

### Default JSON Format

```go
money := moneykit.New(2550, "USD")

data, err := json.Marshal(money)
// {"amount":2550,"currency":"USD"}

var loaded moneykit.Money
err = json.Unmarshal(data, &loaded)
```

### Custom JSON Format

```go
// Override marshaling behavior
moneykit.MarshalJSON = func(m moneykit.Money) ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "value":    m.AsMajorUnits(),
        "currency": m.Currency().Code,
    })
}

money := moneykit.New(2550, "USD")
data, _ := json.Marshal(money)
// {"value":25.5,"currency":"USD"}
```

## Error Handling

### Currency Mismatch

```go
usd := moneykit.New(100, "USD")
eur := moneykit.New(100, "EUR")

_, err := usd.Add(eur)
if errors.Is(err, moneykit.ErrCurrencyMismatch) {
    fmt.Println("Cannot add different currencies")
}
```

### Common Patterns

```go
// Safe arithmetic with error checking
func calculateTotal(prices ...*moneykit.Money) (*moneykit.Money, error) {
    if len(prices) == 0 {
        return moneykit.New(0, "USD"), nil
    }
    
    total := prices[0]
    for _, price := range prices[1:] {
        var err error
        total, err = total.Add(price)
        if err != nil {
            return nil, fmt.Errorf("failed to add price: %w", err)
        }
    }
    
    return total, nil
}
```

## Best Practices

### 1. Always Use Smallest Units

```go
// ✅ Good - store in cents
price := moneykit.New(2550, "USD") // $25.50

// ❌ Avoid - floating point precision issues  
price := moneykit.NewFromFloat(25.50, "USD") // Use sparingly
```

### 2. Handle Currency Mismatches

```go
// ✅ Good - explicit currency checking
func addPrices(a, b *moneykit.Money) (*moneykit.Money, error) {
    if !a.SameCurrency(b) {
        return nil, fmt.Errorf("currency mismatch: %s vs %s", 
            a.Currency().Code, b.Currency().Code)
    }
    return a.Add(b)
}
```

### 3. Use Allocation for Fair Distribution

```go
// ✅ Good - handles remainders properly
shares, err := total.Allocate(1, 1, 1) // Equal shares

// ❌ Avoid - may lose pennies
third := total.Amount() / 3
```

### 4. Validate Before Operations

```go
func processPayment(amount *moneykit.Money) error {
    if amount.IsNegative() {
        return errors.New("payment amount cannot be negative")
    }
    if amount.IsZero() {
        return errors.New("payment amount cannot be zero")  
    }
    // Process payment...
    return nil
}
```

## Examples

### E-commerce Cart

```go
type CartItem struct {
    Price    *moneykit.Money
    Quantity int
}

type Cart struct {
    Items    []CartItem
    Currency string
}

func (c *Cart) Total() (*moneykit.Money, error) {
    total := moneykit.New(0, c.Currency)
    
    for _, item := range c.Items {
        itemTotal := item.Price.Multiply(int64(item.Quantity))
        var err error
        total, err = total.Add(itemTotal)
        if err != nil {
            return nil, err
        }
    }
    
    return total, nil
}

func (c *Cart) SplitBill(people int) ([]*moneykit.Money, error) {
    total, err := c.Total()
    if err != nil {
        return nil, err
    }
    
    return total.Split(people)
}
```

### Investment Portfolio

```go
type Holding struct {
    Symbol string
    Value  *moneykit.Money
}

type Portfolio struct {
    Holdings []Holding
    Currency string
}

func (p *Portfolio) TotalValue() (*moneykit.Money, error) {
    total := moneykit.New(0, p.Currency)
    
    for _, holding := range p.Holdings {
        var err error
        total, err = total.Add(holding.Value)
        if err != nil {
            return nil, err
        }
    }
    
    return total, nil
}

func (p *Portfolio) AllocateByWeight(weights []int) ([]*moneykit.Money, error) {
    total, err := p.TotalValue()
    if err != nil {
        return nil, err
    }
    
    return total.Allocate(weights...)
}
```
## 🤝 Contributing

Contributions to MoneyKit are welcome! Here are some ways you can help improve the project:

- **Report bugs and suggest features** by opening issues on GitHub
- **Submit pull requests** with bug fixes or new features
- **Improve documentation** to help other users and developers
- **Share your custom strategies** with the community

## 📄 License

MoneyKit is distributed under the **GNU General Public License v3.0**.  
For complete license terms and conditions, see the [LICENSE](LICENSE.md) file in the repository.

Copyright © [Raykavin Meireles](https://github.com/raykavin)

---

## 📬 Contact

For support, collaboration, or questions about MoneyKit:

**Email**: [raykavin.meireles@gmail.com](mailto:raykavin.meireles@gmail.com)  
**GitHub**: [@raykavin](https://github.com/raykavin)  
**LinkedIn**: [@raykavin.dev](https://www.linkedin.com/in/raykavin-dev)  
**Instagram**: [@raykavin.dev](https://www.instagram.com/raykavin.dev)