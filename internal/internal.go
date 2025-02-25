package internal

import (
	"errors"
)

type Debtor interface {
	WrOffDebt() error
}

type Customer struct {
	Name         string
	Age          int
	balance      int
	debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func (c *Customer) WrOffDebt() error {
	if c.debt >= c.balance {
		return errors.New("not possible write off")
	}

	c.balance -= c.debt
	c.debt = 0

	return nil
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		debt:     debt,
		Discount: discount,
	}
}
