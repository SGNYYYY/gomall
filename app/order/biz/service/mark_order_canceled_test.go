package service

import (
	"context"
	"testing"
	order "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
)

func TestMarkOrderCanceled_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMarkOrderCanceledService(ctx)
	// init req and assert value

	req := &order.MarkOrderCanceledReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
