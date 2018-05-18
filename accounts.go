package brokerage_server_tda

import (
	"context"

	"github.com/dimfeld/brokerage_server/types"
)

func (tda *Tda) AccountList(ctx context.Context) ([]*types.Account, error) {
	return nil, types.ERR_PLUGIN_NOT_IMPLEMENTED
}
