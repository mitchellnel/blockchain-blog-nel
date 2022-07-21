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

	ErrPostDoesNotExist = sdkerrors.Register(
		ModuleName,
		1401,
		"Post does not exist",
	)

	ErrCommentDoesNotExists = sdkerrors.Register(
		ModuleName,
		1402,
		"Comment does not exist",
	)

	ErrCommentNotFoundOnPost = sdkerrors.Register(
		ModuleName,
		1403,
		"Comment not found on a Post",
	)
)
