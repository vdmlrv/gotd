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

// PageBlockVerticalAlignmentTop represents TL type `pageBlockVerticalAlignmentTop#ba719a6`.
type PageBlockVerticalAlignmentTop struct {
}

// PageBlockVerticalAlignmentTopTypeID is TL type id of PageBlockVerticalAlignmentTop.
const PageBlockVerticalAlignmentTopTypeID = 0xba719a6

// construct implements constructor of PageBlockVerticalAlignmentClass.
func (p PageBlockVerticalAlignmentTop) construct() PageBlockVerticalAlignmentClass { return &p }

// Ensuring interfaces in compile-time for PageBlockVerticalAlignmentTop.
var (
	_ bin.Encoder     = &PageBlockVerticalAlignmentTop{}
	_ bin.Decoder     = &PageBlockVerticalAlignmentTop{}
	_ bin.BareEncoder = &PageBlockVerticalAlignmentTop{}
	_ bin.BareDecoder = &PageBlockVerticalAlignmentTop{}

	_ PageBlockVerticalAlignmentClass = &PageBlockVerticalAlignmentTop{}
)

func (p *PageBlockVerticalAlignmentTop) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageBlockVerticalAlignmentTop) String() string {
	if p == nil {
		return "PageBlockVerticalAlignmentTop(nil)"
	}
	type Alias PageBlockVerticalAlignmentTop
	return fmt.Sprintf("PageBlockVerticalAlignmentTop%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageBlockVerticalAlignmentTop) TypeID() uint32 {
	return PageBlockVerticalAlignmentTopTypeID
}

// TypeName returns name of type in TL schema.
func (*PageBlockVerticalAlignmentTop) TypeName() string {
	return "pageBlockVerticalAlignmentTop"
}

// TypeInfo returns info about TL type.
func (p *PageBlockVerticalAlignmentTop) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageBlockVerticalAlignmentTop",
		ID:   PageBlockVerticalAlignmentTopTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageBlockVerticalAlignmentTop) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockVerticalAlignmentTop#ba719a6 as nil")
	}
	b.PutID(PageBlockVerticalAlignmentTopTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageBlockVerticalAlignmentTop) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockVerticalAlignmentTop#ba719a6 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageBlockVerticalAlignmentTop) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockVerticalAlignmentTop#ba719a6 to nil")
	}
	if err := b.ConsumeID(PageBlockVerticalAlignmentTopTypeID); err != nil {
		return fmt.Errorf("unable to decode pageBlockVerticalAlignmentTop#ba719a6: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageBlockVerticalAlignmentTop) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockVerticalAlignmentTop#ba719a6 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (p *PageBlockVerticalAlignmentTop) EncodeTDLibJSON(b jsontd.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockVerticalAlignmentTop#ba719a6 as nil")
	}
	b.ObjStart()
	b.PutID("pageBlockVerticalAlignmentTop")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (p *PageBlockVerticalAlignmentTop) DecodeTDLibJSON(b jsontd.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockVerticalAlignmentTop#ba719a6 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("pageBlockVerticalAlignmentTop"); err != nil {
				return fmt.Errorf("unable to decode pageBlockVerticalAlignmentTop#ba719a6: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PageBlockVerticalAlignmentMiddle represents TL type `pageBlockVerticalAlignmentMiddle#81741df5`.
type PageBlockVerticalAlignmentMiddle struct {
}

// PageBlockVerticalAlignmentMiddleTypeID is TL type id of PageBlockVerticalAlignmentMiddle.
const PageBlockVerticalAlignmentMiddleTypeID = 0x81741df5

// construct implements constructor of PageBlockVerticalAlignmentClass.
func (p PageBlockVerticalAlignmentMiddle) construct() PageBlockVerticalAlignmentClass { return &p }

// Ensuring interfaces in compile-time for PageBlockVerticalAlignmentMiddle.
var (
	_ bin.Encoder     = &PageBlockVerticalAlignmentMiddle{}
	_ bin.Decoder     = &PageBlockVerticalAlignmentMiddle{}
	_ bin.BareEncoder = &PageBlockVerticalAlignmentMiddle{}
	_ bin.BareDecoder = &PageBlockVerticalAlignmentMiddle{}

	_ PageBlockVerticalAlignmentClass = &PageBlockVerticalAlignmentMiddle{}
)

func (p *PageBlockVerticalAlignmentMiddle) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageBlockVerticalAlignmentMiddle) String() string {
	if p == nil {
		return "PageBlockVerticalAlignmentMiddle(nil)"
	}
	type Alias PageBlockVerticalAlignmentMiddle
	return fmt.Sprintf("PageBlockVerticalAlignmentMiddle%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageBlockVerticalAlignmentMiddle) TypeID() uint32 {
	return PageBlockVerticalAlignmentMiddleTypeID
}

// TypeName returns name of type in TL schema.
func (*PageBlockVerticalAlignmentMiddle) TypeName() string {
	return "pageBlockVerticalAlignmentMiddle"
}

// TypeInfo returns info about TL type.
func (p *PageBlockVerticalAlignmentMiddle) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageBlockVerticalAlignmentMiddle",
		ID:   PageBlockVerticalAlignmentMiddleTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageBlockVerticalAlignmentMiddle) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockVerticalAlignmentMiddle#81741df5 as nil")
	}
	b.PutID(PageBlockVerticalAlignmentMiddleTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageBlockVerticalAlignmentMiddle) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockVerticalAlignmentMiddle#81741df5 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageBlockVerticalAlignmentMiddle) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockVerticalAlignmentMiddle#81741df5 to nil")
	}
	if err := b.ConsumeID(PageBlockVerticalAlignmentMiddleTypeID); err != nil {
		return fmt.Errorf("unable to decode pageBlockVerticalAlignmentMiddle#81741df5: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageBlockVerticalAlignmentMiddle) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockVerticalAlignmentMiddle#81741df5 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (p *PageBlockVerticalAlignmentMiddle) EncodeTDLibJSON(b jsontd.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockVerticalAlignmentMiddle#81741df5 as nil")
	}
	b.ObjStart()
	b.PutID("pageBlockVerticalAlignmentMiddle")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (p *PageBlockVerticalAlignmentMiddle) DecodeTDLibJSON(b jsontd.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockVerticalAlignmentMiddle#81741df5 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("pageBlockVerticalAlignmentMiddle"); err != nil {
				return fmt.Errorf("unable to decode pageBlockVerticalAlignmentMiddle#81741df5: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PageBlockVerticalAlignmentBottom represents TL type `pageBlockVerticalAlignmentBottom#7cb97dd6`.
type PageBlockVerticalAlignmentBottom struct {
}

// PageBlockVerticalAlignmentBottomTypeID is TL type id of PageBlockVerticalAlignmentBottom.
const PageBlockVerticalAlignmentBottomTypeID = 0x7cb97dd6

// construct implements constructor of PageBlockVerticalAlignmentClass.
func (p PageBlockVerticalAlignmentBottom) construct() PageBlockVerticalAlignmentClass { return &p }

// Ensuring interfaces in compile-time for PageBlockVerticalAlignmentBottom.
var (
	_ bin.Encoder     = &PageBlockVerticalAlignmentBottom{}
	_ bin.Decoder     = &PageBlockVerticalAlignmentBottom{}
	_ bin.BareEncoder = &PageBlockVerticalAlignmentBottom{}
	_ bin.BareDecoder = &PageBlockVerticalAlignmentBottom{}

	_ PageBlockVerticalAlignmentClass = &PageBlockVerticalAlignmentBottom{}
)

func (p *PageBlockVerticalAlignmentBottom) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageBlockVerticalAlignmentBottom) String() string {
	if p == nil {
		return "PageBlockVerticalAlignmentBottom(nil)"
	}
	type Alias PageBlockVerticalAlignmentBottom
	return fmt.Sprintf("PageBlockVerticalAlignmentBottom%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageBlockVerticalAlignmentBottom) TypeID() uint32 {
	return PageBlockVerticalAlignmentBottomTypeID
}

// TypeName returns name of type in TL schema.
func (*PageBlockVerticalAlignmentBottom) TypeName() string {
	return "pageBlockVerticalAlignmentBottom"
}

// TypeInfo returns info about TL type.
func (p *PageBlockVerticalAlignmentBottom) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageBlockVerticalAlignmentBottom",
		ID:   PageBlockVerticalAlignmentBottomTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageBlockVerticalAlignmentBottom) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockVerticalAlignmentBottom#7cb97dd6 as nil")
	}
	b.PutID(PageBlockVerticalAlignmentBottomTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageBlockVerticalAlignmentBottom) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockVerticalAlignmentBottom#7cb97dd6 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageBlockVerticalAlignmentBottom) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockVerticalAlignmentBottom#7cb97dd6 to nil")
	}
	if err := b.ConsumeID(PageBlockVerticalAlignmentBottomTypeID); err != nil {
		return fmt.Errorf("unable to decode pageBlockVerticalAlignmentBottom#7cb97dd6: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageBlockVerticalAlignmentBottom) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockVerticalAlignmentBottom#7cb97dd6 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (p *PageBlockVerticalAlignmentBottom) EncodeTDLibJSON(b jsontd.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockVerticalAlignmentBottom#7cb97dd6 as nil")
	}
	b.ObjStart()
	b.PutID("pageBlockVerticalAlignmentBottom")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (p *PageBlockVerticalAlignmentBottom) DecodeTDLibJSON(b jsontd.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockVerticalAlignmentBottom#7cb97dd6 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("pageBlockVerticalAlignmentBottom"); err != nil {
				return fmt.Errorf("unable to decode pageBlockVerticalAlignmentBottom#7cb97dd6: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PageBlockVerticalAlignmentClass represents PageBlockVerticalAlignment generic type.
//
// Example:
//  g, err := tdapi.DecodePageBlockVerticalAlignment(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.PageBlockVerticalAlignmentTop: // pageBlockVerticalAlignmentTop#ba719a6
//  case *tdapi.PageBlockVerticalAlignmentMiddle: // pageBlockVerticalAlignmentMiddle#81741df5
//  case *tdapi.PageBlockVerticalAlignmentBottom: // pageBlockVerticalAlignmentBottom#7cb97dd6
//  default: panic(v)
//  }
type PageBlockVerticalAlignmentClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PageBlockVerticalAlignmentClass

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

// DecodePageBlockVerticalAlignment implements binary de-serialization for PageBlockVerticalAlignmentClass.
func DecodePageBlockVerticalAlignment(buf *bin.Buffer) (PageBlockVerticalAlignmentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PageBlockVerticalAlignmentTopTypeID:
		// Decoding pageBlockVerticalAlignmentTop#ba719a6.
		v := PageBlockVerticalAlignmentTop{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageBlockVerticalAlignmentClass: %w", err)
		}
		return &v, nil
	case PageBlockVerticalAlignmentMiddleTypeID:
		// Decoding pageBlockVerticalAlignmentMiddle#81741df5.
		v := PageBlockVerticalAlignmentMiddle{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageBlockVerticalAlignmentClass: %w", err)
		}
		return &v, nil
	case PageBlockVerticalAlignmentBottomTypeID:
		// Decoding pageBlockVerticalAlignmentBottom#7cb97dd6.
		v := PageBlockVerticalAlignmentBottom{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageBlockVerticalAlignmentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PageBlockVerticalAlignmentClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONPageBlockVerticalAlignment implements binary de-serialization for PageBlockVerticalAlignmentClass.
func DecodeTDLibJSONPageBlockVerticalAlignment(buf jsontd.Decoder) (PageBlockVerticalAlignmentClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "pageBlockVerticalAlignmentTop":
		// Decoding pageBlockVerticalAlignmentTop#ba719a6.
		v := PageBlockVerticalAlignmentTop{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageBlockVerticalAlignmentClass: %w", err)
		}
		return &v, nil
	case "pageBlockVerticalAlignmentMiddle":
		// Decoding pageBlockVerticalAlignmentMiddle#81741df5.
		v := PageBlockVerticalAlignmentMiddle{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageBlockVerticalAlignmentClass: %w", err)
		}
		return &v, nil
	case "pageBlockVerticalAlignmentBottom":
		// Decoding pageBlockVerticalAlignmentBottom#7cb97dd6.
		v := PageBlockVerticalAlignmentBottom{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageBlockVerticalAlignmentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PageBlockVerticalAlignmentClass: %w", jsontd.NewUnexpectedID(id))
	}
}

// PageBlockVerticalAlignment boxes the PageBlockVerticalAlignmentClass providing a helper.
type PageBlockVerticalAlignmentBox struct {
	PageBlockVerticalAlignment PageBlockVerticalAlignmentClass
}

// Decode implements bin.Decoder for PageBlockVerticalAlignmentBox.
func (b *PageBlockVerticalAlignmentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PageBlockVerticalAlignmentBox to nil")
	}
	v, err := DecodePageBlockVerticalAlignment(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PageBlockVerticalAlignment = v
	return nil
}

// Encode implements bin.Encode for PageBlockVerticalAlignmentBox.
func (b *PageBlockVerticalAlignmentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PageBlockVerticalAlignment == nil {
		return fmt.Errorf("unable to encode PageBlockVerticalAlignmentClass as nil")
	}
	return b.PageBlockVerticalAlignment.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for PageBlockVerticalAlignmentBox.
func (b *PageBlockVerticalAlignmentBox) DecodeTDLibJSON(buf jsontd.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode PageBlockVerticalAlignmentBox to nil")
	}
	v, err := DecodeTDLibJSONPageBlockVerticalAlignment(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PageBlockVerticalAlignment = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for PageBlockVerticalAlignmentBox.
func (b *PageBlockVerticalAlignmentBox) EncodeTDLibJSON(buf jsontd.Encoder) error {
	if b == nil || b.PageBlockVerticalAlignment == nil {
		return fmt.Errorf("unable to encode PageBlockVerticalAlignmentClass as nil")
	}
	return b.PageBlockVerticalAlignment.EncodeTDLibJSON(buf)
}
