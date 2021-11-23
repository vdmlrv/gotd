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

// Call represents TL type `call#59a64c86`.
type Call struct {
	// Call identifier, not persistent
	ID int32
	// Peer user identifier
	UserID int32
	// True, if the call is outgoing
	IsOutgoing bool
	// True, if the call is a video call
	IsVideo bool
	// Call state
	State CallStateClass
}

// CallTypeID is TL type id of Call.
const CallTypeID = 0x59a64c86

// Ensuring interfaces in compile-time for Call.
var (
	_ bin.Encoder     = &Call{}
	_ bin.Decoder     = &Call{}
	_ bin.BareEncoder = &Call{}
	_ bin.BareDecoder = &Call{}
)

func (c *Call) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ID == 0) {
		return false
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.IsOutgoing == false) {
		return false
	}
	if !(c.IsVideo == false) {
		return false
	}
	if !(c.State == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *Call) String() string {
	if c == nil {
		return "Call(nil)"
	}
	type Alias Call
	return fmt.Sprintf("Call%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Call) TypeID() uint32 {
	return CallTypeID
}

// TypeName returns name of type in TL schema.
func (*Call) TypeName() string {
	return "call"
}

// TypeInfo returns info about TL type.
func (c *Call) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "call",
		ID:   CallTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "IsOutgoing",
			SchemaName: "is_outgoing",
		},
		{
			Name:       "IsVideo",
			SchemaName: "is_video",
		},
		{
			Name:       "State",
			SchemaName: "state",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *Call) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode call#59a64c86 as nil")
	}
	b.PutID(CallTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *Call) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode call#59a64c86 as nil")
	}
	b.PutInt32(c.ID)
	b.PutInt32(c.UserID)
	b.PutBool(c.IsOutgoing)
	b.PutBool(c.IsVideo)
	if c.State == nil {
		return fmt.Errorf("unable to encode call#59a64c86: field state is nil")
	}
	if err := c.State.Encode(b); err != nil {
		return fmt.Errorf("unable to encode call#59a64c86: field state: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *Call) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode call#59a64c86 to nil")
	}
	if err := b.ConsumeID(CallTypeID); err != nil {
		return fmt.Errorf("unable to decode call#59a64c86: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *Call) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode call#59a64c86 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode call#59a64c86: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode call#59a64c86: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode call#59a64c86: field is_outgoing: %w", err)
		}
		c.IsOutgoing = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode call#59a64c86: field is_video: %w", err)
		}
		c.IsVideo = value
	}
	{
		value, err := DecodeCallState(b)
		if err != nil {
			return fmt.Errorf("unable to decode call#59a64c86: field state: %w", err)
		}
		c.State = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (c *Call) EncodeTDLibJSON(b jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode call#59a64c86 as nil")
	}
	b.ObjStart()
	b.PutID("call")
	b.FieldStart("id")
	b.PutInt32(c.ID)
	b.FieldStart("user_id")
	b.PutInt32(c.UserID)
	b.FieldStart("is_outgoing")
	b.PutBool(c.IsOutgoing)
	b.FieldStart("is_video")
	b.PutBool(c.IsVideo)
	b.FieldStart("state")
	if c.State == nil {
		return fmt.Errorf("unable to encode call#59a64c86: field state is nil")
	}
	if err := c.State.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode call#59a64c86: field state: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (c *Call) DecodeTDLibJSON(b jsontd.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode call#59a64c86 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("call"); err != nil {
				return fmt.Errorf("unable to decode call#59a64c86: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode call#59a64c86: field id: %w", err)
			}
			c.ID = value
		case "user_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode call#59a64c86: field user_id: %w", err)
			}
			c.UserID = value
		case "is_outgoing":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode call#59a64c86: field is_outgoing: %w", err)
			}
			c.IsOutgoing = value
		case "is_video":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode call#59a64c86: field is_video: %w", err)
			}
			c.IsVideo = value
		case "state":
			value, err := DecodeTDLibJSONCallState(b)
			if err != nil {
				return fmt.Errorf("unable to decode call#59a64c86: field state: %w", err)
			}
			c.State = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (c *Call) GetID() (value int32) {
	return c.ID
}

// GetUserID returns value of UserID field.
func (c *Call) GetUserID() (value int32) {
	return c.UserID
}

// GetIsOutgoing returns value of IsOutgoing field.
func (c *Call) GetIsOutgoing() (value bool) {
	return c.IsOutgoing
}

// GetIsVideo returns value of IsVideo field.
func (c *Call) GetIsVideo() (value bool) {
	return c.IsVideo
}

// GetState returns value of State field.
func (c *Call) GetState() (value CallStateClass) {
	return c.State
}
