package txmgr_test

import (
	"fmt"
	"math/big"
	"testing"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	gasmocks "github.com/smartcontractkit/chainlink-evm/pkg/gas/mocks"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys/keystest"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
)

func NewEvmAddress() gethcommon.Address {
	return testutils.NewAddress()
}

type feeConfig struct {
	eip1559DynamicFees bool
	tipCapMin          *assets.Wei
	priceMin           *assets.Wei
	priceMax           *assets.Wei
	limitDefault       uint64
}

func newFeeConfig() *feeConfig {
	return &feeConfig{
		tipCapMin: assets.NewWeiI(0),
		priceMin:  assets.NewWeiI(0),
		priceMax:  assets.NewWeiI(0),
	}
}

func (g *feeConfig) EIP1559DynamicFees() bool                        { return g.eip1559DynamicFees }
func (g *feeConfig) TipCapMin() *assets.Wei                          { return g.tipCapMin }
func (g *feeConfig) PriceMin() *assets.Wei                           { return g.priceMin }
func (g *feeConfig) PriceMaxKey(addr gethcommon.Address) *assets.Wei { return g.priceMax }
func (g *feeConfig) LimitDefault() uint64                            { return g.limitDefault }

func TestTxm_SignTx(t *testing.T) {
	t.Parallel()

	addr := gethcommon.HexToAddress("0xb921F7763960b296B9cbAD586ff066A18D749724")
	to := gethcommon.HexToAddress("0xb921F7763960b296B9cbAD586ff066A18D749724")
	tx := gethtypes.NewTx(&gethtypes.LegacyTx{
		Nonce:    42,
		To:       &to,
		Value:    big.NewInt(142),
		Gas:      242,
		GasPrice: big.NewInt(342),
		Data:     []byte{1, 2, 3},
	})

	t.Run("returns correct hash for non-okex chains", func(t *testing.T) {
		chainID := big.NewInt(1)
		kst := keystest.TxSigner(nil)
		cks := txmgr.NewEvmTxAttemptBuilder(*chainID, newFeeConfig(), kst, nil)
		hash, rawBytes, err := cks.SignTx(t.Context(), addr, tx)
		require.NoError(t, err)
		require.NotNil(t, rawBytes)
		require.Equal(t, "0xdd68f554373fdea7ec6713a6e437e7646465d553a6aa0b43233093366cc87ef0", hash.String())
	})
	// okex used to have a custom hash but now this just verifies that is it the same
	t.Run("returns correct hash for okex chains", func(t *testing.T) {
		chainID := big.NewInt(1)
		kst := keystest.TxSigner(nil)
		cks := txmgr.NewEvmTxAttemptBuilder(*chainID, newFeeConfig(), kst, nil)
		hash, rawBytes, err := cks.SignTx(t.Context(), addr, tx)
		require.NoError(t, err)
		require.NotNil(t, rawBytes)
		require.Equal(t, "0xdd68f554373fdea7ec6713a6e437e7646465d553a6aa0b43233093366cc87ef0", hash.String())
	})
	t.Run("can properly encoded and decode raw transaction for LegacyTx", func(t *testing.T) {
		chainID := big.NewInt(1)
		kst := keystest.TxSigner(nil)
		cks := txmgr.NewEvmTxAttemptBuilder(*chainID, newFeeConfig(), kst, nil)

		_, rawBytes, err := cks.SignTx(t.Context(), addr, tx)
		require.NoError(t, err)
		require.NotNil(t, rawBytes)
		require.Equal(t, "0xe42a82015681f294b921f7763960b296b9cbad586ff066a18d749724818e83010203808080", hexutil.Encode(rawBytes))

		var decodedTx *gethtypes.Transaction
		decodedTx, err = txmgr.GetGethSignedTx(rawBytes)
		require.NoError(t, err)
		require.Equal(t, tx.Hash(), decodedTx.Hash())
	})
	t.Run("can properly encoded and decode raw transaction for DynamicFeeTx", func(t *testing.T) {
		chainID := big.NewInt(1)
		kst := keystest.TxSigner(nil)
		typedTx := gethtypes.NewTx(&gethtypes.DynamicFeeTx{
			Nonce: 42,
			To:    &to,
			Value: big.NewInt(142),
			Gas:   242,
			Data:  []byte{1, 2, 3},
		})
		cks := txmgr.NewEvmTxAttemptBuilder(*chainID, newFeeConfig(), kst, nil)
		_, rawBytes, err := cks.SignTx(t.Context(), addr, typedTx)
		require.NoError(t, err)
		require.NotNil(t, rawBytes)
		require.Equal(t, "0xa702e5802a808081f294b921f7763960b296b9cbad586ff066a18d749724818e83010203c0808080", hexutil.Encode(rawBytes))

		var decodedTx *gethtypes.Transaction
		decodedTx, err = txmgr.GetGethSignedTx(rawBytes)
		require.NoError(t, err)
		require.Equal(t, typedTx.Hash(), decodedTx.Hash())
	})
}

