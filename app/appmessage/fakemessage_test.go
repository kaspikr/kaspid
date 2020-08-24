// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

import "io"

// fakeMessage implements the Message interface and is used to force encode
// errors in messages.
type fakeMessage struct {
	command        MessageCommand
	payload        []byte
	forceEncodeErr bool
	forceLenErr    bool
}

// KaspaDecode doesn't do anything. It just satisfies the appmessage.Message
// interface.
func (msg *fakeMessage) KaspaDecode(r io.Reader, pver uint32) error {
	return nil
}

// KaspaEncode writes the payload field of the fake message or forces an error
// if the forceEncodeErr flag of the fake message is set. It also satisfies the
// appmessage.Message interface.
func (msg *fakeMessage) KaspaEncode(w io.Writer, pver uint32) error {
	if msg.forceEncodeErr {
		err := &MessageError{
			Func:        "fakeMessage.KaspaEncode",
			Description: "intentional error",
		}
		return err
	}

	_, err := w.Write(msg.payload)
	return err
}

// Command returns the command field of the fake message and satisfies the
// Message interface.
func (msg *fakeMessage) Command() MessageCommand {
	return msg.command
}

// MaxPayloadLength returns the length of the payload field of fake message
// or a smaller value if the forceLenErr flag of the fake message is set. It
// satisfies the Message interface.
func (msg *fakeMessage) MaxPayloadLength(pver uint32) uint32 {
	lenp := uint32(len(msg.payload))
	if msg.forceLenErr {
		return lenp - 1
	}

	return lenp
}