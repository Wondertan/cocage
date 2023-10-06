package v1

import "google.golang.org/protobuf/proto"

func (v *VoteExtensionData) Marshal() ([]byte, error) {
	return proto.Marshal(v)
}
