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

// ClosedVectorPath represents TL type `closedVectorPath#2f9276b9`.
type ClosedVectorPath struct {
	// List of vector path commands
	Commands []VectorPathCommandClass
}

// ClosedVectorPathTypeID is TL type id of ClosedVectorPath.
const ClosedVectorPathTypeID = 0x2f9276b9

// Ensuring interfaces in compile-time for ClosedVectorPath.
var (
	_ bin.Encoder     = &ClosedVectorPath{}
	_ bin.Decoder     = &ClosedVectorPath{}
	_ bin.BareEncoder = &ClosedVectorPath{}
	_ bin.BareDecoder = &ClosedVectorPath{}
)

func (c *ClosedVectorPath) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Commands == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ClosedVectorPath) String() string {
	if c == nil {
		return "ClosedVectorPath(nil)"
	}
	type Alias ClosedVectorPath
	return fmt.Sprintf("ClosedVectorPath%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ClosedVectorPath) TypeID() uint32 {
	return ClosedVectorPathTypeID
}

// TypeName returns name of type in TL schema.
func (*ClosedVectorPath) TypeName() string {
	return "closedVectorPath"
}

// TypeInfo returns info about TL type.
func (c *ClosedVectorPath) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "closedVectorPath",
		ID:   ClosedVectorPathTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Commands",
			SchemaName: "commands",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ClosedVectorPath) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode closedVectorPath#2f9276b9 as nil")
	}
	b.PutID(ClosedVectorPathTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ClosedVectorPath) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode closedVectorPath#2f9276b9 as nil")
	}
	b.PutInt(len(c.Commands))
	for idx, v := range c.Commands {
		if v == nil {
			return fmt.Errorf("unable to encode closedVectorPath#2f9276b9: field commands element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare closedVectorPath#2f9276b9: field commands element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ClosedVectorPath) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode closedVectorPath#2f9276b9 to nil")
	}
	if err := b.ConsumeID(ClosedVectorPathTypeID); err != nil {
		return fmt.Errorf("unable to decode closedVectorPath#2f9276b9: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ClosedVectorPath) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode closedVectorPath#2f9276b9 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode closedVectorPath#2f9276b9: field commands: %w", err)
		}

		if headerLen > 0 {
			c.Commands = make([]VectorPathCommandClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeVectorPathCommand(b)
			if err != nil {
				return fmt.Errorf("unable to decode closedVectorPath#2f9276b9: field commands: %w", err)
			}
			c.Commands = append(c.Commands, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (c *ClosedVectorPath) EncodeTDLibJSON(b jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode closedVectorPath#2f9276b9 as nil")
	}
	b.ObjStart()
	b.PutID("closedVectorPath")
	b.FieldStart("commands")
	b.ArrStart()
	for idx, v := range c.Commands {
		if v == nil {
			return fmt.Errorf("unable to encode closedVectorPath#2f9276b9: field commands element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode closedVectorPath#2f9276b9: field commands element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (c *ClosedVectorPath) DecodeTDLibJSON(b jsontd.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode closedVectorPath#2f9276b9 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("closedVectorPath"); err != nil {
				return fmt.Errorf("unable to decode closedVectorPath#2f9276b9: %w", err)
			}
		case "commands":
			if err := b.Arr(func(b jsontd.Decoder) error {
				value, err := DecodeTDLibJSONVectorPathCommand(b)
				if err != nil {
					return fmt.Errorf("unable to decode closedVectorPath#2f9276b9: field commands: %w", err)
				}
				c.Commands = append(c.Commands, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode closedVectorPath#2f9276b9: field commands: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCommands returns value of Commands field.
func (c *ClosedVectorPath) GetCommands() (value []VectorPathCommandClass) {
	return c.Commands
}
