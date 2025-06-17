package moneykit

import "math"

// calculator implements the Calculator interface
type calculator struct{}

// NewCalculator creates and returns a new Calculator instance
func NewCalculator() *calculator {
	return &calculator{}
}

// Add returns the sum of two amounts
func (c *calculator) add(a, b Amount) Amount {
	return a + b
}

// Subtract returns the difference of two amounts (a - b)
func (c *calculator) subtract(a, b Amount) Amount {
	return a - b
}

// Multiply returns the product of an amount and a multiplier
func (c *calculator) multiply(a Amount, m int64) Amount {
	return a * Amount(m)
}

// Divide returns the quotient of an amount divided by a divisor
// Performs integer division which truncates towards zero
// Panics if divisor is 0 (standard Go behavior for division by zero)
func (c *calculator) divide(a Amount, d int64) Amount {
	return a / Amount(d)
}

// Modulus returns the remainder of an amount divided by a divisor
// Panics if divisor is 0 (standard Go behavior for modulus by zero)
func (c *calculator) modulus(a Amount, d int64) Amount {
	return a % Amount(d)
}

// Allocate distributes an amount proportionally based on ratio and shares
// Formula: (amount * ratio) / shares
// This is useful for proportional distribution of costs, taxes, or revenues
// Returns 0 if amount is 0 or shares is 0 to avoid division by zero
func (c *calculator) allocate(a Amount, r, s int64) Amount {
	if a == 0 || s == 0 {
		return 0
	}

	return a * Amount(r) / Amount(s)
}

// Absolute returns the absolute value of an amount
func (c *calculator) absolute(a Amount) Amount {
	if a < 0 {
		return -a
	}
	return a
}

// Negative returns the negative value of an amount
func (c *calculator) negative(a Amount) Amount {
	return -a
}

// Round rounds an amount to the specified precision (number of decimal places)
// Uses "round half up" strategy where 0.5 rounds up to 1
//
// Examples:
//
//	Round(1235, 2) with amount representing 12.35 rounds to 12.40 (1240)
//	Round(1234, 2) with amount representing 12.34 rounds to 12.30 (1230)
//	Round(1250, 1) with amount representing 12.50 rounds to 13.0 (1300)
func (c *calculator) round(a Amount, precision int) Amount {
	if a == 0 {
		return 0
	}

	// Work with absolute value and preserve sign
	absAmount := c.absolute(a)
	factor := int64(math.Pow(10, float64(precision)))
	remainder := absAmount % Amount(factor)

	// Round up if remainder is greater than or equal to half the factor
	if remainder >= Amount(factor)/2 {
		absAmount += Amount(factor)
	}

	// Truncate to desired precision
	rounded := (absAmount / Amount(factor)) * Amount(factor)

	// Restore original sign
	if a < 0 {
		return -rounded
	}
	return rounded
}
