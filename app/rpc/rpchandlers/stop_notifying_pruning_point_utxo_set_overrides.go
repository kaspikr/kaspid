package rpchandlers

import (
	"github.com/kaspikr/kaspid/app/appmessage"
	"github.com/kaspikr/kaspid/app/rpc/rpccontext"
	"github.com/kaspikr/kaspid/infrastructure/network/netadapter/router"
)

// HandleStopNotifyingPruningPointUTXOSetOverrideRequest handles the respectively named RPC command
func HandleStopNotifyingPruningPointUTXOSetOverrideRequest(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.StopPropagatingPruningPointUTXOSetOverrideNotifications()

	response := appmessage.NewStopNotifyingPruningPointUTXOSetOverrideResponseMessage()
	return response, nil
}
