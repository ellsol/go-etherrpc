package processed

import (
	"github.com/ellsol/go-etherrpc"
	"testing"
)

const InfuraEndpoint = "https://mainnet.infura.io/3l5dxBOP3wPspnRDdG1u"
const RPCEndpointLocalHost = "http://localhost:8545"

type TestConfig struct {
	address string
}

func config() *TestConfig {
	return LocalhostConfig()
}

func InfuraConfig() *TestConfig {
	return &TestConfig{
		address: InfuraEndpoint,
	}
}

func LocalhostConfig() *TestConfig {
	return &TestConfig{
		address: RPCEndpointLocalHost,
	}
}

func TestParseTransactionsFromChainByFrom(t *testing.T) {
	err := ParseTransactionsFromChainByTo(IconomiTokenAddress, 2400000, 25000001, &rpc.NewRPCClient(config().address).Eth)

	if err != nil {
		t.Error(err)
		return
	}
	t.Error("")

}