func TestTxm_NewDynamicFeeTx(t *testing.T) {
	addr := NewEvmAddress()
	kst := keystest.TxSigner(nil)
	var n evmtypes.Nonce
	lggr := logger.Test(t)

	t.Run("creates attempt with fields", func(t *testing.T) {
		feeCfg := newFeeConfig()
		feeCfg.priceMax = assets.GWei(200)
		cks := txmgr.NewEvmTxAttemptBuilder(*big.NewInt(1), feeCfg, kst, nil)
		dynamicFee := gas.DynamicFee{GasTipCap: assets.GWei(100), GasFeeCap: assets.GWei(200)}
		a, _, err := cks.NewCustomTxAttempt(t.Context(), txmgr.Tx{Sequence: &n, FromAddress: addr}, gas.EvmFee{
			DynamicFee: gas.DynamicFee{GasTipCap: dynamicFee.GasTipCap, GasFeeCap: dynamicFee.GasFeeCap},
		}, 100, 0x2, lggr)
		require.NoError(t, err)
		assert.Equal(t, 100, int(a.ChainSpecificFeeLimit))
		assert.Nil(t, a.TxFee.GasPrice)
		assert.NotNil(t, a.TxFee.GasTipCap)
		assert.Equal(t, assets.GWei(100).String(), a.TxFee.GasTipCap.String())
		assert.NotNil(t, a.TxFee.GasFeeCap)
		assert.Equal(t, assets.GWei(200).String(), a.TxFee.GasFeeCap.String())
	})

	t.Run("verifies gas tip and fees", func(t *testing.T) {
		cases := []struct {
			name        string
			tipcap      *assets.Wei
			feecap      *assets.Wei
			setCfg      func(c *toml.EVMConfig)
			expectError string
		}{
			{"gas tip = fee cap", assets.GWei(5), assets.GWei(5), nil, ""},
			{"gas tip < fee cap", assets.GWei(4), assets.GWei(5), nil, ""},
			{"gas tip > fee cap", assets.GWei(6), assets.GWei(5), nil, "gas fee cap must be greater than or equal to gas tip cap (fee cap: 5 gwei, tip cap: 6 gwei)"},
			{"fee cap exceeds max allowed", assets.GWei(5), assets.GWei(5), func(c *toml.EVMConfig) {
				c.GasEstimator.PriceMax = assets.GWei(4)
			}, "specified gas fee cap of 5 gwei would exceed max configured gas price of 4 gwei"},
			{"ignores global min gas price", assets.GWei(5), assets.GWei(5), func(c *toml.EVMConfig) {
				c.GasEstimator.PriceMin = assets.GWei(6)
			}, ""},
		}

		for _, tt := range cases {
			test := tt
			t.Run(test.name, func(t *testing.T) {
				cfg := testutils.NewTestChainScopedConfig(t, test.setCfg)
				cks := txmgr.NewEvmTxAttemptBuilder(*big.NewInt(1), cfg.EVM().GasEstimator(), kst, nil)
				dynamicFee := gas.DynamicFee{GasTipCap: test.tipcap, GasFeeCap: test.feecap}
				_, _, err := cks.NewCustomTxAttempt(t.Context(), txmgr.Tx{Sequence: &n, FromAddress: addr}, gas.EvmFee{
					DynamicFee: gas.DynamicFee{GasTipCap: dynamicFee.GasTipCap, GasFeeCap: dynamicFee.GasFeeCap},
				}, 100, 0x2, lggr)
				if test.expectError == "" {
					require.NoError(t, err)
				} else {
					require.ErrorContains(t, err, test.expectError)
				}
			})
		}
	})
}

