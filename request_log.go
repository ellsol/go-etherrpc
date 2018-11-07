package rpc

import (
	"github.com/ellsol/go-ethertypes"
	"github.com/ellsol/go-ethertypes/converters"
)

func (client *Client) RequestEtherLogList(method string, params ...interface{}) ([]ethertypes.EtherLog, error) {
	response, err := client.Call(method, params...)

	if err != nil {
		return nil, err
	}

	js, err := response.ToJSONBytes()

	if err != nil {
		return nil, err
	}

	converters.DecodeEtherLogList(js)
}
