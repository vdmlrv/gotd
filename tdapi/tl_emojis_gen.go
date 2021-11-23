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

// Emojis represents TL type `emojis#77274a16`.
type Emojis struct {
	// List of emojis
	Emojis []string
}

// EmojisTypeID is TL type id of Emojis.
const EmojisTypeID = 0x77274a16

// Ensuring interfaces in compile-time for Emojis.
var (
	_ bin.Encoder     = &Emojis{}
	_ bin.Decoder     = &Emojis{}
	_ bin.BareEncoder = &Emojis{}
	_ bin.BareDecoder = &Emojis{}
)

func (e *Emojis) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Emojis == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *Emojis) String() string {
	if e == nil {
		return "Emojis(nil)"
	}
	type Alias Emojis
	return fmt.Sprintf("Emojis%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Emojis) TypeID() uint32 {
	return EmojisTypeID
}

// TypeName returns name of type in TL schema.
func (*Emojis) TypeName() string {
	return "emojis"
}

// TypeInfo returns info about TL type.
func (e *Emojis) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojis",
		ID:   EmojisTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Emojis",
			SchemaName: "emojis",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *Emojis) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojis#77274a16 as nil")
	}
	b.PutID(EmojisTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *Emojis) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojis#77274a16 as nil")
	}
	b.PutInt(len(e.Emojis))
	for _, v := range e.Emojis {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *Emojis) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojis#77274a16 to nil")
	}
	if err := b.ConsumeID(EmojisTypeID); err != nil {
		return fmt.Errorf("unable to decode emojis#77274a16: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *Emojis) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojis#77274a16 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode emojis#77274a16: field emojis: %w", err)
		}

		if headerLen > 0 {
			e.Emojis = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojis#77274a16: field emojis: %w", err)
			}
			e.Emojis = append(e.Emojis, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (e *Emojis) EncodeTDLibJSON(b jsontd.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode emojis#77274a16 as nil")
	}
	b.ObjStart()
	b.PutID("emojis")
	b.FieldStart("emojis")
	b.ArrStart()
	for _, v := range e.Emojis {
		b.PutString(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (e *Emojis) DecodeTDLibJSON(b jsontd.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode emojis#77274a16 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("emojis"); err != nil {
				return fmt.Errorf("unable to decode emojis#77274a16: %w", err)
			}
		case "emojis":
			if err := b.Arr(func(b jsontd.Decoder) error {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode emojis#77274a16: field emojis: %w", err)
				}
				e.Emojis = append(e.Emojis, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode emojis#77274a16: field emojis: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEmojis returns value of Emojis field.
func (e *Emojis) GetEmojis() (value []string) {
	return e.Emojis
}
