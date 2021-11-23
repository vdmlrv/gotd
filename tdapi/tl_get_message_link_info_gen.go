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

// GetMessageLinkInfoRequest represents TL type `getMessageLinkInfo#d63eb458`.
type GetMessageLinkInfoRequest struct {
	// The message link
	URL string
}

// GetMessageLinkInfoRequestTypeID is TL type id of GetMessageLinkInfoRequest.
const GetMessageLinkInfoRequestTypeID = 0xd63eb458

// Ensuring interfaces in compile-time for GetMessageLinkInfoRequest.
var (
	_ bin.Encoder     = &GetMessageLinkInfoRequest{}
	_ bin.Decoder     = &GetMessageLinkInfoRequest{}
	_ bin.BareEncoder = &GetMessageLinkInfoRequest{}
	_ bin.BareDecoder = &GetMessageLinkInfoRequest{}
)

func (g *GetMessageLinkInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessageLinkInfoRequest) String() string {
	if g == nil {
		return "GetMessageLinkInfoRequest(nil)"
	}
	type Alias GetMessageLinkInfoRequest
	return fmt.Sprintf("GetMessageLinkInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageLinkInfoRequest) TypeID() uint32 {
	return GetMessageLinkInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageLinkInfoRequest) TypeName() string {
	return "getMessageLinkInfo"
}

// TypeInfo returns info about TL type.
func (g *GetMessageLinkInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessageLinkInfo",
		ID:   GetMessageLinkInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMessageLinkInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageLinkInfo#d63eb458 as nil")
	}
	b.PutID(GetMessageLinkInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageLinkInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageLinkInfo#d63eb458 as nil")
	}
	b.PutString(g.URL)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageLinkInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageLinkInfo#d63eb458 to nil")
	}
	if err := b.ConsumeID(GetMessageLinkInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessageLinkInfo#d63eb458: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageLinkInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageLinkInfo#d63eb458 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageLinkInfo#d63eb458: field url: %w", err)
		}
		g.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (g *GetMessageLinkInfoRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageLinkInfo#d63eb458 as nil")
	}
	b.ObjStart()
	b.PutID("getMessageLinkInfo")
	b.FieldStart("url")
	b.PutString(g.URL)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (g *GetMessageLinkInfoRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageLinkInfo#d63eb458 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("getMessageLinkInfo"); err != nil {
				return fmt.Errorf("unable to decode getMessageLinkInfo#d63eb458: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageLinkInfo#d63eb458: field url: %w", err)
			}
			g.URL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (g *GetMessageLinkInfoRequest) GetURL() (value string) {
	return g.URL
}

// GetMessageLinkInfo invokes method getMessageLinkInfo#d63eb458 returning error if any.
func (c *Client) GetMessageLinkInfo(ctx context.Context, url string) (*MessageLinkInfo, error) {
	var result MessageLinkInfo

	request := &GetMessageLinkInfoRequest{
		URL: url,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
