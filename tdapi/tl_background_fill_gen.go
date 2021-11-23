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

// BackgroundFillSolid represents TL type `backgroundFillSolid#3c3dbc1d`.
type BackgroundFillSolid struct {
	// A color of the background in the RGB24 format
	Color int32
}

// BackgroundFillSolidTypeID is TL type id of BackgroundFillSolid.
const BackgroundFillSolidTypeID = 0x3c3dbc1d

// construct implements constructor of BackgroundFillClass.
func (b BackgroundFillSolid) construct() BackgroundFillClass { return &b }

// Ensuring interfaces in compile-time for BackgroundFillSolid.
var (
	_ bin.Encoder     = &BackgroundFillSolid{}
	_ bin.Decoder     = &BackgroundFillSolid{}
	_ bin.BareEncoder = &BackgroundFillSolid{}
	_ bin.BareDecoder = &BackgroundFillSolid{}

	_ BackgroundFillClass = &BackgroundFillSolid{}
)

func (b *BackgroundFillSolid) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Color == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BackgroundFillSolid) String() string {
	if b == nil {
		return "BackgroundFillSolid(nil)"
	}
	type Alias BackgroundFillSolid
	return fmt.Sprintf("BackgroundFillSolid%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BackgroundFillSolid) TypeID() uint32 {
	return BackgroundFillSolidTypeID
}

// TypeName returns name of type in TL schema.
func (*BackgroundFillSolid) TypeName() string {
	return "backgroundFillSolid"
}

// TypeInfo returns info about TL type.
func (b *BackgroundFillSolid) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "backgroundFillSolid",
		ID:   BackgroundFillSolidTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Color",
			SchemaName: "color",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BackgroundFillSolid) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundFillSolid#3c3dbc1d as nil")
	}
	buf.PutID(BackgroundFillSolidTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BackgroundFillSolid) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundFillSolid#3c3dbc1d as nil")
	}
	buf.PutInt32(b.Color)
	return nil
}

// Decode implements bin.Decoder.
func (b *BackgroundFillSolid) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundFillSolid#3c3dbc1d to nil")
	}
	if err := buf.ConsumeID(BackgroundFillSolidTypeID); err != nil {
		return fmt.Errorf("unable to decode backgroundFillSolid#3c3dbc1d: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BackgroundFillSolid) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundFillSolid#3c3dbc1d to nil")
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundFillSolid#3c3dbc1d: field color: %w", err)
		}
		b.Color = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (b *BackgroundFillSolid) EncodeTDLibJSON(buf jsontd.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundFillSolid#3c3dbc1d as nil")
	}
	buf.ObjStart()
	buf.PutID("backgroundFillSolid")
	buf.FieldStart("color")
	buf.PutInt32(b.Color)
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (b *BackgroundFillSolid) DecodeTDLibJSON(buf jsontd.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundFillSolid#3c3dbc1d to nil")
	}

	return buf.Obj(func(buf jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := buf.ConsumeID("backgroundFillSolid"); err != nil {
				return fmt.Errorf("unable to decode backgroundFillSolid#3c3dbc1d: %w", err)
			}
		case "color":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundFillSolid#3c3dbc1d: field color: %w", err)
			}
			b.Color = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetColor returns value of Color field.
func (b *BackgroundFillSolid) GetColor() (value int32) {
	return b.Color
}

// BackgroundFillGradient represents TL type `backgroundFillGradient#925ff17f`.
type BackgroundFillGradient struct {
	// A top color of the background in the RGB24 format
	TopColor int32
	// A bottom color of the background in the RGB24 format
	BottomColor int32
	// Clockwise rotation angle of the gradient, in degrees; 0-359. Should be always
	// divisible by 45
	RotationAngle int32
}

// BackgroundFillGradientTypeID is TL type id of BackgroundFillGradient.
const BackgroundFillGradientTypeID = 0x925ff17f

// construct implements constructor of BackgroundFillClass.
func (b BackgroundFillGradient) construct() BackgroundFillClass { return &b }

// Ensuring interfaces in compile-time for BackgroundFillGradient.
var (
	_ bin.Encoder     = &BackgroundFillGradient{}
	_ bin.Decoder     = &BackgroundFillGradient{}
	_ bin.BareEncoder = &BackgroundFillGradient{}
	_ bin.BareDecoder = &BackgroundFillGradient{}

	_ BackgroundFillClass = &BackgroundFillGradient{}
)

func (b *BackgroundFillGradient) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.TopColor == 0) {
		return false
	}
	if !(b.BottomColor == 0) {
		return false
	}
	if !(b.RotationAngle == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BackgroundFillGradient) String() string {
	if b == nil {
		return "BackgroundFillGradient(nil)"
	}
	type Alias BackgroundFillGradient
	return fmt.Sprintf("BackgroundFillGradient%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BackgroundFillGradient) TypeID() uint32 {
	return BackgroundFillGradientTypeID
}

// TypeName returns name of type in TL schema.
func (*BackgroundFillGradient) TypeName() string {
	return "backgroundFillGradient"
}

// TypeInfo returns info about TL type.
func (b *BackgroundFillGradient) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "backgroundFillGradient",
		ID:   BackgroundFillGradientTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TopColor",
			SchemaName: "top_color",
		},
		{
			Name:       "BottomColor",
			SchemaName: "bottom_color",
		},
		{
			Name:       "RotationAngle",
			SchemaName: "rotation_angle",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BackgroundFillGradient) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundFillGradient#925ff17f as nil")
	}
	buf.PutID(BackgroundFillGradientTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BackgroundFillGradient) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundFillGradient#925ff17f as nil")
	}
	buf.PutInt32(b.TopColor)
	buf.PutInt32(b.BottomColor)
	buf.PutInt32(b.RotationAngle)
	return nil
}

