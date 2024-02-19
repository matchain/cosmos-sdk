package keeper

import (
	"context"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/params/types"
	"cosmossdk.io/x/params/types/proposal"

	"github.com/cosmos/cosmos-sdk/codec"
)

// Keeper of the global paramstore
type Keeper struct {
	cdc         codec.BinaryCodec
	environment appmodule.Environment
	legacyAmino *codec.LegacyAmino
	key         storetypes.StoreKey
	tkey        storetypes.StoreKey
	spaces      map[string]*types.Subspace
}

// NewKeeper constructs a params keeper
func NewKeeper(cdc codec.BinaryCodec, env appmodule.Environment, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) Keeper {
	return Keeper{
		cdc:         cdc,
		environment: env,
		legacyAmino: legacyAmino,
		key:         key,
		tkey:        tkey,
		spaces:      make(map[string]*types.Subspace),
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx context.Context) log.Logger {
	return k.environment.Logger.With("module", "x/"+proposal.ModuleName)
}

// Allocate subspace used for keepers
func (k Keeper) Subspace(s string) types.Subspace {
	_, ok := k.spaces[s]
	if ok {
		panic("subspace already occupied")
	}

	if s == "" {
		panic("cannot use empty string for subspace")
	}

	space := types.NewSubspace(k.cdc, k.environment, k.legacyAmino, k.key, k.tkey, s)
	k.spaces[s] = &space

	return space
}

// Get existing substore from keeper
func (k Keeper) GetSubspace(s string) (types.Subspace, bool) {
	space, ok := k.spaces[s]
	if !ok {
		return types.Subspace{}, false
	}
	return *space, ok
}

// GetSubspaces returns all the registered subspaces.
func (k Keeper) GetSubspaces() []types.Subspace {
	spaces := make([]types.Subspace, len(k.spaces))
	i := 0
	for _, ss := range k.spaces {
		spaces[i] = *ss
		i++
	}

	return spaces
}
