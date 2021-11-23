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

// SearchCallMessagesRequest represents TL type `searchCallMessages#bfcac31c`.
type SearchCallMessagesRequest struct {
	// Identifier of the message from which to search; use 0 to get results from the last
	// message
	FromMessageID int64
	// The maximum number of messages to be returned; up to 100. For optimal performance, the
	// number of returned messages is chosen by TDLib and can be smaller than the specified
	// limit
	Limit int32
	// If true, returns only messages with missed calls
	OnlyMissed bool
}

// SearchCallMessagesRequestTypeID is TL type id of SearchCallMessagesRequest.
const SearchCallMessagesRequestTypeID = 0xbfcac31c

// Ensuring interfaces in compile-time for SearchCallMessagesRequest.
var (
	_ bin.Encoder     = &SearchCallMessagesRequest{}
	_ bin.Decoder     = &SearchCallMessagesRequest{}
	_ bin.BareEncoder = &SearchCallMessagesRequest{}
	_ bin.BareDecoder = &SearchCallMessagesRequest{}
)

func (s *SearchCallMessagesRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.FromMessageID == 0) {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}
	if !(s.OnlyMissed == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchCallMessagesRequest) String() string {
	if s == nil {
		return "SearchCallMessagesRequest(nil)"
	}
	type Alias SearchCallMessagesRequest
	return fmt.Sprintf("SearchCallMessagesRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchCallMessagesRequest) TypeID() uint32 {
	return SearchCallMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchCallMessagesRequest) TypeName() string {
	return "searchCallMessages"
}

// TypeInfo returns info about TL type.
func (s *SearchCallMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchCallMessages",
		ID:   SearchCallMessagesRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FromMessageID",
			SchemaName: "from_message_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "OnlyMissed",
			SchemaName: "only_missed",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchCallMessagesRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchCallMessages#bfcac31c as nil")
	}
	b.PutID(SearchCallMessagesRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchCallMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchCallMessages#bfcac31c as nil")
	}
	b.PutLong(s.FromMessageID)
	b.PutInt32(s.Limit)
	b.PutBool(s.OnlyMissed)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchCallMessagesRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchCallMessages#bfcac31c to nil")
	}
	if err := b.ConsumeID(SearchCallMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchCallMessages#bfcac31c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchCallMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchCallMessages#bfcac31c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode searchCallMessages#bfcac31c: field from_message_id: %w", err)
		}
		s.FromMessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchCallMessages#bfcac31c: field limit: %w", err)
		}
		s.Limit = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode searchCallMessages#bfcac31c: field only_missed: %w", err)
		}
		s.OnlyMissed = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (s *SearchCallMessagesRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchCallMessages#bfcac31c as nil")
	}
	b.ObjStart()
	b.PutID("searchCallMessages")
	b.FieldStart("from_message_id")
	b.PutLong(s.FromMessageID)
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.FieldStart("only_missed")
	b.PutBool(s.OnlyMissed)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (s *SearchCallMessagesRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchCallMessages#bfcac31c to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("searchCallMessages"); err != nil {
				return fmt.Errorf("unable to decode searchCallMessages#bfcac31c: %w", err)
			}
		case "from_message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode searchCallMessages#bfcac31c: field from_message_id: %w", err)
			}
			s.FromMessageID = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchCallMessages#bfcac31c: field limit: %w", err)
			}
			s.Limit = value
		case "only_missed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode searchCallMessages#bfcac31c: field only_missed: %w", err)
			}
			s.OnlyMissed = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFromMessageID returns value of FromMessageID field.
func (s *SearchCallMessagesRequest) GetFromMessageID() (value int64) {
	return s.FromMessageID
}

// GetLimit returns value of Limit field.
func (s *SearchCallMessagesRequest) GetLimit() (value int32) {
	return s.Limit
}

// GetOnlyMissed returns value of OnlyMissed field.
func (s *SearchCallMessagesRequest) GetOnlyMissed() (value bool) {
	return s.OnlyMissed
}

// SearchCallMessages invokes method searchCallMessages#bfcac31c returning error if any.
func (c *Client) SearchCallMessages(ctx context.Context, request *SearchCallMessagesRequest) (*Messages, error) {
	var result Messages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
