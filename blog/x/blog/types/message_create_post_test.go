package types

import (
	"testing"

	"blog/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreatePost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreatePost
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreatePost{
				Creator: "invalid_address",
				Title:   "title",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreatePost{
				Creator: sample.AccAddress(),
				Title:   "title",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
