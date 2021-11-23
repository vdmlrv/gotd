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

// InputChatPhotoPrevious represents TL type `inputChatPhotoPrevious#160e9d1`.
type InputChatPhotoPrevious struct {
	// Identifier of the current user's profile photo to reuse
	ChatPhotoID int64
}

// InputChatPhotoPreviousTypeID is TL type id of InputChatPhotoPrevious.
const InputChatPhotoPreviousTypeID = 0x160e9d1

// construct implements constructor of InputChatPhotoClass.
func (i InputChatPhotoPrevious) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatPhotoPrevious.
var (
	_ bin.Encoder     = &InputChatPhotoPrevious{}
	_ bin.Decoder     = &InputChatPhotoPrevious{}
	_ bin.BareEncoder = &InputChatPhotoPrevious{}
	_ bin.BareDecoder = &InputChatPhotoPrevious{}

	_ InputChatPhotoClass = &InputChatPhotoPrevious{}
)

func (i *InputChatPhotoPrevious) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChatPhotoID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChatPhotoPrevious) String() string {
	if i == nil {
		return "InputChatPhotoPrevious(nil)"
	}
	type Alias InputChatPhotoPrevious
	return fmt.Sprintf("InputChatPhotoPrevious%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChatPhotoPrevious) TypeID() uint32 {
	return InputChatPhotoPreviousTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChatPhotoPrevious) TypeName() string {
	return "inputChatPhotoPrevious"
}

// TypeInfo returns info about TL type.
func (i *InputChatPhotoPrevious) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChatPhotoPrevious",
		ID:   InputChatPhotoPreviousTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatPhotoID",
			SchemaName: "chat_photo_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChatPhotoPrevious) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoPrevious#160e9d1 as nil")
	}
	b.PutID(InputChatPhotoPreviousTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChatPhotoPrevious) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoPrevious#160e9d1 as nil")
	}
	b.PutLong(i.ChatPhotoID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChatPhotoPrevious) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoPrevious#160e9d1 to nil")
	}
	if err := b.ConsumeID(InputChatPhotoPreviousTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatPhotoPrevious#160e9d1: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChatPhotoPrevious) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoPrevious#160e9d1 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputChatPhotoPrevious#160e9d1: field chat_photo_id: %w", err)
		}
		i.ChatPhotoID = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (i *InputChatPhotoPrevious) EncodeTDLibJSON(b jsontd.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoPrevious#160e9d1 as nil")
	}
	b.ObjStart()
	b.PutID("inputChatPhotoPrevious")
	b.FieldStart("chat_photo_id")
	b.PutLong(i.ChatPhotoID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (i *InputChatPhotoPrevious) DecodeTDLibJSON(b jsontd.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoPrevious#160e9d1 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("inputChatPhotoPrevious"); err != nil {
				return fmt.Errorf("unable to decode inputChatPhotoPrevious#160e9d1: %w", err)
			}
		case "chat_photo_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode inputChatPhotoPrevious#160e9d1: field chat_photo_id: %w", err)
			}
			i.ChatPhotoID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatPhotoID returns value of ChatPhotoID field.
func (i *InputChatPhotoPrevious) GetChatPhotoID() (value int64) {
	return i.ChatPhotoID
}

// InputChatPhotoStatic represents TL type `inputChatPhotoStatic#75f7e2b3`.
type InputChatPhotoStatic struct {
	// Photo to be set as profile photo. Only inputFileLocal and inputFileGenerated are
	// allowed
	Photo InputFileClass
}

// InputChatPhotoStaticTypeID is TL type id of InputChatPhotoStatic.
const InputChatPhotoStaticTypeID = 0x75f7e2b3

// construct implements constructor of InputChatPhotoClass.
func (i InputChatPhotoStatic) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatPhotoStatic.
var (
	_ bin.Encoder     = &InputChatPhotoStatic{}
	_ bin.Decoder     = &InputChatPhotoStatic{}
	_ bin.BareEncoder = &InputChatPhotoStatic{}
	_ bin.BareDecoder = &InputChatPhotoStatic{}

	_ InputChatPhotoClass = &InputChatPhotoStatic{}
)

