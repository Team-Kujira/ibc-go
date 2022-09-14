package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
)

// GetSendEnabled retrieves the send enabled boolean from the paramstore
func (k Keeper) GetSendEnabled(ctx sdk.Context) bool {
	var res bool
	k.paramSpace.Get(ctx, types.KeySendEnabled, &res)
	return res
}

// GetReceiveEnabled retrieves the receive enabled boolean from the paramstore
func (k Keeper) GetReceiveEnabled(ctx sdk.Context) bool {
	var res bool
	k.paramSpace.Get(ctx, types.KeyReceiveEnabled, &res)
	return res
}

// GetSlashPrefix retrieves the slash prefix from the paramstore
func (k Keeper) GetSlashPrefix(ctx sdk.Context) string {
	var res string
	k.paramSpace.Get(ctx, types.KeySlashPrefix, &res)
	return res
}

// GetParams returns the total set of ibc-transfer parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(k.GetSendEnabled(ctx), k.GetReceiveEnabled(ctx), k.GetSlashPrefix(ctx))
}

// SetParams sets the total set of ibc-transfer parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
