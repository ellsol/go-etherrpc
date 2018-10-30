package rpc

import (
	"github.com/ellsol/gox-ethertypes"
)

func (eth Eth) GetLatestBalance(address string, quantity *ethertypes.Quantity) (*ethertypes.EtherValue, error) {
	return eth.GetBalance(address, ethertypes.QuantityLatest())
}
