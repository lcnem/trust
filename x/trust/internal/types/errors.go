package types

import (

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
var (
	ErrInvalidTopicID = sdkerrors.Register(ModuleName, 1, "topic id must start with [a-z] and available characters are [a-z],[0-9],-")
	ErrInvalidWeight = sdkerrors.Register(ModuleName, 2, "weight must be positive")
)
