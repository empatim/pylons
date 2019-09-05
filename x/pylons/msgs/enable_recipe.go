package msgs

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgEnableRecipe defines a SetName message
type MsgEnableRecipe struct {
	RecipeID string
	Sender   sdk.AccAddress
}

// NewMsgEnableRecipe a constructor for EnableCookbook msg
func NewMsgEnableRecipe(recipeID string, sender sdk.AccAddress) MsgEnableRecipe {
	return MsgEnableRecipe{
		RecipeID: recipeID,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg MsgEnableRecipe) Route() string { return "pylons" }

// Type should return the action
func (msg MsgEnableRecipe) Type() string { return "enable_recipe" }

// ValidateBasic validates the Msg
func (msg MsgEnableRecipe) ValidateBasic() sdk.Error {

	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress(msg.Sender.String())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEnableRecipe) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners gets the signer who should have signed the message
func (msg MsgEnableRecipe) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
