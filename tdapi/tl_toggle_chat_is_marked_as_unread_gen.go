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

// ToggleChatIsMarkedAsUnreadRequest represents TL type `toggleChatIsMarkedAsUnread#c538dadf`.
type ToggleChatIsMarkedAsUnreadRequest struct {
	// Chat identifier
	ChatID int64
	// New value of is_marked_as_unread
	IsMarkedAsUnread bool
}

// ToggleChatIsMarkedAsUnreadRequestTypeID is TL type id of ToggleChatIsMarkedAsUnreadRequest.
const ToggleChatIsMarkedAsUnreadRequestTypeID = 0xc538dadf

// Ensuring interfaces in compile-time for ToggleChatIsMarkedAsUnreadRequest.
var (
	_ bin.Encoder     = &ToggleChatIsMarkedAsUnreadRequest{}
	_ bin.Decoder     = &ToggleChatIsMarkedAsUnreadRequest{}
	_ bin.BareEncoder = &ToggleChatIsMarkedAsUnreadRequest{}
	_ bin.BareDecoder = &ToggleChatIsMarkedAsUnreadRequest{}
)

func (t *ToggleChatIsMarkedAsUnreadRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ChatID == 0) {
		return false
	}
	if !(t.IsMarkedAsUnread == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleChatIsMarkedAsUnreadRequest) String() string {
	if t == nil {
		return "ToggleChatIsMarkedAsUnreadRequest(nil)"
	}
	type Alias ToggleChatIsMarkedAsUnreadRequest
	return fmt.Sprintf("ToggleChatIsMarkedAsUnreadRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleChatIsMarkedAsUnreadRequest) TypeID() uint32 {
	return ToggleChatIsMarkedAsUnreadRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleChatIsMarkedAsUnreadRequest) TypeName() string {
	return "toggleChatIsMarkedAsUnread"
}

// TypeInfo returns info about TL type.
func (t *ToggleChatIsMarkedAsUnreadRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleChatIsMarkedAsUnread",
		ID:   ToggleChatIsMarkedAsUnreadRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "IsMarkedAsUnread",
			SchemaName: "is_marked_as_unread",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleChatIsMarkedAsUnreadRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatIsMarkedAsUnread#c538dadf as nil")
	}
	b.PutID(ToggleChatIsMarkedAsUnreadRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleChatIsMarkedAsUnreadRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatIsMarkedAsUnread#c538dadf as nil")
	}
	b.PutLong(t.ChatID)
	b.PutBool(t.IsMarkedAsUnread)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleChatIsMarkedAsUnreadRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatIsMarkedAsUnread#c538dadf to nil")
	}
	if err := b.ConsumeID(ToggleChatIsMarkedAsUnreadRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleChatIsMarkedAsUnread#c538dadf: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleChatIsMarkedAsUnreadRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatIsMarkedAsUnread#c538dadf to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatIsMarkedAsUnread#c538dadf: field chat_id: %w", err)
		}
		t.ChatID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatIsMarkedAsUnread#c538dadf: field is_marked_as_unread: %w", err)
		}
		t.IsMarkedAsUnread = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (t *ToggleChatIsMarkedAsUnreadRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatIsMarkedAsUnread#c538dadf as nil")
	}
	b.ObjStart()
	b.PutID("toggleChatIsMarkedAsUnread")
	b.FieldStart("chat_id")
	b.PutLong(t.ChatID)
	b.FieldStart("is_marked_as_unread")
	b.PutBool(t.IsMarkedAsUnread)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (t *ToggleChatIsMarkedAsUnreadRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatIsMarkedAsUnread#c538dadf to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("toggleChatIsMarkedAsUnread"); err != nil {
				return fmt.Errorf("unable to decode toggleChatIsMarkedAsUnread#c538dadf: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatIsMarkedAsUnread#c538dadf: field chat_id: %w", err)
			}
			t.ChatID = value
		case "is_marked_as_unread":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatIsMarkedAsUnread#c538dadf: field is_marked_as_unread: %w", err)
			}
			t.IsMarkedAsUnread = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (t *ToggleChatIsMarkedAsUnreadRequest) GetChatID() (value int64) {
	return t.ChatID
}

// GetIsMarkedAsUnread returns value of IsMarkedAsUnread field.
func (t *ToggleChatIsMarkedAsUnreadRequest) GetIsMarkedAsUnread() (value bool) {
	return t.IsMarkedAsUnread
}

// ToggleChatIsMarkedAsUnread invokes method toggleChatIsMarkedAsUnread#c538dadf returning error if any.
func (c *Client) ToggleChatIsMarkedAsUnread(ctx context.Context, request *ToggleChatIsMarkedAsUnreadRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
