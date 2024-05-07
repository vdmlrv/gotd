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

// ReaddQuickReplyShortcutMessagesRequest represents TL type `readdQuickReplyShortcutMessages#652518f7`.
type ReaddQuickReplyShortcutMessagesRequest struct {
	// Name of the target shortcut
	ShortcutName string
	// Identifiers of the quick reply messages to readd. Message identifiers must be in a
	// strictly increasing order
	MessageIDs []int64
}

// ReaddQuickReplyShortcutMessagesRequestTypeID is TL type id of ReaddQuickReplyShortcutMessagesRequest.
const ReaddQuickReplyShortcutMessagesRequestTypeID = 0x652518f7

// Ensuring interfaces in compile-time for ReaddQuickReplyShortcutMessagesRequest.
var (
	_ bin.Encoder     = &ReaddQuickReplyShortcutMessagesRequest{}
	_ bin.Decoder     = &ReaddQuickReplyShortcutMessagesRequest{}
	_ bin.BareEncoder = &ReaddQuickReplyShortcutMessagesRequest{}
	_ bin.BareDecoder = &ReaddQuickReplyShortcutMessagesRequest{}
)

func (r *ReaddQuickReplyShortcutMessagesRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ShortcutName == "") {
		return false
	}
	if !(r.MessageIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReaddQuickReplyShortcutMessagesRequest) String() string {
	if r == nil {
		return "ReaddQuickReplyShortcutMessagesRequest(nil)"
	}
	type Alias ReaddQuickReplyShortcutMessagesRequest
	return fmt.Sprintf("ReaddQuickReplyShortcutMessagesRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReaddQuickReplyShortcutMessagesRequest) TypeID() uint32 {
	return ReaddQuickReplyShortcutMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReaddQuickReplyShortcutMessagesRequest) TypeName() string {
	return "readdQuickReplyShortcutMessages"
}

// TypeInfo returns info about TL type.
func (r *ReaddQuickReplyShortcutMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "readdQuickReplyShortcutMessages",
		ID:   ReaddQuickReplyShortcutMessagesRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortcutName",
			SchemaName: "shortcut_name",
		},
		{
			Name:       "MessageIDs",
			SchemaName: "message_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReaddQuickReplyShortcutMessagesRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readdQuickReplyShortcutMessages#652518f7 as nil")
	}
	b.PutID(ReaddQuickReplyShortcutMessagesRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReaddQuickReplyShortcutMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readdQuickReplyShortcutMessages#652518f7 as nil")
	}
	b.PutString(r.ShortcutName)
	b.PutInt(len(r.MessageIDs))
	for _, v := range r.MessageIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReaddQuickReplyShortcutMessagesRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readdQuickReplyShortcutMessages#652518f7 to nil")
	}
	if err := b.ConsumeID(ReaddQuickReplyShortcutMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode readdQuickReplyShortcutMessages#652518f7: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReaddQuickReplyShortcutMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readdQuickReplyShortcutMessages#652518f7 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode readdQuickReplyShortcutMessages#652518f7: field shortcut_name: %w", err)
		}
		r.ShortcutName = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode readdQuickReplyShortcutMessages#652518f7: field message_ids: %w", err)
		}

		if headerLen > 0 {
			r.MessageIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode readdQuickReplyShortcutMessages#652518f7: field message_ids: %w", err)
			}
			r.MessageIDs = append(r.MessageIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReaddQuickReplyShortcutMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode readdQuickReplyShortcutMessages#652518f7 as nil")
	}
	b.ObjStart()
	b.PutID("readdQuickReplyShortcutMessages")
	b.Comma()
	b.FieldStart("shortcut_name")
	b.PutString(r.ShortcutName)
	b.Comma()
	b.FieldStart("message_ids")
	b.ArrStart()
	for _, v := range r.MessageIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReaddQuickReplyShortcutMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode readdQuickReplyShortcutMessages#652518f7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("readdQuickReplyShortcutMessages"); err != nil {
				return fmt.Errorf("unable to decode readdQuickReplyShortcutMessages#652518f7: %w", err)
			}
		case "shortcut_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode readdQuickReplyShortcutMessages#652518f7: field shortcut_name: %w", err)
			}
			r.ShortcutName = value
		case "message_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode readdQuickReplyShortcutMessages#652518f7: field message_ids: %w", err)
				}
				r.MessageIDs = append(r.MessageIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode readdQuickReplyShortcutMessages#652518f7: field message_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetShortcutName returns value of ShortcutName field.
func (r *ReaddQuickReplyShortcutMessagesRequest) GetShortcutName() (value string) {
	if r == nil {
		return
	}
	return r.ShortcutName
}

// GetMessageIDs returns value of MessageIDs field.
func (r *ReaddQuickReplyShortcutMessagesRequest) GetMessageIDs() (value []int64) {
	if r == nil {
		return
	}
	return r.MessageIDs
}

// ReaddQuickReplyShortcutMessages invokes method readdQuickReplyShortcutMessages#652518f7 returning error if any.
func (c *Client) ReaddQuickReplyShortcutMessages(ctx context.Context, request *ReaddQuickReplyShortcutMessagesRequest) (*QuickReplyMessages, error) {
	var result QuickReplyMessages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}