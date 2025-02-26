package main

import (
	"fmt"
	"sync"
	"time"
)

type BankCell struct {
	DollarCell float64
	mu         sync.Mutex
}

func (b *BankCell) GetBalance() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.DollarCell
}

func (b *BankCell) SubBalance(value float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.DollarCell-value >= 0 {
		time.Sleep(time.Millisecond * 1)
		b.DollarCell -= value
	}
}

func (b *BankCell) AddBalance(value float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.DollarCell += value
}

func main() {
	bc := &BankCell{
		DollarCell: 10.0,
	}

	go bc.SubBalance(7.0)
	bc.SubBalance(7.0)
	time.Sleep(time.Second)

	fmt.Println(bc.GetBalance())
}
