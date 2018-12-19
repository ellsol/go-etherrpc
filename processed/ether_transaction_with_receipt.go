package processed

import (
	"fmt"
	"github.com/ellsol/go-etherrpc"
	"github.com/ellsol/go-ethertypes"
)

func LoadTransactionReceiptAndMerge(et *ethertypes.EtherTransaction, eth *rpc.Eth) (*ethertypes.EtherTransactionWithReceipt, error) {
	receipt, err := eth.GetTransactionReceipt(et.Hash.String())
	if err != nil {
		return nil, fmt.Errorf("GetTransactionReceipt: %v", err.Error())
	}

	return ethertypes.MergeTransactionWithReceipt(et, receipt)
}

