package brokerage_server_tda

import (
	"context"

	"github.com/dimfeld/brokerage_server/types"
)

func (tda *Tda) GetOptionsChain(ctx context.Context, symbol string) (types.OptionChain, error) {
	return types.OptionChain{}, types.ERR_PLUGIN_NOT_IMPLEMENTED
}

func (tda *Tda) GetOptionsQuotes(ctx context.Context, params types.OptionsQuoteParams) ([]*types.OptionQuote, error) {
	return nil, types.ERR_PLUGIN_NOT_IMPLEMENTED
}
