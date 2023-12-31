package da

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/Wondertan/cocage/modules/da"
	v1 "github.com/Wondertan/cocage/modules/da/v1"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func PrepareProposalHandler(da da.Keeper, cfg client.TxConfig, client Client) sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		resp := &abci.ResponsePrepareProposal{}
		if len(req.LocalLastCommit.Votes) == 0 {
			return resp, nil
		}

		latestHeight, err := da.LatestDataCommitmentHeight(ctx)
		if err != nil {
			return resp, fmt.Errorf("getting latest commitment DA height: %w", err)
		}

		daHeights := make(map[uint64][][]byte)
		for _, vt := range req.LocalLastCommit.Votes {
			height := binary.LittleEndian.Uint64(vt.VoteExtension)
			if height <= latestHeight {
				// skip votes that are already committed
				continue
			}

			daHeights[height] = append(daHeights[height], vt.Validator.Address)
		}

		endHeight, err := da.HighestHeightWithMajority(ctx.Context(), daHeights)
		if err != nil {
			return resp, fmt.Errorf("getting highest height with majority: %w", err)
		}

		timeoutCtx, cancel := context.WithTimeout(ctx.Context(), time.Second) // ensure we don't block for too long
		defer cancel()

		roots, err := client.DataCommitments(timeoutCtx, latestHeight+1, endHeight)
		if err != nil {
			return resp, err
		}

		bldr := cfg.NewTxBuilder()
		err = bldr.SetMsgs(&v1.MsgAttestDataCommitment{
			DataCommitments: roots,
		})
		if err != nil {
			return resp, fmt.Errorf("setting DA message to builder: %w", err)
		}

		data, err := cfg.TxEncoder()(bldr.GetTx())
		if err != nil {
			return resp, fmt.Errorf("encoding DA transaction: %w", err)
		}

		return &abci.ResponsePrepareProposal{
			Txs: append(req.Txs, data),
		}, nil
	}
}
