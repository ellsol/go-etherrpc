package rpc

import (
	"github.com/ellsol/go-ethereum-rpc/rpcutils"
	"github.com/ellsol/gox-ethertypes"
)

type EthCallParams struct {
	From     string `json:"from,omitempty"`      // (optional) 20 Bytes - The url the transaction is send from.
	To       string `json:"to,omitempty"`        // (optional when creating new contract) 20 Bytes - The url the transaction is directed to.
	Gas      int64  `json:"gas,omitempty"`       // (optional) Integer of the gas provided for the transaction execution. eth_call consumes zero gas, but this parameter may be needed by some executions.
	GasPrice int64  `json:"gas_price,omitempty"` // (optional) Integer of the gas price used for each paid gas.
	Value    int64  `json:"value,omitempty"`     // (optional) Integer of the value sent with this transaction.
	Data     string `json:"data,omitempty"`      // (optional) 4 byte hash of the method signature followed by encoded parameters.
}

func (ecp *EthCallParams) ToMap() map[string]interface{} {
	m := make(map[string]interface{}, 0)

	m["to"] = ecp.To

	if ecp.Data != "" {
		m["data"] = ecp.Data
	}

	return m
}

func (ecp *EthCallParams) ToContract(address string, functionKeccak string) *EthCallParams {
	return &EthCallParams{
		Data: functionKeccak,
		To:   address,
	}
}

func (ecp *EthCallParams) ToContractWithArgument(address string, functionSignature string, arg []byte) (*EthCallParams, error) {
	keccak, err := rpcutils.Signature2MethodId(functionSignature)

	if err != nil {
		return nil, err
	}

	hsf, err := ethertypes.NewHexString().SetString(keccak)

	if err != nil {
		return nil, err
	}

	hsa := ethertypes.NewHexString().SetBytes(arg).LeftPadTo(32)

	data := hsf.Concat(hsa)

	return &EthCallParams{
		Data: data.String(),
		To:   address,
	}, nil
}

func (ecp *EthCallParams) ToContractWithArgumentsOfLength32(address string, functionSignature string, arg ...[]byte) (*EthCallParams, error) {
	keccak, err := rpcutils.Signature2MethodId(functionSignature)

	if err != nil {
		return nil, err
	}

	hsf, err := ethertypes.NewHexString().SetString(keccak)

	if err != nil {
		return nil, err
	}

	for _, v := range arg {
		hsf = hsf.Concat(ethertypes.NewHexString().SetBytes(v).LeftPadTo(32))
	}

	//hsa := ethertypes.NewHexString().SetStringFromBytes(arg).PadTo(32)


	return &EthCallParams{
		Data: hsf.String(),
		To:   address,
	}, nil
}

func (ecp *EthCallParams) ToContractWithValue(address string, functionKeccak string) *EthCallParams {
	return &EthCallParams{
		Data: functionKeccak,
		To:   address,
	}
}

type EthEstimateGasParams struct {
	From     string            `json:"from"`      // (optional) 20 Bytes - The url the transaction is send from.
	To       string            `json:"to"`        // (optional when creating new contract) 20 Bytes - The url the transaction is directed to.
	Gas      int64             `json:"gas"`       // (optional) Integer of the gas provided for the transaction execution. eth_call consumes zero gas, but this parameter may be needed by some executions.
	GasPrice int64             `json:"gas_price"` // (optional) Integer of the gas price used for each paid gas.
	Value    int64             `json:"value"`     // (optional) Integer of the value sent with this transaction.
	Data     string            `json:"data"`      // (optional) 4 byte hash of the method signature followed by encoded parameters.
	Quantity ethertypes.Quantity `json:"quantity"`  // (optional) Integer block number, or the string 'latest', 'earliest' or 'pending', see the default block parameter.
}

type NewFilterParams struct {
	FromBlock ethertypes.Quantity `json:"fromBlock"` // Tag - (optional) (default: latest) Integer block number, or 'latest' for the last mined block or 'pending', 'earliest' for not yet mined transactions.
	ToBlock   ethertypes.Quantity `json:"toBlock"`   // Tag - (optional) (default: latest) Integer block number, or 'latest' for the last mined block or 'pending', 'earliest' for not yet mined transactions.
	Address   string            `json:"url"`
	Topics    [3][]string       `json:"topics"`
}

func (p *NewFilterParams) ToMap() map[string]interface{} {
	m := make(map[string]interface{})

	m["fromBlock"] = p.FromBlock.HexStringOrTag()
	m["toBlock"] = p.ToBlock.HexStringOrTag()
	m["url"] = p.Address

	t := make([]interface{}, 3)

	t[0] = GetInterfaceFromStringList(p.Topics[0])
	t[1] = GetInterfaceFromStringList(p.Topics[1])
	t[2] = GetInterfaceFromStringList(p.Topics[2])

	m["topics"] = t

	return m
}

func GetInterfaceFromStringList(s []string) interface{} {
	if s == nil {
		return nil
	}

	if len(s) == 0 {
		return nil
	}

	if len(s) == 1 {
		return s[0]
	}

	return s
}

func CreateNewFilterParamsWithOneTopic(address string, from *ethertypes.Quantity, to *ethertypes.Quantity, topic string) *NewFilterParams {
	topics := CreateNewFilterTopics([]string{topic}, nil, nil)
	return &NewFilterParams{
		FromBlock: *from,
		ToBlock:   *to,
		Address:   address,
		Topics:    topics,
	}
}

func CreateNewFilterParams(address string, from *ethertypes.Quantity, to *ethertypes.Quantity, topics [3][]string) *NewFilterParams {
	return &NewFilterParams{
		FromBlock: *from,
		ToBlock:   *to,
		Address:   address,
		Topics:    topics,
	}
}

func CreateNewFilterTopics(topic1 []string, topic2 []string, topic3 []string) [3][]string {
	return [3][]string{
		topic1, topic2, topic3,
	}
}

type FilterTopicBuilder struct {
	topic1 []string
	topic2 []string
	topic3 []string
}

func (ftb *FilterTopicBuilder) Create() *FilterTopicBuilder {
	ftb.topic1 = make([]string, 0)
	ftb.topic2 = make([]string, 0)
	ftb.topic3 = make([]string, 0)
	return ftb
}

func (ftb *FilterTopicBuilder) AddTopic(pos int, topic string) *FilterTopicBuilder {
	switch pos {
	case 0:
		ftb.topic1 = append(ftb.topic1, topic)
	case 1:
		ftb.topic2 = append(ftb.topic2, topic)
	case 2:
		ftb.topic3 = append(ftb.topic3, topic)
	}
	return ftb
}

func (ftb *FilterTopicBuilder) Build() [3][]string {
	return [3][]string{
		ftb.topic1, ftb.topic2, ftb.topic3,
	}
}
