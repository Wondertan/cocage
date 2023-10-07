package da

// import (
// 	"context"
// 	"fmt"

// 	v1 "github.com/Wondertan/cocage/modules/da/v1"
// )

// var _ v1.MsgServer = msgServer{}

// type msgServer struct {
// 	v1.UnimplementedMsgServer

// 	keeper Keeper
// }

// func NewMsgServer(keeper Keeper) v1.MsgServer {
// 	return msgServer{keeper: keeper}
// }

// func (s msgServer) AttestDataRoot(ctx context.Context, msg *v1.MsgAttestDataCommitment) (*v1.MsgAttestDataCommitmentResponse, error) {
// 	latestHeight, err := s.keeper.LatestDataCommitmentHeight(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting latest commitment DA height: %w", err)
// 	}

// 	for idx, dataCommitment := range msg.DataCommitments {
// 		height := uint64(idx) + latestHeight + 1
// 		err = s.keeper.SetDataCommitment(ctx, height, dataCommitment)
// 		if err != nil {
// 			return nil, fmt.Errorf("setting data commitment at DA height %d: %w", height, err)
// 		}
// 	}

// 	return &v1.MsgAttestDataCommitmentResponse{}, nil
// }
