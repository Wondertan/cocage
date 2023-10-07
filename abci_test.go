package da_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/log"
	da "github.com/Wondertan/cocage"
	"github.com/Wondertan/cocage/test"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func Test_DA_E2E(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	db := dbm.NewMemDB()
	logger := log.NewTestLogger(t)

	dataCommitments := map[uint64][]byte{
		1: crypto.CRandBytes(32),
		2: crypto.CRandBytes(32),
		3: crypto.CRandBytes(32),
	}

	daClient := NewMockClient(dataCommitments)
	testApp := test.NewSimApp(logger, db, nil, true, daClient, test.EmptyAppOptions{})

	genesis := testApp.DefaultGenesis()
	stateBytes, err := json.MarshalIndent(genesis, "", " ")
	require.NoError(t, err)

	_, err = testApp.InitChain(&abci.RequestInitChain{
		AppStateBytes: stateBytes,
	})
	require.NoError(t, err)

	// assert that the app responds with the height of the last
	// sampled header on the DA network
	extendVoteResp, err := testApp.ExtendVote(ctx, &abci.RequestExtendVote{})
	require.NoError(t, err)
	require.NotNil(t, extendVoteResp.VoteExtension)
	// binary.PutUint64(extendVoteResp.VoteExtension)

	// Move on to height 2
	_, err = testApp.Commit()
	require.NoError(t, err)

	// // generate a local commit info from the vote extensions
	// prepareProposalResp, err := testApp.PrepareProposal(&abci.RequestPrepareProposal{})
	// require.NoError(t, err)

	// // assert that the node only votes for the block if the data commitments match
	// processProposalResp, err := testApp.ProcessProposal(&abci.RequestProcessProposal{})

	// // assert that when the transaction is submitted that the da module
	// // correctly updates its state
	// _, err := testApp.FinalizeBlock(&abci.RequestFinalizeBlock{})

}

var _ da.Client = mockClient{}

type mockClient struct {
	dataCommitments map[uint64][]byte
}

func NewMockClient(dataCommitmets map[uint64][]byte) da.Client {
	return &mockClient{
		dataCommitments: dataCommitmets,
	}
}

func (c mockClient) DataCommitments(
	ctx context.Context,
	startHeight,
	endHeight uint64,
) ([][]byte, error) {
	roots := make([][]byte, endHeight-startHeight+1)
	i := 0
	for height := startHeight; height <= endHeight; height++ {
		dataCommitment, exists := c.dataCommitments[height]
		if !exists {
			return nil, fmt.Errorf("date commitment at height %d not found", height)
		}
		roots[i] = dataCommitment
		i++
	}
	return roots, nil
}

func (c mockClient) LatestSampledHeight(ctx context.Context) (uint64, error) {
	highestHeight := uint64(0)
	for height := range c.dataCommitments {
		if height > highestHeight {
			highestHeight = height
		}
	}
	return highestHeight, nil
}
