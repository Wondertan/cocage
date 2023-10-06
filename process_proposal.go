package da

import (
	"bytes"
	"context"
	"time"

	"github.com/Wondertan/da/modules/da"
	v1 "github.com/Wondertan/da/modules/da/v1"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	celestia "github.com/rollkit/celestia-openrpc"
)

// ProcessProposalHandler is required for using the DA module. It loops to
// find the first MsgAttestDataCommitment in the block. It then verifies the
// transaction by taking the height for the last committed data root and
// checking that each height listed in the message has been sampled and is
// "available" to the node. Futhermore it checks that the data root matches
// it's local data root. If this passes, it approves the proposal
func ProcessProposalHandler(dec sdk.TxDecoder, da da.Keeper, client *celestia.Client) sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		msg := findDataCommitmentMsg(dec, req.Txs)
		if msg == nil {
			return accept(), nil
		}

		latestHeight, err := da.LatestDataCommitmentHeight(ctx)
		if err != nil {
			return reject(), err
		}

		timeoutCtx, cancel := context.WithTimeout(ctx.Context(), time.Second) // ensure we don't block for too long
		defer cancel()

		stats, err := client.DAS.SamplingStats(timeoutCtx)
		if err != nil {
			return reject(), nil
		}

		for idx, dataCommitment := range msg.DataCommitments {
			height := uint64(idx) + latestHeight + 1
			if stats.SampledChainHead < height {
				return reject(), nil
			}
			// TODO: these should theoretically be all the heights
			// that the node has already sampled as it would have
			// already indicated the sampled heights in the vote extensions.
			// What would be more reliable is to cache the sampled header's
			// data roots so we don't have to make a second remote call
			header, err := client.Header.GetByHeight(timeoutCtx, height)
			if err != nil {
				return reject(), nil
			}
			if !bytes.Equal(header.DataHash, dataCommitment) {
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
