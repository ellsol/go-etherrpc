package rpcutils

import (
	"fmt"
	"github.com/ellsol/go-ethertypes"
	"log"
)

const (
	EthereumStandardByteLength = 32
)

var FPTAddress = FunctionParamType{"address", false, 0}
var FPTBool = FunctionParamType{"bool", false, 0}

var FPTUInt8 = FunctionParamType{"uint8", false, 0}
var FPTUInt16 = FunctionParamType{"uint16", false, 0}
var FPTUInt32 = FunctionParamType{"uint32", false, 0}
var FPTUInt64 = FunctionParamType{"uint64", false, 0}
var FPTUInt128 = FunctionParamType{"uint128", false, 0}
var FPTUInt256 = FunctionParamType{"uint256", false, 0}

var FPTBytes = FunctionParamType{"bytes", true, 0}
var FPTBytes32 = FunctionParamType{"bytes32", false, 0}
var FPTString = FunctionParamType{"string", true, 0}
var FPTByte32Array = FunctionParamType{"bytes32[]", true, 0}

type FunctionParamType struct {
	Type        string
	IsDynamic   bool
	ArrayLength int
}

type FunctionParam struct {
	Type  string
	Value interface{}
}

func NewFunctionParamType(t string, arraylength int) *FunctionParamType {
	fmt.Println(t)
	fpt := &FunctionParamType{}
	switch t {
	case FPTString.Type:
		fpt = &FPTString
	case FPTBytes.Type:
		fpt = &FPTBytes
	case FPTBytes32.Type:
		fpt = &FPTBytes32
	case FPTAddress.Type:
		fpt = &FPTAddress
	case FPTBool.Type:
		fpt = &FPTBool
	case FPTUInt8.Type:
		fpt = &FPTUInt8
	case FPTUInt16.Type:
		fpt = &FPTUInt16
	case FPTUInt32.Type:
		fpt = &FPTUInt32
	case FPTUInt64.Type:
		fpt = &FPTUInt64
	case FPTUInt128.Type:
		fpt = &FPTUInt128
	case FPTUInt256.Type:
		fpt = &FPTUInt256
	case FPTByte32Array.Type:
		fpt = &FPTByte32Array
	}

	fpt.ArrayLength = arraylength

	return fpt
}

type FunctionSignature struct {
	Params []FunctionParamType
}

func NewFunctionSignature(values ...FunctionParamType) *FunctionSignature {
	params := make([]FunctionParamType, 0)

	for _, v := range values {
		params = append(params, v)
	}

	return &FunctionSignature{
		Params: params,
	}
}

func (fs *FunctionSignature) DecodeEventData(input string) ([]FunctionParam, error) {
	//if len(input) < 2+64*fs.Len() {
	//	return nil, fmt.Errorf("input length should be at least methodsignature 10 + %v * 64 (32byte in hex) long", fs.Len())
	//}

	hsinput, err := ethertypes.NewHexString().SetString(input)

	if err != nil {
		return nil, err
	}

	return fs.DecodeEventDataFromHex(hsinput)
}
func (fs *FunctionSignature) DecodeFunctionInput(input string) ([]FunctionParam, error) {
	fmt.Println(fs)
	fmt.Println(input)
	if len(input) < 10+64*fs.Len() {
		return nil, fmt.Errorf("input length should be at least methodsignature 10 + %v * 64 (32byte in hex) long", fs.Len())
	}

	hsinput, err := ethertypes.NewHexString().SetString(input)

	if err != nil {
		return nil, err
	}
	b := hsinput.Bytes()

	return fs.DecodeFunctionInputFromHex(ethertypes.NewHexString().SetBytes(b[4:]))
}

func (fs *FunctionSignature) DecodeEventDataFromHex(input *ethertypes.HexString) ([]FunctionParam, error) {
	fp := make([]FunctionParam, len(fs.Params))

	log.Println("input: ", input.Plain())

	head := fs.ReadHead(input)
	//log.Println("Head: ")
	//for k,v := range head {
	//	log.Printf("%v: %v",k, v.Hash())
	//}

	body, err := fs.ReadBody(input)
	if err != nil {
		return nil, err
	}
	headLength := len(fs.Params) * EthereumStandardByteLength

	for k, h := range head {
		p := fs.Params[k]
		pType := p.Type

		//log.Println("Type: ", pType)
		//log.Println("Value: ", h.Hash())
		//log.Println("..........................", h.Hash())
		if !p.IsDynamic {
			value, err := fromNonDynamicValue(&h, pType)

			if err != nil {
				return nil, err
			}

			fp[k] = FunctionParam{pType, value}
		} else {
			f := FunctionParam{pType, ""}
			location := int(h.Int64()) - headLength

			switch pType {
			case FPTString.Type:
				s, err := decodeString(body, location)
				if err != nil {
					return nil, err
				}
				f = FunctionParam{pType, s}
			case FPTString.Type:
				s, err := decodeBytes(body, location)
				if err != nil {
					return nil, err
				}
				f = FunctionParam{pType, s}
			case FPTByte32Array.Type:
				s, err := decodeBytesArray(body, location)
				if err != nil {
					return nil, err
				}
				f = FunctionParam{pType, s}
			}

			fp[k] = f
		}
	}

	return fp, nil
}