func (i *InputChatPhotoStatic) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Photo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChatPhotoStatic) String() string {
	if i == nil {
		return "InputChatPhotoStatic(nil)"
	}
	type Alias InputChatPhotoStatic
	return fmt.Sprintf("InputChatPhotoStatic%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChatPhotoStatic) TypeID() uint32 {
	return InputChatPhotoStaticTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChatPhotoStatic) TypeName() string {
	return "inputChatPhotoStatic"
}

// TypeInfo returns info about TL type.
func (i *InputChatPhotoStatic) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChatPhotoStatic",
		ID:   InputChatPhotoStaticTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChatPhotoStatic) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoStatic#75f7e2b3 as nil")
	}
	b.PutID(InputChatPhotoStaticTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChatPhotoStatic) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoStatic#75f7e2b3 as nil")
	}
	if i.Photo == nil {
		return fmt.Errorf("unable to encode inputChatPhotoStatic#75f7e2b3: field photo is nil")
	}
	if err := i.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChatPhotoStatic#75f7e2b3: field photo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChatPhotoStatic) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoStatic#75f7e2b3 to nil")
	}
	if err := b.ConsumeID(InputChatPhotoStaticTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatPhotoStatic#75f7e2b3: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChatPhotoStatic) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoStatic#75f7e2b3 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatPhotoStatic#75f7e2b3: field photo: %w", err)
		}
		i.Photo = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (i *InputChatPhotoStatic) EncodeTDLibJSON(b jsontd.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoStatic#75f7e2b3 as nil")
	}
	b.ObjStart()
	b.PutID("inputChatPhotoStatic")
	b.FieldStart("photo")
	if i.Photo == nil {
		return fmt.Errorf("unable to encode inputChatPhotoStatic#75f7e2b3: field photo is nil")
	}
	if err := i.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputChatPhotoStatic#75f7e2b3: field photo: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (i *InputChatPhotoStatic) DecodeTDLibJSON(b jsontd.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoStatic#75f7e2b3 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("inputChatPhotoStatic"); err != nil {
				return fmt.Errorf("unable to decode inputChatPhotoStatic#75f7e2b3: %w", err)
			}
		case "photo":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputChatPhotoStatic#75f7e2b3: field photo: %w", err)
			}
			i.Photo = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPhoto returns value of Photo field.
func (i *InputChatPhotoStatic) GetPhoto() (value InputFileClass) {
	return i.Photo
}

// InputChatPhotoAnimation represents TL type `inputChatPhotoAnimation#56a3422`.
type InputChatPhotoAnimation struct {
	// Animation to be set as profile photo. Only inputFileLocal and inputFileGenerated are
	// allowed
	Animation InputFileClass
	// Timestamp of the frame, which will be used as static chat photo
	MainFrameTimestamp float64
}

// InputChatPhotoAnimationTypeID is TL type id of InputChatPhotoAnimation.
const InputChatPhotoAnimationTypeID = 0x56a3422

// construct implements constructor of InputChatPhotoClass.
func (i InputChatPhotoAnimation) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatPhotoAnimation.
var (
	_ bin.Encoder     = &InputChatPhotoAnimation{}
	_ bin.Decoder     = &InputChatPhotoAnimation{}
	_ bin.BareEncoder = &InputChatPhotoAnimation{}
	_ bin.BareDecoder = &InputChatPhotoAnimation{}

	_ InputChatPhotoClass = &InputChatPhotoAnimation{}
)

func (i *InputChatPhotoAnimation) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Animation == nil) {
		return false
	}
	if !(i.MainFrameTimestamp == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChatPhotoAnimation) String() string {
	if i == nil {
		return "InputChatPhotoAnimation(nil)"
	}
	type Alias InputChatPhotoAnimation
	return fmt.Sprintf("InputChatPhotoAnimation%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChatPhotoAnimation) TypeID() uint32 {
	return InputChatPhotoAnimationTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChatPhotoAnimation) TypeName() string {
	return "inputChatPhotoAnimation"
}

// TypeInfo returns info about TL type.
func (i *InputChatPhotoAnimation) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChatPhotoAnimation",
		ID:   InputChatPhotoAnimationTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Animation",
			SchemaName: "animation",
		},
		{
			Name:       "MainFrameTimestamp",
			SchemaName: "main_frame_timestamp",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChatPhotoAnimation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoAnimation#56a3422 as nil")
	}
	b.PutID(InputChatPhotoAnimationTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChatPhotoAnimation) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoAnimation#56a3422 as nil")
	}
	if i.Animation == nil {
		return fmt.Errorf("unable to encode inputChatPhotoAnimation#56a3422: field animation is nil")
	}
	if err := i.Animation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChatPhotoAnimation#56a3422: field animation: %w", err)
	}
	b.PutDouble(i.MainFrameTimestamp)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChatPhotoAnimation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoAnimation#56a3422 to nil")
	}
	if err := b.ConsumeID(InputChatPhotoAnimationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatPhotoAnimation#56a3422: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChatPhotoAnimation) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoAnimation#56a3422 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatPhotoAnimation#56a3422: field animation: %w", err)
		}
		i.Animation = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputChatPhotoAnimation#56a3422: field main_frame_timestamp: %w", err)
		}
		i.MainFrameTimestamp = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (i *InputChatPhotoAnimation) EncodeTDLibJSON(b jsontd.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoAnimation#56a3422 as nil")
	}
	b.ObjStart()
	b.PutID("inputChatPhotoAnimation")
	b.FieldStart("animation")
	if i.Animation == nil {
		return fmt.Errorf("unable to encode inputChatPhotoAnimation#56a3422: field animation is nil")
	}
	if err := i.Animation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputChatPhotoAnimation#56a3422: field animation: %w", err)
	}
	b.FieldStart("main_frame_timestamp")
	b.PutDouble(i.MainFrameTimestamp)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (i *InputChatPhotoAnimation) DecodeTDLibJSON(b jsontd.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoAnimation#56a3422 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("inputChatPhotoAnimation"); err != nil {
				return fmt.Errorf("unable to decode inputChatPhotoAnimation#56a3422: %w", err)
			}
		case "animation":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputChatPhotoAnimation#56a3422: field animation: %w", err)
			}
			i.Animation = value
		case "main_frame_timestamp":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode inputChatPhotoAnimation#56a3422: field main_frame_timestamp: %w", err)
			}
			i.MainFrameTimestamp = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAnimation returns value of Animation field.
