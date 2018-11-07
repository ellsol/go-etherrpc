package rpc

import (
	"github.com/ellsol/go-ethertypes"
)

func (eth Eth) GetLatestBalance(address string, quantity *ethertypes.Quantity) (*ethertypes.EtherValue, error) {
	return eth.GetBalance(address, ethertypes.QuantityLatest())
}
