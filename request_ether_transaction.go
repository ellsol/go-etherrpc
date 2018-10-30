package rpc

import (
	"fmt"
	"github.com/ellsol/gox-ethertypes"
	"encoding/json"
	"github.com/ellsol/gox-ethertypes/converters"
)

func (client *Client) RequestEtherTransaction(method string, params ...interface{}) (*ethertypes.EtherTransaction, error) {
	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no transaction found for %v", params)
	}

	js, err := json.Marshal(response)

	if err != nil {
		return nil, err
	}

	return converters.DecodeEtherTransactions(js)
}


// transaction receipt

func (client *Client) RequestEtherTransactionReceipt(method string, params ...interface{}) (*ethertypes.EtherTransactionReceipt, error) {
	response, err := checkRPCError(client.Call(method, params...))
	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no transaction found for %v", params)
	}

	js, err := json.Marshal(response.Result)

	if err != nil {
		return nil, err
	}

	return new(ethertypes.TransactionReceiptRaw).FromJSON(js)
}
