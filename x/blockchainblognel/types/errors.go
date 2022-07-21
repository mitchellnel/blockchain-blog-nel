package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/blockchainblognel module sentinel errors
var (
	ErrCommentOld = sdkerrors.Register(
		ModuleName,
		1300,
		"Comment more than 100 blocks older than Post",
	)
)
