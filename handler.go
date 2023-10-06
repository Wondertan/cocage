package da

import (
	"context"
	"fmt"
	"time"

	"github.com/Wondertan/da/modules/da"
	v1 "github.com/Wondertan/da/modules/da/v1"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	celestia "github.com/rollkit/celestia-openrpc"
)

func VoteExtensionHandler(da da.Keeper, client celestia.Client) sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
		timeoutCtx, cancel := context.WithTimeout(ctx.Context(), time.Second)
		defer cancel()

		resp := &abci.ResponseExtendVote{}
		latestHeight, err := da.LatestDataCommitmentHeight(ctx)
		if err != nil {
			return resp, fmt.Errorf("getting latest commitment DA height: %w", err)
		}

		stats, err := client.DAS.SamplingStats(timeoutCtx)
		if err != nil {
			return resp, fmt.Errorf("getting DA sampling stats: %w", err)
		}

		if latestHeight >= stats.SampleChainHead {
			// no new heights sampled, skip...
			return resp, nil
		}
		toSample := stats.SampledChainHead - latestHeight

		latestHeader, err := client.Header.GetByHeight(timeoutCtx, latestHeight)
		if err != nil {
			return resp, fmt.Errorf("getting latest commitment DA header height %d: %w", latestHeight, err)
		}

		hdrs, err := client.Header.GetVerifiedRangeByHeight(timeoutCtx, latestHeader, toSample)
		if err != nil {
			return resp, fmt.Errorf("getting DA header range(%d;%d] for sampling: %w", latestHeader.Height(), stats.SampledChainHead, err)
		}

		tuples := make([]*v1.DataTuple, len(hdrs))
		for i, hdr := range hdrs {
			tuples[i] = &v1.DataTuple{
				Height:   hdr.Height(),
				DataRoot: hdr.DataHash,
			}
		}

		voteExtensionData := v1.VoteExtensionData{
			DataTuples: tuples,
		}

		bz, err := voteExtensionData.Marshal()
		if err != nil {
			return resp, err
		}

		resp.VoteExtension = bz
		return resp, nil
	}
}

func PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		return &abci.ResponsePrepareProposal{Txs: req.Txs}, nil
	}
}

func ProcessProposalHandler() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		return &abci.ResponseProcessProposal{}, nil
	}
}
