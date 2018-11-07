package rpc

import (
	"math/big"
	"github.com/ellsol/go-ethertypes"
)

type SendTransaction struct {
	From     *ethertypes.EtherAddress `json:"from"`
	To       *ethertypes.EtherAddress `json:"to"`
	Gas      *big.Int                 `json:"gas"`
	GasPrice *big.Int                 `json:"gasPrice"`
	Value    *big.Int                 `json:"value"`
	Data     *ethertypes.HexString    `json:"data"`
}

func (it *SendTransaction) ToMap() map[string]interface{} {
	m := make(map[string]interface{})

	m["from"] = it.From.HexString().String()
	m["to"] = it.To.HexString().String()
	m["gas"] = new(ethertypes.HexString).SetBytes(it.Gas.Bytes())
	m["gasPrice"] = new(ethertypes.HexString).SetBytes(it.GasPrice.Bytes())
	m["value"] = new(ethertypes.HexString).SetBytes(it.Value.Bytes())
	m["data"] = it.Data.String()

	return m
}

func (it *SendTransaction) WithFrom(from *ethertypes.EtherAddress) *SendTransaction {
	it.From = from
	return it
}
