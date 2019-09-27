package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodespace is the Module Name
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeInvalidTopicID sdk.CodeType = 1
	CodeInvalidWeight  sdk.CodeType = 2
)

// ErrInvalidTopicID is the error for invalid topic id
func ErrInvalidTopicID() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidTopicID, "Invalid topic ID")
}

// ErrInvalidWeight is the error for invalid weight
func ErrInvalidWeight() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidWeight, "Invalid weight")
}
