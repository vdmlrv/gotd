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

// GetCountryCodeRequest represents TL type `getCountryCode#5bd398f2`.
type GetCountryCodeRequest struct {
}

// GetCountryCodeRequestTypeID is TL type id of GetCountryCodeRequest.
const GetCountryCodeRequestTypeID = 0x5bd398f2

// Ensuring interfaces in compile-time for GetCountryCodeRequest.
var (
	_ bin.Encoder     = &GetCountryCodeRequest{}
	_ bin.Decoder     = &GetCountryCodeRequest{}
	_ bin.BareEncoder = &GetCountryCodeRequest{}
	_ bin.BareDecoder = &GetCountryCodeRequest{}
)

func (g *GetCountryCodeRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCountryCodeRequest) String() string {
	if g == nil {
		return "GetCountryCodeRequest(nil)"
	}
	type Alias GetCountryCodeRequest
	return fmt.Sprintf("GetCountryCodeRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCountryCodeRequest) TypeID() uint32 {
	return GetCountryCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCountryCodeRequest) TypeName() string {
	return "getCountryCode"
}

// TypeInfo returns info about TL type.
func (g *GetCountryCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCountryCode",
		ID:   GetCountryCodeRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCountryCodeRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCountryCode#5bd398f2 as nil")
	}
	b.PutID(GetCountryCodeRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCountryCodeRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCountryCode#5bd398f2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCountryCodeRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCountryCode#5bd398f2 to nil")
	}
	if err := b.ConsumeID(GetCountryCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCountryCode#5bd398f2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCountryCodeRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCountryCode#5bd398f2 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (g *GetCountryCodeRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getCountryCode#5bd398f2 as nil")
	}
	b.ObjStart()
	b.PutID("getCountryCode")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (g *GetCountryCodeRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getCountryCode#5bd398f2 to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("getCountryCode"); err != nil {
				return fmt.Errorf("unable to decode getCountryCode#5bd398f2: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCountryCode invokes method getCountryCode#5bd398f2 returning error if any.
func (c *Client) GetCountryCode(ctx context.Context) (*Text, error) {
	var result Text

	request := &GetCountryCodeRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
