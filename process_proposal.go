package da

import (
	"bytes"

	"github.com/Wondertan/da/modules/da"
	v1 "github.com/Wondertan/da/modules/da/v1"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	celestia "github.com/rollkit/celestia-openrpc"
)

func ProcessProposalHandler(dec sdk.TxDecoder, da da.Keeper, client celestia.Client) sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		latestHeight, err := da.LatestDataCommitmentHeight(ctx)
		if err != nil {
			return nil, err
		}

		msg := findDataCommitmentMsg(dec, req.Txs)
		if msg == nil {
			return accept(), nil
		}

		for idx, dataCommitment := range msg.DataCommitments {
			height := uint64(idx) + latestHeight + 1
			header, err := client.Header.GetByHeight(ctx, height)
			if err != nil {
				return reject(), nil
			}
			if !bytes.Equal(header.DataHash, dataCommitment) {
				return reject(), nil
			}
		}

		return &abci.ResponseProcessProposal{}, nil
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
