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
	celestia "github.com/rollkit/celestia-openrpc"
)

func PrepareProposalHandler(da da.Keeper, cfg client.TxConfig, client *celestia.Client) sdk.PrepareProposalHandler {
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

		latestHeader, err := client.Header.GetByHeight(timeoutCtx, latestHeight)
		if err != nil {
			return resp, fmt.Errorf("getting latest commitment DA header height %d: %w", latestHeight, err)
		}

		hdrs, err := client.Header.GetVerifiedRangeByHeight(timeoutCtx, latestHeader, endHeight+1)
		if err != nil {
			return resp, fmt.Errorf("getting DA header range(%d;%d] for sampling: %w", latestHeader.Height(), endHeight, err)
		}

		roots := make([][]byte, len(hdrs))
		for i, hdr := range hdrs {
			roots[i] = hdr.DataHash
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
