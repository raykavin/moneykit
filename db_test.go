package moneykit

import (
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney_Value(t *testing.T) {
	tests := []struct {
		have      *Money
		separator string
		want      string
		wantErr   bool
	}{
		{
			have:      New(10, CAD),
			separator: "|",
			want:      "10|CAD",
		},
		{
			have:      New(-10, USD),
			separator: "+-+",
			want:      "-10+-+USD",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#v", tt.have), func(t *testing.T) {
			want := driver.Value(tt.want)
			DBMoneyValueSeparator = tt.separator
			got, err := tt.have.Value()

			assert.NoError(t, err, "Value() should not return an error")
			assert.Equal(t, want, got, "Value() should return expected driver.Value")
		})
	}
}

func TestMoney_Scan(t *testing.T) {
	tests := []struct {
		src       interface{}
		separator string
		want      *Money
		wantErr   bool
	}{
		{
			src:  "10|CAD",
			want: New(10, CAD),
		},
		{
			src:  "20|USD",
			want: New(20, USD),
		},
		{
			src:       "30000,IDR",
			separator: ",",
			want:      New(30000, IDR),
		},
		{
			src:     "10|",
			wantErr: true,
		},
		{
			src:     "|SAR",
			wantErr: true,
		},
		{
			src:     "10",
			wantErr: true,
		},
		{
			src:     "USD",
			wantErr: true,
		},
		{
			src:     "USD|10",
			wantErr: true,
		},
		{
			src:     "",
			wantErr: true,
		},
		{
			src:     "a|b|c",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#v", tt.src), func(t *testing.T) {
			if tt.separator != "" {
				DBMoneyValueSeparator = tt.separator
			} else {
				DBMoneyValueSeparator = DefaultDBMoneyValueSeparator
			}
			got := &Money{}
			err := got.Scan(tt.src)

			if tt.wantErr {
				assert.Error(t, err, "Scan() should return an error for invalid input")
				return
			}

			assert.NoError(t, err, "Scan() should not return an error for valid input")
			assert.NotNil(t, got, "Scan() result should not be nil")

			eq, err := tt.want.Equals(got)
			assert.NoError(t, err, "Equals() should not return an error")
			assert.True(t, eq, "Scanned money should equal expected value. Got: %s %s, Want: %s %s",
				got.Display(), got.Currency().Code, tt.want.Display(), tt.want.Currency().Code)
		})
	}
}

func TestCurrency_Value(t *testing.T) {
	for code, cc := range currencies {
		t.Run(code, func(t *testing.T) {
			want := driver.Value(code)

			got, err := cc.Value()
			assert.NoError(t, err, "Value() should not return an error")
			assert.Equal(t, want, got, "Value() should return expected driver.Value")
		})
	}
}

func TestCurrency_Scan(t *testing.T) {
	for code, want := range currencies {
		t.Run(code, func(t *testing.T) {
			src := interface{}(code)

			got := &Currency{}
			err := got.Scan(src)
			assert.NoError(t, err, "Scan() should not return an error")
			assert.Equal(t, want, got, "Scan() should return expected currency")
		})
	}
}
