package da

import (
	"context"

	celestia "github.com/rollkit/celestia-openrpc"
	"github.com/rollkit/celestia-openrpc/types/das"
	"github.com/rollkit/celestia-openrpc/types/header"
)

type Client interface {
	GetVerifiedRangeByHeight(
		context.Context,
		*header.ExtendedHeader,
		uint64,
	) ([]*header.ExtendedHeader, error)

	GetByHeight(context.Context, uint64) (*header.ExtendedHeader, error)

	SamplingStats(ctx context.Context) (das.SamplingStats, error)
}

type celestiaClient struct {
	*celestia.Client
}

func NewClient(client *celestia.Client) Client {
	return celestiaClient{client}
}

func (c celestiaClient) GetVerifiedRangeByHeight(
	ctx context.Context,
	hdr *header.ExtendedHeader,
	endHeight uint64,
) ([]*header.ExtendedHeader, error) {
	return c.Client.Header.GetVerifiedRangeByHeight(ctx, hdr, endHeight)
}

func (c celestiaClient) GetByHeight(ctx context.Context, height uint64) (*header.ExtendedHeader, error) {
	return c.Client.Header.GetByHeight(ctx, height)
}

func (c celestiaClient) SamplingStats(ctx context.Context) (das.SamplingStats, error) {
	return c.Client.DAS.SamplingStats(ctx)
}