// Decode implements bin.Decoder.
func (b *BackgroundFillGradient) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundFillGradient#925ff17f to nil")
	}
	if err := buf.ConsumeID(BackgroundFillGradientTypeID); err != nil {
		return fmt.Errorf("unable to decode backgroundFillGradient#925ff17f: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BackgroundFillGradient) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundFillGradient#925ff17f to nil")
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundFillGradient#925ff17f: field top_color: %w", err)
		}
		b.TopColor = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundFillGradient#925ff17f: field bottom_color: %w", err)
		}
		b.BottomColor = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundFillGradient#925ff17f: field rotation_angle: %w", err)
		}
		b.RotationAngle = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (b *BackgroundFillGradient) EncodeTDLibJSON(buf jsontd.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundFillGradient#925ff17f as nil")
	}
	buf.ObjStart()
	buf.PutID("backgroundFillGradient")
	buf.FieldStart("top_color")
	buf.PutInt32(b.TopColor)
	buf.FieldStart("bottom_color")
	buf.PutInt32(b.BottomColor)
	buf.FieldStart("rotation_angle")
	buf.PutInt32(b.RotationAngle)
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (b *BackgroundFillGradient) DecodeTDLibJSON(buf jsontd.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundFillGradient#925ff17f to nil")
	}

	return buf.Obj(func(buf jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := buf.ConsumeID("backgroundFillGradient"); err != nil {
				return fmt.Errorf("unable to decode backgroundFillGradient#925ff17f: %w", err)
			}
		case "top_color":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundFillGradient#925ff17f: field top_color: %w", err)
			}
			b.TopColor = value
		case "bottom_color":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundFillGradient#925ff17f: field bottom_color: %w", err)
			}
			b.BottomColor = value
		case "rotation_angle":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundFillGradient#925ff17f: field rotation_angle: %w", err)
			}
			b.RotationAngle = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetTopColor returns value of TopColor field.
func (b *BackgroundFillGradient) GetTopColor() (value int32) {
	return b.TopColor
}

