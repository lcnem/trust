package types

import (
	"regexp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// verify interface at compile time
var _ sdk.Msg = &MsgEvaluate{}

// MsgEvaluate - struct for unjailing jailed validator
type MsgEvaluate struct {
	TopicID     string         `json:"topic_id"`
	FromAddress sdk.AccAddress `json:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address"`
	MilliWeight  float32       `json:"milli_weight"`
}

// NewMsgEvaluate creates a new MsgEvaluate instance
func NewMsgEvaluate(topicID string, fromAddress sdk.AccAddress, toAddress sdk.AccAddress, milliWeight float32) MsgEvaluate {
	return MsgEvaluate{
		TopicID:     topicID,
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		MilliWeight:  milliWeight,
	}
}

const evaluateConst = "evaluate"

// nolint
func (msg MsgEvaluate) Route() string { return RouterKey }
func (msg MsgEvaluate) Type() string  { return evaluateConst }
func (msg MsgEvaluate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgEvaluate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgEvaluate) ValidateBasic() error {
	if !validateTopicID(msg.TopicID) {
		return ErrInvalidTopicID
	}
	if msg.FromAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing from_address")
	}
	if msg.ToAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing to_address")
	}
	if msg.MilliWeight <= 0 {
		return ErrInvalidWeight
	}
	return nil
}

// verify interface at compile time
var _ sdk.Msg = &MsgDistributeByScore{}

// MsgDistributeByScore - struct for unjailing jailed validator
type MsgDistributeByScore struct {
	TopicID     string         `json:"topic_id"`
	FromAddress sdk.AccAddress `json:"from_address"`
	Amount      sdk.Coin       `json:"amount"`
}

// NewMsgDistributeByScore creates a new MsgDistributeByScore instance
func NewMsgDistributeByScore(topicID string, fromAddress sdk.AccAddress, amount sdk.Coin) MsgDistributeByScore {
	return MsgDistributeByScore{
		TopicID:     topicID,
		FromAddress: fromAddress,
		Amount:      amount,
	}
}

const distributeByScoreConst = "distributeByScore"

// nolint
func (msg MsgDistributeByScore) Route() string { return RouterKey }
func (msg MsgDistributeByScore) Type() string  { return distributeByScoreConst }
func (msg MsgDistributeByScore) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDistributeByScore) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgDistributeByScore) ValidateBasic() error {
	if !validateTopicID(msg.TopicID) {
		return ErrInvalidTopicID
	}
	if msg.FromAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing from_address")
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	return nil
}

// verify interface at compile time
var _ sdk.Msg = &MsgDistributeByEvaluation{}

// MsgDistributeByEvaluation - struct for unjailing jailed validator
type MsgDistributeByEvaluation struct {
	TopicID     string         `json:"topic_id"`
	FromAddress sdk.AccAddress `json:"from_address"`
	EvaluatedAddress     sdk.AccAddress `json:"evaluated_address"`
	Amount      sdk.Coin       `json:"amount"`
}

// NewMsgDistributeByEvaluation creates a new MsgDistributeByEvaluation instance
func NewMsgDistributeByEvaluation(topicID string, fromAddress sdk.AccAddress, evaluatedAddress sdk.AccAddress, amount sdk.Coin) MsgDistributeByEvaluation {
	return MsgDistributeByEvaluation{
		TopicID:     topicID,
		FromAddress: fromAddress,
		EvaluatedAddress:     evaluatedAddress,
		Amount:      amount,
	}
}

const distributeByEvaluationConst = "distributeByEvaluation"

// nolint
func (msg MsgDistributeByEvaluation) Route() string { return RouterKey }
func (msg MsgDistributeByEvaluation) Type() string  { return distributeByEvaluationConst }
func (msg MsgDistributeByEvaluation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDistributeByEvaluation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgDistributeByEvaluation) ValidateBasic() error {
	if !validateTopicID(msg.TopicID) {
		return ErrInvalidTopicID
	}
	if msg.FromAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing from_address")
	}
	if msg.EvaluatedAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing evaluated_address")
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	return nil
}

func validateTopicID(topicID string) bool {
	return regexp.MustCompile("^[a-z_][a-z_0-9]*$").Match([]byte(topicID))
}
