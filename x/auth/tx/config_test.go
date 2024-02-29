package tx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	_ "cosmossdk.io/api/cosmos/crypto/secp256k1"
	"cosmossdk.io/x/auth/tx"
	txtestutil "cosmossdk.io/x/auth/tx/testutil"
	"cosmossdk.io/x/tx/signing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := testutil.CodecOptions{}.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	signingCtx := interfaceRegistry.SigningContext()
	suite.Run(t, txtestutil.NewTxConfigTestSuite(tx.NewTxConfig(protoCodec, signingCtx.AddressCodec(), signingCtx.ValidatorAddressCodec(), tx.DefaultSignModes)))
}

func TestConfigOptions(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	configOptions := tx.ConfigOptions{SigningOptions: &signing.Options{
		AddressCodec:          protoCodec.InterfaceRegistry().SigningContext().AddressCodec(),
		ValidatorAddressCodec: protoCodec.InterfaceRegistry().SigningContext().ValidatorAddressCodec(),
	}}
	txConfig, err := tx.NewTxConfigWithOptions(protoCodec, configOptions)
	require.NoError(t, err)
	require.NotNil(t, txConfig)
	handler := txConfig.SignModeHandler()
	require.NotNil(t, handler)
}
