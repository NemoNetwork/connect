package cryptodotcom_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	providertypes "github.com/skip-mev/slinky/providers/types"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/skip-mev/slinky/oracle/constants"
	"github.com/skip-mev/slinky/oracle/types"
	"github.com/skip-mev/slinky/providers/base/websocket/handlers"
	"github.com/skip-mev/slinky/providers/websockets/cryptodotcom"
)

var (
	btc_usd = cryptodotcom.DefaultMarketConfig.MustGetProviderTicker(constants.BITCOIN_USD)
	eth_usd = cryptodotcom.DefaultMarketConfig.MustGetProviderTicker(constants.ETHEREUM_USD)
	sol_usd = cryptodotcom.DefaultMarketConfig.MustGetProviderTicker(constants.SOLANA_USD)
	logger  = zap.NewExample()
)

func TestHandleMessage(t *testing.T) {
	testCases := []struct {
		name         string
		msg          func() []byte
		resp         types.PriceResponse
		expUpdateMsg func() []handlers.WebsocketEncodedMessage
		expErr       bool
	}{
		{
			name: "cannot unmarshal to base message",
			msg: func() []byte {
				return []byte(`no rizz message`)
			},
			resp:         types.PriceResponse{},
			expUpdateMsg: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:       true,
		},
		{
			name: "unknown method type",
			msg: func() []byte {
				msg := cryptodotcom.InstrumentResponseMessage{
					Method: "unknown",
				}
				bz, err := json.Marshal(msg)
				require.NoError(t, err)
				return bz
			},
			resp:         types.PriceResponse{},
			expUpdateMsg: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:       true,
		},
		{
			name: "unknown status code",
			msg: func() []byte {
				msg := cryptodotcom.InstrumentResponseMessage{
					Method: string(cryptodotcom.InstrumentMethod),
					Code:   1,
				}
				bz, err := json.Marshal(msg)
				require.NoError(t, err)
				return bz
			},
			resp:         types.PriceResponse{},
			expUpdateMsg: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:       true,
		},
		{
			name: "heartbeat",
			msg: func() []byte {
				msg := cryptodotcom.InstrumentResponseMessage{
					ID:     42069,
					Method: string(cryptodotcom.HeartBeatRequestMethod),
					Code:   0,
				}
				bz, err := json.Marshal(msg)
				require.NoError(t, err)
				return bz
			},
			resp: types.PriceResponse{},
			expUpdateMsg: func() []handlers.WebsocketEncodedMessage {
				msg := cryptodotcom.HeartBeatResponseMessage{
					ID:     42069,
					Method: string(cryptodotcom.HeartBeatResponseMethod),
				}
				bz, err := json.Marshal(msg)
				require.NoError(t, err)
				return []handlers.WebsocketEncodedMessage{bz}
			},
			expErr: false,
		},
		{
			name: "instrument response with no data",
			msg: func() []byte {
				msg := cryptodotcom.InstrumentResponseMessage{
					ID:     42069,
					Method: string(cryptodotcom.InstrumentMethod),
					Code:   0,
					Result: cryptodotcom.InstrumentResult{
						Data: []cryptodotcom.InstrumentData{},
					},
				}
				bz, err := json.Marshal(msg)
				require.NoError(t, err)
				return bz
			},
			resp: types.PriceResponse{},
			expUpdateMsg: func() []handlers.WebsocketEncodedMessage {
				return nil
			},
			expErr: true,
		},
		{
			name: "instrument response with one instrument",
			msg: func() []byte {
				msg := cryptodotcom.InstrumentResponseMessage{
					ID:     42069,
					Method: string(cryptodotcom.InstrumentMethod),
					Code:   0,
					Result: cryptodotcom.InstrumentResult{
						Data: []cryptodotcom.InstrumentData{
							{
								Name:             "BTCUSD-PERP",
								LatestTradePrice: "42069",
							},
						},
					},
				}
				bz, err := json.Marshal(msg)
				require.NoError(t, err)
				return bz
			},
			resp: types.PriceResponse{
				Resolved: types.ResolvedPrices{
					btc_usd: types.NewPriceResult(big.NewFloat(42069e18), time.Now()),
				},
				UnResolved: types.UnResolvedPrices{},
			},
			expUpdateMsg: func() []handlers.WebsocketEncodedMessage {
				return nil
			},
			expErr: false,
		},
		{
			name: "unknown instrument",
			msg: func() []byte {
				msg := cryptodotcom.InstrumentResponseMessage{
					ID:     42069,
					Method: string(cryptodotcom.InstrumentMethod),
					Code:   0,
					Result: cryptodotcom.InstrumentResult{
						Data: []cryptodotcom.InstrumentData{
							{
								Name:             "MOGUSD-PERP",
								LatestTradePrice: "42069",
							},
						},
					},
				}
				bz, err := json.Marshal(msg)
				require.NoError(t, err)
				return bz
			},
			resp: types.PriceResponse{},
			expUpdateMsg: func() []handlers.WebsocketEncodedMessage {
				return nil
			},
			expErr: false,
		},
		{
			name: "instrument response with multiple instruments",
			msg: func() []byte {
				msg := cryptodotcom.InstrumentResponseMessage{
					ID:     42069,
					Method: string(cryptodotcom.InstrumentMethod),
					Code:   0,
					Result: cryptodotcom.InstrumentResult{
						Data: []cryptodotcom.InstrumentData{
							{
								Name:             "BTCUSD-PERP",
								LatestTradePrice: "42069",
							},
							{
								Name:             "ETHUSD-PERP",
								LatestTradePrice: "2000",
							},
							{
								Name:             "SOLUSD-PERP",
								LatestTradePrice: "1000",
							},
						},
					},
				}
				bz, err := json.Marshal(msg)
				require.NoError(t, err)
				return bz
			},
			resp: types.PriceResponse{
				Resolved: types.ResolvedPrices{
					btc_usd: types.NewPriceResult(big.NewFloat(42069e18), time.Now()),
					eth_usd: types.NewPriceResult(big.NewFloat(2000e18), time.Now()),
					sol_usd: types.NewPriceResult(big.NewFloat(1000e18), time.Now()),
				},
				UnResolved: types.UnResolvedPrices{},
			},
			expUpdateMsg: func() []handlers.WebsocketEncodedMessage {
				return nil
			},
			expErr: false,
		},
		{
			name: "instrument response with one instrument and one bad price instrument",
			msg: func() []byte {
				msg := cryptodotcom.InstrumentResponseMessage{
					ID:     42069,
					Method: string(cryptodotcom.InstrumentMethod),
					Code:   0,
					Result: cryptodotcom.InstrumentResult{
						Data: []cryptodotcom.InstrumentData{
							{
								Name:             "BTCUSD-PERP",
								LatestTradePrice: "42069",
							},
							{
								Name:             "SOLUSD-PERP",
								LatestTradePrice: "$42,069.00",
							},
						},
					},
				}
				bz, err := json.Marshal(msg)
				require.NoError(t, err)
				return bz
			},
			resp: types.PriceResponse{
				Resolved: types.ResolvedPrices{
					btc_usd: types.NewPriceResult(big.NewFloat(42069e18), time.Now()),
				},
				UnResolved: types.UnResolvedPrices{
					sol_usd: providertypes.UnresolvedResult{
						ErrorWithCode: providertypes.NewErrorWithCode(fmt.Errorf("failed to parse price $42,069.00: invalid syntax"), providertypes.ErrorWebSocketGeneral),
					},
				},
			},
			expUpdateMsg: func() []handlers.WebsocketEncodedMessage {
				return nil
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wsHandler, err := cryptodotcom.NewWebSocketDataHandler(logger, cryptodotcom.DefaultWebSocketConfig)
			require.NoError(t, err)

			// Update the cache since it is assumed that CreateMessages is executed before anything else.
			_, err = wsHandler.CreateMessages([]types.ProviderTicker{btc_usd, eth_usd, sol_usd})
			require.NoError(t, err)

			resp, updateMsg, err := wsHandler.HandleMessage(tc.msg())
			if tc.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expUpdateMsg(), updateMsg)

			require.Equal(t, len(tc.resp.Resolved), len(resp.Resolved))
			require.Equal(t, len(tc.resp.UnResolved), len(resp.UnResolved))

			for cp, result := range tc.resp.Resolved {
				require.Contains(t, resp.Resolved, cp)
				require.Equal(t, result.Value.SetPrec(types.DefaultTickerDecimals), resp.Resolved[cp].Value.SetPrec(types.DefaultTickerDecimals))
			}

			for cp := range tc.resp.UnResolved {
				require.Contains(t, resp.UnResolved, cp)
				require.Error(t, resp.UnResolved[cp])
			}
		})
	}
}

