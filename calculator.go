package money

import (
	"github.com/Rhymond/go-money/core/money"
	"math"
)

type calculator struct{}

func (c *calculator) add(a, b money.Amount) money.Amount {
	return a + b
}

func (c *calculator) subtract(a, b money.Amount) money.Amount {
	return a - b
}

func (c *calculator) multiply(a money.Amount, m int64) money.Amount {
	return a * m
}

func (c *calculator) divide(a money.Amount, d int64) money.Amount {
	return a / d
}

func (c *calculator) modulus(a money.Amount, d int64) money.Amount {
	return a % d
}

func (c *calculator) allocate(a money.Amount, r, s uint) money.Amount {
	if a == 0 || s == 0 {
		return 0
	}

	return a * int64(r) / int64(s)
}

func (c *calculator) absolute(a money.Amount) money.Amount {
	if a < 0 {
		return -a
	}

	return a
}

func (c *calculator) negative(a money.Amount) money.Amount {
	if a > 0 {
		return -a
	}

	return a
}

func (c *calculator) round(a money.Amount, e int) money.Amount {
	if a == 0 {
		return 0
	}

	absam := c.absolute(a)
	exp := int64(math.Pow(10, float64(e)))
	m := absam % exp

	if m > (exp / 2) {
		absam += exp
	}

	absam = (absam / exp) * exp

	if a < 0 {
		a = -absam
	} else {
		a = absam
	}

	return a
}
