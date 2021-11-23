// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/jsontd"
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
	_ = jsontd.Encoder{}
)

// AddChatMemberRequest represents TL type `addChatMember#46805eaa`.
type AddChatMemberRequest struct {
	// Chat identifier
	ChatID int64
	// Identifier of the user
	UserID int32
	// The number of earlier messages from the chat to be forwarded to the new member; up to
	// 100. Ignored for supergroups and channels
	ForwardLimit int32
}

// AddChatMemberRequestTypeID is TL type id of AddChatMemberRequest.
const AddChatMemberRequestTypeID = 0x46805eaa

// Ensuring interfaces in compile-time for AddChatMemberRequest.
var (
	_ bin.Encoder     = &AddChatMemberRequest{}
	_ bin.Decoder     = &AddChatMemberRequest{}
	_ bin.BareEncoder = &AddChatMemberRequest{}
	_ bin.BareDecoder = &AddChatMemberRequest{}
)

func (a *AddChatMemberRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ChatID == 0) {
		return false
	}
	if !(a.UserID == 0) {
		return false
	}
	if !(a.ForwardLimit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddChatMemberRequest) String() string {
	if a == nil {
		return "AddChatMemberRequest(nil)"
	}
	type Alias AddChatMemberRequest
	return fmt.Sprintf("AddChatMemberRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddChatMemberRequest) TypeID() uint32 {
	return AddChatMemberRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddChatMemberRequest) TypeName() string {
	return "addChatMember"
}

// TypeInfo returns info about TL type.
func (a *AddChatMemberRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addChatMember",
		ID:   AddChatMemberRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "ForwardLimit",
			SchemaName: "forward_limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddChatMemberRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatMember#46805eaa as nil")
	}
	b.PutID(AddChatMemberRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddChatMemberRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatMember#46805eaa as nil")
	}
	b.PutLong(a.ChatID)
	b.PutInt32(a.UserID)
	b.PutInt32(a.ForwardLimit)
	return nil
}

// Decode implements bin.Decoder.
func (a *AddChatMemberRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatMember#46805eaa to nil")
	}
	if err := b.ConsumeID(AddChatMemberRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addChatMember#46805eaa: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddChatMemberRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatMember#46805eaa to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode addChatMember#46805eaa: field chat_id: %w", err)
		}
		a.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode addChatMember#46805eaa: field user_id: %w", err)
		}
		a.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode addChatMember#46805eaa: field forward_limit: %w", err)
		}
		a.ForwardLimit = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (a *AddChatMemberRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatMember#46805eaa as nil")
	}
	b.ObjStart()
	b.PutID("addChatMember")
	b.FieldStart("chat_id")
	b.PutLong(a.ChatID)
	b.FieldStart("user_id")
	b.PutInt32(a.UserID)
	b.FieldStart("forward_limit")
	b.PutInt32(a.ForwardLimit)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (a *AddChatMemberRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatMember#46805eaa to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("addChatMember"); err != nil {
				return fmt.Errorf("unable to decode addChatMember#46805eaa: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode addChatMember#46805eaa: field chat_id: %w", err)
			}
			a.ChatID = value
		case "user_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode addChatMember#46805eaa: field user_id: %w", err)
			}
			a.UserID = value
		case "forward_limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode addChatMember#46805eaa: field forward_limit: %w", err)
			}
			a.ForwardLimit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (a *AddChatMemberRequest) GetChatID() (value int64) {
	return a.ChatID
}

// GetUserID returns value of UserID field.
func (a *AddChatMemberRequest) GetUserID() (value int32) {
	return a.UserID
}

// GetForwardLimit returns value of ForwardLimit field.
func (a *AddChatMemberRequest) GetForwardLimit() (value int32) {
	return a.ForwardLimit
}

// AddChatMember invokes method addChatMember#46805eaa returning error if any.
func (c *Client) AddChatMember(ctx context.Context, request *AddChatMemberRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
