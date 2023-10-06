package da

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/Wondertan/da/modules/da"
	v1 "github.com/Wondertan/da/modules/da/v1"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	celestia "github.com/rollkit/celestia-openrpc"
	"google.golang.org/protobuf/proto"
)

func VoteExtensionHandler(da da.Keeper, client celestia.Client) sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
		timeoutCtx, cancel := context.WithTimeout(ctx.Context(), time.Second) // ensure we don't block for too long
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

		if latestHeight >= stats.SampledChainHead {
			// no new heights sampled, skip...
			return resp, nil
		}

		latestHeader, err := client.Header.GetByHeight(timeoutCtx, latestHeight)
		if err != nil {
			return resp, fmt.Errorf("getting latest commitment DA header height %d: %w", latestHeight, err)
		}

		hdrs, err := client.Header.GetVerifiedRangeByHeight(timeoutCtx, latestHeader, stats.SampledChainHead+1)
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

		voteExtensionData := v1.VoteExtensionData{DataTuples: tuples}
		resp.VoteExtension, err = voteExtensionData.Marshal()
		return resp, err
	}
}

func PrepareProposalHandler(da da.Keeper) sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		resp := &abci.ResponsePrepareProposal{}


		var maxTuples int
		votes := map[string]*v1.VoteExtensionData{} // address -> vote extension
		attests := make([]*v1.Attestation, len(req.LocalLastCommit.Votes))

		for i, vt := range req.LocalLastCommit.Votes {
			data := &v1.VoteExtensionData{}
			err := proto.Unmarshal(vt.VoteExtension, data)
			if err != nil {
				return resp, fmt.Errorf("unmarshalling vote extension: %w", err)
			}

			votes[string(vt.Validator.Address)] = data
			attests[i] = &v1.Attestation{
				ValidatorAddress: string(vt.Validator.Address), // TODO change to bytes
				Size: uint32(len(data.DataTuples)),
				Signature:        vt.ExtensionSignature,
			}
			if len(data.DataTuples) > maxTuples {
				maxTuples = len(data.DataTuples)
			}
		}

		// TODO: We have to cross check validity of signed datatuples
		//  If height is adequate and matches the dataroot

		var minHeight uint64
		daHeights := make(map[uint64][][]byte, maxTuples) // height -> list of validator addresses that signed over the height
		roots := make(map[uint64][]byte, maxTuples)       // height -> dataroot

		for i := 0; i < maxTuples; i++ {
			var height uint64
			for _, vt := range votes {
				// get the first vote that has the ith height
				if len(vt.DataTuples) <= i {
					height = vt.DataTuples[i].Height
					roots[height] = vt.DataTuples[i].DataRoot
					break
				}
			}

			for addr, vt := range votes {
				if  vt.DataTuples[i].Height == height && bytes.Equal(vt.DataTuples[i].DataRoot, roots[height]) {
					// only add validator address if it signed over the same height and dataroot
					daHeights[height] = append(daHeights[height], []byte(addr))
				}
			}

			if minHeight == 0 || height < minHeight {
				minHeight = height
			}
		}

		endHeight, err := da.HighestHeightWithMajority(ctx.Context(), daHeights)
		if err != nil {
			return nil, err
		}

		commitments := make([][]byte, endHeight-minHeight+1)
		for i := minHeight; i <= endHeight; i++ {
			commitments[i-minHeight] = roots[i]
		}

		_ = &v1.MsgAttestDataCommitment{
			DataCommitments: commitments,
			Attestations: attests,
			EndHeight:    endHeight,
		}

		// TODO: Pack msg into the response

		// create the MsgAttestDataCommitment from the req.LocalLastCommit.Votes (which contain the vote extensions)
		// insert the transaction as the first transaction in the response
		// check that the max size is not reached
		return &abci.ResponsePrepareProposal{
			Txs: req.Txs,
		}, nil
	}
}

func ProcessProposalHandler() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		// check if there is a MsgAttestDataCommitment in the proposal
		// verify it using the da keeper logic
		return &abci.ResponseProcessProposal{}, nil
	}
}
