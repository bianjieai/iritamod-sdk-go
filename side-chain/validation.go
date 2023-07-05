package side_chain

import (
	"fmt"
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	nftClassIdString = `[a-z][a-zA-Z0-9/]{2,100}`
	nftTokenIdString = `[a-zA-Z0-9/]{1,100}`

	regexpNftClassId = regexp.MustCompile(fmt.Sprintf(`^%s$`, nftClassIdString)).MatchString
	regexpNftTokenId = regexp.MustCompile(fmt.Sprintf(`^%s$`, nftTokenIdString)).MatchString
)

func ValidateSpaceId(spaceId uint64) error {
	if spaceId == 0 {
		return sdkerrors.Wrapf(ErrInvalidSpace, "space id cannot be zero")
	}
	return nil
}

func ValidateClassIdForNFT(classId string) error {
	if !regexpNftClassId(classId) {
		return sdkerrors.Wrapf(ErrInvalidClassIdForNFT, "class id can only accept characters that match the regular expression: (%s),but got (%s)", nftClassIdString, classId)
	}
	return nil
}

func ValidateTokenIdForNFT(tokenId string) error {
	if !regexpNftTokenId(tokenId) {
		return sdkerrors.Wrapf(ErrInvalidTokenIdForNFT, "token id can only accept characters that match the regular expression: (%s),but got (%s)", nftClassIdString, tokenId)
	}
	return nil
}
