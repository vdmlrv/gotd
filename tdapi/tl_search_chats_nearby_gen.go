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

// SearchChatsNearbyRequest represents TL type `searchChatsNearby#f445c81f`.
type SearchChatsNearbyRequest struct {
	// Current user location
	Location Location
}

// SearchChatsNearbyRequestTypeID is TL type id of SearchChatsNearbyRequest.
const SearchChatsNearbyRequestTypeID = 0xf445c81f

// Ensuring interfaces in compile-time for SearchChatsNearbyRequest.
var (
	_ bin.Encoder     = &SearchChatsNearbyRequest{}
	_ bin.Decoder     = &SearchChatsNearbyRequest{}
	_ bin.BareEncoder = &SearchChatsNearbyRequest{}
	_ bin.BareDecoder = &SearchChatsNearbyRequest{}
)

func (s *SearchChatsNearbyRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Location.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchChatsNearbyRequest) String() string {
	if s == nil {
		return "SearchChatsNearbyRequest(nil)"
	}
	type Alias SearchChatsNearbyRequest
	return fmt.Sprintf("SearchChatsNearbyRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchChatsNearbyRequest) TypeID() uint32 {
	return SearchChatsNearbyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchChatsNearbyRequest) TypeName() string {
	return "searchChatsNearby"
}

// TypeInfo returns info about TL type.
func (s *SearchChatsNearbyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchChatsNearby",
		ID:   SearchChatsNearbyRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Location",
			SchemaName: "location",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchChatsNearbyRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchChatsNearby#f445c81f as nil")
	}
	b.PutID(SearchChatsNearbyRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchChatsNearbyRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchChatsNearby#f445c81f as nil")
	}
	if err := s.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode searchChatsNearby#f445c81f: field location: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchChatsNearbyRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchChatsNearby#f445c81f to nil")
	}
	if err := b.ConsumeID(SearchChatsNearbyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchChatsNearby#f445c81f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchChatsNearbyRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchChatsNearby#f445c81f to nil")
	}
	{
		if err := s.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode searchChatsNearby#f445c81f: field location: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (s *SearchChatsNearbyRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchChatsNearby#f445c81f as nil")
	}
	b.ObjStart()
	b.PutID("searchChatsNearby")
	b.FieldStart("location")
	if err := s.Location.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode searchChatsNearby#f445c81f: field location: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (s *SearchChatsNearbyRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchChatsNearby#f445c81f to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("searchChatsNearby"); err != nil {
				return fmt.Errorf("unable to decode searchChatsNearby#f445c81f: %w", err)
			}
		case "location":
			if err := s.Location.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode searchChatsNearby#f445c81f: field location: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLocation returns value of Location field.
func (s *SearchChatsNearbyRequest) GetLocation() (value Location) {
	return s.Location
}

// SearchChatsNearby invokes method searchChatsNearby#f445c81f returning error if any.
func (c *Client) SearchChatsNearby(ctx context.Context, location Location) (*ChatsNearby, error) {
	var result ChatsNearby

	request := &SearchChatsNearbyRequest{
		Location: location,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
