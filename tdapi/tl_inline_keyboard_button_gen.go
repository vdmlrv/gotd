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

// InlineKeyboardButton represents TL type `inlineKeyboardButton#e9d21e18`.
type InlineKeyboardButton struct {
	// Text of the button
	Text string
	// Type of the button
	Type InlineKeyboardButtonTypeClass
}

// InlineKeyboardButtonTypeID is TL type id of InlineKeyboardButton.
const InlineKeyboardButtonTypeID = 0xe9d21e18

// Ensuring interfaces in compile-time for InlineKeyboardButton.
var (
	_ bin.Encoder     = &InlineKeyboardButton{}
	_ bin.Decoder     = &InlineKeyboardButton{}
	_ bin.BareEncoder = &InlineKeyboardButton{}
	_ bin.BareDecoder = &InlineKeyboardButton{}
)

func (i *InlineKeyboardButton) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Text == "") {
		return false
	}
	if !(i.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineKeyboardButton) String() string {
	if i == nil {
		return "InlineKeyboardButton(nil)"
	}
	type Alias InlineKeyboardButton
	return fmt.Sprintf("InlineKeyboardButton%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineKeyboardButton) TypeID() uint32 {
	return InlineKeyboardButtonTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineKeyboardButton) TypeName() string {
	return "inlineKeyboardButton"
}

// TypeInfo returns info about TL type.
func (i *InlineKeyboardButton) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineKeyboardButton",
		ID:   InlineKeyboardButtonTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineKeyboardButton) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButton#e9d21e18 as nil")
	}
	b.PutID(InlineKeyboardButtonTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineKeyboardButton) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButton#e9d21e18 as nil")
	}
	b.PutString(i.Text)
	if i.Type == nil {
		return fmt.Errorf("unable to encode inlineKeyboardButton#e9d21e18: field type is nil")
	}
	if err := i.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inlineKeyboardButton#e9d21e18: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineKeyboardButton) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButton#e9d21e18 to nil")
	}
	if err := b.ConsumeID(InlineKeyboardButtonTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineKeyboardButton#e9d21e18: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineKeyboardButton) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButton#e9d21e18 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButton#e9d21e18: field text: %w", err)
		}
		i.Text = value
	}
	{
		value, err := DecodeInlineKeyboardButtonType(b)
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButton#e9d21e18: field type: %w", err)
		}
		i.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (i *InlineKeyboardButton) EncodeTDLibJSON(b jsontd.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButton#e9d21e18 as nil")
	}
	b.ObjStart()
	b.PutID("inlineKeyboardButton")
	b.FieldStart("text")
	b.PutString(i.Text)
	b.FieldStart("type")
	if i.Type == nil {
		return fmt.Errorf("unable to encode inlineKeyboardButton#e9d21e18: field type is nil")
	}
	if err := i.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inlineKeyboardButton#e9d21e18: field type: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (i *InlineKeyboardButton) DecodeTDLibJSON(b jsontd.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButton#e9d21e18 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("inlineKeyboardButton"); err != nil {
				return fmt.Errorf("unable to decode inlineKeyboardButton#e9d21e18: %w", err)
			}
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inlineKeyboardButton#e9d21e18: field text: %w", err)
			}
			i.Text = value
		case "type":
			value, err := DecodeTDLibJSONInlineKeyboardButtonType(b)
			if err != nil {
				return fmt.Errorf("unable to decode inlineKeyboardButton#e9d21e18: field type: %w", err)
			}
			i.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (i *InlineKeyboardButton) GetText() (value string) {
	return i.Text
}

// GetType returns value of Type field.
func (i *InlineKeyboardButton) GetType() (value InlineKeyboardButtonTypeClass) {
	return i.Type
}
