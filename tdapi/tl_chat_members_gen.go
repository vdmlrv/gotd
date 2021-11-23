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

// ChatMembers represents TL type `chatMembers#8ecbb24`.
type ChatMembers struct {
	// Approximate total count of chat members found
	TotalCount int32
	// A list of chat members
	Members []ChatMember
}

// ChatMembersTypeID is TL type id of ChatMembers.
const ChatMembersTypeID = 0x8ecbb24

// Ensuring interfaces in compile-time for ChatMembers.
var (
	_ bin.Encoder     = &ChatMembers{}
	_ bin.Decoder     = &ChatMembers{}
	_ bin.BareEncoder = &ChatMembers{}
	_ bin.BareDecoder = &ChatMembers{}
)

func (c *ChatMembers) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.TotalCount == 0) {
		return false
	}
	if !(c.Members == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMembers) String() string {
	if c == nil {
		return "ChatMembers(nil)"
	}
	type Alias ChatMembers
	return fmt.Sprintf("ChatMembers%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMembers) TypeID() uint32 {
	return ChatMembersTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMembers) TypeName() string {
	return "chatMembers"
}

// TypeInfo returns info about TL type.
func (c *ChatMembers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMembers",
		ID:   ChatMembersTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Members",
			SchemaName: "members",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMembers) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembers#8ecbb24 as nil")
	}
	b.PutID(ChatMembersTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMembers) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembers#8ecbb24 as nil")
	}
	b.PutInt32(c.TotalCount)
	b.PutInt(len(c.Members))
	for idx, v := range c.Members {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatMembers#8ecbb24: field members element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMembers) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembers#8ecbb24 to nil")
	}
	if err := b.ConsumeID(ChatMembersTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMembers#8ecbb24: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMembers) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembers#8ecbb24 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatMembers#8ecbb24: field total_count: %w", err)
		}
		c.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatMembers#8ecbb24: field members: %w", err)
		}

		if headerLen > 0 {
			c.Members = make([]ChatMember, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatMember
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatMembers#8ecbb24: field members: %w", err)
			}
			c.Members = append(c.Members, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (c *ChatMembers) EncodeTDLibJSON(b jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembers#8ecbb24 as nil")
	}
	b.ObjStart()
	b.PutID("chatMembers")
	b.FieldStart("total_count")
	b.PutInt32(c.TotalCount)
	b.FieldStart("members")
	b.ArrStart()
	for idx, v := range c.Members {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatMembers#8ecbb24: field members element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (c *ChatMembers) DecodeTDLibJSON(b jsontd.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembers#8ecbb24 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("chatMembers"); err != nil {
				return fmt.Errorf("unable to decode chatMembers#8ecbb24: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatMembers#8ecbb24: field total_count: %w", err)
			}
			c.TotalCount = value
		case "members":
			if err := b.Arr(func(b jsontd.Decoder) error {
				var value ChatMember
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode chatMembers#8ecbb24: field members: %w", err)
				}
				c.Members = append(c.Members, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatMembers#8ecbb24: field members: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (c *ChatMembers) GetTotalCount() (value int32) {
	return c.TotalCount
}

// GetMembers returns value of Members field.
func (c *ChatMembers) GetMembers() (value []ChatMember) {
	return c.Members
}
