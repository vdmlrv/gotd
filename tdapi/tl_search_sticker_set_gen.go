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

// SearchStickerSetRequest represents TL type `searchStickerSet#45049cee`.
type SearchStickerSetRequest struct {
	// Name of the sticker set
	Name string
}

// SearchStickerSetRequestTypeID is TL type id of SearchStickerSetRequest.
const SearchStickerSetRequestTypeID = 0x45049cee

// Ensuring interfaces in compile-time for SearchStickerSetRequest.
var (
	_ bin.Encoder     = &SearchStickerSetRequest{}
	_ bin.Decoder     = &SearchStickerSetRequest{}
	_ bin.BareEncoder = &SearchStickerSetRequest{}
	_ bin.BareDecoder = &SearchStickerSetRequest{}
)

func (s *SearchStickerSetRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchStickerSetRequest) String() string {
	if s == nil {
		return "SearchStickerSetRequest(nil)"
	}
	type Alias SearchStickerSetRequest
	return fmt.Sprintf("SearchStickerSetRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchStickerSetRequest) TypeID() uint32 {
	return SearchStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchStickerSetRequest) TypeName() string {
	return "searchStickerSet"
}

// TypeInfo returns info about TL type.
func (s *SearchStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchStickerSet",
		ID:   SearchStickerSetRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchStickerSetRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchStickerSet#45049cee as nil")
	}
	b.PutID(SearchStickerSetRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchStickerSet#45049cee as nil")
	}
	b.PutString(s.Name)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchStickerSetRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchStickerSet#45049cee to nil")
	}
	if err := b.ConsumeID(SearchStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchStickerSet#45049cee: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchStickerSet#45049cee to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchStickerSet#45049cee: field name: %w", err)
		}
		s.Name = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (s *SearchStickerSetRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchStickerSet#45049cee as nil")
	}
	b.ObjStart()
	b.PutID("searchStickerSet")
	b.FieldStart("name")
	b.PutString(s.Name)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (s *SearchStickerSetRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchStickerSet#45049cee to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("searchStickerSet"); err != nil {
				return fmt.Errorf("unable to decode searchStickerSet#45049cee: %w", err)
			}
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchStickerSet#45049cee: field name: %w", err)
			}
			s.Name = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetName returns value of Name field.
func (s *SearchStickerSetRequest) GetName() (value string) {
	return s.Name
}

// SearchStickerSet invokes method searchStickerSet#45049cee returning error if any.
func (c *Client) SearchStickerSet(ctx context.Context, name string) (*StickerSet, error) {
	var result StickerSet

	request := &SearchStickerSetRequest{
		Name: name,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
