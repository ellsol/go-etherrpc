package processed

import (
	"github.com/ellsol/go-etherrpc"
	"github.com/ellsol/go-ethertypes"
)

type ERC20TransfersParam struct {
	Address   string
	FromBlock *ethertypes.Quantity
	ToBlock   *ethertypes.Quantity
	Sender    string
	Receiver  string
}

func RequestERC20TransferByPair(address string, from int64, to int64, sender string, receiver string, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   ethertypes.QuantityBlockInt64(to),
		FromBlock: ethertypes.QuantityBlockInt64(from),
		Receiver:  receiver,
		Sender:    sender,
	}, client)
}

func RequestERC20TransfersByReceiver(address string, from int64, to int64, receiver string, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   ethertypes.QuantityBlockInt64(to),
		FromBlock: ethertypes.QuantityBlockInt64(from),
		Receiver:  receiver,
	}, client)
}

func RequestERC20TransfersBySender(address string, from int64, to int64, sender string, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   ethertypes.QuantityBlockInt64(to),
		FromBlock: ethertypes.QuantityBlockInt64(from),
		Sender:    sender,
	}, client)
}

func RequestERC20TransfersDefault(address string, from int64, to int64, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   ethertypes.QuantityBlockInt64(to),
		FromBlock: ethertypes.QuantityBlockInt64(from),
	}, client)
}

func RequestERC20TransfersByBlock(address string, block int64, client *rpc.Client) ([]ERC20Transfer, error) {
	return RequestERC20Transfers(&ERC20TransfersParam{
		Address:   address,
		ToBlock:   ethertypes.QuantityBlockInt64(block),
		FromBlock: ethertypes.QuantityBlockInt64(block),
	}, client)
}

func RequestERC20Transfers(p *ERC20TransfersParam, client *rpc.Client) ([]ERC20Transfer, error) {

	ftb := new(rpc.FilterTopicBuilder).AddTopic(0, ERC20TransferTopic)

	if p.Sender != "" {
		ftb.AddTopic(1, p.Sender)
	}

	if p.Receiver != "" {
		ftb.AddTopic(2, p.Receiver)
	}

	logParam := rpc.CreateNewFilterParams(p.Address, p.FromBlock, p.ToBlock, ftb.Build())
	logs, err := client.Eth.GetLogs(logParam)

	if err != nil {
		return nil, err
	}

	result := make([]ERC20Transfer, 0)

	for _, v := range logs {
		t, err := new(ERC20Transfer).FromEtherLog(&v)

		if err != nil {
			return nil, err
		}

		result = append(result, *t)
	}

	return result, nil
}
