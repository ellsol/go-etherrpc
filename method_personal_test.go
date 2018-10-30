package rpc

import (
	"testing"
	"github.com/ellsol/gox-ethertypes"
)

func TestPersonal_ListAccounts(t *testing.T) {
	accounts, err := NewRPCClient(config().url).Personal.ListAccounts()

	if err != nil {
		t.Error(err)
		return
	}

	expected1, _ := ethertypes.NewHexString().SetString("0x44a139cc0aed5eb5dbc6838b284fb051cad72dcb")
	expected2, _ := ethertypes.NewHexString().SetString("0xe96f31db85aa516b5a6ab2973d333f0406ddcb9b")

	err = ethertypes.CompareHexStringList([]ethertypes.HexString{*expected1, *expected2}, accounts)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestPersonal_NewAccount(t *testing.T) {
	// stop testing unless on private test net
	//test := "password"
	//_, err := NewRPCClient(config().url).Personal.NewAccount(test)
	//
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
}
