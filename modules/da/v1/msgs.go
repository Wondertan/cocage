package v1

import (
	"fmt"

	"crypto/sha256"
)

const HashSize = sha256.Size

func (msg MsgAttestDataCommitment) ValidateBasic() error {
	for idx, dataCommitment := range msg.DataCommitments {
		if len(dataCommitment) != HashSize {
			return fmt.Errorf("data commitment at index %d has invalid length. Expected: %d, got: %d", idx, len(dataCommitment), HashSize)
		}
	}

	if msg.EndHeight == 0 {
		return fmt.Errorf("end height cannot be zero")
	}

	for idx, attestation := range msg.Attestations {
		if attestation == nil {
			return fmt.Errorf("attestation at index %d is nil", idx)
		}
	}

	return nil
}