func TestTxm_NewLegacyAttempt(t *testing.T) {
	addr := NewEvmAddress()
	kst := keystest.TxSigner(nil)
	gc := newFeeConfig()
	gc.priceMin = assets.NewWeiI(10)
	gc.priceMax = assets.NewWeiI(50)
	cks := txmgr.NewEvmTxAttemptBuilder(*big.NewInt(1), gc, kst, nil)
	lggr := logger.Test(t)

	t.Run("creates attempt with fields", func(t *testing.T) {
		var n evmtypes.Nonce
		a, _, err := cks.NewCustomTxAttempt(t.Context(), txmgr.Tx{Sequence: &n, FromAddress: addr}, gas.EvmFee{GasPrice: assets.NewWeiI(25)}, 100, 0x0, lggr)
		require.NoError(t, err)
		assert.Equal(t, 100, int(a.ChainSpecificFeeLimit))
		assert.NotNil(t, a.TxFee.GasPrice)
		assert.Equal(t, "25 wei", a.TxFee.GasPrice.String())
		assert.Nil(t, a.TxFee.GasTipCap)
		assert.Nil(t, a.TxFee.GasFeeCap)
	})

	t.Run("verifies max gas price", func(t *testing.T) {
		_, _, err := cks.NewCustomTxAttempt(t.Context(), txmgr.Tx{FromAddress: addr}, gas.EvmFee{GasPrice: assets.NewWeiI(100)}, 100, 0x0, lggr)
		require.Error(t, err)
		assert.Contains(t, err.Error(), fmt.Sprintf("specified gas price of 100 wei would exceed max configured gas price of 50 wei for key %s", addr.String()))
	})
}

func TestTxm_NewPurgeAttempt(t *testing.T) {
	addr := NewEvmAddress()
	kst := keystest.TxSigner(nil)
	gc := newFeeConfig()
	gc.priceMin = assets.GWei(10)
	gc.priceMax = assets.GWei(50)
	gc.limitDefault = uint64(10)
	est := gasmocks.NewEvmFeeEstimator(t)
	bumpedLegacy := assets.GWei(30)
	bumpedDynamicFee := assets.GWei(15)
	bumpedDynamicTip := assets.GWei(10)
	bumpedFee := gas.EvmFee{GasPrice: bumpedLegacy, DynamicFee: gas.DynamicFee{GasTipCap: bumpedDynamicTip, GasFeeCap: bumpedDynamicFee}}
	est.On("BumpFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(bumpedFee, uint64(10_000), nil)
	cks := txmgr.NewEvmTxAttemptBuilder(*big.NewInt(1), gc, kst, est)
	lggr := logger.Test(t)

	t.Run("creates legacy purge attempt with fields if previous attempt is legacy", func(t *testing.T) {
		n := evmtypes.Nonce(0)
		etx := txmgr.Tx{Sequence: &n, FromAddress: addr, EncodedPayload: []byte{1, 2, 3}}
		prevAttempt, _, err := cks.NewCustomTxAttempt(t.Context(), etx, gas.EvmFee{GasPrice: bumpedLegacy.Sub(assets.GWei(1))}, 100, 0x0, lggr)
		require.NoError(t, err)
		etx.TxAttempts = append(etx.TxAttempts, prevAttempt)
		a, err := cks.NewPurgeTxAttempt(t.Context(), etx, lggr)
		require.NoError(t, err)
		// The fee limit is overridden with LimitDefault since purge attempts are just empty attempts
		require.Equal(t, gc.limitDefault, a.ChainSpecificFeeLimit)
		require.NotNil(t, a.TxFee.GasPrice)
		require.Equal(t, bumpedLegacy.String(), a.TxFee.GasPrice.String())
		require.Nil(t, a.TxFee.GasTipCap)
		require.Nil(t, a.TxFee.GasFeeCap)
		require.Equal(t, true, a.IsPurgeAttempt)
		require.Equal(t, []byte{}, a.Tx.EncodedPayload)
		require.Equal(t, *big.NewInt(0), a.Tx.Value)
	})

	t.Run("creates dynamic purge attempt with fields if previous attempt is dynamic", func(t *testing.T) {
		n := evmtypes.Nonce(0)
		etx := txmgr.Tx{Sequence: &n, FromAddress: addr, EncodedPayload: []byte{1, 2, 3}}
		prevAttempt, _, err := cks.NewCustomTxAttempt(t.Context(), etx, gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: bumpedDynamicTip.Sub(assets.GWei(1)), GasFeeCap: bumpedDynamicFee.Sub(assets.GWei(1))}}, 100, 0x2, lggr)
		require.NoError(t, err)
		etx.TxAttempts = append(etx.TxAttempts, prevAttempt)
		a, err := cks.NewPurgeTxAttempt(t.Context(), etx, lggr)
		require.NoError(t, err)
		// The fee limit is overridden with LimitDefault since purge attempts are just empty attempts
		require.Equal(t, gc.limitDefault, a.ChainSpecificFeeLimit)
		require.Nil(t, a.TxFee.GasPrice)
		require.NotNil(t, a.TxFee.GasTipCap)
		require.NotNil(t, a.TxFee.GasFeeCap)
		require.Equal(t, bumpedDynamicTip.String(), a.TxFee.GasTipCap.String())
		require.Equal(t, bumpedDynamicFee.String(), a.TxFee.GasFeeCap.String())
		require.Equal(t, true, a.IsPurgeAttempt)
		require.Equal(t, []byte{}, a.Tx.EncodedPayload)
		require.Equal(t, *big.NewInt(0), a.Tx.Value)
	})

	t.Run("creates bump purge attempt with fields", func(t *testing.T) {
		n := evmtypes.Nonce(0)
		etx := txmgr.Tx{Sequence: &n, FromAddress: addr, EncodedPayload: []byte{1, 2, 3}}
		prevAttempt, _, err := cks.NewCustomTxAttempt(t.Context(), etx, gas.EvmFee{GasPrice: bumpedLegacy.Sub(assets.GWei(1))}, 100, 0x0, lggr)
		require.NoError(t, err)
		etx.TxAttempts = append(etx.TxAttempts, prevAttempt)
		purgeAttempt, err := cks.NewPurgeTxAttempt(t.Context(), etx, lggr)
		require.NoError(t, err)
		etx.TxAttempts = append(etx.TxAttempts, purgeAttempt)
		bumpAttempt, _, _, _, err := cks.NewBumpTxAttempt(t.Context(), etx, purgeAttempt, etx.TxAttempts, lggr)
		require.NoError(t, err)
		// The fee limit is overridden with LimitDefault since purge attempts are just empty attempts
		require.Equal(t, gc.limitDefault, bumpAttempt.ChainSpecificFeeLimit)
		require.NotNil(t, bumpAttempt.TxFee.GasPrice)
		require.Equal(t, bumpedLegacy.String(), bumpAttempt.TxFee.GasPrice.String())
		require.Nil(t, bumpAttempt.TxFee.GasTipCap)
		require.Nil(t, bumpAttempt.TxFee.GasFeeCap)
		require.Equal(t, true, bumpAttempt.IsPurgeAttempt)
		require.Equal(t, []byte{}, bumpAttempt.Tx.EncodedPayload)
		require.Equal(t, *big.NewInt(0), bumpAttempt.Tx.Value)
	})
}

