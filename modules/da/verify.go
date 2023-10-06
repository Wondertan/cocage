package da

import (
	"context"

	"github.com/Wondertan/da/modules/da/v1"
)

func (k Keeper) VerifyDataCommitment(ctx context.Context, msg *v1.MsgAttestDataCommitment) error {
	latestHeight := k.LatestDataCommitmentHeight()
	dataTupleMap := make(map[int64][]byte)
	for _, attestation := range msg.Attestations {
		if attestation.Size > msg.EndHeight - latestHeight {
			
		}

		voteExtensionData := v1.VoteExtensionData{
			DataTuples: make([]*v1.DataTuple, attestation.Size),
		}
		for height := latestHeight + 1; height <= msg.EndHeight; height++ {

	}

	return nil
}

