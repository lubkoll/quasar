package keeper

import (
	"github.com/abag/quasarnode/x/intergamm/types"
)

type msgServer struct {
	k *Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

var _ types.MsgServer = msgServer{}
