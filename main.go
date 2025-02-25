package main

import (
	"fmt"
	"time"
)

type Checker interface {
	CheckDate() bool
}

type Card struct {
	Balance     int
	ExpiredDate string
	CVV         int
	Num         int64
	Owner       string
}

func (c *Card) CheckDate() bool {
	ex, _ := time.Parse("2006.01.02", c.ExpiredDate)
	return time.Now().Before(ex)
}

type CreditCard struct {
	Checker
	Limit int
}

func main() {
	c := &Card{
		Balance:     10000,
		ExpiredDate: "01.02.2023",
		CVV:         132,
		Num:         4234536475474656,
		Owner:       "Vasily Ivanov",
	}
	cc := &CreditCard{
		c,
		100000,
	}

	fmt.Printf("%t", cc.CheckDate())
}
