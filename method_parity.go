package rpc

const (
	MethodEnode    = "parity_enode"
	MethodNodeName = "parity_nodeName"
)

type Parity struct {
	client *Client
}

// Node Settings

/*
	rpc method: "parity_enode"
	curl --data '{"method":"parity_enode","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
	Returns the node enode URI.
 */
func (it *Parity) Enode() (string, error) {
	return it.client.RequestString(MethodEnode)
}

//parity_mode
//parity_nodeKind
/*
	rpc method: "parity_nodeName"
	curl --data '{"method":"parity_nodeName","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
	Returns node name, set when starting parity with --identity NAME.
 */
func (it *Parity) NodeName() (string, error) {
	return it.client.RequestString(MethodNodeName)
}
//parity_wsUrl