// GetBottomColor returns value of BottomColor field.
func (b *BackgroundFillGradient) GetBottomColor() (value int32) {
	return b.BottomColor
}

// GetRotationAngle returns value of RotationAngle field.
func (b *BackgroundFillGradient) GetRotationAngle() (value int32) {
	return b.RotationAngle
}

// BackgroundFillFreeformGradient represents TL type `backgroundFillFreeformGradient#fa31756a`.
type BackgroundFillFreeformGradient struct {
	// A list of 3 or 4 colors of the freeform gradients in the RGB24 format
	Colors []int32
}

// BackgroundFillFreeformGradientTypeID is TL type id of BackgroundFillFreeformGradient.
const BackgroundFillFreeformGradientTypeID = 0xfa31756a

// construct implements constructor of BackgroundFillClass.
func (b BackgroundFillFreeformGradient) construct() BackgroundFillClass { return &b }

// Ensuring interfaces in compile-time for BackgroundFillFreeformGradient.
var (
	_ bin.Encoder     = &BackgroundFillFreeformGradient{}
	_ bin.Decoder     = &BackgroundFillFreeformGradient{}
	_ bin.BareEncoder = &BackgroundFillFreeformGradient{}
	_ bin.BareDecoder = &BackgroundFillFreeformGradient{}

	_ BackgroundFillClass = &BackgroundFillFreeformGradient{}
)

