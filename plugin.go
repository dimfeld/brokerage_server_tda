package brokerage_server_tda

import (
	"context"

	"github.com/dimfeld/brokerage_server/types"
)

type Tda struct {
	debugLevel types.DebugLevel
}

func New() (*Tda, error) {
	return nil, types.ERR_PLUGIN_NOT_IMPLEMENTED
}

func (tda *Tda) Connect() error {
	return nil
}

func (tda *Tda) Close() error {
	// Nothing to do
	return nil
}

func (tda *Tda) Status() *types.ConnectionStatus {
	return &types.ConnectionStatus{
		Connected: true,
		Error:     nil,
	}
}

func (tda *Tda) Error() error {
	return nil
}

func (tda *Tda) SetDebugLevel(level types.DebugLevel) {
	tda.debugLevel = level
}

func (tda *Tda) GetHistoricalData(ctx context.Context, params types.HistoricalDataParams) ([]*types.Quote, error) {

	return nil, types.ERR_PLUGIN_NOT_IMPLEMENTED
}
