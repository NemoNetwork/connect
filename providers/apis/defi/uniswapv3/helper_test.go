package uniswapv3_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/skip-mev/slinky/oracle/types"
	"github.com/skip-mev/slinky/providers/apis/defi/uniswapv3"
	"github.com/skip-mev/slinky/providers/apis/defi/uniswapv3/mocks"
	"github.com/skip-mev/slinky/providers/base/api/metrics"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewDevelopment()

	// PoolConfigs used for testing.
	wethusdcCfg = uniswapv3.PoolConfig{
		Address:       "0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8",
		BaseDecimals:  18,
		QuoteDecimals: 6,
		Invert:        true,
	}

	// Tickers used for testing.
	wethusdcTicker = types.NewProviderTicker(uniswapv3.Name, "WETH/USDC", wethusdcCfg.MustToJSON())
)

func createPriceFetcher(
	t *testing.T,
) *uniswapv3.PriceFetcher {
	t.Helper()

	client := mocks.NewEVMClient(t)
	fetcher, err := uniswapv3.NewPriceFetcher(
		logger,
		uniswapv3.DefaultAPIConfig,
		client,
	)
	require.NoError(t, err)

	return fetcher
}

func createPriceFetcherWithClient(
	t *testing.T,
	client uniswapv3.EVMClient,
) *uniswapv3.PriceFetcher {
	t.Helper()

	fetcher, err := uniswapv3.NewPriceFetcher(
		logger,
		uniswapv3.DefaultAPIConfig,
		client,
	)
	require.NoError(t, err)

	return fetcher
}

func createEVMClientWithResponse(
	t *testing.T,
	failedRequestErr error,
	responses []string,
	errs []error,
) uniswapv3.EVMClient {
	t.Helper()

	c := mocks.NewEVMClient(t)
	if failedRequestErr != nil {
		c.On("BatchCallContext", mock.Anything, mock.Anything).Return(failedRequestErr)
	} else {
		c.On("BatchCallContext", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			elems, ok := args.Get(1).([]rpc.BatchElem)
			require.True(t, ok)

			require.True(t, ok)
			require.Equal(t, len(elems), len(responses))
			require.Equal(t, len(elems), len(errs))

			for i, elem := range elems {
				elem.Result = &responses[i]
				elem.Error = errs[i]
				elems[i] = elem
			}
		})
	}

	return c
}
