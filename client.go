package da

import (
	"context"
	"fmt"

	celestia "github.com/rollkit/celestia-openrpc"
)

type Client interface {
	// NOTE: range is end inclusive at the moment
	DataCommitments(context.Context, uint64, uint64) ([][]byte, error)

	LatestSampledHeight(ctx context.Context) (uint64, error)
}

type celestiaClient struct {
	*celestia.Client
}

func NewClient(client *celestia.Client) Client {
	return celestiaClient{client}
}

func (c celestiaClient) DataCommitments(
	ctx context.Context,
	startHeight,
	endHeight uint64,
) ([][]byte, error) {
	latestHeader, err := c.Client.Header.GetByHeight(ctx, startHeight)
	if err != nil {
		return nil, fmt.Errorf("getting latest commitment DA header height %d: %w", startHeight, err)
	}

	hdrs, err := c.Client.Header.GetVerifiedRangeByHeight(ctx, latestHeader, endHeight+1)
	if err != nil {
		return nil, fmt.Errorf("getting DA header range(%d;%d] for sampling: %w", latestHeader.Height(), endHeight, err)
	}

	roots := make([][]byte, len(hdrs))
	for i, hdr := range hdrs {
		roots[i] = hdr.DataHash
	}

	return roots, nil
}

func (c celestiaClient) LatestSampledHeight(ctx context.Context) (uint64, error) {
	stats, err := c.Client.DAS.SamplingStats(ctx)
	if err != nil {
		return 0, err
	}
	return stats.SampledChainHead, nil
}
