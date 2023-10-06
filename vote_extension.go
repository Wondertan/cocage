package da

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/Wondertan/cocage/modules/da"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	celestia "github.com/rollkit/celestia-openrpc"
)

func ExtendVoteHandler(da da.Keeper, client *celestia.Client) sdk.ExtendVoteHandler {
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

		resp.VoteExtension = binary.LittleEndian.AppendUint64(nil, stats.SampledChainHead)
		return resp, err
	}
}
