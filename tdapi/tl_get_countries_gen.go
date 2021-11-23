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

// GetCountriesRequest represents TL type `getCountries#fce8099e`.
type GetCountriesRequest struct {
}

// GetCountriesRequestTypeID is TL type id of GetCountriesRequest.
const GetCountriesRequestTypeID = 0xfce8099e

// Ensuring interfaces in compile-time for GetCountriesRequest.
var (
	_ bin.Encoder     = &GetCountriesRequest{}
	_ bin.Decoder     = &GetCountriesRequest{}
	_ bin.BareEncoder = &GetCountriesRequest{}
	_ bin.BareDecoder = &GetCountriesRequest{}
)

func (g *GetCountriesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCountriesRequest) String() string {
	if g == nil {
		return "GetCountriesRequest(nil)"
	}
	type Alias GetCountriesRequest
	return fmt.Sprintf("GetCountriesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCountriesRequest) TypeID() uint32 {
	return GetCountriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCountriesRequest) TypeName() string {
	return "getCountries"
}

// TypeInfo returns info about TL type.
func (g *GetCountriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCountries",
		ID:   GetCountriesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCountriesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCountries#fce8099e as nil")
	}
	b.PutID(GetCountriesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCountriesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCountries#fce8099e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCountriesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCountries#fce8099e to nil")
	}
	if err := b.ConsumeID(GetCountriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCountries#fce8099e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCountriesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCountries#fce8099e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements jsontd.TDLibEncoder.
func (g *GetCountriesRequest) EncodeTDLibJSON(b jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getCountries#fce8099e as nil")
	}
	b.ObjStart()
	b.PutID("getCountries")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements jsontd.TDLibDecoder.
func (g *GetCountriesRequest) DecodeTDLibJSON(b jsontd.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getCountries#fce8099e to nil")
	}

	return b.Obj(func(b jsontd.Decoder, key []byte) error {
		switch string(key) {
		case jsontd.TypeField:
			if err := b.ConsumeID("getCountries"); err != nil {
				return fmt.Errorf("unable to decode getCountries#fce8099e: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCountries invokes method getCountries#fce8099e returning error if any.
func (c *Client) GetCountries(ctx context.Context) (*Countries, error) {
	var result Countries

	request := &GetCountriesRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
