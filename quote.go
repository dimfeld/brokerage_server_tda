package brokerage_server_tda

import (
	"context"

	"github.com/dimfeld/brokerage_server/types"
)

func (tda *Tda) GetStockQuote(ctx context.Context, symbol string) (*types.Quote, error) {
	return nil, types.ERR_PLUGIN_NOT_IMPLEMENTED
}
