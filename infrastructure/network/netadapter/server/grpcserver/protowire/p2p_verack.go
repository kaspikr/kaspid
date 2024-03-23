package protowire

import (
	"github.com/kaspikr/kaspid/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KaspidMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KaspidMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *KaspidMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
