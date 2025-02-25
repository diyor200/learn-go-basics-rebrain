package internal

type Partner struct {
	Name    string
	Age     int
	balance int
	debt    int
}

func NewPartner(name string, age int, balance int, debt int) *Partner {
	return &Partner{
		Name:    name,
		Age:     age,
		balance: balance,
		debt:    debt,
	}
}

func (c *Partner) WrOffDebt() error {
	c.debt = 0

	return nil
}
