package da

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types"

	v1 "github.com/Wondertan/da/modules/da/v1"
)

func (k Keeper) VerifyDataCommitment(ctx context.Context, msg *v1.MsgAttestDataCommitment) error {
	latestHeight, err := k.LatestDataCommitmentHeight(ctx)
	if err != nil {
		return err
	}
	voteExtensionMap := make(map[uint32][]byte)
	for _, attestation := range msg.Attestations {

		voteExtension, ok := voteExtensionMap[attestation.Size]
		if !ok {
			voteExtension, err = makeVoteExtension(msg.DataCommitments, latestHeight, attestation.Size)
			if err != nil {
				return err
			}
			voteExtensionMap[attestation.Size] = voteExtension
		}

		addr, err := types.ConsAddressFromBech32(attestation.ValidatorAddress)
		if err != nil {
			return err
		}
		validator, err := k.stakingKeeper.GetValidatorByConsAddr(ctx, addr)
		if err != nil {
			return err
		}

	}

	return nil
}

func makeVoteExtension(dataCommitments [][]byte, latestHeight uint64, size uint32) ([]byte, error) {
	voteExtensionData := v1.VoteExtensionData{
		DataTuples: make([]*v1.DataTuple, size),
	}
	for i := 0; i < int(size); i++ {
		height := latestHeight + uint64(i)
		dataTuple := &v1.DataTuple{
			Height:   height,
			DataRoot: dataCommitments[i],
		}
		voteExtensionData.DataTuples[i] = dataTuple
	}

	return voteExtensionData.Marshal()
}
