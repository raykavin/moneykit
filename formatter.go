package moneykit

import (
	"math"
	"strconv"
	"strings"
)

// Formatter handles the formatting of monetary amounts according to currency-specific rules.
// It provides methods to format amounts as strings and convert to major units.
type Formatter struct {
	Fraction int    // Number of decimal places
	Decimal  string // Decimal separator
	Thousand string // Thousands separator  
	Grapheme string // Currency symbol
	Template string // Formatting template
}

// NewFormatter creates a new Formatter with the specified formatting rules.
//
// Parameters:
//   - fraction: Number of decimal places
//   - decimal: Decimal separator ("." or ",")
//   - thousand: Thousands separator ("," or "." or "")
//   - grapheme: Currency symbol
//   - template: Format template ("$1" or "1 $")
//
// Example:
//
//	formatter := moneykit.NewFormatter(2, ".", ",", "€", "1 $")
//	result := formatter.Format(123456) // 1,234.56 €
func NewFormatter(fraction int, decimal, thousand, grapheme, template string) *Formatter {
	return &Formatter{
		Fraction: fraction,
		Decimal:  decimal,
		Thousand: thousand,
		Grapheme: grapheme,
		Template: template,
	}
}

// Format converts an integer amount to a formatted string using the formatter's rules.
// The amount should be in the currency's smallest unit.
//
// Parameters:
//   - amount: Amount in smallest currency unit (e.g., cents)
//
// Example:
//
//	formatter := moneykit.NewFormatter(2, ".", ",", "$", "$1")
//	result := formatter.Format(123456) // $1,234.56
//	result = formatter.Format(-500)    // -$5.00
func (f *Formatter) Format(amount int64) string {
	// Work with absolute amount value
	sa := strconv.FormatInt(f.abs(amount), 10)

	if len(sa) <= f.Fraction {
		sa = strings.Repeat("0", f.Fraction-len(sa)+1) + sa
	}

	if f.Thousand != "" {
		for i := len(sa) - f.Fraction - 3; i > 0; i -= 3 {
			sa = sa[:i] + f.Thousand + sa[i:]
		}
	}

	if f.Fraction > 0 {
		sa = sa[:len(sa)-f.Fraction] + f.Decimal + sa[len(sa)-f.Fraction:]
	}
	sa = strings.Replace(f.Template, "1", sa, 1)
	sa = strings.Replace(sa, "$", f.Grapheme, 1)

	// Add minus sign for negative amount.
	if amount < 0 {
		sa = "-" + sa
	}

	return sa
}

// ToMajorUnits converts an integer amount to a floating-point number in major units.
// This is useful when you need the decimal representation of the amount.
//
// Parameters:
//   - amount: Amount in smallest currency unit
//
// Example:
//
//	formatter := moneykit.NewFormatter(2, ".", ",", "$", "$1")
//	result := formatter.ToMajorUnits(123456) // 1234.56
//	result = formatter.ToMajorUnits(500)     // 5.00
func (f *Formatter) ToMajorUnits(amount int64) float64 {
	if f.Fraction == 0 {
		return float64(amount)
	}

	return float64(amount) / float64(math.Pow10(f.Fraction))
}

// abs return absolute value of given integer.
func (f Formatter) abs(amount int64) int64 {
	if amount < 0 {
		return -amount
	}

	return amount
}
