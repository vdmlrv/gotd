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

// ChatAffiliatePrograms represents TL type `chatAffiliatePrograms#74f423e2`.
type ChatAffiliatePrograms struct {
	// The total number of affiliate programs that were connected to the chat
	TotalCount int32
	// The list of connected affiliate programs
	Programs []ChatAffiliateProgram
	// The offset for the next request. If empty, then there are no more results
	NextOffset string
}

// ChatAffiliateProgramsTypeID is TL type id of ChatAffiliatePrograms.
const ChatAffiliateProgramsTypeID = 0x74f423e2

// Ensuring interfaces in compile-time for ChatAffiliatePrograms.
var (
	_ bin.Encoder     = &ChatAffiliatePrograms{}
	_ bin.Decoder     = &ChatAffiliatePrograms{}
	_ bin.BareEncoder = &ChatAffiliatePrograms{}
	_ bin.BareDecoder = &ChatAffiliatePrograms{}
)

func (c *ChatAffiliatePrograms) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.TotalCount == 0) {
		return false
	}
	if !(c.Programs == nil) {
		return false
	}
	if !(c.NextOffset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatAffiliatePrograms) String() string {
	if c == nil {
		return "ChatAffiliatePrograms(nil)"
	}
	type Alias ChatAffiliatePrograms
	return fmt.Sprintf("ChatAffiliatePrograms%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatAffiliatePrograms) TypeID() uint32 {
	return ChatAffiliateProgramsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatAffiliatePrograms) TypeName() string {
	return "chatAffiliatePrograms"
}

// TypeInfo returns info about TL type.
func (c *ChatAffiliatePrograms) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatAffiliatePrograms",
		ID:   ChatAffiliateProgramsTypeID,
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
			Name:       "Programs",
			SchemaName: "programs",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatAffiliatePrograms) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAffiliatePrograms#74f423e2 as nil")
	}
	b.PutID(ChatAffiliateProgramsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatAffiliatePrograms) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAffiliatePrograms#74f423e2 as nil")
	}
	b.PutInt32(c.TotalCount)
	b.PutInt(len(c.Programs))
	for idx, v := range c.Programs {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatAffiliatePrograms#74f423e2: field programs element with index %d: %w", idx, err)
		}
	}
	b.PutString(c.NextOffset)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatAffiliatePrograms) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAffiliatePrograms#74f423e2 to nil")
	}
	if err := b.ConsumeID(ChatAffiliateProgramsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAffiliatePrograms#74f423e2: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatAffiliatePrograms) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAffiliatePrograms#74f423e2 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatAffiliatePrograms#74f423e2: field total_count: %w", err)
		}
		c.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatAffiliatePrograms#74f423e2: field programs: %w", err)
		}

		if headerLen > 0 {
			c.Programs = make([]ChatAffiliateProgram, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatAffiliateProgram
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatAffiliatePrograms#74f423e2: field programs: %w", err)
			}
			c.Programs = append(c.Programs, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatAffiliatePrograms#74f423e2: field next_offset: %w", err)
		}
		c.NextOffset = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatAffiliatePrograms) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAffiliatePrograms#74f423e2 as nil")
	}
	b.ObjStart()
	b.PutID("chatAffiliatePrograms")
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(c.TotalCount)
	b.Comma()
	b.FieldStart("programs")
	b.ArrStart()
	for idx, v := range c.Programs {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatAffiliatePrograms#74f423e2: field programs element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("next_offset")
	b.PutString(c.NextOffset)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatAffiliatePrograms) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAffiliatePrograms#74f423e2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatAffiliatePrograms"); err != nil {
				return fmt.Errorf("unable to decode chatAffiliatePrograms#74f423e2: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatAffiliatePrograms#74f423e2: field total_count: %w", err)
			}
			c.TotalCount = value
		case "programs":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ChatAffiliateProgram
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode chatAffiliatePrograms#74f423e2: field programs: %w", err)
				}
				c.Programs = append(c.Programs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatAffiliatePrograms#74f423e2: field programs: %w", err)
			}
		case "next_offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatAffiliatePrograms#74f423e2: field next_offset: %w", err)
			}
			c.NextOffset = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (c *ChatAffiliatePrograms) GetTotalCount() (value int32) {
	if c == nil {
		return
	}
	return c.TotalCount
}

// GetPrograms returns value of Programs field.
func (c *ChatAffiliatePrograms) GetPrograms() (value []ChatAffiliateProgram) {
	if c == nil {
		return
	}
	return c.Programs
}

// GetNextOffset returns value of NextOffset field.
func (c *ChatAffiliatePrograms) GetNextOffset() (value string) {
	if c == nil {
		return
	}
	return c.NextOffset
}
