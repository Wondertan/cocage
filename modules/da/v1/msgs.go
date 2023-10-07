package dav1

import (
	"errors"
	"fmt"

	"crypto/sha256"
)

const HashSize = sha256.Size

func (msg MsgAttestDataCommitment) ValidateBasic() error {
	if len(msg.DataCommitments) == 0 {
		return errors.New("No data commitments")
	}

	for idx, dataCommitment := range msg.DataCommitments {
		if len(dataCommitment) != HashSize {
			return fmt.Errorf("data commitment at index %d has invalid length. Expected: %d, got: %d", idx, len(dataCommitment), HashSize)
		}
	}

	return nil
}
