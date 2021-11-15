package wasm

import "github.com/irisnet/core-sdk-go/types/errors"

const (
	MaxWasmSize = 500 * 1024

	// MaxLabelSize is the longest label that can be used when Instantiating a contract
	MaxLabelSize = 128
)

func validateWasmCode(s []byte) error {
	if len(s) == 0 {
		return errors.Wrap(ErrEmpty, "is required")
	}
	if len(s) > MaxWasmSize {
		return errors.Wrapf(ErrLimit, "cannot be longer than %d bytes", MaxWasmSize)
	}
	return nil
}

func validateLabel(label string) error {
	if label == "" {
		return errors.Wrap(ErrEmpty, "is required")
	}
	if len(label) > MaxLabelSize {
		return errors.Wrap(ErrLimit, "cannot be longer than 128 characters")
	}
	return nil
}
