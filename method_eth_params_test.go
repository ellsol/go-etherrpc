package rpc

import (
	"testing"
	"github.com/ellsol/gox-ethertypes"
	"log"
)

func TestNewFilterParams_ToMap(t *testing.T) {
	params := CreateNewFilterParams("url", ethertypes.QuantityLatest(), ethertypes.QuantityLatest(), CreateNewFilterTopics([]string{"t11", "t12"}, []string{"t21"}, []string{}))
//	t.Errorf("%+v", params.ToMap())
	log.Println(params)
}
