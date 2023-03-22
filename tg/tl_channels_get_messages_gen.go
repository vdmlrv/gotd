// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// ChannelsGetMessagesRequest represents TL type `channels.getMessages#ad8c9a23`.
// Get channel/supergroup¹ messages
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getMessages for reference.
type ChannelsGetMessagesRequest struct {
	// Channel/supergroup
	Channel InputChannelClass
	// IDs of messages to get
	ID []InputMessageClass
}

// ChannelsGetMessagesRequestTypeID is TL type id of ChannelsGetMessagesRequest.
const ChannelsGetMessagesRequestTypeID = 0xad8c9a23

// Ensuring interfaces in compile-time for ChannelsGetMessagesRequest.
var (
	_ bin.Encoder     = &ChannelsGetMessagesRequest{}
	_ bin.Decoder     = &ChannelsGetMessagesRequest{}
	_ bin.BareEncoder = &ChannelsGetMessagesRequest{}
	_ bin.BareDecoder = &ChannelsGetMessagesRequest{}
)

func (g *ChannelsGetMessagesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetMessagesRequest) String() string {
	if g == nil {
		return "ChannelsGetMessagesRequest(nil)"
	}
	type Alias ChannelsGetMessagesRequest
	return fmt.Sprintf("ChannelsGetMessagesRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetMessagesRequest from given interface.
func (g *ChannelsGetMessagesRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetID() (value []InputMessageClass)
}) {
	g.Channel = from.GetChannel()
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetMessagesRequest) TypeID() uint32 {
	return ChannelsGetMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetMessagesRequest) TypeName() string {
	return "channels.getMessages"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getMessages",
		ID:   ChannelsGetMessagesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChannelsGetMessagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getMessages#ad8c9a23 as nil")
	}
	b.PutID(ChannelsGetMessagesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getMessages#ad8c9a23 as nil")
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getMessages#ad8c9a23: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getMessages#ad8c9a23: field channel: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for idx, v := range g.ID {
		if v == nil {
			return fmt.Errorf("unable to encode channels.getMessages#ad8c9a23: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.getMessages#ad8c9a23: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetMessagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getMessages#ad8c9a23 to nil")
	}
	if err := b.ConsumeID(ChannelsGetMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getMessages#ad8c9a23: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getMessages#ad8c9a23 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getMessages#ad8c9a23: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getMessages#ad8c9a23: field id: %w", err)
		}

		if headerLen > 0 {
			g.ID = make([]InputMessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.getMessages#ad8c9a23: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// GetChannel returns value of Channel field.
func (g *ChannelsGetMessagesRequest) GetChannel() (value InputChannelClass) {
	if g == nil {
		return
	}
	return g.Channel
}

// GetID returns value of ID field.
func (g *ChannelsGetMessagesRequest) GetID() (value []InputMessageClass) {
	if g == nil {
		return
	}
	return g.ID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (g *ChannelsGetMessagesRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return g.Channel.AsNotEmpty()
}

// MapID returns field ID wrapped in InputMessageClassArray helper.
func (g *ChannelsGetMessagesRequest) MapID() (value InputMessageClassArray) {
	return InputMessageClassArray(g.ID)
}

// ChannelsGetMessages invokes method channels.getMessages#ad8c9a23 returning error if any.
// Get channel/supergroup¹ messages
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	406 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 MESSAGE_IDS_EMPTY: No message ids were provided.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	500 NEED_DOC_INVALID:
//	400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels.
//
// See https://core.telegram.org/method/channels.getMessages for reference.
// Can be used by bots.
func (c *Client) ChannelsGetMessages(ctx context.Context, request *ChannelsGetMessagesRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
