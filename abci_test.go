package da_test

import (
	"context"
	"testing"

	"cosmossdk.io/log"
	"github.com/Wondertan/da"
	"github.com/Wondertan/da/test"
	abci "github.com/cometbft/cometbft/abci/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/rollkit/celestia-openrpc/types/das"
	"github.com/rollkit/celestia-openrpc/types/header"
	"github.com/stretchr/testify/require"
)

func Test_DA_E2E(t *testing.T) {
	db := dbm.NewMemDB()
	logger := log.NewTestLogger(t)
	daClient := &mockClient{}
	testApp := test.NewSimApp(logger, db, nil, true, daClient, test.EmptyAppOptions{})

	_, err := testApp.InitChain(&abci.RequestInitChain{})
	require.NoError(t, err)
}

var _ da.Client = mockClient{}

type mockClient struct {
}

func (c mockClient) GetVerifiedRangeByHeight(
	ctx context.Context,
	hdr *header.ExtendedHeader,
	endHeight uint64,
) ([]*header.ExtendedHeader, error) {
	panic("not implemented")
}

func (c mockClient) GetByHeight(ctx context.Context, height uint64) (*header.ExtendedHeader, error) {
	panic("not implemented")
}

func (c mockClient) SamplingStats(ctx context.Context) (das.SamplingStats, error) {
	panic("not implemented")
}
