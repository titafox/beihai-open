package tests

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFloat(t *testing.T) {
	f := 109.3 * 4 / 100

	s := Decimal(f)
	value := strconv.FormatFloat(s, 'f', 2, 64)
	t.Log(value)
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
