package rpc

import (
	"fmt"
	"github.com/ellsol/gox/typex"
	"github.com/ellsol/go-ethertypes"
)

/*
	Request EtherValue
 */
func (client *Client) RequestEtherValue(method string, params ...interface{}) (*ethertypes.EtherValue, error) {
	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	val, ok := response.Result.(string)

	if !ok {
		return nil, fmt.Errorf("could not parse string from %s", response.Result)
	}

	r, err := ethertypes.HexToBigInt(val)

	if err != nil {
		return nil, err
	}

	return new(ethertypes.EtherValue).SetBigInt(r), nil
}

/*
	Request Int64
 */
func (client *Client) RequestInt64(method string, params ...interface{}) (int64, error) {
	response, err := checkRPCError(client.Call(method, params...))
	if err != nil {
		return -1, err
	}

	if response.Result == nil {
		return -1, fmt.Errorf("m: %v, p: %v didn't return error but also no response", method, params)
	}

	val, ok := response.Result.(string)

	if !ok {
		return 0, fmt.Errorf("could not parse string from %s", response.Result)
	}

	hs, err := ethertypes.NewHexString().SetString(val)
	if err != nil {
		return -1, err
	}
	return hs.Int64(), nil
}

func (client *Client) RequestString(method string, params ...interface{}) (string, error) {
	response, err := checkRPCError(client.Call(method, params...))
	if err != nil {
		return "", err
	}

	val, ok := response.Result.(string)

	if !ok {
		return "", fmt.Errorf("could not parse string from %s", response.Result)
	}

	return val, nil
}

/*
	Request Bool
 */
func (client *Client) RequestBool(method string, params ...interface{}) (bool, error) {
	response, err := checkRPCError(client.Call(method, params...))
	if err != nil {
		return false, err
	}

	if response.Result == nil {
		return false, fmt.Errorf("m: %v, p: %v didn't return error but also no response", method, params)
	}

	val, ok := response.Result.(bool)

	if !ok {
		return false, fmt.Errorf("could not parse bool from %v", response.Result)
	}

	return val, nil
}

/*
	Request HexString
 */
func (client *Client) RequestHexString(method string, params ...interface{}) (*ethertypes.HexString, error) {
	response, err := checkRPCError(client.Call(method, params...))
	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("m: %v, p: %v didn't return error but also no response", method, params)
	}

	val, ok := response.Result.(string)

	if !ok {
		return nil, fmt.Errorf("could not parse string from %v", response.Result)
	}

	return ethertypes.NewHexString().SetString(val)
}

/*
	Request HexStringList
 */
func (client *Client) RequestHexStringList(method string, params ...interface{}) ([]ethertypes.HexString, error) {
	response, err := checkRPCError(client.Call(method, params...))
	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("m: %v, p: %v didn't return error but also no response", method, params)
	}

	val, ok := response.Result.([]interface{})

	if !ok {
		return nil, fmt.Errorf("could not parse string list from %v", response.Result)
	}

	sl, err := typex.StringListFromInterfaceList(val)

	if err != nil {
		return nil, err
	}

	return ethertypes.StringListToHexStringList(sl)
}
