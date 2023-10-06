package da

import (
	"fmt"

	"github.com/Wondertan/da/modules/da"
	v1 "github.com/Wondertan/da/modules/da/v1"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	celestia "github.com/rollkit/celestia-openrpc"
)

func VoteExtensionHandler(da da.Keeper, client celestia.Client) sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
		resp := &abci.ResponseExtendVote{}
		latestHeight, err := da.LatestDataCommitmentHeight(ctx)
		if err != nil {
			return resp, fmt.Errorf("getting latest commitment DA height: %w", err)
		}

		latestHeader, err := client.Header.GetByHeight(ctx.Context(), latestHeight)
		if err != nil {
			return resp, fmt.Errorf("getting latest commitment DA header height %d: %w", latestHeight, err)
		}

		networkHead, err := client.Header.NetworkHead(ctx.Context())
		if err != nil {
			return resp, fmt.Errorf("getting DA network head: %w", err)
		}

		toSample := networkHead.Height() - latestHeader.Height()
		if toSample == 0 {
			// TODO: Figure out what to send in response as it is required
			return resp, nil
		}
	
		hdrs, err := client.Header.GetVerifiedRangeByHeight(ctx.Context(), latestHeader, toSample)
		if err != nil {
			return resp, fmt.Errorf("getting DA header range(%d;%d] for sampling: %w", latestHeader.Height(), networkHead.Height(), err)
		}

		tuples := make([]*v1.DataTuple, len(hdrs))

		// NOTE: Not much value in doing this concurently, as DASer already does it for us.  
		for i, hdr := range hdrs {
			err := client.Share.SharesAvailable(ctx.Context(), hdr.DAH)
			if err != nil {
				return resp, fmt.Errorf("checking DA availability for height %d: %w", hdr.Height(), err)
			}

			tuples[i] = &v1.DataTuple{
				Height: hdr.Height(),
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
