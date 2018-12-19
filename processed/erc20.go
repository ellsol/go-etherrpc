package processed

import (
	"fmt"
	"github.com/ellsol/go-etherrpc"
	"github.com/ellsol/go-ethertypes"
)

const (
	IconomiTokenAddress = "0x888666CA69E0f178DED6D75b5726Cee99A87D698"
	ERC20TransferTopic  = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

type ERC20Transfer struct {
	TransactionHash *ethertypes.HexString    `json:"transactionHash"`
	BlockNumber     int64                    `json:"blockNumber"`
	Date            int64                    `json:"date"`
	From            *ethertypes.EtherAddress `json:"from"`
	To              *ethertypes.EtherAddress `json:"to"`
	TokenValue      *ethertypes.EtherValue   `json:"tokenValue"`
}

func (erc *ERC20Transfer) Log() string {
	return fmt.Sprintf("(%v) %v -> %v [%v]", erc.BlockNumber, erc.From.ShortFormat(), erc.To.ShortFormat(), erc.TokenValue)
}

func (erc *ERC20Transfer) FromReceipt(transReceipt *ethertypes.EtherTransactionWithReceipt, date int64) (*ERC20Transfer, error) {
	var err error
	erc.BlockNumber = transReceipt.BlockNumber
	erc.TransactionHash = &transReceipt.Hash
	erc.Date = date
	logs := transReceipt.Logs

	for _, v := range logs {
		l := v
		if len(l.Topics) != 3 || l.Topics[0].String() != ERC20TransferTopic {
			continue
		}

		erc.From, err = new(ethertypes.EtherAddress).Set32ByteString(l.Topics[1].String())
		if err != nil {
			return nil, err
		}
		erc.To, err = new(ethertypes.EtherAddress).Set32ByteString(l.Topics[2].String())
		if err != nil {
			return nil, err
		}
		erc.TokenValue, err = new(ethertypes.EtherValue).SetHexString(l.Data.String())
		if err != nil {
			return nil, err
		}
		return erc, nil
	}

	return nil, fmt.Errorf("cannot create erc20 token, log of %v/%v doesnt contain Transfer information", transReceipt.Hash, transReceipt.BlockNumber)
}

func (erc *ERC20Transfer) FromEtherLog(log *ethertypes.EtherLog) (*ERC20Transfer, error) {
	var err error
	erc.BlockNumber = log.BlockNumber
	erc.TransactionHash = &log.TransactionHash
	erc.Date = 0

	erc.From, err = new(ethertypes.EtherAddress).Set32ByteString(log.Topics[1].String())

	if err != nil {
		return nil, err
	}
	erc.To, err = new(ethertypes.EtherAddress).Set32ByteString(log.Topics[2].String())

	if err != nil {
		return nil, err
	}

	erc.TokenValue, err = new(ethertypes.EtherValue).SetHexString(log.Data.String())

	if err != nil {
		return nil, err
	}

	return erc, nil
}

func ERC20BalanceOf(tokenAddress string, toAddress string, eth rpc.Eth) (*ethertypes.EtherValue, error) {
	to, err := ethertypes.NewHexString().SetString(toAddress)

	if err != nil {
		return nil, err
	}

	params, err := new(rpc.EthCallParams).ToContractWithArgument(tokenAddress, "balanceOf(address)", to.Bytes())

	if err != nil {
		return nil, err
	}

	hex, err := eth.Call(params, ethertypes.QuantityLatest())

	if err != nil {
		return nil, err
	}

	return ethertypes.NewEtherValue().SetHexString(hex.String())
}

func GetERC20BalanceOf(tokenAddress string, toAddress string, eth rpc.Eth) (*ethertypes.EtherValue, error) {
	to, err := ethertypes.NewHexString().SetString(toAddress)

	if err != nil {
		return nil, err
	}

	params, err := new(rpc.EthCallParams).ToContractWithArgument(tokenAddress, "getBalanceOf(address)", to.Bytes())

	if err != nil {
		return nil, err
	}

	hex, err := eth.Call(params, ethertypes.QuantityLatest())

	if err != nil {
		return nil, err
	}

	return ethertypes.NewEtherValue().SetHexString(hex.String())
}

func GetERC20BalanceOfWithQuantity(tokenAddress string, toAddress string, quantity *ethertypes.Quantity, eth rpc.Eth) (*ethertypes.EtherValue, error) {
	to, err := ethertypes.NewHexString().SetString(toAddress)

	if err != nil {
		return nil, err
	}

	params, err := new(rpc.EthCallParams).ToContractWithArgument(tokenAddress, "getBalanceOf(address)", to.Bytes())

	if err != nil {
		return nil, err
	}

	hex, err := eth.Call(params, quantity)

	if err != nil {
		return nil, err
	}

	return ethertypes.NewEtherValue().SetHexString(hex.String())
}
