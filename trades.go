package brokerage_server_tda

import (
	"context"

	"github.com/dimfeld/brokerage_server/types"
)

func (tda *Tda) GetPositions(ctx context.Context) ([]*types.Position, error) {
	return nil, types.ERR_PLUGIN_NOT_IMPLEMENTED
}

func (tda *Tda) GetTrades(ctx context.Context) ([]*types.Trade, error) {
	return nil, types.ERR_PLUGIN_NOT_IMPLEMENTED
}
