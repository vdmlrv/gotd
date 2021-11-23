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

// CreateChatFilterRequest represents TL type `createChatFilter#2ecaca6`.
type CreateChatFilterRequest struct {
	// Chat filter
	Filter ChatFilter
}

// CreateChatFilterRequestTypeID is TL type id of CreateChatFilterRequest.
const CreateChatFilterRequestTypeID = 0x2ecaca6

// Ensuring interfaces in compile-time for CreateChatFilterRequest.
var (
	_ bin.Encoder     = &CreateChatFilterRequest{}
	_ bin.Decoder     = &CreateChatFilterRequest{}
	_ bin.BareEncoder = &CreateChatFilterRequest{}
	_ bin.BareDecoder = &CreateChatFilterRequest{}
)

func (c *CreateChatFilterRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Filter.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreateChatFilterRequest) String() string {
	if c == nil {
		return "CreateChatFilterRequest(nil)"
	}
	type Alias CreateChatFilterRequest
	return fmt.Sprintf("CreateChatFilterRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreateChatFilterRequest) TypeID() uint32 {
	return CreateChatFilterRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreateChatFilterRequest) TypeName() string {
	return "createChatFilter"
}

// TypeInfo returns info about TL type.
func (c *CreateChatFilterRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createChatFilter",
		ID:   CreateChatFilterRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreateChatFilterRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createChatFilter#2ecaca6 as nil")
	}
	b.PutID(CreateChatFilterRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreateChatFilterRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createChatFilter#2ecaca6 as nil")
	}
	if err := c.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode createChatFilter#2ecaca6: field filter: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CreateChatFilterRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createChatFilter#2ecaca6 to nil")
	}
	if err := b.ConsumeID(CreateChatFilterRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createChatFilter#2ecaca6: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreateChatFilterRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createChatFilter#2ecaca6 to nil")
	}
	{
		if err := c.Filter.Decode(b); err != nil {
			return fmt.Errorf("unable to decode createChatFilter#2ecaca6: field filter: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (c *CreateChatFilterRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode createChatFilter#2ecaca6 as nil")
	}
	b.ObjStart()
	b.PutID("createChatFilter")
	b.FieldStart("filter")
	if err := c.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode createChatFilter#2ecaca6: field filter: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (c *CreateChatFilterRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode createChatFilter#2ecaca6 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("createChatFilter"); err != nil {
				return fmt.Errorf("unable to decode createChatFilter#2ecaca6: %w", err)
			}
		case "filter":
			if err := c.Filter.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode createChatFilter#2ecaca6: field filter: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFilter returns value of Filter field.
func (c *CreateChatFilterRequest) GetFilter() (value ChatFilter) {
	return c.Filter
}

// CreateChatFilter invokes method createChatFilter#2ecaca6 returning error if any.
func (c *Client) CreateChatFilter(ctx context.Context, filter ChatFilter) (*ChatFilterInfo, error) {
	var result ChatFilterInfo

	request := &CreateChatFilterRequest{
		Filter: filter,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
