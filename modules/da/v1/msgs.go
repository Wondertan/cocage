package v1

import (
	"fmt"

	"crypto/sha256"
)

const (
	HashSize = sha256.Size

	// There can only be 10 data root attestations per block
	MaxAttestations = 10
)

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

		if int(attestation.Size) > len(msg.DataCommitments) {
			return fmt.Errorf("attestation size exceeded number of data commitments: %d > %d", attestation.Size, len(msg.DataCommitments))
		}

		if int(attestation.Size) > MaxAttestations {
			return fmt.Errorf("attestation size exceeded max allowed: %d > %d", attestation.Size, MaxAttestations)
		}
	}

	return nil
}
