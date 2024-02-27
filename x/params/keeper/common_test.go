package keeper_test

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/params"
	paramskeeper "cosmossdk.io/x/params/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
)

func testComponents() (*codec.LegacyAmino, sdk.Context, storetypes.StoreKey, storetypes.StoreKey, paramskeeper.Keeper) {
	encodingConfig := moduletestutil.MakeTestEncodingConfig(params.AppModule{})
	cdc := encodingConfig.Codec

	legacyAmino := createTestCodec()
	mkey := storetypes.NewKVStoreKey("test")
	env := runtime.NewEnvironment(runtime.NewKVStoreService(mkey), log.NewNopLogger())
	tkey := storetypes.NewTransientStoreKey("transient_test")
	ctx := sdktestutil.DefaultContext(mkey, tkey)
	keeper := paramskeeper.NewKeeper(cdc, env, legacyAmino, mkey, tkey)

	return legacyAmino, ctx, mkey, tkey, keeper
}

type invalid struct{}

type s struct {
	I int
}

func createTestCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cdc.RegisterConcrete(s{}, "test/s", nil)
	cdc.RegisterConcrete(invalid{}, "test/invalid", nil)
	return cdc
}
