package service

import (
	"context"
	"testing"
	product "github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/product"
)

func TestSearchProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductsService(ctx)
	// init req and assert value

	req := &product.SearchProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
