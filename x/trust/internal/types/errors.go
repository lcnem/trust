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
	return sdk.NewError(DefaultCodespace, CodeInvalidTopicID, "Topic ID must start with [a-z] and available characters are [a-z],[0-9],-")
}

// ErrInvalidWeight is the error for invalid weight
func ErrInvalidWeight() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidWeight, "1000times weight must be positive and less than or equal to 1000.")
}
