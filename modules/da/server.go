package da

import (
	"context"

	v1 "github.com/Wondertan/da/modules/da/v1"
)

var _ v1.MsgServer = msgServer{}

type msgServer struct {
	v1.UnimplementedMsgServer

	keeper Keeper
}

func NewMsgServer(keeper Keeper) v1.MsgServer {
	return msgServer{keeper: keeper}
}

func (s msgServer) AttestDataRoot(ctx context.Context, msg *v1.MsgAttestDataCommitment) (*v1.MsgAttestDataCommitmentResponse, error) {
	return nil, nil
}
