package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetOrders(t *testing.T) {
	req := require.New(t)
	orders := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	req.Len(orders, 9)
	req.NotEqual(orders[0], "2")
}
