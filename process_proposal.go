package da

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/Wondertan/cocage/modules/da"
	v1 "github.com/Wondertan/cocage/modules/da/v1"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ProcessProposalHandler is required for using the DA module. It loops to
// find the first MsgAttestDataCommitment in the block. It then verifies the
// transaction by taking the height for the last committed data root and
// checking that each height listed in the message has been sampled and is
// "available" to the node. Futhermore it checks that the data root matches
// it's local data root. If this passes, it approves the proposal
func ProcessProposalHandler(dec sdk.TxDecoder, da da.Keeper, client Client) sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		msg := findDataCommitmentMsg(dec, req.Txs)
		if msg == nil {
			return accept(), nil
		}

		latestHeight, err := da.LatestDataCommitmentHeight(ctx)
		if err != nil {
			return reject(), fmt.Errorf("getting latest commitment DA height: %w", err)
		}
		startHeight := latestHeight + 1
		endHeight := latestHeight + uint64(len(msg.DataCommitments))

		timeoutCtx, cancel := context.WithTimeout(ctx.Context(), time.Second) // ensure we don't block for too long
		defer cancel()

		latestSampledHeight, err := client.LatestSampledHeight(timeoutCtx)
		if err != nil {
			return reject(), nil
		}

		// Check if the data commitments are more than what the node has locally sampled.
		// If so then either the proposer has been malicious or has seen 2/3+ of the other
		// validators indicate that they have a later height
		if latestSampledHeight < endHeight {
			return reject(), nil
		}

		ourDataCommitments, err := client.DataCommitments(timeoutCtx, startHeight, endHeight)
		if err != nil {
			return reject(), nil
		}

		for idx, dataCommitment := range msg.DataCommitments {
			// if the proposer has proposed a different data commitment for any of the heights, we reject
			if !bytes.Equal(ourDataCommitments[idx], dataCommitment) {
				return reject(), nil
			}
		}

		return accept(), nil
	}
}

func reject() *abci.ResponseProcessProposal {
	return &abci.ResponseProcessProposal{
		Status: abci.ResponseProcessProposal_REJECT,
	}
}

func accept() *abci.ResponseProcessProposal {
	return &abci.ResponseProcessProposal{
		Status: abci.ResponseProcessProposal_ACCEPT,
	}
}

func findDataCommitmentMsg(decoder sdk.TxDecoder, Txs [][]byte) *v1.MsgAttestDataCommitment {
	for _, rawTx := range Txs {
		tx, err := decoder(rawTx)
		if err != nil {
			continue
		}
		msgs := tx.GetMsgs()
		if len(msgs) != 1 {
			continue
		}
		if msg, ok := msgs[0].(*v1.MsgAttestDataCommitment); ok {
			return msg
		}
	}
	return nil
}
