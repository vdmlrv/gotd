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

// Text represents TL type `text#22765898`.
type Text struct {
	// Text
	Text string
}

// TextTypeID is TL type id of Text.
const TextTypeID = 0x22765898

// Ensuring interfaces in compile-time for Text.
var (
	_ bin.Encoder     = &Text{}
	_ bin.Decoder     = &Text{}
	_ bin.BareEncoder = &Text{}
	_ bin.BareDecoder = &Text{}
)

func (t *Text) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *Text) String() string {
	if t == nil {
		return "Text(nil)"
	}
	type Alias Text
	return fmt.Sprintf("Text%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Text) TypeID() uint32 {
	return TextTypeID
}

// TypeName returns name of type in TL schema.
func (*Text) TypeName() string {
	return "text"
}

// TypeInfo returns info about TL type.
func (t *Text) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "text",
		ID:   TextTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *Text) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode text#22765898 as nil")
	}
	b.PutID(TextTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *Text) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode text#22765898 as nil")
	}
	b.PutString(t.Text)
	return nil
}

// Decode implements bin.Decoder.
func (t *Text) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode text#22765898 to nil")
	}
	if err := b.ConsumeID(TextTypeID); err != nil {
		return fmt.Errorf("unable to decode text#22765898: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *Text) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode text#22765898 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode text#22765898: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (t *Text) EncodeTDLibJSON(b jsontd.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode text#22765898 as nil")
	}
	b.ObjStart()
	b.PutID("text")
	b.FieldStart("text")
	b.PutString(t.Text)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (t *Text) DecodeTDLibJSON(b jsontd.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode text#22765898 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("text"); err != nil {
				return fmt.Errorf("unable to decode text#22765898: %w", err)
			}
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode text#22765898: field text: %w", err)
			}
			t.Text = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (t *Text) GetText() (value string) {
	return t.Text
}