func (i *InputChatPhotoAnimation) GetAnimation() (value InputFileClass) {
	return i.Animation
}

// GetMainFrameTimestamp returns value of MainFrameTimestamp field.
func (i *InputChatPhotoAnimation) GetMainFrameTimestamp() (value float64) {
	return i.MainFrameTimestamp
}

// InputChatPhotoClass represents InputChatPhoto generic type.
//
// Example:
//  g, err := tdapi.DecodeInputChatPhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.InputChatPhotoPrevious: // inputChatPhotoPrevious#160e9d1
//  case *tdapi.InputChatPhotoStatic: // inputChatPhotoStatic#75f7e2b3
//  case *tdapi.InputChatPhotoAnimation: // inputChatPhotoAnimation#56a3422
//  default: panic(v)
//  }
type InputChatPhotoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputChatPhotoClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b jsontd.Encoder) error
	DecodeTDLibJSON(b jsontd.Decoder) error
}

// DecodeInputChatPhoto implements binary de-serialization for InputChatPhotoClass.
func DecodeInputChatPhoto(buf *bin.Buffer) (InputChatPhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputChatPhotoPreviousTypeID:
		// Decoding inputChatPhotoPrevious#160e9d1.
		v := InputChatPhotoPrevious{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	case InputChatPhotoStaticTypeID:
		// Decoding inputChatPhotoStatic#75f7e2b3.
		v := InputChatPhotoStatic{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	case InputChatPhotoAnimationTypeID:
		// Decoding inputChatPhotoAnimation#56a3422.
		v := InputChatPhotoAnimation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInputChatPhoto implements binary de-serialization for InputChatPhotoClass.
func DecodeTDLibJSONInputChatPhoto(buf jsontd.Decoder) (InputChatPhotoClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inputChatPhotoPrevious":
		// Decoding inputChatPhotoPrevious#160e9d1.
		v := InputChatPhotoPrevious{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	case "inputChatPhotoStatic":
		// Decoding inputChatPhotoStatic#75f7e2b3.
		v := InputChatPhotoStatic{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	case "inputChatPhotoAnimation":
		// Decoding inputChatPhotoAnimation#56a3422.
		v := InputChatPhotoAnimation{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", jsontd.NewUnexpectedID(id))
	}
}

// InputChatPhoto boxes the InputChatPhotoClass providing a helper.
type InputChatPhotoBox struct {
	InputChatPhoto InputChatPhotoClass
}

// Decode implements bin.Decoder for InputChatPhotoBox.
func (b *InputChatPhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputChatPhotoBox to nil")
	}
	v, err := DecodeInputChatPhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputChatPhoto = v
	return nil
}

// Encode implements bin.Encode for InputChatPhotoBox.
func (b *InputChatPhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputChatPhoto == nil {
		return fmt.Errorf("unable to encode InputChatPhotoClass as nil")
	}
	return b.InputChatPhoto.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InputChatPhotoBox.
func (b *InputChatPhotoBox) DecodeTDLibJSON(buf jsontd.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputChatPhotoBox to nil")
	}
	v, err := DecodeTDLibJSONInputChatPhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputChatPhoto = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InputChatPhotoBox.
func (b *InputChatPhotoBox) EncodeTDLibJSON(buf jsontd.Encoder) error {
	if b == nil || b.InputChatPhoto == nil {
		return fmt.Errorf("unable to encode InputChatPhotoClass as nil")
	}
	return b.InputChatPhoto.EncodeTDLibJSON(buf)
}