func (fs *FunctionSignature) DecodeFunctionInputFromHex(input *ethertypes.HexString) ([]FunctionParam, error) {
	fp := make([]FunctionParam, len(fs.Params))

	//log.Println("input: ", input.Plain())

	head := fs.ReadHead(input)
	//log.Println("Head: ")
	//for k,v := range head {
	//	log.Printf("%v: %v",k, v.Hash())
	//}

	body, err := fs.ReadBody(input)
	if err != nil {
		return nil, err
	}
	headLength := len(fs.Params) * EthereumStandardByteLength

	for k, h := range head {
		p := fs.Params[k]
		pType := p.Type

		//log.Println("Type: ", pType)
		//log.Println("Value: ", h.Hash())
		//log.Println("..........................", h.Hash())
		if !p.IsDynamic {
			value, err := fromNonDynamicValue(&h, pType)

			if err != nil {
				return nil, err
			}

			fp[k] = FunctionParam{pType, value}
		} else {
			f := FunctionParam{pType, ""}
			location := int(h.Int64()) - headLength

			switch pType {
			case FPTString.Type:
				s, err := decodeString(body, location)
				if err != nil {
					return nil, err
				}
				f = FunctionParam{pType, s}
			case FPTString.Type:
				s, err := decodeBytes(body, location)
				if err != nil {
					return nil, err
				}
				f = FunctionParam{pType, s}
			case FPTByte32Array.Type:
				s, err := decodeBytesArray(body, location)
				if err != nil {
					return nil, err
				}
				f = FunctionParam{pType, s}
			}

			fp[k] = f
		}
	}

	return fp, nil
}

func (fs *FunctionSignature) Len() int {
	return len(fs.Params)
}

func (fs *FunctionSignature) ReadHead(input *ethertypes.HexString) []ethertypes.HexString {
	b := input.Bytes()

	result := make([]ethertypes.HexString, fs.Len())

	for k := range result {
		from := k * EthereumStandardByteLength
		to := EthereumStandardByteLength + from
		result[k] = *new(ethertypes.HexString).SetBytes(b[from:to])
	}

	return result
}

func (fs *FunctionSignature) ReadBody(input *ethertypes.HexString) (*ethertypes.HexString, error) {

	b := input.Bytes()[fs.Len()*EthereumStandardByteLength:]

	if len(b)%EthereumStandardByteLength != 0 {
		return nil, fmt.Errorf("function input body is not factor of EthereumStandardByteLength bytes")
	}

	return new(ethertypes.HexString).SetBytes(b), nil
}

func fromNonDynamicValue(val *ethertypes.HexString, paramType string) (interface{}, error) {
	switch paramType {
	case FPTAddress.Type:
		return new(ethertypes.EtherAddress).Set32ByteString(val.String())
	case FPTBytes.Type:
		return val.Bytes(), nil
	case FPTBytes32.Type:
		return val.String(), nil
	case FPTUInt8.Type:
		return int(val.BigInt().Int64()), nil
	case FPTUInt16.Type:
		return int(val.BigInt().Int64()), nil
	case FPTUInt32.Type:
		return int(val.BigInt().Int64()), nil
	case FPTUInt64.Type:
		return val.BigInt().Int64(), nil
	case FPTUInt128.Type:
		return val.BigInt(), nil
	case FPTUInt256.Type:
		return val.BigInt(), nil
	case FPTBool.Type:
		return val.Int64() > 0, nil
	}

	return "", nil
}

func decodeString(body *ethertypes.HexString, location int) (string, error) {
	if len(body.Bytes()) < location {
		return "", fmt.Errorf("function input body too short")
	}

	dataLengthPart := body.Bytes()[location : location+EthereumStandardByteLength]
	length := int(new(ethertypes.HexString).SetBytes(dataLengthPart).Int64())
	dataPart := body.Bytes()[location+EthereumStandardByteLength : location+EthereumStandardByteLength+length]
	result := new(ethertypes.HexString).SetBytes(dataPart).Ascii()
	return result, nil
}

func decodeBytes(body *ethertypes.HexString, location int) ([]byte, error) {
	if len(body.Bytes()) < location {
		return nil, fmt.Errorf("function input body too short")
	}

	dataLengthPart := body.Bytes()[location : location+EthereumStandardByteLength]
	length := int(new(ethertypes.HexString).SetBytes(dataLengthPart).Int64())
	dataPart := body.Bytes()[location+EthereumStandardByteLength : location+EthereumStandardByteLength+length]

	return dataPart, nil
}

func decodeBytesArray(body *ethertypes.HexString, location int) ([][]byte, error) {
	if len(body.Bytes()) < location {
		return nil, fmt.Errorf("function input body too short")
	}

	dataLengthPart := body.Bytes()[location : location+EthereumStandardByteLength]
	length := int(new(ethertypes.HexString).SetBytes(dataLengthPart).Int64())

	result := make([][]byte, length)

	for i := 0; i < length; i++ {
		result[i] = body.Bytes()[location+(i+1)*EthereumStandardByteLength : location+(i+2)*EthereumStandardByteLength]
	}

	return result, nil
}