func (b *BackgroundFillFreeformGradient) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Colors == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BackgroundFillFreeformGradient) String() string {
	if b == nil {
		return "BackgroundFillFreeformGradient(nil)"
	}
	type Alias BackgroundFillFreeformGradient
	return fmt.Sprintf("BackgroundFillFreeformGradient%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BackgroundFillFreeformGradient) TypeID() uint32 {
	return BackgroundFillFreeformGradientTypeID
}

// TypeName returns name of type in TL schema.
func (*BackgroundFillFreeformGradient) TypeName() string {
	return "backgroundFillFreeformGradient"
}

// TypeInfo returns info about TL type.
func (b *BackgroundFillFreeformGradient) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "backgroundFillFreeformGradient",
		ID:   BackgroundFillFreeformGradientTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Colors",
			SchemaName: "colors",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BackgroundFillFreeformGradient) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundFillFreeformGradient#fa31756a as nil")
	}
	buf.PutID(BackgroundFillFreeformGradientTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BackgroundFillFreeformGradient) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundFillFreeformGradient#fa31756a as nil")
	}
	buf.PutInt(len(b.Colors))
	for _, v := range b.Colors {
		buf.PutInt32(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BackgroundFillFreeformGradient) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundFillFreeformGradient#fa31756a to nil")
	}
	if err := buf.ConsumeID(BackgroundFillFreeformGradientTypeID); err != nil {
		return fmt.Errorf("unable to decode backgroundFillFreeformGradient#fa31756a: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BackgroundFillFreeformGradient) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundFillFreeformGradient#fa31756a to nil")
	}
	{
		headerLen, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundFillFreeformGradient#fa31756a: field colors: %w", err)
		}

		if headerLen > 0 {
			b.Colors = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundFillFreeformGradient#fa31756a: field colors: %w", err)
			}
			b.Colors = append(b.Colors, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (b *BackgroundFillFreeformGradient) EncodeTDLibJSON(buf jsontd.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundFillFreeformGradient#fa31756a as nil")
	}
	buf.ObjStart()
	buf.PutID("backgroundFillFreeformGradient")
	buf.FieldStart("colors")
	buf.ArrStart()
	for _, v := range b.Colors {
		buf.PutInt32(v)
	}
	buf.ArrEnd()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (b *BackgroundFillFreeformGradient) DecodeTDLibJSON(buf jsontd.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundFillFreeformGradient#fa31756a to nil")
	}

	return buf.Obj(func(buf jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := buf.ConsumeID("backgroundFillFreeformGradient"); err != nil {
				return fmt.Errorf("unable to decode backgroundFillFreeformGradient#fa31756a: %w", err)
			}
		case "colors":
			if err := buf.Arr(func(buf jsontd.Decoder) error {
				value, err := buf.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode backgroundFillFreeformGradient#fa31756a: field colors: %w", err)
				}
				b.Colors = append(b.Colors, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode backgroundFillFreeformGradient#fa31756a: field colors: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetColors returns value of Colors field.
func (b *BackgroundFillFreeformGradient) GetColors() (value []int32) {
	return b.Colors
}

// BackgroundFillClass represents BackgroundFill generic type.
//
// Example:
//  g, err := tdapi.DecodeBackgroundFill(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.BackgroundFillSolid: // backgroundFillSolid#3c3dbc1d
//  case *tdapi.BackgroundFillGradient: // backgroundFillGradient#925ff17f
//  case *tdapi.BackgroundFillFreeformGradient: // backgroundFillFreeformGradient#fa31756a
//  default: panic(v)
//  }
type BackgroundFillClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BackgroundFillClass

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

// DecodeBackgroundFill implements binary de-serialization for BackgroundFillClass.
func DecodeBackgroundFill(buf *bin.Buffer) (BackgroundFillClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BackgroundFillSolidTypeID:
		// Decoding backgroundFillSolid#3c3dbc1d.
		v := BackgroundFillSolid{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundFillClass: %w", err)
		}
		return &v, nil
	case BackgroundFillGradientTypeID:
		// Decoding backgroundFillGradient#925ff17f.
		v := BackgroundFillGradient{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundFillClass: %w", err)
		}
		return &v, nil
	case BackgroundFillFreeformGradientTypeID:
		// Decoding backgroundFillFreeformGradient#fa31756a.
		v := BackgroundFillFreeformGradient{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundFillClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BackgroundFillClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONBackgroundFill implements binary de-serialization for BackgroundFillClass.
func DecodeTDLibJSONBackgroundFill(buf jsontd.Decoder) (BackgroundFillClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "backgroundFillSolid":
		// Decoding backgroundFillSolid#3c3dbc1d.
		v := BackgroundFillSolid{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundFillClass: %w", err)
		}
		return &v, nil
	case "backgroundFillGradient":
		// Decoding backgroundFillGradient#925ff17f.
		v := BackgroundFillGradient{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundFillClass: %w", err)
		}
		return &v, nil
	case "backgroundFillFreeformGradient":
		// Decoding backgroundFillFreeformGradient#fa31756a.
		v := BackgroundFillFreeformGradient{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundFillClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BackgroundFillClass: %w", jsontd.NewUnexpectedID(id))
	}
}

// BackgroundFill boxes the BackgroundFillClass providing a helper.
type BackgroundFillBox struct {
	BackgroundFill BackgroundFillClass
}

// Decode implements bin.Decoder for BackgroundFillBox.
func (b *BackgroundFillBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BackgroundFillBox to nil")
	}
	v, err := DecodeBackgroundFill(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BackgroundFill = v
	return nil
}

// Encode implements bin.Encode for BackgroundFillBox.
func (b *BackgroundFillBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BackgroundFill == nil {
		return fmt.Errorf("unable to encode BackgroundFillClass as nil")
	}
	return b.BackgroundFill.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for BackgroundFillBox.
func (b *BackgroundFillBox) DecodeTDLibJSON(buf jsontd.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode BackgroundFillBox to nil")
	}
	v, err := DecodeTDLibJSONBackgroundFill(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BackgroundFill = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for BackgroundFillBox.
func (b *BackgroundFillBox) EncodeTDLibJSON(buf jsontd.Encoder) error {
	if b == nil || b.BackgroundFill == nil {
		return fmt.Errorf("unable to encode BackgroundFillClass as nil")
	}
	return b.BackgroundFill.EncodeTDLibJSON(buf)
}
