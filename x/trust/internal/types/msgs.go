package types

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgEvaluate defines a Evaluate message
type MsgEvaluate struct {
	TopicID     string         `json:"topic_id"`
	FromAddress sdk.AccAddress `json:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address"`
	Weight1000  sdk.Int        `json:"weight1000"`
}

// NewMsgEvaluate is a constructor function for MsgEvaluate
func NewMsgEvaluate(topicID string, fromAddress sdk.AccAddress, toAddress sdk.AccAddress, weight1000 sdk.Int) MsgEvaluate {
	return MsgEvaluate{
		TopicID:     topicID,
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Weight1000:  weight1000,
	}
}

// Route should return the name of the module
func (msg MsgEvaluate) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEvaluate) Type() string { return "set_evaluation" }

// ValidateBasic runs stateless checks on the message
func (msg MsgEvaluate) ValidateBasic() sdk.Error {
	if !validateTopicID(msg.TopicID) {
		return ErrInvalidTopicID()
	}
	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress(msg.FromAddress.String())
	}
	if msg.ToAddress.Empty() {
		return sdk.ErrInvalidAddress(msg.ToAddress.String())
	}
	if msg.Weight1000.IsNegative() || msg.Weight1000.GT(sdk.NewInt(1000)) {
		return ErrInvalidWeight()
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEvaluate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEvaluate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// MsgDistributeTokenByScore defines a DistributeTokenByScore message
type MsgDistributeTokenByScore struct {
	TopicID     string         `json:"topic_id"`
	FromAddress sdk.AccAddress `json:"from_address"`
	Amount      sdk.Coin       `json:"amount"`
}

// NewMsgDistributeTokenByScore is a constructor function for MsgMsgDistributeTokenByScore
func NewMsgDistributeTokenByScore(topicID string, fromAddress sdk.AccAddress, amount sdk.Coin) MsgDistributeTokenByScore {
	return MsgDistributeTokenByScore{
		TopicID:     topicID,
		FromAddress: fromAddress,
		Amount:      amount,
	}
}

// Route should return the name of the module
func (msg MsgDistributeTokenByScore) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDistributeTokenByScore) Type() string { return "distribute_token_by_score" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDistributeTokenByScore) ValidateBasic() sdk.Error {
	if !validateTopicID(msg.TopicID) {
		return ErrInvalidTopicID()
	}
	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress(msg.FromAddress.String())
	}
	if msg.Amount.IsNegative() {
		return sdk.ErrInvalidCoins(msg.Amount.String())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDistributeTokenByScore) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDistributeTokenByScore) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// MsgDistributeTokenByEvaluation defines a DistributeTokenByEvaluation message
type MsgDistributeTokenByEvaluation struct {
	TopicID     string         `json:"topic_id"`
	Address     sdk.AccAddress `json:"address"`
	FromAddress sdk.AccAddress `json:"from_address"`
	Amount      sdk.Coin       `json:"amount"`
}

// NewMsgDistributeTokenByEvaluation is a constructor function for MsgMsgDistributeTokenByEvaluation
func NewMsgDistributeTokenByEvaluation(topicID string, address sdk.AccAddress, fromAddress sdk.AccAddress, amount sdk.Coin) MsgDistributeTokenByEvaluation {
	return MsgDistributeTokenByEvaluation{
		TopicID:     topicID,
		Address:     address,
		FromAddress: fromAddress,
		Amount:      amount,
	}
}

// Route should return the name of the module
func (msg MsgDistributeTokenByEvaluation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDistributeTokenByEvaluation) Type() string { return "distribute_token_by_evaluation" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDistributeTokenByEvaluation) ValidateBasic() sdk.Error {
	if !validateTopicID(msg.TopicID) {
		return ErrInvalidTopicID()
	}
	if msg.Address.Empty() {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}
	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress(msg.FromAddress.String())
	}
	if msg.Amount.IsNegative() {
		return sdk.ErrInvalidCoins(msg.Amount.String())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDistributeTokenByEvaluation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDistributeTokenByEvaluation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

func validateTopicID(topicID string) bool {
	return regexp.MustCompile("^[0-9a-z]{32}$").Match([]byte(topicID))
}
