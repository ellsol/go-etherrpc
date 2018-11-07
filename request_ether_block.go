package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/ellsol/go-ethertypes"
	"github.com/ellsol/go-ethertypes/converters"
)

func (client *Client) RequestEtherBlock(method string, params ...interface{}) (*ethertypes.EtherBlock, error) {
	p := []interface{}(params)

	if len(p) != 2 {
		return nil, fmt.Errorf("wrong params in reqeusting ether block: %v", p)
	}

	full := p[1].(bool)

	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no block found for %v", params)
	}


	js, err := json.Marshal(response)

	if err != nil {
		return nil, err
	}


	return converters.DecodeEtherBlock(js, full)
}

