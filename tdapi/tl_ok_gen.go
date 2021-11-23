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

// Ok represents TL type `ok#d4edbe69`.
type Ok struct {
}

// OkTypeID is TL type id of Ok.
const OkTypeID = 0xd4edbe69

// Ensuring interfaces in compile-time for Ok.
var (
	_ bin.Encoder     = &Ok{}
	_ bin.Decoder     = &Ok{}
	_ bin.BareEncoder = &Ok{}
	_ bin.BareDecoder = &Ok{}
)

func (o *Ok) Zero() bool {
	if o == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (o *Ok) String() string {
	if o == nil {
		return "Ok(nil)"
	}
	type Alias Ok
	return fmt.Sprintf("Ok%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Ok) TypeID() uint32 {
	return OkTypeID
}

// TypeName returns name of type in TL schema.
func (*Ok) TypeName() string {
	return "ok"
}

// TypeInfo returns info about TL type.
func (o *Ok) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "ok",
		ID:   OkTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (o *Ok) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode ok#d4edbe69 as nil")
	}
	b.PutID(OkTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *Ok) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode ok#d4edbe69 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (o *Ok) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode ok#d4edbe69 to nil")
	}
	if err := b.ConsumeID(OkTypeID); err != nil {
		return fmt.Errorf("unable to decode ok#d4edbe69: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *Ok) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode ok#d4edbe69 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (o *Ok) EncodeTDLibJSON(b jsontd.Encoder) error {
	if o == nil {
		return fmt.Errorf("can't encode ok#d4edbe69 as nil")
	}
	b.ObjStart()
	b.PutID("ok")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (o *Ok) DecodeTDLibJSON(b jsontd.Decoder) error {
	if o == nil {
		return fmt.Errorf("can't decode ok#d4edbe69 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("ok"); err != nil {
				return fmt.Errorf("unable to decode ok#d4edbe69: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}
