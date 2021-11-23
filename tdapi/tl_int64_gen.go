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

// Int64 represents TL type `int64#5d9ed744`.
type Int64 struct {
}

// Int64TypeID is TL type id of Int64.
const Int64TypeID = 0x5d9ed744

// Ensuring interfaces in compile-time for Int64.
var (
	_ bin.Encoder     = &Int64{}
	_ bin.Decoder     = &Int64{}
	_ bin.BareEncoder = &Int64{}
	_ bin.BareDecoder = &Int64{}
)

func (i *Int64) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *Int64) String() string {
	if i == nil {
		return "Int64(nil)"
	}
	type Alias Int64
	return fmt.Sprintf("Int64%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Int64) TypeID() uint32 {
	return Int64TypeID
}

// TypeName returns name of type in TL schema.
func (*Int64) TypeName() string {
	return "int64"
}

// TypeInfo returns info about TL type.
func (i *Int64) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "int64",
		ID:   Int64TypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *Int64) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode int64#5d9ed744 as nil")
	}
	b.PutID(Int64TypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *Int64) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode int64#5d9ed744 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *Int64) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode int64#5d9ed744 to nil")
	}
	if err := b.ConsumeID(Int64TypeID); err != nil {
		return fmt.Errorf("unable to decode int64#5d9ed744: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *Int64) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode int64#5d9ed744 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (i *Int64) EncodeTDLibJSON(b jsontd.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode int64#5d9ed744 as nil")
	}
	b.ObjStart()
	b.PutID("int64")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (i *Int64) DecodeTDLibJSON(b jsontd.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode int64#5d9ed744 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("int64"); err != nil {
				return fmt.Errorf("unable to decode int64#5d9ed744: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}
