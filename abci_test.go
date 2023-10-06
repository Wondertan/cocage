package da_test

import (
	"context"
	"testing"
	"time"

	"cosmossdk.io/log"
	da "github.com/Wondertan/cocage"
	"github.com/Wondertan/cocage/test"
	abci "github.com/cometbft/cometbft/abci/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/rollkit/celestia-openrpc/types/das"
	"github.com/rollkit/celestia-openrpc/types/header"
	"github.com/stretchr/testify/require"
)

func Test_DA_E2E(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	db := dbm.NewMemDB()
	logger := log.NewTestLogger(t)
	daClient := &mockClient{}
	testApp := test.NewSimApp(logger, db, nil, true, daClient, test.EmptyAppOptions{})

	_, err := testApp.InitChain(&abci.RequestInitChain{})
	require.NoError(t, err)

	// assert that the app responds with the height of the last
	// sampled header on the DA network
	extendVoteResp, err := testApp.ExtendVote(ctx, &abci.RequestExtendVote{})

	// Move on to height 2
	_, err = testApp.Commit()
	require.NoError(t, err)

	// generate a local commit info from the vote extensions
	prepareProposalResp, err := testApp.PrepareProposal(&abci.RequestPrepareProposal{})
	require.NoError(t, err)

	// assert that the node only votes for the block if the data commitments match
	processProposalResp, err := testApp.ProcessProposal(&abci.RequestProcessProposal{})

	// assert that when the transaction is submitted that the da module
	// correctly updates its state
	_, err := testApp.FinalizeBlock(&abci.RequestFinalizeBlock{})

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