func TestCreateMessage(t *testing.T) {
	testCases := []struct {
		name        string
		cps         []types.ProviderTicker
		msg         cryptodotcom.InstrumentRequestMessage
		expectedErr bool
	}{
		{
			name:        "no currency pairs",
			cps:         []types.ProviderTicker{},
			msg:         cryptodotcom.InstrumentRequestMessage{},
			expectedErr: true,
		},
		{
			name: "one currency pair",
			cps: []types.ProviderTicker{
				btc_usd,
			},
			msg: cryptodotcom.InstrumentRequestMessage{
				Method: "subscribe",
				Params: cryptodotcom.InstrumentParams{
					Channels: []string{"ticker.BTCUSD-PERP"},
				},
			},
			expectedErr: false,
		},
		{
			name: "multiple currency pairs",
			cps: []types.ProviderTicker{
				btc_usd,
				eth_usd,
				sol_usd,
			},
			msg: cryptodotcom.InstrumentRequestMessage{
				Method: "subscribe",
				Params: cryptodotcom.InstrumentParams{
					Channels: []string{
						"ticker.BTCUSD-PERP",
						"ticker.ETHUSD-PERP",
						"ticker.SOLUSD-PERP",
					},
				},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wsHandler, err := cryptodotcom.NewWebSocketDataHandler(logger, cryptodotcom.DefaultWebSocketConfig)
			require.NoError(t, err)

			msgs, err := wsHandler.CreateMessages(tc.cps)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			expectedBz, err := json.Marshal(tc.msg)
			require.NoError(t, err)
			require.Equal(t, 1, len(msgs))
			require.EqualValues(t, expectedBz, msgs[0])
		})
	}
}
