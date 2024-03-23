package protowire

import (
	"github.com/kaspikr/kaspid/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KaspidMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KaspidMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *KaspidMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