func TestTxm_NewCustomTxAttempt_NonRetryableErrors(t *testing.T) {
	t.Parallel()

	kst := keystest.TxSigner(nil)
	lggr := logger.Test(t)
	cks := txmgr.NewEvmTxAttemptBuilder(*big.NewInt(1), newFeeConfig(), kst, nil)

	dynamicFee := gas.DynamicFee{GasTipCap: assets.GWei(100), GasFeeCap: assets.GWei(200)}
	legacyFee := assets.NewWeiI(100)

	t.Run("dynamic fee with legacy tx type", func(t *testing.T) {
		_, retryable, err := cks.NewCustomTxAttempt(t.Context(), txmgr.Tx{}, gas.EvmFee{
			DynamicFee: dynamicFee,
		}, 100, 0x0, lggr)
		require.Error(t, err)
		assert.False(t, retryable)
	})
	t.Run("legacy fee with dynamic tx type", func(t *testing.T) {
		_, retryable, err := cks.NewCustomTxAttempt(t.Context(), txmgr.Tx{}, gas.EvmFee{GasPrice: legacyFee}, 100, 0x2, lggr)
		require.Error(t, err)
		assert.False(t, retryable)
	})

	t.Run("invalid type", func(t *testing.T) {
		_, retryable, err := cks.NewCustomTxAttempt(t.Context(), txmgr.Tx{}, gas.EvmFee{}, 100, 0xA, lggr)
		require.Error(t, err)
		assert.False(t, retryable)
	})
}

func TestTxm_EvmTxAttemptBuilder_RetryableEstimatorError(t *testing.T) {
	est := gasmocks.NewEvmFeeEstimator(t)
	est.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(gas.EvmFee{}, uint64(0), pkgerrors.New("fail"))
	est.On("BumpFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(gas.EvmFee{}, uint64(0), pkgerrors.New("fail"))

	kst := keystest.TxSigner(nil)
	lggr := logger.Test(t)
	ctx := t.Context()
	cks := txmgr.NewEvmTxAttemptBuilder(*big.NewInt(1), &feeConfig{eip1559DynamicFees: true}, kst, est)

	t.Run("NewAttempt", func(t *testing.T) {
		_, _, _, retryable, err := cks.NewTxAttempt(ctx, txmgr.Tx{}, lggr)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get fee")
		assert.True(t, retryable)
	})
	t.Run("NewAttemptWithType", func(t *testing.T) {
		_, _, _, retryable, err := cks.NewTxAttemptWithType(ctx, txmgr.Tx{}, lggr, 0x0)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get fee")
		assert.True(t, retryable)
	})
	t.Run("NewBumpAttempt", func(t *testing.T) {
		_, _, _, retryable, err := cks.NewBumpTxAttempt(ctx, txmgr.Tx{}, txmgr.TxAttempt{}, nil, lggr)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to bump fee")
		assert.True(t, retryable)
	})
}
