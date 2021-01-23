// Code generated by gotdgen, DO NOT EDIT.

package e2e

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// InputStickerSetShortName represents TL type `inputStickerSetShortName#861cc8a0`.
// Stickerset by short name, from tg://addstickers?set=short_name
//
// See https://core.telegram.org/constructor/inputStickerSetShortName for reference.
type InputStickerSetShortName struct {
	// From tg://addstickers?set=short_name
	ShortName string
}

// InputStickerSetShortNameTypeID is TL type id of InputStickerSetShortName.
const InputStickerSetShortNameTypeID = 0x861cc8a0

func (i *InputStickerSetShortName) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ShortName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetShortName) String() string {
	if i == nil {
		return "InputStickerSetShortName(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputStickerSetShortName")
	sb.WriteString("{\n")
	sb.WriteString("\tShortName: ")
	sb.WriteString(fmt.Sprint(i.ShortName))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputStickerSetShortName) TypeID() uint32 {
	return InputStickerSetShortNameTypeID
}

// Encode implements bin.Encoder.
func (i *InputStickerSetShortName) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetShortName#861cc8a0 as nil")
	}
	b.PutID(InputStickerSetShortNameTypeID)
	b.PutString(i.ShortName)
	return nil
}

// GetShortName returns value of ShortName field.
func (i *InputStickerSetShortName) GetShortName() (value string) {
	return i.ShortName
}

// Decode implements bin.Decoder.
func (i *InputStickerSetShortName) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetShortName#861cc8a0 to nil")
	}
	if err := b.ConsumeID(InputStickerSetShortNameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetShortName#861cc8a0: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetShortName#861cc8a0: field short_name: %w", err)
		}
		i.ShortName = value
	}
	return nil
}

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetShortName) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetShortName.
var (
	_ bin.Encoder = &InputStickerSetShortName{}
	_ bin.Decoder = &InputStickerSetShortName{}

	_ InputStickerSetClass = &InputStickerSetShortName{}
)

// InputStickerSetEmpty represents TL type `inputStickerSetEmpty#ffb62b95`.
// Empty constructor
//
// See https://core.telegram.org/constructor/inputStickerSetEmpty for reference.
type InputStickerSetEmpty struct {
}

// InputStickerSetEmptyTypeID is TL type id of InputStickerSetEmpty.
const InputStickerSetEmptyTypeID = 0xffb62b95

func (i *InputStickerSetEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetEmpty) String() string {
	if i == nil {
		return "InputStickerSetEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputStickerSetEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputStickerSetEmpty) TypeID() uint32 {
	return InputStickerSetEmptyTypeID
}

// Encode implements bin.Encoder.
func (i *InputStickerSetEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetEmpty#ffb62b95 as nil")
	}
	b.PutID(InputStickerSetEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerSetEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetEmpty#ffb62b95 to nil")
	}
	if err := b.ConsumeID(InputStickerSetEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetEmpty#ffb62b95: %w", err)
	}
	return nil
}

// construct implements constructor of InputStickerSetClass.
func (i InputStickerSetEmpty) construct() InputStickerSetClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetEmpty.
var (
	_ bin.Encoder = &InputStickerSetEmpty{}
	_ bin.Decoder = &InputStickerSetEmpty{}

	_ InputStickerSetClass = &InputStickerSetEmpty{}
)

// InputStickerSetClass represents InputStickerSet generic type.
//
// See https://core.telegram.org/type/InputStickerSet for reference.
//
// Example:
//  g, err := DecodeInputStickerSet(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputStickerSetShortName: // inputStickerSetShortName#861cc8a0
//  case *InputStickerSetEmpty: // inputStickerSetEmpty#ffb62b95
//  default: panic(v)
//  }
type InputStickerSetClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputStickerSetClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputStickerSet implements binary de-serialization for InputStickerSetClass.
func DecodeInputStickerSet(buf *bin.Buffer) (InputStickerSetClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStickerSetShortNameTypeID:
		// Decoding inputStickerSetShortName#861cc8a0.
		v := InputStickerSetShortName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", err)
		}
		return &v, nil
	case InputStickerSetEmptyTypeID:
		// Decoding inputStickerSetEmpty#ffb62b95.
		v := InputStickerSetEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStickerSetClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputStickerSet boxes the InputStickerSetClass providing a helper.
type InputStickerSetBox struct {
	InputStickerSet InputStickerSetClass
}

// Decode implements bin.Decoder for InputStickerSetBox.
func (b *InputStickerSetBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStickerSetBox to nil")
	}
	v, err := DecodeInputStickerSet(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStickerSet = v
	return nil
}

// Encode implements bin.Encode for InputStickerSetBox.
func (b *InputStickerSetBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputStickerSet == nil {
		return fmt.Errorf("unable to encode InputStickerSetClass as nil")
	}
	return b.InputStickerSet.Encode(buf)
}