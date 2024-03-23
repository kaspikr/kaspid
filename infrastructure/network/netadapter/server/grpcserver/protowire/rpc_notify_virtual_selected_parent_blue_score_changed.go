package protowire

import (
	"github.com/kaspikr/kaspid/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KaspidMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KaspidMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest is nil")
	}
	return &appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage{}, nil
}

func (x *KaspidMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest) fromAppMessage(_ *appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage) error {
	x.NotifyVirtualSelectedParentBlueScoreChangedRequest = &NotifyVirtualSelectedParentBlueScoreChangedRequestMessage{}
	return nil
}

func (x *KaspidMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KaspidMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse is nil")
	}
	return x.NotifyVirtualSelectedParentBlueScoreChangedResponse.toAppMessage()
}

func (x *KaspidMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse) fromAppMessage(message *appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.NotifyVirtualSelectedParentBlueScoreChangedResponse = &NotifyVirtualSelectedParentBlueScoreChangedResponseMessage{
		Error: err,
	}
	return nil
}

func (x *NotifyVirtualSelectedParentBlueScoreChangedResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NotifyVirtualSelectedParentBlueScoreChangedResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *KaspidMessage_VirtualSelectedParentBlueScoreChangedNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KaspidMessage_VirtualSelectedParentBlueScoreChangedNotification is nil")
	}
	return x.VirtualSelectedParentBlueScoreChangedNotification.toAppMessage()
}

func (x *KaspidMessage_VirtualSelectedParentBlueScoreChangedNotification) fromAppMessage(message *appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage) error {
	x.VirtualSelectedParentBlueScoreChangedNotification = &VirtualSelectedParentBlueScoreChangedNotificationMessage{
		VirtualSelectedParentBlueScore: message.VirtualSelectedParentBlueScore,
	}
	return nil
}

func (x *VirtualSelectedParentBlueScoreChangedNotificationMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "VirtualSelectedParentBlueScoreChangedNotificationMessage is nil")
	}
	return &appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage{
		VirtualSelectedParentBlueScore: x.VirtualSelectedParentBlueScore,
	}, nil
}
