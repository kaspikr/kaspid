package rpchandlers

import (
	"github.com/kaspikr/kaspid/app/appmessage"
	"github.com/kaspikr/kaspid/app/rpc/rpccontext"
	"github.com/kaspikr/kaspid/infrastructure/network/netadapter/router"
)

// HandleGetSubnetwork handles the respectively named RPC command
func HandleGetSubnetwork(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetSubnetworkResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
